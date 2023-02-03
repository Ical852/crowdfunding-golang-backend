package helper

import "log"

func ErrorHelper(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}