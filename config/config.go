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

type Config struct {
	SizeOfFileMB         int           `mapstructure:"APP_SIZE_OF_FILES_MB"`
	SizeOfPVCGB          int           `mapstructure:"APP_SIZE_OF_PVC_GB"`
	ChurnPercentage      float64       `mapstructure:"APP_CHURN_PERCENTAGE"`
	ChurnIntervalMinutes time.Duration `mapstructure:"APP_CHURN_INTERVAL_MINUTES"`
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	// Read configuration from environment variables
	v.AutomaticEnv()
	v.BindEnv("APP_SIZE_OF_FILES_MB")
	v.BindEnv("APP_SIZE_OF_PVC_GB")
	v.BindEnv("APP_CHURN_PERCENTAGE")
	v.BindEnv("APP_CHURN_INTERVAL_MINUTES")

	// Read default configuration from embedded file
	err := v.ReadConfig(bytes.NewBuffer(defaultConfiguration))
	if err != nil {
		return nil, err
	}

	// Unmarshal the configuration
	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
