package config

import (
	"bytes"
	"github.com/spf13/viper"
	"io"
)

type Config struct {
	Redis struct {
		Addr         string `toml:"addr"`
		Pass         string `toml:"pass"`
		Db           int    `toml:"db"`
		MaxRetries   int    `toml:"maxRetries"`
		PoolSize     int    `toml:"poolSize"`
		MinIdleConns int    `toml:"minIdleConns"`
	} `toml:"redis"`
}

var (
	devConfigs []byte
)

var config = new(Config)

func init() {
	var r io.Reader

	r = bytes.NewReader(devConfigs)

	viper.SetConfigType("toml")

	if err := viper.ReadConfig(r); err != nil {
		// error
	}

	viper.SetConfigName("dev_configs")
	viper.AddConfigPath("./config")
}

func Get() Config {
	return *config
}
