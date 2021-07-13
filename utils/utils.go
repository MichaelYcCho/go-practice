package utils

import "log"

func HadlerErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
