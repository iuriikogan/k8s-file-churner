package config

import (
<<<<<<< HEAD
=======
	"bytes"
	_ "embed" // embed is needed here to embed the config.yaml file into the binary

>>>>>>> parent of 2305d96 (something broken)
	"github.com/spf13/viper"
)

// defaultConfiguration is from config.yaml, embedded in the binary

//go:embed config.yaml
var defaultConfiguration []byte

// Config struct is exported to make is easier to work with the vars in main.go
type Config struct {
<<<<<<< HEAD
	SizeOfFileMB         int `mapstructure:"APP_SIZE_OF_FILES_MB"`
	SizeOfPVCGB          int `mapstructure:"APP_SIZE_OF_PVC_GB"`
	ChurnPercentage      int `mapstructure:"APP_CHURN_PERCENTAGE"`
	ChurnIntervalMinutes int `mapstructure:"APP_CHURN_INTERVAL_MINUTES"`
	ChurnDurationHours   int `mapstructure:"APP_CHURN_DURATION_HOURS"`
=======
	SizeOfFileMB         int     `mapstructure:"APP_SIZE_OF_FILES_MB"`
	SizeOfPVCGB          int     `mapstructure:"APP_SIZE_OF_PVC_GB"`
	ChurnPercentage      float64 `mapstructure:"APP_CHURN_PERCENTAGE"`
	ChurnIntervalMinutes int64   `mapstructure:"APP_CHURN_INTERVAL_MINUTES"`
>>>>>>> parent of 2305d96 (something broken)
}

// LoadConfig loads the configuration from the environment variables and sets default values
func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetDefault("APP_SIZE_OF_FILES_MB", 10)
	v.SetDefault("APP_SIZE_OF_PVC_GB", 1)
<<<<<<< HEAD
	v.SetDefault("APP_CHURN_PERCENTAGE", 70)
	v.SetDefault("APP_CHURN_INTERVAL_MINUTES", 5)
	v.SetDefault("APP_CHURN_DURATION_HOURS", 2)
=======
	v.SetDefault("APP_CHURN_PERCENTAGE", 0.2)
	v.SetDefault("APP_CHURN_INTERVAL_MINUTES", 10)
>>>>>>> parent of 2305d96 (something broken)

	// Read configuration from environment variables
	v.BindEnv("APP_SIZE_OF_FILES_MB")
	v.BindEnv("APP_SIZE_OF_PVC_GB")
	v.BindEnv("APP_CHURN_PERCENTAGE")
	v.BindEnv("APP_CHURN_INTERVAL_MINUTES")
<<<<<<< HEAD
	v.BindEnv("APP_CHURN_DURATION_HOURS")
	v.AutomaticEnv()

	// Unmarshal the configuration
	var config *Config
	err := v.Unmarshal(&config)
=======

	// Read default configuration from embedded file
	err := v.ReadConfig(bytes.NewBuffer(defaultConfiguration))
>>>>>>> parent of 2305d96 (something broken)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
=======
	// Unmarshal the configuration
	var config *Config
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

>>>>>>> parent of 2305d96 (something broken)
	return config, nil
}
