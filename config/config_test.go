package config

import (
	_ "embed"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Save the current environment variables to restore later
	backupEnv := os.Environ()

	// Test cases with different environment variables and expected results
	testCases := []struct {
		name           string
		envVars        map[string]string
		expectedConfig *Config
		expectedError  error
	}{
		{
			name: "Custom Configuration",
			envVars: map[string]string{
				"APP_SIZE_OF_FILES_MB":       "200",
				"APP_SIZE_OF_PVC_GB":         "10",
				"APP_CHURN_PERCENTAGE":       "0.5",
				"APP_CHURN_INTERVAL_MINUTES": "60",
			},
			expectedConfig: &Config{
				SizeOfFileMB:         200,
				SizeOfPVCGB:          10,
				ChurnPercentage:      0.5,
				ChurnIntervalMinutes: 60,
			},
			expectedError: nil,
		},
		{
			name: "Custom Configuration 2",
			envVars: map[string]string{
				"APP_SIZE_OF_FILES_MB":       "500",
				"APP_SIZE_OF_PVC_GB":         "5",
				"APP_CHURN_PERCENTAGE":       "20",
				"APP_CHURN_INTERVAL_MINUTES": "5",
				"APP_CHURN_DURATION_HOURS":   "2",
			},
			expectedConfig: &Config{
				SizeOfFileMB:         500,
				SizeOfPVCGB:          5,
				ChurnPercentage:      0.2,
				ChurnIntervalMinutes: 40,
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set up environment variables for the test case
			os.Clearenv()
			for key, value := range tc.envVars {
				os.Setenv(key, value)
			}

			// Perform the test
			config, err := LoadConfig()

			// Check the result
			assert.Equal(t, tc.expectedConfig, config)
			assert.Equal(t, tc.expectedError, err)
		})
	}

	// Restore the original environment variables
	os.Clearenv()
	for _, envVar := range backupEnv {
		pair := strings.SplitN(envVar, "=", 2)
		os.Setenv(pair[0], pair[1])
	}
}
