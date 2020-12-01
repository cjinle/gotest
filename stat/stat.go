package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	TxId     string `json:txid`
	OrderId  string `json:orderid`
	Status   int    `json:status`
	Detail   string `json:detail`
	Channel  string `json:channel`
	Amount   int    `json:amount`
	Currency string `json:currency`
}

/*
20201201 14:36:31|status err REQUEST:{"d":"{\"sid\":2,\"fc\":\"pay.complete\",\"mid\":\"16069309\",\"api\":16830976,\"t\":1606808193,\"v\":\"1.9.5\",\"l\":2,\"param\":{\"cardno\":\"10\",\"orderId\":\"BF16069309_64760882_50_104\",\"paymode\":10}}","s":"f0a662ec3a921431dac82f42d57dc776"}|data:{"txid":"20201201143629859062","orderid":"BF16069309_64760882_50_104","status":609,"detail":"Invalid PIN","channel":"12call","amount":0,"currency":"THB"}
*/
func main() {
	fmt.Println("stat start ... ")
	file, err := os.Open("cashcard.log")
	checkErr(err, "")
	defer file.Close()

	statInfo := make(map[string]map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		arr := strings.Split(scanner.Text(), "|")
		if len(arr) < 3 {
			continue
		}
		month := arr[0][:6]
		if _, ok := statInfo[month]; !ok {
			statInfo[month] = make(map[string]int)
		}
		result := &Result{}
		err = json.Unmarshal([]byte(strings.TrimLeft(arr[2], "data:")), result)
		if err != nil {
			continue
		}
		// fmt.Println(result)
		statInfo[month][result.Channel+"|"+result.Detail]++

	}
	// fmt.Println(statInfo)
	str, _ := json.Marshal(statInfo)
	fmt.Println(string(str))
}

func checkErr(err error, str string) {
	if err != nil {
		fmt.Println(str)
		panic(err)
	}
}
