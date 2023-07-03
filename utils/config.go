package utils

import (
	"time"

	"github.com/spf13/viper"
)

// Config
type Config struct {
	SizeOfFileMB         int           `mapstructure:"APP_SIZE_OF_FILE_MB"`        // Size of the file in GB
	SizeOfPVCGB          int           `mapstructure:"APP_SIZE_OF_PVC_GB"`         // Number of files to create
	ChurnPercentage      float64       `mapstructure:"APP_CHURN_PERCENTAGE"`       // Percentage of files to churn must be passed as float64 (0.5 = 50%)
	ChurnIntervalMinutes time.Duration `mapstructure:"APP_CHURN_INTERVAL_MINUTES"` // Churn interval in minutes
}

func LoadConfig() (config *Config, err error) {
	// Set default values
	viper.SetDefault("APP_SIZE_OF_FILE_MB", 999)
	viper.SetDefault("APP_SIZE_OF_PVC_GB", 30)
	viper.SetDefault("APP_CHURN_PERCENTAGE", 0.20)
	viper.SetDefault("APP_CHURN_INTERVAL_MINUTES", 3600)

	// read environment variables
	viper.SetEnvPrefix("APP")
	viper.BindEnv("APP_SIZE_OF_FILE_MB")
	viper.BindEnv("APP_SIZE_OF_PVC_GB")
	viper.BindEnv("APP_CHURN_PERCENTAGE")
	viper.BindEnv("APP_CHURN_INTERVAL_MINUTES")

	// Unmarshal the configuration into the Config struct
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
