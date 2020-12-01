package service

import (
	"Week02/dao"
	"context"

	"github.com/pkg/errors"
)

type UserInfoDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var (
	ErrUserNotExisted = errors.New("user is existed")
)

type UserService interface {
	//用户查询
	QueryUser(ctx context.Context, email string) (*UserInfoDTO, error)
}

type UserServiceImpl struct {
	userDAO dao.UserDAO
}

func MakeUserServiceImpl(userDAO dao.UserDAO) UserService {
	return &UserServiceImpl{
		userDAO: userDAO,
	}
}

func (userService *UserServiceImpl) QueryUser(ctx context.Context, email string) (*UserInfoDTO, error) {
	user, err := userService.userDAO.SelectByEmail(email)
	if err == nil {
		return &UserInfoDTO{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}, nil
	} else {
		return nil, errors.Wrap(err, "Service query failed!")
	}

}
