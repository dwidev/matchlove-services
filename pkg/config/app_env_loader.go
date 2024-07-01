package config

import (
	"log"

	"github.com/spf13/viper"
)

func YamlLoader() *viperConfigLoader {
	return &viperConfigLoader{
		configPath: "./pkg/config/",
		configName: "config",
		configType: "yaml",
	}
}

type viperConfigLoader struct {
	configPath string
	configName string
	configType string
}

func (v *viperConfigLoader) Run() (*Schema, error) {
	viper.AddConfigPath(v.configPath)
	viper.SetConfigName(v.configName)
	viper.SetConfigType(v.configType)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file :%s", err)
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error parse to Config :%s", err)
		return nil, err
	}

	return config, nil
}
