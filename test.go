package test

import (
	"log"
)

func init() {
	// log.SetFlags(log.Lshortfile | log.LstdFlags)
}

// DivLine print the div line
func DivLine(str string) {
	log.Println("===============", str, "===============")
}

// CheckErr is check for error
func CheckErr(err error, opts ...bool) {
	if err != nil {
		if len(opts) == 0 || !opts[0] {
			log.Fatal(err)
		} else {
			log.Println(err)
		}
	}

}

// Print is print some values
func Print(v ...interface{}) {
	log.Println(v...)
}

// Fatal is fatal some values
// the other comments
func Fatal(v ...interface{}) {
	log.Fatal(v...)
}
