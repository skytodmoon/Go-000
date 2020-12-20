package biz

import (
	"fmt"

	"gorm.io/gorm"

	"Week04/pkg/setting"

	"gorm.io/driver/mysql"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	// db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
	// 	databaseSetting.UserName,
	// 	databaseSetting.Password,
	// 	databaseSetting.Host,
	// 	databaseSetting.Port,
	// 	databaseSetting.DBName,
	// 	databaseSetting.Charset,
	// 	databaseSetting.ParseTime,
	// ))
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
