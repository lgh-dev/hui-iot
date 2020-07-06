package utils

import "log"

func CheckErrorLn(prefix string, err error, v ...interface{}) bool {
	if err != nil {
		log.Printf(prefix+"\n", err, v)
		return false
	}
	return true
}

func CheckError(prefix string, err error, v ...interface{}) bool {
	if err != nil {
		log.Printf(prefix+"\n", err, v)
		return false
	}
	return true
}
