package logger

import "github.com/sahandPgr/car-selling-service/config"

// Logger is an interface for logger servvices
type Logger interface {
	Init()
	Info(cat Category, subCat SubCategory, extra map[ExtraKey]interface{}, msg string)
	Debug(cat Category, subCat SubCategory, extra map[ExtraKey]interface{}, msg string)
	Warn(cat Category, subCat SubCategory, extra map[ExtraKey]interface{}, msg string)
	Error(cat Category, subCat SubCategory, extra map[ExtraKey]interface{}, msg string)
	Fatal(cat Category, subCat SubCategory, extra map[ExtraKey]interface{}, msg string)
}

// NewLogger is a factory method to create a new logger instance
func NewLogger(config *config.Config) Logger {
	return newZeroLogger(config)
}
