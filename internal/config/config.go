// Package config loads application configuration.
package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Load reads the application configuration file and environment overrides.
func Load(cfgFile, env string) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./configs/")
		if env == "" {
			viper.SetConfigName("config")
		} else {
			viper.SetConfigName("config." + env)
		}
	}
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config failed!: %w", err)
	}
	fmt.Println("loaded config file:", viper.ConfigFileUsed())
	return nil
}
