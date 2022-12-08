package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type server struct {
	Addr    string `yaml:"addr"`
	Timeout struct {
		Server int `yaml:"server"`
		Read   int `yaml:"read"`
		Write  int `yaml:"write"`
		Idle   int `yaml:"idle"`
	} `yaml:"timeout"`
}

type database struct {
	Name string `yaml:"name"`
}

type trafficAPI struct {
	AuthKey string `yaml:"authKey" envconfig:"TRAFFIC_API_AUTH_KEY"`
	URL     string `yaml:"url"`
}

type mail struct {
	From     string `yaml:"from"`
	Password string `yaml:"password" envconfig:"MAIL_PASSWORD"`
	SmtpHost string `yaml:"smtpHost"`
	SmtpPort int    `yaml:"smtpPort"`
}

type Config struct {
	Server     server     `yaml:"server"`
	Database   database   `yaml:"database"`
	TrafficAPI trafficAPI `yaml:"traffic_api"`
	Mail       mail       `yaml:"mail"`
}

func Load(file *string) *Config {
	var cfg Config
	loadCfgFile(file, &cfg)
	loadEnv(&cfg)
	return &cfg
}

func loadCfgFile(file *string, c *Config) {
	f, _ := ioutil.ReadFile(*file)
	must(yaml.Unmarshal(f, c))
}

func loadEnv(c *Config) {
	must(envconfig.Process("", c))
}

func must(err error) {
	if err != nil {
		log.Fatal("Unable to load cfg ", err)
	}
}
