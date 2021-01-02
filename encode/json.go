package encode

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
)

func JsonEncode() {
	log.Println("json encode")
	user := &User{1000, "Zhang San", 92.8, 1, 20}
	str, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("json bytes: %s, len: %d", str, len(str))
	log.Println(hex.EncodeToString(str))

	var buf bytes.Buffer

	err = json.NewEncoder(&buf).Encode(user)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("json bytes: %s, len: %d", buf.Bytes(), len(buf.Bytes()))
	log.Println(hex.EncodeToString(buf.Bytes()))

}

func JsonDecode() {
	var err error
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		log.Println("========= recover err ============", e.(error))
	}()
	user := &User{1000, "Zhang San", 92.8, 1, 20}
	str, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}

	var newUser User
	err = json.Unmarshal(str, &newUser)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(newUser)

	err = json.NewDecoder(bytes.NewReader(str)).Decode(&newUser)
	log.Println(newUser)

	confFile := "/data/wwwroot/go/src/github.com/cjinle/ip2regionserver/conf/app.json"
	file, err := os.Open(confFile)
	if err != nil {
		log.Fatal(err)
	}
	var v interface{}
	err = json.NewDecoder(file).Decode(&v)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(v, v.(map[string]interface{})["listen"])
	// file, _ = os.Open(confFile)
	// log.Println(ioutil.ReadAll(file))
}
