package utils

import "log"

func CheckAndExit(e error) {
	if e != nil {
		log.Fatal("\t[ERROR]\t|\t", e)
	}
}
