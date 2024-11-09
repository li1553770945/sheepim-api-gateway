package config

import (
	"fmt"
	"github.com/li1553770945/sheepim-api-gateway/biz/constant"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type ServerConfig struct {
	ServiceName   string `yaml:"service-name"`
	ListenAddress string `yaml:"listen-address"`
}

type OpenTelemetryConfig struct {
	Endpoint string `yaml:"endpoint"`
}

type RpcConfig struct {
	UserServiceName    string `yaml:"user-service-name"`
	AuthServiceName    string `yaml:"auth-service-name"`
	ProjectServiceName string `yaml:"project-service-name"`
}

type EtcdConfig struct {
	Endpoint []string `yaml:"endpoint"`
}

type Config struct {
	Env                 string
	ServerConfig        ServerConfig        `yaml:"server"`
	OpenTelemetryConfig OpenTelemetryConfig `yaml:"open-telemetry"`
	RpcConfig           RpcConfig           `yaml:"rpc"`
	EtcdConfig          EtcdConfig          `yaml:"etcd"`
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
