package helper

import (
	"log"
	"strings"
)

func HandleError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func IsBlank(str string) bool {
	str = strings.TrimSpace(str)
	return len(str) == 0
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}
