package main


import (
	"fmt"
)

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

func main() {
	f, _ := Create("aaa")

	// w := NewWriter(f)
	// fmt.Printf("%T,%v", w, w)
	var w Writer = f
	// var w Writer = &File{}
	fmt.Println(w.Write([]byte("hello")))
	
}