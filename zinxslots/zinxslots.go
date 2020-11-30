package zinxslots

import (
	"encoding/json"
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/cjinle/goslots"
)

type BetRouter struct {
	znet.BaseRouter
}

type BetInfo struct {
	Bet int `json:bet`
}

type BetResult struct {
	Win int `json:win`
	Ret []int `json:ret`
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
	betResult := &BetResult{}
	betResult.Win, betResult.Ret = goslots.Bet(betInfo.Bet)
	// fmt.Println("betResult = ", string(BetResult))
	bytes, _ := json.Marshal(betResult)
	err = request.GetConnection().SendMsg(0, []byte(bytes))
	if err != nil {
		fmt.Println(err)
	}
}

func Start() {
	fmt.Println("zinxslots start ... ")
	s := znet.NewServer()
	s.AddRouter(0, &BetRouter{})
	s.Serve()
}
