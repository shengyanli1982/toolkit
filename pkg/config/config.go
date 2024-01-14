package config

import (
	"bytes"
	"io"
	"strings"
)

// DefaultConfigReader is the default reader for the config file
var DefaultConfigReader = bytes.NewReader(make([]byte, 0))

// Config is the configuration object
type Config struct {
	paths        []string  // config file search paths
	fileName     string    // config file name
	fileType     string    // config file format
	streamReader io.Reader // config file reader
}

// NewConfig returns a new config with default values
func NewConfig() *Config {
	return &Config{
		paths:        []string{DefaultConfigPath},
		fileName:     DefaultConfigName,
		fileType:     DefaultConfigType,
		streamReader: DefaultConfigReader,
	}
}

// SetSearchPaths sets the search paths for the config
func (c *Config) SetSearchPaths(paths []string) *Config {
	c.paths = append(c.paths, paths...)
	return c
}

// SetFileName sets the file name for the config
func (c *Config) SetFileName(fileName string) *Config {
	c.fileName = fileName
	return c
}

// SetFileFormat sets the file format for the config
func (c *Config) SetFileFormat(fileFormat string) *Config {
	if !isConfigTypeSupported(fileFormat) {
		return c
	}
	c.fileType = fileFormat
	return c
}

// SetReader sets the reader which contain the config data for the config
func (c *Config) SetReader(reader io.Reader) *Config {
	c.streamReader = reader
	return c
}

// DefaultConfig returns a new config with default values
func DefaultConfig() *Config {
	return NewConfig()
}

// isConfigTypeSupported checks if the config type is supported
func isConfigTypeSupported(fileFormat string) bool {
	fileFormat = strings.ToLower(strings.TrimSpace(fileFormat))
	switch fileFormat {
	case JSONType, YAMLType, TOMLType:
		return true
	}
	return false
}

// isConfigValid checks if the config is valid, if not, it will return a new config with default values
func isConfigValid(conf *Config) *Config {
	if conf != nil {
		if strings.TrimSpace(conf.fileName) == "" {
			conf.fileName = DefaultConfigName
		}
		if !isConfigTypeSupported(conf.fileType) {
			conf.fileType = DefaultConfigType
		}
		if len(conf.paths) == 0 {
			conf.paths = []string{DefaultConfigPath}
		}
	} else {
		conf = NewConfig()
	}

	return conf
}
