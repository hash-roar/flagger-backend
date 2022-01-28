package appconfig

import (
	"fmt"
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

type DB struct {
	User   string `yaml:"user"`
	Host   string `yaml:"host"`
	Dbname string `yaml:"dbname"`
	Sec    string `yaml:"sec"`
}

type Conf struct {
	AppBase `yaml:"app_base"`
	JwtInfo `yaml:"jwt_info"`
	Dsn     string `yaml:"dsn"`
	DB      `yaml:"db"`
}

var AppConfig *Conf

func init() {
	AppConfig = new(Conf)
	confFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(confFile, AppConfig); err != nil {
		log.Fatal(err)
	}
	AppConfig.Dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		AppConfig.User, AppConfig.Sec, AppConfig.Host, AppConfig.Dbname)
}
