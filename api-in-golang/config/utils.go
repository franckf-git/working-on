package config

import "log"

func ErrorLogg(message ...interface{}) {
	if debug {
		log.Printf("Error: %#+v\n", message...)
	}
}
