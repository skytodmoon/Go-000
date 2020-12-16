package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name      string `yaml:"SiteName"`
	Addr      string `yaml:"SiteAddr"`
	HTTPS     bool   `yaml:"Https"`
	SiteNginx Nginx  `yaml:"Nginx"`
}

func main() {
	var setting Config
	config, err := ioutil.ReadFile("./configs/config.yaml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &setting)
}
