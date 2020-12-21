package config

import (
	"fmt"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Config struct {
	Http  HttpConfig  `mapstructure:"http"`
	Mysql MysqlConfig `mapstructure:"mysql"`
}

type HttpConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	RunModel string `mapstructure:"run_model"`
	Version  string `mapstructure:"version"`
}

type MysqlConfig struct {
	URL         string `mapstructure:"url"`
	MaxIdle     int    `mapstructure:"max_idle"`
	MaxOpen     int    `mapstructure:"max_open"`
	MaxLeftTime int    `mapstructure:"max_leftTime"`
	Debug       bool   `mapstructure:"debug"`
}

func NewConfig(path string) *Config {
	v := viper.New()
	v.SetConfigName("dev")
	v.SetConfigType("yaml")
	v.AddConfigPath("/config/")
	v.AddConfigPath(".")
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	var conf Config
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("Could not unmarshal config: %s", err))
	}
	return &conf
}

var WireSet = wire.NewSet(
	NewConfig,
)
