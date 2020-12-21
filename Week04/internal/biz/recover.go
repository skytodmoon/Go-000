package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Recover struct {
	logger *logrus.Logger
}

func NewRecover(logger *logrus.Logger) *Recover {
	return &Recover{logger: logger}
}

func (r *Recover) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
