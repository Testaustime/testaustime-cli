package config

import (
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/romeq/testaustime-cli/utils"
)

type Config struct {
	file   string
	Token  string
	ApiUrl string
}

// New returns a new Config struct with given parameters.
// if apiUrl is empty string, it will be later replaced with current
// production server.
func New(file, token, apiUrl string) Config {
	return Config{
		file,
		token,
		apiUrl,
	}
}

// UpdateField will update field in current Config struct and
// in configuration file which was used to launch testaustime-cli
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

// GetConfiguration will read configuration from configFileOverride.
// if configFileOverride is empty string, GetConfiguration will resolve
// configuration path using standard enviroment variables.
func GetConfiguration(configFileOverride string) (config Config) {
	configFile := utils.StringOr(configFileOverride, resolveConfigPath())

	_, err := toml.DecodeFile(configFile, &config)
	utils.Check(err)

	return New(configFile, config.Token, config.ApiUrl)
}

// resolveConfigPath will resolve configuration path using standards
func resolveConfigPath() string {
	globalConfigDir := os.ExpandEnv("$HOME/.config")
	xdg_cfg_home := os.Getenv("XDG_CONFIG_HOME")
	if xdg_cfg_home != "" {
		globalConfigDir = xdg_cfg_home
	}
	return path.Join(globalConfigDir, "testaustime-cli/config.toml")
}
