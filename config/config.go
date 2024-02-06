package config

import (
	"github.com/spf13/viper"
)

// Config struct is exported to make is easier to work with the vars in main.go
type Config struct {
	SizeOfFileMB         int `mapstructure:"APP_SIZE_OF_FILES_MB"`
	SizeOfPVCGB          int `mapstructure:"APP_SIZE_OF_PVC_GB"`
	ChurnPercentage      int `mapstructure:"APP_CHURN_PERCENTAGE"`
	ChurnIntervalMinutes int `mapstructure:"APP_CHURN_INTERVAL_MINUTES"`
	ChurnDurationHours   int `mapstructure:"APP_CHURN_DURATION_HOURS"`
}

// LoadConfig loads the configuration from the environment variables and sets default values
func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetDefault("APP_SIZE_OF_FILES_MB", 10)
	v.SetDefault("APP_SIZE_OF_PVC_GB", 1)
	v.SetDefault("APP_CHURN_PERCENTAGE", 70)
	v.SetDefault("APP_CHURN_INTERVAL_MINUTES", 5)
	v.SetDefault("APP_CHURN_DURATION_HOURS", 2)

	// Read configuration from environment variables
	v.BindEnv("APP_SIZE_OF_FILES_MB")
	v.BindEnv("APP_SIZE_OF_PVC_GB")
	v.BindEnv("APP_CHURN_PERCENTAGE")
	v.BindEnv("APP_CHURN_INTERVAL_MINUTES")
	v.BindEnv("APP_CHURN_DURATION_HOURS")
	v.AutomaticEnv()

	// Unmarshal the configuration
	var config *Config
	err := v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
