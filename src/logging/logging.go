package logging

import "log"

var logger *log.Logger

func init() {
	logger = log.Default()
}

func LogError(err error) {
	Log(err.Error())
}

func Log(message string) {
	logger.Println(message)
}
