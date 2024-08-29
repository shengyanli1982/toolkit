package server

import (
	"log"
)

// Logger 是一个接口，定义了一个 Errorf 方法，该方法接受一个格式化字符串和可变参数
// Logger is an interface that defines an Errorf method, which accepts a format string and variadic arguments
type Logger interface {
	Errorf(format string, args ...interface{})
}

// defaultLogger 是 Logger 接口的一个实现
// defaultLogger is an implementation of the Logger interface
type defaultLogger struct{}

// Errorf 是 defaultLogger 的方法，它使用 log.Printf 打印格式化的错误信息
// Errorf is a method of defaultLogger, it uses log.Printf to print formatted error messages
func (l *defaultLogger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}
