package logging

import (
	"io"
	"runtime"

	"github.com/rs/zerolog"
)

var Verbose bool = false

var InfoLogger zerolog.Logger
var ErrorLogger zerolog.Logger

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

	InfoLogger = zerolog.New(infoHandle).With().Timestamp().Logger()

	ErrorLogger = zerolog.New(errorHandle).With().Timestamp().Logger()

}

func Info(msg string, fields ...Pair) {
	if Verbose {

		fieldsMap := make(map[string]interface{})

		for _, field := range fields {
			fieldsMap[field.Key] = field.Value
		}

		_, caller, _, ok := runtime.Caller(1)
		if ok {
			InfoLogger.Info().Str(CALLER_FIELD, caller).Fields(fieldsMap).Msg(msg)
		} else {
			InfoLogger.Info().Fields(fieldsMap).Msg(msg)
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
			ErrorLogger.Info().Str(CALLER_FIELD, caller).Fields(fieldsMap).Msg(msg)
		} else {
			ErrorLogger.Info().Fields(fieldsMap).Msg(msg)
		}

	}
}
