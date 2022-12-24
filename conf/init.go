package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

var C *Config

type Config struct {
	DB struct {
		UserName string
		Password string
		Host     string
		Schema   string
	}
}

func init() {
	C = &Config{}
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		logrus.Fatal("[init] read DB config error %v", err)
	}
	err = json.Unmarshal(data, C)
	if err != nil {
		logrus.Fatal("[init] init json parse %v", err)
	}
}
