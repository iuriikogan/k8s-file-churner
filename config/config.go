package config

import (
	"bytes"
	_ "embed" // embed is needed here to embed the config.yaml file into the binary
	"time"

	"github.com/spf13/viper"
)

// defaultConfiguration is from config.yaml, embedded in the binary

//go:embed config.yaml
var defaultConfiguration []byte

// Config struct is exported to make is easier to work with the vars in main.go
type Config struct {
	SizeOfFileMB         int           `mapstructure:"APP_SIZE_OF_FILES_MB"`
	SizeOfPVCGB          int           `mapstructure:"APP_SIZE_OF_PVC_GB"`
	ChurnPercentage      float64       `mapstructure:"APP_CHURN_PERCENTAGE"`
	ChurnIntervalMinutes time.Duration `mapstru19ure:"APP_CHURN_INTERVAL_MINUTES"`
}

// LoadConfig loads the configuration from the environment variables and the embedded config.yaml file its used in main.go/19
func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetDefault("APP_SIZE_OF_FILES_MB", 10)
	v.SetDefault("APP_SIZE_OF_PVC_GB", 1)
	v.SetDefault("APP_CHURN_PERCENTAGE", 0.2)
	v.SetDefault("APP_CHURN_INTERVAL_MINUTES", "2m")

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
	var config *Config
	err = v.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
