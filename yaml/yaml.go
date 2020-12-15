package yaml

import (
	"bufio"
	"io/ioutil"
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

type M map[interface{}]interface{}

func GetCfg() {
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

func GetCfg2() {
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

func ParseGceCfg() {
	DivLine("ParseGceCfg Read File")
	bytes, err := ioutil.ReadFile("gce-tw2.yml")
	CheckErr(err)

	DivLine("v")
	var v interface{}
	err = yaml.Unmarshal(bytes, &v)
	CheckErr(err)
	Print(v)

	DivLine("done")
}
