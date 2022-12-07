package customlogger

import (
	"log"
	"os"
	"sync"
)

type logger struct {
	filename string
	*log.Logger
}

var logge *logger
var once sync.Once

func GetInstance() *logger {
	once.Do(func() {
		logge = createLogger("logging.log")
	})
	return logge
}

func createLogger(fname string) *logger {

	return &logger{
		filename: fname,
		Logger:   log.New(os.Stdout, "", log.LstdFlags|log.Llongfile),
	}
}
