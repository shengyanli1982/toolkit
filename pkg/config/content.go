package config

import (
	"io"
	"strings"

	"github.com/spf13/viper"
)

type Content struct {
	config *Config      // Config is the configuration object
	viper  *viper.Viper // Viper is the viper object
}

// NewContent creates a new Content instance
func NewContent(config *Config) *Content {
	// validate config
	config = isConfigValid(config)
	// create viper instance
	viper := viper.New()
	// set config file type
	viper.SetConfigType(config.fileType)
	// set config file name
	viper.SetConfigFile(config.fileName)
	// set config file search paths
	for _, path := range config.paths {
		viper.AddConfigPath(path)
	}
	// return Content instance
	return &Content{
		config: config,
		viper:  viper,
	}
}

// GetViper returns the viper object
func (c *Content) GetViper() *viper.Viper {
	return c.viper
}

// LoadFromFile loads configuration data from a file
func (c *Content) LoadFromFile(data any, opts ...viper.DecoderConfigOption) error {
	// read config file
	if err := c.viper.ReadInConfig(); err != nil {
		return err
	}
	// unmarshal config file data
	if err := c.viper.Unmarshal(data, opts...); err != nil {
		return err
	}
	return nil
}

// SaveToFile saves the configuration to a file
func (c *Content) SaveToFile() error {
	return c.viper.WriteConfigAs(c.config.fileName)
}

// SaveToFileWithName saves the configuration to a file with the given name
func (c *Content) SaveToFileWithName(fileName string) error {
	return c.viper.WriteConfigAs(strings.TrimSpace(fileName))
}

type StreamContent struct {
	Content
}

// NewStreamContent creates a new StreamContent instance
func NewStreamContent(config *Config) *StreamContent {
	config = isConfigValid(config)
	return &StreamContent{
		Content: *NewContent(config),
	}
}

// GetViper returns the viper object
func (c *StreamContent) GetViper() *viper.Viper {
	return c.viper
}

// LoadFromStream loads configuration data from a stream
func (c *StreamContent) LoadFromStream(data any, reader io.Reader, opts ...viper.DecoderConfigOption) error {
	// set config file type
	c.viper.SetConfigType(c.config.fileType)
	// set reader
	if reader == nil {
		reader = c.config.streamReader
	}
	// read content from io.Reader
	if err := c.viper.ReadConfig(reader); err != nil {
		return err
	}
	// unmarshal config file data
	if err := c.viper.Unmarshal(data, opts...); err != nil {
		return err
	}
	return nil
}

// SaveToFile saves the configuration to a file
func (c *StreamContent) SaveToFile() error {
	return c.viper.WriteConfigAs(c.config.fileName)
}

// SaveToFileWithName saves the configuration to a file with the given name
func (c *StreamContent) SaveToFileWithName(fileName string) error {
	return c.viper.WriteConfigAs(strings.TrimSpace(fileName))
}
