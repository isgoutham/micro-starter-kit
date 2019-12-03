package main

import (
	"path/filepath"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"

	"github.com/micro/go-micro/service/grpc"

	log "github.com/sirupsen/logrus"

	myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	logger "github.com/xmlking/micro-starter-kit/shared/log"
	logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
	transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
	validatorWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/validator"

	"github.com/xmlking/micro-starter-kit/shared/constants"
	"github.com/xmlking/micro-starter-kit/shared/util"
	"github.com/xmlking/micro-starter-kit/srv/account/handler"
	accountPB "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
	"github.com/xmlking/micro-starter-kit/srv/account/registry"
	"github.com/xmlking/micro-starter-kit/srv/account/repository"
	greeterPB "github.com/xmlking/micro-starter-kit/srv/greeter/proto/greeter"
)

const (
	serviceName = "accountsrv"
)

var (
	configDir  string
	configFile string
	cfg        myConfig.ServiceConfiguration
)

func main() {
	// New Service
	service := grpc.NewService(
		// optional cli flag to override config.
		// comment out if you don't need to override any base config via CLI
		micro.Flags(
			cli.StringFlag{
				Name:        "configDir, d",
				Value:       "/config",
				Usage:       "Path to the config directory. Defaults to 'config'",
				EnvVar:      "CONFIG_DIR",
				Destination: &configDir,
			},
			cli.StringFlag{
				Name:        "configFile, f",
				Value:       "config.yaml",
				Usage:       "Config file in configDir. Defaults to 'config.yaml'",
				EnvVar:      "CONFIG_FILE",
				Destination: &configFile,
			}),
		micro.Name(serviceName),
		micro.Version(myConfig.Version),
	)

	// Initialize service
	service.Init(
		micro.Action(func(c *cli.Context) {
			// load config
			myConfig.InitConfig(configDir, configFile)
			config.Scan(&cfg)
			logger.InitLogger(cfg.Log)
		}),
	)

	// Initialize Features
	// Wrappers are invoked in the order as they added
	var options []micro.Option
	if cfg.Features["mtls"].Enabled {
		// tlsConf, _ := util.GetSelfSignedTLSConfig("localhost")
		if tlsConf, err := util.GetTLSConfig(
			filepath.Join(configDir, config.Get("features", "mtls", "certfile").String("")),
			filepath.Join(configDir, config.Get("features", "mtls", "keyfile").String("")),
			filepath.Join(configDir, config.Get("features", "mtls", "cafile").String("")),
			filepath.Join(configDir, config.Get("features", "mtls", "servername").String("")),
		); err != nil {
			log.WithError(err).Fatal("unable to load certs")
		} else {
			options = append(options,
				// https://github.com/ykumar-rb/ZTP/blob/master/pnp/server.go
				grpc.WithTLS(tlsConf),
			)
		}
	}

	if cfg.Features["reqlogs"].Enabled {
		options = append(options,
			micro.WrapHandler(logWrapper.NewHandlerWrapper()),
			micro.WrapClient(logWrapper.NewClientWrapper()),
		)
	}
	if cfg.Features["validator"].Enabled {
		options = append(options,
			micro.WrapHandler(validatorWrapper.NewHandlerWrapper()),
		)
	}
	if cfg.Features["translogs"].Enabled {
		topic := config.Get("features", "translogs", "topic").String("recordersrv")
		publisher := micro.NewPublisher(topic, service.Client())
		options = append(options,
			micro.WrapHandler(transWrapper.NewHandlerWrapper(publisher)),
		)
	}

	// Initialize Features
	service.Init(
		options...,
	)

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	log.Debugf("Client type: grpc or regular? %T\n", service.Client()) // FIXME: expected *grpc.grpcClient but got *micro.clientWrapper

	// Publisher publish to "emailersrv"
	emailerSrvEp := config.Get("services", constants.EMAILERSRV, "endpoint").String(constants.EMAILERSRV)
	publisher := micro.NewPublisher(emailerSrvEp, service.Client())
	// greeterSrv Client to call "greetersrv"
	greeterSrvEp := config.Get("services", constants.GREETERSRV, "endpoint").String(constants.GREETERSRV)
	greeterSrvClient := greeterPB.NewGreeterService(greeterSrvEp, service.Client())

	// // Handlers
	userHandler := handler.NewUserHandler(ctn.Resolve("user-repository").(repository.UserRepository), publisher, greeterSrvClient)
	profileHandler := ctn.Resolve("profile-handler").(accountPB.ProfileServiceHandler)

	// Register Handlers
	accountPB.RegisterUserServiceHandler(service.Server(), userHandler)
	accountPB.RegisterProfileServiceHandler(service.Server(), profileHandler)

	myConfig.PrintBuildInfo()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
