package push

import (
	"log"
	"testing"
)

func TestOne(t *testing.T) {
	push := New()
	log.Println(push, push.config)
	push.Start()
}
