package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/aceld/zinx/znet"
	"github.com/cjinle/test/zinxslots"
)

func main() {

	fmt.Println("Client Test ... start")

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		dp := znet.NewDataPack()
		betInfo := &zinxslots.BetInfo{5000}
		bytes, _ := json.Marshal(betInfo)
		msg, _ := dp.Pack(znet.NewMsgPackage(0, []byte(bytes)))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			fmt.Println("read head error")
			break
		}

		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}

			var v interface{}
			// var v zinxslots.BetResult
			err = json.Unmarshal(msg.Data, &v)
			if err != nil {
				fmt.Println(string(msg.Data))
				continue
			}
			fmt.Println(v, v.(map[string]interface{})["Ret"])
			// fmt.Println(v)

			// fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
		}

		time.Sleep(1 * time.Second)
	}
}
