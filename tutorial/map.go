package tutorial

import "fmt"

type La struct {
	x, y int
}

func Map() {
	var m = map[string]La{
		"a": {1, 2},
		"b": {3, 4},
	}

	fmt.Println(m["a"].x + m["b"].y)
}
