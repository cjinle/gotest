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

func GetProxyEnv() string {
	return "proxy env: " + os.Getenv("http_proxy")
}
