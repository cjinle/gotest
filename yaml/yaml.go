package yaml

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	. "github.com/cjinle/test"
	"gopkg.in/yaml.v2"
)

type RedisCfg struct {
	ID   int    `yaml:"id"`
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}
type Cfg struct {
	Common struct {
		MongoDB  string `yaml:"mongodb"`
		Expire   int    `yaml:"expire"`
		TableNum int    `yaml:"tableNum"`
	} `yaml:"common"`
	Redis []struct {
		ID   int    `yaml:"id"`
		IP   string `yaml:"ip"`
		Port int    `yaml:"port"`
	} `yaml:"redis"`
}

type M map[interface{}]interface{}

func getCfg() {
	Print("GetCfg ... ")
	bytes, err := ioutil.ReadFile("config.yaml")
	CheckErr(err)

	DivLine("v")
	// v := make(map[string]RedisCfg{})
	// var v map[string][]map[string]interface{}
	var v map[string][]RedisCfg
	// var v interface{}
	err = yaml.Unmarshal(bytes, &v)
	CheckErr(err)

	Print(v["redis"][0].ID, v["redis"][0].IP, v["redis"][0].Port)

	DivLine("v2")
	v2 := map[string]interface{}{"canonical": 685230}
	Print(v2)

	DivLine("v3")
	var v3 map[string]interface{}
	err = yaml.Unmarshal(bytes, &v3)
	CheckErr(err)

	// log.Println(v3)
	// redisCfgArr := v3["redis"]
	Print(v3["redis"].([]interface{})[0].(M)["id"])
	// log.Println(redisCfgArr.([]RedisCfg)[0].(RedisCfg).ID)

	Print(reflect.ValueOf(v3["redis"]).Type())

}

func getCfg2() {
	DivLine("GetCfg2")
	file, err := os.Open("config.yaml")
	CheckErr(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	// var v interface{}
	var v map[interface{}]interface{}
	err = yaml.NewDecoder(reader).Decode(&v)
	CheckErr(err)
	Print(v["redis"].([]interface{})[0].(map[interface{}]interface{})["port"])
}

func getCfg3() {
	bytes, err := ioutil.ReadFile("config.yaml")
	CheckErr(err)

	DivLine("v")
	// var v map[string]interface{}
	var v Cfg
	err = yaml.Unmarshal(bytes, &v)
	CheckErr(err)
	log.Println(string(bytes), v, v.Common.MongoDB)
}

type Dns struct {
	Enable     bool     `yaml:"enable"`
	NameServer []string `yaml:"nameserver"`
	Fallback   []string `yaml:"fallback"`
}

type Proxy struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Password int    `yaml:"password"`
}

type ProxyGroup struct {
	Name     string   `yaml:"name"`
	Type     string   `yaml:"type"`
	Proxyies []string `proxies`
	Url      string   `yaml:"url"`
	Interval int      `yaml:"interval"`
}

type GceCfg struct {
	Port               int          `yaml:"port"`
	SocksPort          int          `yaml:"socks-port"`
	RedirPort          int          `yaml:"redir-port"`
	AllowLan           bool         `yaml:"allow-lan"`
	Mode               string       `yaml:"mode"`
	LogLevel           string       `yaml:"log-level"`
	ExternalController string       `yaml:"external-controller"`
	Secret             string       `yaml:"secret"`
	Dns                Dns          `yaml:"dns"`
	CfwBypass          []string     `yaml:"cfw-bypass"`
	CfwLatencyTimeout  int          `yaml:"cfw-latency-timeout"`
	Proxy              []Proxy      `yaml:"Proxy"`
	ProxyGroup         []ProxyGroup `yaml:"Proxy Group"`
	Rule               []string     `yaml:"Rule"`
}

func parseGceCfg() {
	DivLine("ParseGceCfg Read File")
	bytes, err := ioutil.ReadFile("gce-tw2.yml")
	CheckErr(err)

	DivLine("v")
	// var v interface{}
	var v GceCfg
	err = yaml.Unmarshal(bytes, &v)
	CheckErr(err)
	Print(v, v.ProxyGroup[0])

	if v.Dns.Enable {
		log.Println("dns is enable")
	}

	DivLine("done")

}
