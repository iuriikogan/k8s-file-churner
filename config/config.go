package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config struct is exported to make is easier to work with the vars in main.go
type Config struct {
	SizeOfFileMB         int           `mapstructure:"APP_SIZE_OF_FILES_MB"`
	SizeOfPVCGB          int           `mapstructure:"APP_SIZE_OF_PVC_GB"`
	ChurnPercentage      float64       `mapstructure:"APP_CHURN_PERCENTAGE"`
	ChurnIntervalMinutes time.Duration `mapstructure:"APP_CHURN_INTERVAL_MINUTES"` // time.Duration is in nanoseconds and is used to calculate the churn interval
	ChurnDurationHours   time.Duration `mapstructure:"APP_CHURN_DURATION_HOURS"`   // time.Duration is in nanoseconds and is used to calculate the churn duration
}

// LoadConfig loads the configuration from the environment variables and the embedded config.yaml file its used in main.go/19
func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetDefault("APP_SIZE_OF_FILES_MB", 10)
	v.SetDefault("APP_SIZE_OF_PVC_GB", 1)
	v.SetDefault("APP_CHURN_PERCENTAGE", 0.2)
	v.SetDefault("APP_CHURN_INTERVAL_MINUTES", "60m")
	v.SetDefault("APP_CHURN_DURATION_HOURS", "1h")

	// Read configuration from environment variables
	v.AutomaticEnv()
	v.BindEnv("APP_SIZE_OF_FILES_MB")
	v.BindEnv("APP_SIZE_OF_PVC_GB")
	v.BindEnv("APP_CHURN_PERCENTAGE")
	v.BindEnv("APP_CHURN_INTERVAL_MINUTES")
	v.BindEnv("APP_CHURN_DURATION_HOURS")

	// // Read default configuration from embedded file
	// err := v.ReadConfig()
	// if err != nil {
	// 	return nil, err
	// }

	// Unmarshal the configuration
	var config *Config
	err := v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	for key, value := range v.AllSettings() {
		println(key, value)
	}
	return config, nil
}
