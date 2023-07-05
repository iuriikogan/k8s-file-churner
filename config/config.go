package config

import (
	"bytes"
	_ "embed"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// defaultConfiguration is the default configuration, embedded in the binary.

//go:embed config.yaml
var defaultConfiguration []byte

// Config
type App struct {
	SizeOfFileMB         int           // Size of the file in GB
	SizeOfPVCGB          int           // Number of files to create
	ChurnPercentage      float64       // Percentage of files to churn must be passed as float64 (0.5 = 50%)
	ChurnIntervalMinutes time.Duration // Churn interval in minutes
}

type Config struct {
	App *App
}

func LoadConfig() (*Config, error) {
	// Environment variables
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

	// Unmarshal the configuration
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
