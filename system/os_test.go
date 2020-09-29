package system

import (
	"testing"
)

func TestHostName(t *testing.T) {

	name, err := GetHostName()
	if err != nil {
		t.Error("get hostname error: ", err)
	}
	if name == "debian" {
		t.Log("hostname is debian")
	} else {
		t.Error("hostname is not debian")
	}
}
