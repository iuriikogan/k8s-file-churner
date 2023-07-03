package utils

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Load the config
	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Verify the loaded config values
	verifyConfigValue(t, config.SizeOfFileMB, 2048, "APP_SIZE_OF_FILE_MB")
	verifyConfigValue(t, config.SizeOfPVCGB, 20, "APP_SIZE_OF_PVC_GB")
	verifyConfigValue(t, config.ChurnPercentage, 75, "APP_CHURN_PERCENTAGE")
	verifyConfigValue(t, config.ChurnIntervalMinutes, 30, "APP_CHURN_INTERVAL_MINUTES")
}

// Helper function to verify a config value
func verifyConfigValue(t *testing.T, actualValue interface{}, expectedValue interface{}, configKey string) {
	if actualValue != expectedValue {
		t.Errorf("Config value mismatch for %s. Expected: %v, Actual: %v", configKey, expectedValue, actualValue)
	}
}
