package member

import (
	"log"
	"net"
	"net/rpc"
	"net/url"
	"os"

	"github.com/cjinle/test/member/gameinfo"
	"github.com/cjinle/test/member/info"
	"github.com/cjinle/test/member/rpcdata"
)

// Member 数据结构
type Member struct {
	GI   *gameinfo.GameInfo
	Info *info.Info
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

// NewMember 初始化对象
func NewMember() *Member {
	member := &Member{}
	member.GI = gameinfo.New()
	member.Info = info.New()
	return member
}

// GetGameInfo 获取用户信息
func (member *Member) GetGameInfo(args *rpcdata.Args, reply *rpcdata.GameInfo) error {
	ugi, err := member.GI.GetCache(args.Mid)
	if err != nil {
		return err
	}
	reply.Ret = 0
	reply.Money = ugi.Money
	return nil
}

// AddMoney 用户加减钱
func (member *Member) AddMoney(args *rpcdata.ArgsAddMoney, reply *rpcdata.ReplyAddMoney) error {
	ugi, err := member.GI.GetCache(args.Mid)
	if err != nil {
		return err
	}
	err = ugi.AddMoney(args.Mode, args.Num)
	if err != nil {
		return err
	}
	reply.Ret = 0
	reply.Mid = ugi.Mid
	reply.Money = ugi.Money
	return nil
}

// RPCServer 启动
func RPCServer() {
	rpcServer := rpc.NewServer()
	member := NewMember()
	rpcServer.Register(member)

	link := "unix:///tmp/rpc2.sock"
	urlInfo, _ := url.Parse(link)
	if urlInfo.Scheme == "unix" {
		os.Remove(urlInfo.Path)
	}
	// listen, err := net.Listen(urlInfo.Scheme, urlInfo.Path)
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal("listen error ", err)
	}
	rpcServer.Accept(listen)
}
