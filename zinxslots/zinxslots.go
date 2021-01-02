package zinxslots

import (
	"encoding/json"
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/cjinle/test/goslots"
)

type BetRouter struct {
	znet.BaseRouter
}

type BetInfo struct {
	Bet int `json:"bet"`
}

type BetResult struct {
	Win int   `json:"win"`
	Ret []int `json:"ret"`
}

func (this *BetRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call BetRouter Handle")
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	betInfo := &BetInfo{}
	err := json.Unmarshal(request.GetData(), betInfo)
	if err != nil {
		fmt.Println("json unmarshal err", err)
		return
	}
	money, err := request.GetConnection().GetProperty("money")
	if err != nil {
		fmt.Println("bet get property money err", err)
		return
	}
	id, _ := request.GetConnection().GetProperty("id")
	if money.(int) < betInfo.Bet {
		errStr := fmt.Sprintf("id=%d, money=%d", id, money)
		fmt.Println(errStr)
		request.GetConnection().SendMsg(1, []byte(errStr))
		request.GetConnection().Stop()
		return
	}
	request.GetConnection().SetProperty("money", money.(int)-betInfo.Bet)
	betResult := &BetResult{}
	betResult.Win, betResult.Ret = goslots.Bet(betInfo.Bet)
	bytes, _ := json.Marshal(betResult)
	fmt.Println("betResult = ", string(bytes), id, money)
	err = request.GetConnection().SendMsg(0, []byte(bytes))
	if err != nil {
		fmt.Println(err)
	}
}

var ID int = 1
var StartMoney int = 50000

func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnecionBegin is Called ... ")

	fmt.Println("Set conn Name, Home done!")
	conn.SetProperty("id", ID)
	conn.SetProperty("money", StartMoney)

	err := conn.SendMsg(2, []byte(fmt.Sprintf("DoConnection BEGIN ID(%d)...", ID)))

	if err != nil {
		fmt.Println(err)
	}
	ID++
}

func DoConnectionLost(conn ziface.IConnection) {
	if id, err := conn.GetProperty("id"); err == nil {
		fmt.Println("id = ", id)
	}

	if money, err := conn.GetProperty("money"); err == nil {
		fmt.Println("Conn Property money = ", money)
	}

	fmt.Println("DoConneciotnLost is Called ... ")
}

func Start() {
	fmt.Println("zinxslots start ... ")
	s := znet.NewServer()

	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	s.AddRouter(0, &BetRouter{})
	s.Serve()
}
