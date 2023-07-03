package utils

import (
	"github.com/spf13/viper"
)

// T
type Config struct {
	sizeOfFileMB         int    `mapstructure:"SIZE_OF_FILE_MB"`        // Size of the file in GBi
	sizeOfPVCGB          int    `mapstructure:"SIZE_OF_PVC_GB"`         // Number of files to create
	storageClassName     string `mapstructure:"STORAGE_CLASS_NAME"`     // Storage class name
	churnPercentage      int    `mapstructure:"CHURN_PERCENTAGE"`       // Percentage of files to churn
	churnIntervalMinutes int    `mapstructure:"CHURN_INTERVAL_MINUTES"` // Interval in minutes to churn files
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
