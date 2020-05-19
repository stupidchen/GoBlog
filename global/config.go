package global

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Sys SystemConfig
	Db DatabaseConfig
}

type SystemConfig struct {
	LogFile string
}

type DatabaseConfig struct {
	Username string
	Password string
	Host string
	Port string
}

func getConfigPath() string {
	path := os.Getenv(ConfigPathEnvKey)
	if path == "" {
		return DefaultConfigFile
	}
	return ""
}


func initConfigByBytes(configBytes []byte) *Config {
	var config Config

	err := json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil
	}
	return &config
}


func initConfigByEnvVar() *Config {
	envConfigStr := os.Getenv(ConfigEnvKey)
	if envConfigStr == "" {
		return nil
	}

	return initConfigByBytes([]byte(envConfigStr))
}

func initConfigByFile() *Config {
	path := getConfigPath()

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	return initConfigByBytes(data)
}

func initConfig() *Config {
	var config *Config

	config = initConfigByEnvVar()
	if config != nil {
		return config
	}

	config = initConfigByFile()
	if config != nil {
		return config
	}

	return nil
}