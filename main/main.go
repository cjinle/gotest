package main

import (
	"flag"
	"fmt"

	//"github.com/cjinle/test/golog"
	"github.com/golang/glog"
)

func main() {
	fmt.Println("test main starting ... ")

	flag.Parse()
	defer glog.Flush()
	glog.Infof("hello")

	// logger := golog.NewLogger(&golog.LogConf{Path: "./log", Level: golog.Debug, Prefix: "aa_"})
	// logger.Printf(golog.Info, "test log info %v", [5]int{1, 2, 3, 4, 5})

	// mylib.ArrOuput()
	// mylib.Http()
	// mylib.LogOutput()
	// mylib.Redis()
	// mylib.Func()
}
