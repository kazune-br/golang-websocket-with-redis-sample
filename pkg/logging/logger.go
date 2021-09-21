package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var logger zerolog.Logger

func InitLogger() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	logger = zerolog.New(os.Stderr).With().Timestamp().Str("tag", "app").Logger()
}

func Fatal(err error, msg string) {
	logger.Fatal().Stack().Caller().Err(err).Msg(msg)
}

func Error(err error, msg string) {
	logger.Error().Stack().Caller().Err(err).Msg(msg)
}

func Info(msg string) {
	logger.Info().Msg(msg)
}

func Infof(msg string, v ...interface{}) {
	logger.Info().Msgf(msg, v...)
}
