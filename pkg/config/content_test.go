package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContent_LoadFromFile_JSON(t *testing.T) {
	// Create a temporary config file for testing
	tmpFile, err := os.CreateTemp("", "config_test")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	// Write test data to the temporary config file
	testData := `
	{
		"key1": "value1",
		"key2": "value2"
	}
	`
	_, err = tmpFile.Write([]byte(testData))
	assert.NoError(t, err)

	// Create a new Content instance
	content := NewContent(&Config{
		fileName: tmpFile.Name(),
		fileType: JSONType,
		paths:    []string{filepath.Dir(tmpFile.Name())},
	})

	// Define the expected data structure for unmarshaling
	var data struct {
		Key1 string `json:"key1"`
		Key2 string `json:"key2"`
	}

	// Call the LoadFromJSONFile method
	err = content.LoadFromFile(&data)
	assert.NoError(t, err)

	// Verify the loaded data
	assert.Equal(t, "value1", data.Key1)
	assert.Equal(t, "value2", data.Key2)
}

func TestContent_LoadFromFile_YAML(t *testing.T) {
	// Create a temporary config file for testing
	tmpFile, err := os.CreateTemp("", "config_test")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	// Write test data to the temporary config file
	testData := `
key1: value1
key2: value2
`
	_, err = tmpFile.Write([]byte(testData))
	assert.NoError(t, err)

	// Create a new Content instance
	content := NewContent(&Config{
		fileName: tmpFile.Name(),
		fileType: YAMLType,
		paths:    []string{filepath.Dir(tmpFile.Name())},
	})

	// Define the expected data structure for unmarshaling
	var data struct {
		Key1 string `yaml:"key1"`
		Key2 string `yaml:"key2"`
	}

	// Call the LoadFromYAMLFile method
	err = content.LoadFromFile(&data)
	assert.NoError(t, err)

	// Verify the loaded data
	assert.Equal(t, "value1", data.Key1)
	assert.Equal(t, "value2", data.Key2)
}

func TestContent_LoadFromFile_TOML(t *testing.T) {
	// Create a temporary config file for testing
	tmpFile, err := os.CreateTemp("", "config_test")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	// Write test data to the temporary config file
	testData := `
	key1 = "value1"	
	key2 = "value2"
	`
	_, err = tmpFile.Write([]byte(testData))
	assert.NoError(t, err)

	// Create a new Content instance
	content := NewContent(&Config{
		fileName: tmpFile.Name(),
		fileType: TOMLType,
		paths:    []string{filepath.Dir(tmpFile.Name())},
	})

	// Define the expected data structure for unmarshaling
	var data struct {
		Key1 string `yaml:"key1"`
		Key2 string `yaml:"key2"`
	}

	// Call the LoadFromYAMLFile method
	err = content.LoadFromFile(&data)
	assert.NoError(t, err)

	// Verify the loaded data
	assert.Equal(t, "value1", data.Key1)
	assert.Equal(t, "value2", data.Key2)
}

func TestContent_SaveToFile(t *testing.T) {
	// Create a temporary config file for testing
	tmpFile, err := os.CreateTemp("", "config_test")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	// Create a new Content instance
	content := NewContent(&Config{
		fileName: tmpFile.Name(),
		fileType: JSONType,
		paths:    []string{filepath.Dir(tmpFile.Name())},
	})

	// Set some test data
	content.GetViper().Set("key1", "value1")
	content.GetViper().Set("key2", "value2")

	// Call the SaveToFile method
	err = content.SaveToFile()
	assert.NoError(t, err)

	// Read the saved config file
	savedData, err := os.ReadFile(tmpFile.Name())
	assert.NoError(t, err)

	// Verify the saved data
	expectedData := "{\n  \"key1\": \"value1\",\n  \"key2\": \"value2\"\n}"
	assert.Equal(t, expectedData, string(savedData))
}

func TestContent_SaveToFileWithName(t *testing.T) {
	// Create a temporary config file for testing
	tmpFile, err := os.CreateTemp("", "config_test")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	// Create a new Content instance
	content := NewContent(&Config{
		fileName: tmpFile.Name(),
		fileType: JSONType,
		paths:    []string{filepath.Dir(tmpFile.Name())},
	})

	// Set some test data
	content.GetViper().Set("key1", "value1")
	content.GetViper().Set("key2", "value2")

	// Call the SaveToFileWithName method
	err = content.SaveToFileWithName("test_config.json")
	assert.NoError(t, err)

	// Read the saved config file
	savedData, err := os.ReadFile("test_config.json")
	assert.NoError(t, err)

	// Verify the saved data
	expectedData := "{\n  \"key1\": \"value1\",\n  \"key2\": \"value2\"\n}"
	assert.Equal(t, expectedData, string(savedData))
}
