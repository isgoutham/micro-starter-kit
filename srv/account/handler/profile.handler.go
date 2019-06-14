package handler

import (
	"context"
	"fmt"
	"time"

	// "github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	"github.com/xmlking/micro-starter-kit/srv/account/model"
	pb "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
	"github.com/xmlking/micro-starter-kit/srv/account/repository"
)

// ProfileHandler struct
type profileHandler struct {
	profileRepository repository.ProfileRepository
}

// NewProfileHandler returns an instance of `ProfileServiceHandler`.
func NewProfileHandler(repo repository.ProfileRepository) pb.ProfileServiceHandler {
	return &profileHandler{
		profileRepository: repo,
	}
}

func (h *profileHandler) List(ctx context.Context, req *pb.ProfileListQuery, rsp *pb.ProfileListResponse) error {
	log.Log("Received ProfileHandler.List request")
	total, profiles, err := h.profileRepository.List(req)
	if err != nil {
		return err
	}
	rsp.Total = total
	newProfiles := make([]*pb.Profile, len(profiles))
	for index, profile := range profiles {
		newProfiles[index] = profileModelToProto(profile)
	}
	rsp.Profiles = newProfiles
	rsp.Msg = fmt.Sprintf("%v Total Profiles Found", total) // "Profiles Found"
	rsp.Code = "200"
	return nil
}

func (h *profileHandler) Get(ctx context.Context, req *pb.ProfileRequest, rsp *pb.ProfileResponse) error {
	log.Log("Received ProfileHandler.Get request")
	if req.Id == 0 {
		return fmt.Errorf("missing Id")
	}

	profile, err := h.profileRepository.Get(req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.Profile = nil
			rsp.Msg = "Profile Not Found"
			rsp.Code = "404"
			return nil
		}
		return err
	}

	rsp.Profile = profileModelToProto(profile)
	rsp.Msg = "Profile Found"
	rsp.Code = "200"
	return nil
}

func (h *profileHandler) Create(ctx context.Context, req *pb.ProfileRequest, rsp *pb.ProfileResponse) error {
	log.Log("Received ProfileHandler.Create request")
	if err := h.profileRepository.Create(req); err != nil {
		return err
	}
	rsp.Msg = "Profile Created"
	rsp.Code = "200"
	return nil
}

func profileModelToProto(profile *model.Profile) (proto *pb.Profile) {
	// birthday, err := ptypes.TimestampProto(profile.Birthday)
	return &pb.Profile{
		Id:     uint32(profile.Model.ID),
		Tz:     profile.TZ,
		Avatar: profile.Avatar,
		Gender: profile.Gender,
		// Birthday:  birthday,
		UserId:    profile.UserID,
		CreatedAt: profile.Model.CreatedAt.Format(time.RFC3339),
		UpdatedAt: profile.Model.UpdatedAt.Format(time.RFC3339),
	}
}