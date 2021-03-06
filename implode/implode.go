package implode

import (
	"fmt"
	"strings"
)

// Implode str...
func Implode(sep string, v ...interface{}) string {
	data := make([]string, len(v))
	for idx, val := range v {
		data[idx] = fmt.Sprint(val)
	}
	return strings.Join(data, sep)
}

// Explode sth...
func Explode(seq, str string) []string {
	return []string{"hello", "word"}
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type File struct {
}

func (f *File) Write(b []byte) (n int, err error) {
	return 0, nil
}

func Create(name string) (file *File, err error) {
	var f File
	fmt.Println(name)
	return &f, nil
}

func NewWriter(w Writer) *Writer {
	return &w
}

func Run() {
	f, _ := Create("aaa")

	// w := NewWriter(f)
	// fmt.Printf("%T,%v", w, w)
	var w Writer = f
	// var w Writer = &File{}
	fmt.Println(w.Write([]byte("hello")))

	idBytes := []byte{1, 15: 0}
	fmt.Println(len(idBytes), cap(idBytes))
}
