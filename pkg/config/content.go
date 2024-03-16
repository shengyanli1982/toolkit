package config

import (
	"bytes"
	"io"
	"strings"

	"github.com/spf13/viper"
)

// Content 结构体包含了配置对象和 viper 对象
// Content struct contains the configuration object and the viper object
type Content struct {
	// config 是配置对象
	// config is the configuration object
	config *Config

	// viper 是 viper 对象，用于处理配置文件
	// viper is the viper object, used for handling configuration files
	viper *viper.Viper
}

// NewContent 创建一个新的 Content 实例
// NewContent creates a new Content instance
func NewContent(config *Config) *Content {
	// 验证配置是否有效
	// validate the config
	config = isConfigValid(config)

	// 创建一个新的 viper 实例
	// create a new viper instance
	viper := viper.New()

	// 设置配置文件的类型
	// set the type of the config file
	viper.SetConfigType(config.fileType)

	// 设置配置文件的名称
	// set the name of the config file
	viper.SetConfigFile(config.fileName)

	// 设置配置文件的搜索路径
	// set the search paths of the config file
	for _, path := range config.paths {
		viper.AddConfigPath(path)
	}

	// 返回一个新的 Content 实例
	// return a new Content instance
	return &Content{
		config: config,
		viper:  viper,
	}
}

// GetViper 返回 viper 对象
// GetViper returns the viper object
func (c *Content) GetViper() *viper.Viper {
	return c.viper
}

// LoadFromFile 从文件中加载配置数据
// LoadFromFile loads configuration data from a file
func (c *Content) LoadFromFile(data any, opts ...viper.DecoderConfigOption) error {
	// 读取配置文件
	// read config file
	if err := c.viper.ReadInConfig(); err != nil {
		return err
	}

	// 反序列化配置文件数据
	// unmarshal config file data
	if err := c.viper.Unmarshal(data, opts...); err != nil {
		return err
	}

	// 成功
	// success
	return nil
}

// SaveToFile 将配置保存到文件
// SaveToFile saves the configuration to a file
func (c *Content) SaveToFile() error {
	return c.viper.WriteConfigAs(c.config.fileName)
}

// SaveToFileWithName 使用给定的名称将配置保存到文件
// SaveToFileWithName saves the configuration to a file with the given name
func (c *Content) SaveToFileWithName(fileName string) error {
	return c.viper.WriteConfigAs(strings.TrimSpace(fileName))
}

// StreamContent 结构体继承了 Content 结构体
// StreamContent struct inherits from Content struct
type StreamContent struct {
	// config 是配置对象
	// config is the configuration object
	config *Config

	// viper 是 viper 对象，用于处理配置文件
	// viper is the viper object, used for handling configuration files
	viper *viper.Viper
}

// NewStreamContent 创建一个新的 StreamContent 实例
// NewStreamContent creates a new StreamContent instance
func NewStreamContent(config *Config) *StreamContent {
	// 验证配置是否有效
	// validate the config
	config = isConfigValid(config)

	// 创建一个新的 viper 实例
	// create a new viper instance
	viper := viper.New()

	// 设置配置文件的类型
	// set the type of the config file
	viper.SetConfigType(config.fileType)

	// 返回一个新的 Content 实例
	// return a new Content instance
	return &StreamContent{
		config: config,
		viper:  viper,
	}
}

// GetViper 返回 viper 对象
// GetViper returns the viper object
func (c *StreamContent) GetViper() *viper.Viper {
	return c.viper
}

// LoadFromStream 从流中加载配置数据
// LoadFromStream loads configuration data from a stream
func (c *StreamContent) LoadFromStream(data any, opts ...viper.DecoderConfigOption) error {
	// 从 io.Reader 读取所有字节
	// read all bytes from io.Reader
	content, err := io.ReadAll(c.config.streamReader)
	if err != nil {
		return err
	}

	// 从 io.Reader 读取内容
	// read content from io.Reader
	if err := c.viper.ReadConfig(bytes.NewReader(content)); err != nil {
		return err
	}

	// 反序列化配置文件数据
	// unmarshal config file data
	if err := c.viper.Unmarshal(data, opts...); err != nil {
		return err
	}

	// 重置流读取器
	// reset stream reader
	c.config.streamReader = bytes.NewReader(content)

	// 成功
	// success
	return nil
}

// SaveToFile 将配置保存到文件
// SaveToFile saves the configuration to a file
func (c *StreamContent) SaveToFile() error {
	return c.viper.WriteConfigAs(c.config.fileName)
}

// SaveToFileWithName 使用给定的名称将配置保存到文件
// SaveToFileWithName saves the configuration to a file with the given name
func (c *StreamContent) SaveToFileWithName(fileName string) error {
	return c.viper.WriteConfigAs(strings.TrimSpace(fileName))
}
