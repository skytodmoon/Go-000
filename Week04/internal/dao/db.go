package dao

import (
	"Week04/config"
	"Week04/internal/model"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

func NewDatabaseMYSQL(conf *config.Config, logger *logrus.Logger) (*gorm.DB, error) {
	opt := conf.Mysql
	db, err := gorm.Open("mysql", opt.URL)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(opt.MaxIdle)                                     //初始化数据库连接数
	db.DB().SetMaxOpenConns(opt.MaxOpen)                                     //额外增开
	db.DB().SetConnMaxLifetime(time.Duration(opt.MaxLeftTime) * time.Second) //链接时长
	db.SingularTable(true)

	if opt.Debug {
		db.LogMode(true)
	}
	return db, db.AutoMigrate(new(model.User)).Error
}
