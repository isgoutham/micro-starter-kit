package main

import (
    "net"

    "github.com/micro/cli/v2"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/config"
    "github.com/rs/zerolog/log"

    sgrpc "github.com/micro/go-micro/v2/server/grpc"

    "github.com/xmlking/micro-starter-kit/service/greeter/handler"
    greeterPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/greeter"
    healthPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/health"
    myConfig "github.com/xmlking/micro-starter-kit/shared/config"
    "github.com/xmlking/micro-starter-kit/shared/constants"
    _ "github.com/xmlking/micro-starter-kit/shared/logger"
    "github.com/xmlking/micro-starter-kit/shared/util"
    logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
    transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
)

const (
	serviceName = constants.GREETER_SERVICE
)

var (
    cfg = myConfig.GetServiceConfig()
    ff = myConfig.GetFeatureFlags()
)

func main() {
    // lis, err := net.Listen("tcp", ":0")
    lis, err := net.Listen("unix", "/tmp/greeter.sock") //  you can also use Unix Domain Sockets (UDS) as address
    if err != nil {
        log.Fatal().Msgf("failed to listen: %v", err)
    }
    println(lis.Addr().String())

	// New Service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(myConfig.Version),
        micro.Server(sgrpc.NewServer(sgrpc.Listener(lis))),
	)

	// Initialize service
	service.Init(
		micro.Action(func(c *cli.Context) (err error) {
			// do some life cycle actions
			return
		}),
	)

	// Initialize Features
	var options []micro.Option
	if ff.IsTLSEnabled() {
		if tlsConf, err := myConfig.CreateServerCerts(); err != nil {
			log.Error().Err(err).Msg("unable to load certs")
		} else {
            log.Info().Msg("TLS Enabled")
			options = append(options,
				util.WithTLS(tlsConf),
			)
		}
	}
	// Wrappers are invoked in the order as they added
    if ff.IsReqlogsEnabled() {
		options = append(options, micro.WrapHandler(logWrapper.NewHandlerWrapper()))
	}
    if ff.IsTranslogsEnabled() {
		topic := config.Get("features", "translogs", "topic").String(constants.RECORDER_SERVICE)
		publisher := micro.NewEvent(topic, service.Client())
		options = append(options, micro.WrapHandler(transWrapper.NewHandlerWrapper(publisher)))
	}

	// Initialize Features
	service.Init(
		options...,
	)

	// Register Handler
	_ = greeterPB.RegisterGreeterServiceHandler(service.Server(), handler.NewGreeterHandler())
    _ = healthPB.RegisterHealthHandler(service.Server(), handler.NewHealthHandler())

    println(myConfig.GetBuildInfo())

	// Run service
	if err := service.Run(); err != nil {
        log.Fatal().Err(err).Msg("")
	}
}
