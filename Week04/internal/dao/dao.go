package dao

import (
	"Week04/internal/model"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var Provider = wire.NewSet(NewDao, NewDBEngine)

type Dao interface {
	GetUser(email string) (*model.UserEntity, error)
}

type dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *dao {
	return &dao{db: db}
}
