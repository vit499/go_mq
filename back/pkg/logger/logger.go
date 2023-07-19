package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

var (
	logger Logger
	once   sync.Once
)

func Get() *Logger {

	once.Do(func() {
		zeroLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
		loglevel := "info"
		// Set proper loglevel based on config
		switch loglevel {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "info":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "warn", "warning":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "err", "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		case "panic":
			zerolog.SetGlobalLevel(zerolog.PanicLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel) // log info and above by default
		}
		logger = Logger{&zeroLogger}
	})
	return &logger
}
