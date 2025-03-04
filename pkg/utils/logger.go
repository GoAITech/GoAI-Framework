package utils

import "log"

// LogInfo logs general info
func LogInfo(msg string) {
	log.Println("[INFO]:", msg)
}
