package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/sahandPgr/car-selling-service/config"
)

var once sync.Once
var zeroLogSingle *zerolog.Logger

// zeroLogger is a implementation of zero logger
type zeroLogger struct {
	config *config.Config
	logger *zerolog.Logger
}

// zeroLogLevel is a map of log levels
var zeroLogLevel = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

// NewZeroLogger is a constructor for zero logger
func newZeroLogger(config *config.Config) *zeroLogger {
	logger := &zeroLogger{config: config}
	logger.Init()
	return logger
}

// Init is a method to initialize zero logger
func (log *zeroLogger) Init() {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		file, err := os.OpenFile(log.config.Log.Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
			logger.Fatal().Err(err).Msg("Failed to open log file")
			return
		}

		logger := zerolog.New(file).
			With().
			Timestamp().
			Str(string(AppName), log.config.Log.Logger).
			Logger()

		zerolog.SetGlobalLevel(zeroLogLevel[log.config.Log.Level])
		zeroLogSingle = &logger
	})
	log.logger = zeroLogSingle
}

// Info is a method to log info level
func (log *zeroLogger) Info(cat Category, subCat SubCategory, extra map[ExtraKey]interface{}, msg string) {
	log.logger.Info().
		Str("Category", string(cat)).
		Str("SubCategory", string(subCat)).
		Fields(logPramsToZerolog(extra)).
		Msg(msg)
}

// Debug is a method to log debug level
func (log *zeroLogger) Debug(cat Category, subCat SubCategory, extra map[ExtraKey]interface{}, msg string) {
	log.logger.Debug().
		Str("Category", string(cat)).
		Str("SubCategory", string(subCat)).
		Fields(logPramsToZerolog(extra)).
		Msg(msg)
}

// Warn is a method to log warn level
func (log *zeroLogger) Warn(cat Category, subCat SubCategory, extra map[ExtraKey]interface{}, msg string) {
	log.logger.Warn().
		Str("Category", string(cat)).
		Str("SubCategory", string(subCat)).
		Fields(logPramsToZerolog(extra)).
		Msg(msg)
}

// Error is a method to log error level
func (log *zeroLogger) Error(cat Category, subCat SubCategory, extra map[ExtraKey]interface{}, msg string) {
	log.logger.Error().
		Str("Category", string(cat)).
		Str("SubCategory", string(subCat)).
		Fields(logPramsToZerolog(extra)).
		Msg(msg)
}

// Fatal is a method to log fatal level
func (log *zeroLogger) Fatal(cat Category, subCat SubCategory, extra map[ExtraKey]interface{}, msg string) {
	log.logger.Fatal().
		Str("Category", string(cat)).
		Str("SubCategory", string(subCat)).
		Fields(logPramsToZerolog(extra)).
		Msg(msg)
}
