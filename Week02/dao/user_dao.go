package dao

import (
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserEntity struct {
	ID       int64
	Username string
	Password string
	Email    string
	CreateAt time.Time
}

func (UserEntity) TableName() string {
	return "user"
}

type UserDAO interface {
	SelectByEmail(email string) (*UserEntity, error)
	//Save(user *UserEntity) error
}

type UserDAOImpl struct {
}

func (UserDAO *UserDAOImpl) SelectByEmail(email string) (*UserEntity, error) {
	user := &UserEntity{}
	err := db.Where("email = ?", email).First(&user).Error
	var e = errors.New("")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		e = errors.Wrap(err, "Dao query failed!")
	} else {
		e = err
	}
	return user, e
}
