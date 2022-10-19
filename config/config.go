package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	Server struct {
		Host    string `yaml:"host"`
		Port    int    `yaml:"port"`
		Timeout struct {
			Server int `yaml:"server"`
			Read   int `yaml:"read"`
			Write  int `yaml:"write"`
			Idle   int `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`
	TrafficAPI struct {
		AuthKey string `yaml:"authKey" envconfig:"TRAFFIC_API_AUTH_KEY"`
		URL     string `yaml:"url"`
	} `yaml:"traffic_api"`
}

var Cfg Config

func Load() {
	loadCfgFile()
	loadEnv()
}

func loadCfgFile() {
	f, _ := ioutil.ReadFile("config.yml")
	must(yaml.Unmarshal(f, &Cfg))
}

func loadEnv() {
	must(envconfig.Process("", &Cfg))
}

func must(err error) {
	if err != nil {
		log.Fatal("Unable to load cfg ", err)
	}
}
