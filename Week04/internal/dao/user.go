package dao

import (
	"Week04/global"
	"Week04/internal/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (d *dao) GetUser(email string) (*model.UserEntity, error) {
	user := &model.UserEntity{}
	err := global.DBEngine.Where("email = ?", email).First(&user).Error
	var e = errors.New("")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		e = errors.Wrap(err, "Dao query failed!")
	} else {
		e = err
	}
	return user, e
}
