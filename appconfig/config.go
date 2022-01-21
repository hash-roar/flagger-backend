package appconfig

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type AppBase struct {
	Appid     string `yaml:"appid"`
	AppSecret string `yaml:"app_secret"`
}
type JwtInfo struct {
	JwtSec string `yaml:"jwt_sec"`
	MaxAge uint   `yaml:"jwt_maxage"`
}

type Conf struct {
	AppBase `yaml:"app_base"`
	JwtInfo `yaml:"jwt_info"`
}

var AppConfig *Conf

func init() {
	AppConfig = new(Conf)
	confFile, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(confFile, AppConfig); err != nil {
		log.Fatal(err)
	}
}
