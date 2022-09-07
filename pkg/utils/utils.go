package utils

import "log"

func HandleErr(errMsg string, err error) {
	if err != nil {
		log.Fatal(errMsg, err)
	}
}
