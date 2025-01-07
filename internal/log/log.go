package log

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

func CreateLogger(toFile bool) zerolog.Logger {
	if toFile {
		f, err := os.OpenFile(
			time.Now().Format("./logs/2006_01_02_examly.log"),
			os.O_WRONLY|os.O_CREATE|os.O_APPEND,
			0644,
		)
		if err != nil {
			panic(err)
		}
		return zerolog.New(f).Level(zerolog.TraceLevel).With().Caller().Timestamp().Logger()
	}

	return zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "02.01.2006 15:04:05",
	}).Level(zerolog.TraceLevel).With().Caller().Timestamp().Logger()
}
