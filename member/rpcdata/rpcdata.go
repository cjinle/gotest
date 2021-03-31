package rpcdata

type Args struct {
	Mid uint32
}

type Reply struct {
	Ret  int
	Msg  string
	Data interface{}
}

type GameInfo struct {
	Ret   int
	Mid   uint32
	Money uint32
	Coin  uint32
}

type ArgsAddMoney struct {
	Mid  uint32
	Mode uint32
	Num  int32
}

type ReplyAddMoney struct {
	Ret   int
	Mid   uint32
	Money uint32
}
