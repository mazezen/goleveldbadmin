package sdk

import (
	"github.com/spf13/viper"
	"log"
)

var Cf *Config

func init() {
	Cf = &Config{}
}

type Config struct {
	Source Source `yaml:"source"`
	Ti     Ti     `yaml:"ti"`
	Jwt    Jwt    `yaml:"jwt"`
}
type Source struct {
	ServiceName string `yaml:"service_name"`
	Dir         string `yaml:"dir"`
}

type Ti struct {
	Acc string `yaml:"acc"`
	Pwd string `yaml:"pwd"`
}

type Jwt struct {
	ExpireTime int    `yaml:"expire_time"`
	Secret     string `yaml:"secret"`
}

func ParseConfig(filepath string) {
	var v = viper.New()
	v.SetConfigFile(filepath)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := v.Unmarshal(Cf); err != nil {
		log.Fatalf("Error parsing config file, %s", err)
	}
}
