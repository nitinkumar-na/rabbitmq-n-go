package helper

import "log"

func HandleError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
