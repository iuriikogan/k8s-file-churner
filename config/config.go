package config

import (
	"bytes"
	_ "embed"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// defaultConfiguration is from app-cm.yaml, embedded in the binary
//
//go:embed app-cm.yaml
var defaultConfiguration []byte

type App struct {
	SizeOfFileMB         int           `mapstructure:"APP_SIZE_OF_FILES_MB"`
	SizeOfPVCGB          int           `mapstructure:"APP_SIZE_OF_PVC_GB"`
	ChurnPercentage      float64       `mapstructure:"APP_CHURN_PERCENTAGE"`
	ChurnIntervalMinutes time.Duration `mapstructure:"APP_CHURN_PERCENTAGE"`
}

type Config struct {
	App *App
}

func LoadConfig() (*Config, error) {
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// Configuration file
	viper.SetConfigType("yaml")
	viper.SetConfigName("app-cm")
	viper.AddConfigPath("/etc/config/")
	// Read configuration
	if err := viper.ReadConfig(bytes.NewBuffer(defaultConfiguration)); err != nil {
		return nil, err
	}
	// merge with the external config file if it exists
	viper.MergeInConfig()

	// Unmarshal the configuration
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
