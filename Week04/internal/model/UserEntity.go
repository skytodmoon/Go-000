package model

import "time"

type UserEntity struct {
	ID       int64
	Username string
	Password string
	Email    string
	CreateAt time.Time
}
