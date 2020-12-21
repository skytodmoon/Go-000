package service

import (
	pb "Week04/api"
	"Week04/internal/dao"
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

var Provider = wire.NewSet(NewService, wire.Bind(new()))

type Service struct {
	dao dao.Dao
}

func NewService(d dao.Dao) *Service {
	return &Service{dao: d}
}

func (s *Service) QueryUserInfo(ctx context.Context, req *pb.QueryUserReq) (replay *pb.QueryUserResp, err error) {
	user, err := s.dao.GetUser(req.Email)
	if err == nil {
		replay = &pb.QueryUserResp{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}
	} else {
		err = errors.Wrap(err, "Service query failed!")
	}
	return
}

// Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
// Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
// Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
