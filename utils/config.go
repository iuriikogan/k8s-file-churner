package utils

import (
	"github.com/spf13/viper"
)

// Config
type Config struct {
	sizeOfFileMB        int    `mapstructure:"SIZE_OF_FILE_MB"`        // Size of the file in GB
	sizeOfPVCGB         int    `mapstructure:"SIZE_OF_PVC_GB"`         // Number of files to create
	storageClassName    string `mapstructure:"STORAGE_CLASS_NAME"`     // Storage class name
	churnPercentage     int    `mapstructure:"CHURN_PERCENTAGE"`       // Percentage of files to churn
	hurnIntervalMinutes int    `mapstructure:"CHURN_INTERVAL_MINUTES"` // Interval in minutes to churn files
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("SIZE_OF_FILE_MB", 1024)
	viper.SetDefault("SIZE_OF_PVC_GB", 10)
	viper.SetDefault("STORAGE_CLASS_NAME", "default")
	viper.SetDefault("CHURN_PERCENTAGE", 50)
	viper.SetDefault("CHURN_INTERVAL_MINUTES", 60)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
