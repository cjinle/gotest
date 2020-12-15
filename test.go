package test

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func DivLine(str string) {
	log.Println("===============", str, "===============")
}

func CheckErr(err error, opts ...bool) {
	if err != nil {
		if len(opts) == 0 || !opts[0] {
			log.Fatal(err)
		} else {
			log.Println(err)
		}
	}
}

func Print(v ...interface{}) {
	log.Println(v...)
}

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}
