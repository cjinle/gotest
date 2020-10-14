package sorting

import (
	"fmt"
	"testing"
	"github.com/cjinle/test/utils"
)

func TestBubbleSort(t *testing.T) {
	arr := utils.RandArray(10)
	fmt.Println(BubbleSort(arr))
	arr = utils.RandArray(20)
	fmt.Println(BubbleSort2(arr))
}


