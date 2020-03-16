package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file := "xxx.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	// logger := log.New(logFile, "", log.Ldate|log.Ltime)
	logger := log.New(logFile, "", log.LstdFlags)

	fmt.Println(logger.Flags())

	fmt.Println("--------------")

	fmt.Println(logger.Prefix())

	logger.Println("xcvxcvxcvxcv")

	logger.Output(2, "output log message!")
	logger.Println("output log message!")
	logger.Printf("name: %s, age: %d", "chenjinle", 12)

	std := log.New(os.Stderr, "", log.LstdFlags)
	std.Println("sdfsdf")

}
