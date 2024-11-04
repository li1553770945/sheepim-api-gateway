package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"sheepim-api-gateway/biz/constant"
)

type ServerConfig struct {
	ServiceName   string `yaml:"service-name"`
	ListenAddress string `yaml:"listen-address"`
}

type OpenTelemetryConfig struct {
	Endpoint string `yaml:"endpoint"`
}

type EtcdConfig struct {
	Endpoint string `yaml:"endpoint"`
}

type DatabaseConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Address  string `yaml:"address"`
	Port     int32  `yaml:"port"`
}

type RpcConfig struct {
	UserServiceName string `yaml:"user-service-name"`
	AuthServiceName string `yaml:"auth-service-name"`
}

type Config struct {
	Env                 string
	ServerConfig        ServerConfig        `yaml:"server"`
	OpenTelemetryConfig OpenTelemetryConfig `yaml:"open-telemetry"`
	DatabaseConfig      DatabaseConfig      `yaml:"database"`
	EtcdConfig          EtcdConfig          `yaml:"etcd"`
	RpcConfig           RpcConfig           `yaml:"rpc"`
}

func GetConfig(env string) *Config {
	if env != constant.EnvProduction && env != constant.EnvDevelopment {
		panic(fmt.Sprintf("环境必须是%s或者%s之一", constant.EnvProduction, constant.EnvDevelopment))
	}
	conf := &Config{}
	path := filepath.Join("conf", fmt.Sprintf("%s.yml", env))
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	err = yaml.NewDecoder(f).Decode(conf)
	conf.Env = env
	if err != nil {
		panic(err)
	}

	return conf
}
