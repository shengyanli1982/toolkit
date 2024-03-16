package config

import (
	"bytes"
	"io"
	"strings"
)

// DefaultConfigReader 是配置文件的默认读取器
// DefaultConfigReader is the default reader for the config file
var DefaultConfigReader = bytes.NewReader(make([]byte, 0))

// Config 是配置对象，包含了配置文件的搜索路径、文件名、文件格式和读取器
// Config is the configuration object, which includes the search path, file name, file format, and reader of the configuration file
type Config struct {
	// paths 是配置文件的搜索路径
	// paths is the search path of the configuration file
	paths []string

	// fileName 是配置文件的文件名
	// fileName is the file name of the configuration file
	fileName string

	// fileType 是配置文件的文件格式
	// fileType is the file format of the configuration file
	fileType string

	// streamReader 是配置文件的读取器
	// streamReader is the reader of the configuration file
	streamReader io.Reader
}

// NewConfig 返回一个带有默认值的新配置，包括默认的搜索路径、文件名、文件格式和读取器
// NewConfig returns a new config with default values, including the default search path, file name, file format, and reader
func NewConfig() *Config {
	return &Config{
		// DefaultConfigPath 是默认的配置文件搜索路径
		// DefaultConfigPath is the default search path of the configuration file
		paths: []string{DefaultConfigPath},

		// DefaultConfigName 是默认的配置文件名
		// DefaultConfigName is the default file name of the configuration file
		fileName: DefaultConfigName,

		// DefaultConfigType 是默认的配置文件格式
		// DefaultConfigType is the default file format of the configuration file
		fileType: DefaultConfigType,

		// DefaultConfigReader 是默认的配置文件读取器
		// DefaultConfigReader is the default reader of the configuration file
		streamReader: DefaultConfigReader,
	}
}

// SetSearchPaths 设置配置的搜索路径
// SetSearchPaths sets the search paths for the config
func (c *Config) SetSearchPaths(paths []string) *Config {
	// 将新的搜索路径添加到现有的搜索路径中
	// Add the new search paths to the existing search paths
	c.paths = append(c.paths, paths...)
	return c
}

// SetFileName 设置配置的文件名
// SetFileName sets the file name for the config
func (c *Config) SetFileName(fileName string) *Config {
	// 设置配置的文件名
	// Set the file name of the config
	c.fileName = fileName
	return c
}

// SetFileFormat 设置配置的文件格式
// SetFileFormat sets the file format for the config
func (c *Config) SetFileFormat(fileFormat string) *Config {
	// 如果文件格式不被支持，直接返回
	// If the file format is not supported, return directly
	if !isConfigTypeSupported(fileFormat) {
		return c
	}
	// 设置配置的文件格式
	// Set the file format of the config
	c.fileType = fileFormat
	return c
}

// SetReader 设置包含配置数据的读取器
// SetReader sets the reader which contain the config data for the config
func (c *Config) SetReader(reader io.Reader) *Config {
	// 设置配置的读取器
	// Set the reader of the config
	c.streamReader = reader
	return c
}

// DefaultConfig 返回一个带有默认值的新配置
// DefaultConfig returns a new config with default values
func DefaultConfig() *Config {
	// 返回一个新的配置
	// Return a new config
	return NewConfig()
}

// isConfigTypeSupported 检查配置类型是否被支持
// isConfigTypeSupported checks if the config type is supported
func isConfigTypeSupported(fileFormat string) bool {
	// 将文件格式转换为小写并去除两端的空格
	// Convert the file format to lowercase and trim the spaces at both ends
	fileFormat = strings.ToLower(strings.TrimSpace(fileFormat))

	// 如果文件格式是 JSON、YAML 或 TOML，返回 true
	// If the file format is JSON, YAML, or TOML, return true
	switch fileFormat {
	case JSONType, YAMLType, TOMLType:
		return true
	}

	// 否则，返回 false
	// Otherwise, return false
	return false
}

// isConfigValid 检查配置是否有效，如果不是，它将返回一个带有默认值的新配置
// isConfigValid checks if the config is valid, if not, it will return a new config with default values
func isConfigValid(conf *Config) *Config {
	// 如果配置不为空
	// If the config is not null
	if conf != nil {
		// 如果配置的文件名为空，设置为默认文件名
		// If the file name of the config is empty, set it to the default file name
		if strings.TrimSpace(conf.fileName) == "" {
			conf.fileName = DefaultConfigName
		}

		// 如果配置的文件类型不被支持，设置为默认文件类型
		// If the file type of the config is not supported, set it to the default file type
		if !isConfigTypeSupported(conf.fileType) {
			conf.fileType = DefaultConfigType
		}

		// 如果配置的搜索路径为空，设置为默认搜索路径
		// If the search paths of the config are empty, set them to the default search paths
		if len(conf.paths) == 0 {
			conf.paths = []string{DefaultConfigPath}
		}
	} else {
		// 如果配置为空，返回一个新的带有默认值的配置
		// If the config is null, return a new config with default values
		conf = NewConfig()
	}

	// 返回配置
	// Return the config
	return conf
}
