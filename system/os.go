package system

import (
	"os"
)

func GetHostName() (string, error) {
	return os.Hostname()
}

func GetEnv() string {
	return "PATH: " + os.Getenv("PATH")
}
