package logging

import (
	"io"
	"runtime"

	"github.com/rs/zerolog"
)

var Verbose bool = false

var infoLogger zerolog.Logger
var errorLogger zerolog.Logger

type Pair struct {
	Key   string
	Value interface{}
}

func NewPair(key string, value interface{}) Pair {
	return Pair{
		Key:   key,
		Value: value,
	}
}

const (
	CALLER_FIELD = "caller"
)

func Init(
	infoHandle io.Writer,
	errorHandle io.Writer) {

	infoLogger = zerolog.New(infoHandle).With().Timestamp().Logger()

	errorLogger = zerolog.New(errorHandle).With().Timestamp().Logger()

}

func Info(msg string, fields ...Pair) {
	if Verbose {

		fieldsMap := make(map[string]interface{})

		for _, field := range fields {
			fieldsMap[field.Key] = field.Value
		}

		_, caller, _, ok := runtime.Caller(1)
		if ok {
			infoLogger.Info().Str(CALLER_FIELD, caller).Fields(fieldsMap).Msg(msg)
		} else {
			infoLogger.Info().Fields(fieldsMap).Msg(msg)
		}

	}
}

func Error(msg string, fields ...Pair) {
	if Verbose {

		fieldsMap := make(map[string]interface{})

		for _, field := range fields {
			fieldsMap[field.Key] = field.Value
		}

		_, caller, _, ok := runtime.Caller(1)
		if ok {
			errorLogger.Info().Str(CALLER_FIELD, caller).Fields(fieldsMap).Msg(msg)
		} else {
			errorLogger.Info().Fields(fieldsMap).Msg(msg)
		}

	}
}
