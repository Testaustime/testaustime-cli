package config

import (
	"errors"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/romeq/testaustime-cli/logger"
	"github.com/romeq/testaustime-cli/utils"
)

type Config struct {
	Token  string
	ApiUrl string
}

func New(token string, apiUrl string) Config {
	return Config{
		token,
		apiUrl,
	}
}

func GetConfiguration(alternateConfigFile string) (config Config) {
	configFile := utils.StringOr(alternateConfigFile, resolveConfigPath())

	m, err := toml.DecodeFile(configFile, &config)
	utils.Check(err)
	checkConfiguration(&m)

	return New(config.Token, config.ApiUrl)
}

func checkConfiguration(m *toml.MetaData) {
	if !m.IsDefined("token") {
		logger.Error(errors.New("Authentication token was not found from the configuration file."))
	}
}

func resolveConfigPath() string {
	globalConfigDir := os.ExpandEnv("$HOME/.config")
	xdg_cfg_home := os.Getenv("XDG_CONFIG_HOME")
	if xdg_cfg_home != "" {
		globalConfigDir = xdg_cfg_home
	}
	return path.Join(globalConfigDir, "testaustime-cli/config.toml")
}
