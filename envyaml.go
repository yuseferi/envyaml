package envyaml

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v3"
)

// LoadConfig unmarshals a YAML file into a struct, replacing placeholders with
// environment variables and performing validation.
func LoadConfig(configFilePath string, configStruct interface{}) error {
	// Read the YAML file
	yamlData, err := os.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to read YAML file: %v", err)
	}

	// Unmarshal YAML into the config struct
	err = yaml.Unmarshal(yamlData, configStruct)
	if err != nil {
		return fmt.Errorf("failed to unmarshal YAML: %v", err)
	}

	// Parse environment variables into the config struct
	if err := env.Parse(configStruct); err != nil {
		return fmt.Errorf("failed to parse environment variables: %v", err)
	}

	return nil
}
