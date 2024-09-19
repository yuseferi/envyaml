package envyaml

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Set up test environment variables
	os.Setenv("TEST_PASSWORD", "secret_password")

	// Create a temporary YAML file for testing
	yamlContent := `
host: localhost
port: 3606
password: ${TEST_PASSWORD}
`
	tempFile, err := os.CreateTemp("", "config*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(yamlContent); err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	// Define the config struct
	type TestConfig struct {
		Host     string `yaml:"host" env:"TEST_HOST"`
		Port     int    `yaml:"port" env:"TEST_PORT"`
		Password string `yaml:"password" env:"TEST_PASSWORD"`
	}

	// Load the configuration
	var cfg TestConfig
	err = LoadConfig(tempFile.Name(), &cfg)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "localhost", cfg.Host)
	assert.Equal(t, 3606, cfg.Port)
	assert.Equal(t, "secret_password", cfg.Password)
}

func TestLoadConfig_MissingRequiredEnv(t *testing.T) {

	// Create a temporary YAML file for testing
	yamlContent := `
host: localhost
port: 3606
password: ${TEST_PASSWORD2}
`
	tempFile, err := os.CreateTemp("", "config*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(yamlContent); err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	// Define the config struct
	type TestConfig struct {
		Host     string `yaml:"host" env:"TEST_HOST"`
		Port     int    `yaml:"port" env:"TEST_PORT"`
		Password string `yaml:"password" env:"TEST_PASSWORD2,required"`
	}

	// Load the configuration
	var cfg TestConfig
	err = LoadConfig(tempFile.Name(), &cfg)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse environment variables")
	assert.Contains(t, err.Error(), "required environment variable \"TEST_PASSWORD2\" is not set")
}
func TestLoadConfig_ErrorLoadingFile(t *testing.T) {

	// Define the config struct
	type TestConfig struct {
		Host     string `yaml:"host" env:"TEST_HOST"`
		Port     int    `yaml:"port" env:"TEST_PORT"`
		Password string `yaml:"password" env:"TEST_PASSWORD2,required"`
	}

	// Load the configuration
	var cfg TestConfig
	err := LoadConfig("not_existfile.yaml", &cfg)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read YAML file: open not_existfile.yaml:")
}
func TestLoadConfig_ErrorOnUnmarshalling(t *testing.T) {

	// Create a temporary YAML file for testing
	yamlContent := `
 host: localhost
port:3606
password: ${TEST_PASSWORD2}
`
	tempFile, err := os.CreateTemp("", "config*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(yamlContent); err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}
	// Define the config struct
	type TestConfig struct {
		Host     string `yaml:"host" env:"TEST_HOST"`
		Port     int    `yaml:"port" env:"TEST_PORT"`
		Password string `yaml:"password" env:"TEST_PASSWORD2,required"`
	}

	// Load the configuration
	var cfg TestConfig
	err = LoadConfig(tempFile.Name(), &cfg)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal YAML")
}
