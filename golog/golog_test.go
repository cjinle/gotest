package golog

import (
	"testing"
)

func TestInfo(t *testing.T) {
	logger := NewLogger(&LogConf{Path: "./log", Level: Debug, Prefix: "aa_"})
	defer logger.Close()
	logger.Printf(Debug, "hello world test")
}
