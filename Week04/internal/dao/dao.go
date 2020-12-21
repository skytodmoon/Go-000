package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Dao struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewDao(logger *logrus.Logger, db *gorm.DB) *Dao {
	return &Dao{logger: logger, db: db}
}

func (u *Dao) FindByUid(uid int) {
	uid = 0
	return
}
