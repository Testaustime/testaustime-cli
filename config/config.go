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
	file   string
	Token  string
	ApiUrl string
}

func New(file, token, apiUrl string) Config {
	return Config{
		file,
		token,
		apiUrl,
	}
}

func (c *Config) UpdateField(field *string, newValue string) {
	*field = newValue

	fhandle, err := os.OpenFile(c.file, os.O_WRONLY|os.O_CREATE, 0644)
	utils.Check(err)
	defer fhandle.Close()

	utils.Check(toml.NewEncoder(fhandle).Encode(map[string]string{
		"token":  c.Token,
		"apiurl": c.ApiUrl,
	}))
}

func GetConfiguration(alternateConfigFile string) (config Config) {
	configFile := utils.StringOr(alternateConfigFile, resolveConfigPath())

	m, err := toml.DecodeFile(configFile, &config)
	utils.Check(err)
	checkConfiguration(&m)

	return New(configFile, config.Token, config.ApiUrl)
}

func checkConfiguration(m *toml.MetaData) {
	if !m.IsDefined("token") {
		logger.Error(errors.New("Authentication token was not found in the configuration file."))
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
