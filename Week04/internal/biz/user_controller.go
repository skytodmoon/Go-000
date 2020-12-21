package biz

import (
	"Week04/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	logger  *logrus.Logger
	userSvc *service.UserService
}

func NewUserController(logger *logrus.Logger, userSvc *service.UserService) *UserController {
	return &UserController{
		logger:  logger,
		userSvc: userSvc,
	}
}

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Router /user [post]
func (u *UserController) Create(c *gin.Context) {

}
