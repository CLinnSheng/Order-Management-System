package common

import "syscall"

func GetEnvVar(key string, fallback string) string {
	if val, ok := syscall.Getenv(key); ok {
		return val
	}

	return fallback
}
