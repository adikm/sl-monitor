package config

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"log"
	"os"
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

type Database struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host" envconfig:"DB_HOST"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type cache struct {
	Host string `yaml:"host" envconfig:"CACHE_HOST"`
	Port int    `yaml:"port"`
}

type trafficAPI struct {
	AuthKey string `yaml:"authKey" envconfig:"TRAFFIC_API_AUTH_KEY"`
	URL     string `yaml:"url"`
}

type mail struct {
	From     string `yaml:"from" envconfig:"MAIL_USERNAME"`
	Password string `yaml:"password" envconfig:"MAIL_PASSWORD"`
	SmtpHost string `yaml:"smtpHost"`
	SmtpPort int    `yaml:"smtpPort"`
}

type Config struct {
	Server     server     `yaml:"server"`
	Database   Database   `yaml:"database"`
	Cache      cache      `yaml:"cache"`
	TrafficAPI trafficAPI `yaml:"traffic_api"`
	Mail       mail       `yaml:"mail"`
}

func Load(file *string) *Config {
	var cfg Config
	loadCfgFile(file, &cfg)
	loadEnv(&cfg)
	prettyPrint(&cfg)
	return &cfg
}

func loadCfgFile(file *string, c *Config) {
	f, _ := os.ReadFile(*file)
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

func prettyPrint(data *Config) {
	log.Println("Config loaded: ")
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
