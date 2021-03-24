package encode

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"

	"github.com/ugorji/go/codec"
)

// GobEncode func
func GobEncode() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Gob Encode ... ")
	user := &User{1000, "Zhang San", 92.8, 1, 20}
	str, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("json bytes: %s, len: %d", str, len(str))
	log.Println(hex.EncodeToString(str))

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(user)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("gob byes: %s, len: %d", buf.Bytes(), len(buf.Bytes()))
	log.Println(hex.EncodeToString(buf.Bytes()))

	var w io.Writer
	var mh codec.MsgpackHandle
	var b []byte
	enc2 := codec.NewEncoder(w, &mh)
	enc2 = codec.NewEncoderBytes(&b, &mh)
	err = enc2.Encode(user)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("msgpack byes: %s, len: %d", b, len(b))

	log.Println(hex.EncodeToString(b))
	log.Println(base64.StdEncoding.EncodeToString(b))

}
