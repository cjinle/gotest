package implode

import (
	"crypto/rand"
	"fmt"
	"strings"
	"testing"

	"github.com/cjinle/test/utils"
)

func TestImplode(t *testing.T) {
	fmt.Println(utils.Md5Sum("123456"))

	A := []int{1, 2, 3}
	fmt.Println(utils.Implode2("&", A))

	bytes := make([]byte, 16)
	rand.Read(bytes)
	fmt.Printf("%x", bytes)
	utils.Assert(1 == 1, "1 == 1")
	utils.Assert(false == false, "false is false")

	str := "xxx"
	str = strings.TrimRight(str, "|")
	strings.Split(str, "|")
}

func TestImplode2(t *testing.T) {
	Run()

}
