package logging

import "log"

var logger *log.Logger

func init() {
	logger = log.Default()
}

func LogError(err error) {
	logger.Println(err.Error())
}
