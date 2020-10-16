package tutorial

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// --------- string/[]bype ------------

func Buffer() {
	var b bytes.Buffer

	b.Write([]byte("hello"))

	fmt.Fprintf(&b, "world!")

	io.Copy(os.Stdout, &b)

}


// --------- array/slice --------------

func Append() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, y...)

	for i := 7; i < 20; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
}
