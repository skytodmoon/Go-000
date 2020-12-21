package service

import (
	"Week04/internal/dao"

	"github.com/sirupsen/logrus"
)

type UserService struct {
	dao    *dao.Dao
	logger *logrus.Logger
}

func NewService(d *dao.Dao, logger *logrus.Logger) *UserService {
	return &UserService{dao: d}
}

func (s *UserService) FindByUid(uid int) {
	s.dao.FindByUid(uid)
}
