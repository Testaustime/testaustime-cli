package config

import (
	"github.com/BurntSushi/toml"
	"github.com/romeq/testaustime-cli/utils"
)

type Config struct {
	file   string
	ApiUrl string
    CaseInsensitiveFields []string
}

// New returns a new Config struct with given parameters.
// if apiUrl is empty string, it will be later replaced with current
// production server.
func New(file, apiUrl string, CaseInsensitiveFields []string) Config {
	return Config{
		file,
		apiUrl,
        CaseInsensitiveFields,
	}
}

// GetConfiguration will read configuration from configFileOverride.
// if configFileOverride is empty string, GetConfiguration will resolve
// configuration path using standard enviroment variables.
func GetConfiguration(configFileOverride string) (config Config) {
	configFile := utils.StringOr(configFileOverride, utils.ResolveWantedPath(
        utils.EnvOrString("$XDG_CONFIG_HOME", "$HOME/.config/"),
        "config.toml",
    ))
	_, err := toml.DecodeFile(configFile, &config)
	utils.Check(err)

	return New(configFile, config.ApiUrl, config.CaseInsensitiveFields)
}

