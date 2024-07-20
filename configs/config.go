package configs

import (
	"github.com/spf13/viper"
)

type Conf struct {
	HostName   string `mapstructure:"HOST_NAME_CONF"`
	WsEndPoint string `mapstructure:"WS_ENDPOINT_CONF"`
	Port       string `mapstructure:"PORT_CONF"`
	HostRedis  string `mapstructure:"HOST_REDIS"`
	PortRedis  string `mapstructure:"PORT_REDIS"`
	PassRedis  string `mapstructure:"PASSWORD_REDIS"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	viper.SetConfigName("go_mongo")
	viper.SetConfigType("env")
	viper.SetConfigFile(path + ".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return cfg, err
}
