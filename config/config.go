package config

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/spf13/viper"
)

// defaultConfiguration is from config.yaml, embedded in the binary

//go:embed config.yaml
var defaultConfiguration []byte

type App struct {
	SizeOfFileMB         int           `mapstructure:"SIZE_OF_FILES_MB"`
	SizeOfPVCGB          int           `mapstructure:"SIZE_OF_PVC_GB"`
	ChurnPercentage      float64       `mapstructure:"CHURN_PERCENTAGE"`
	ChurnIntervalMinutes time.Duration `mapstructure:"CHURN_PERCENTAGE"`
}

type Config struct {
	App *App `mapstructure:"APP"`
}

func LoadConfig() (*Config, error) {

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
