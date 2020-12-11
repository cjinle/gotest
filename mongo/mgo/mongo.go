package mongo

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Row struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `bson:"name"`
}

// type M map[string]interface{}

var session *mgo.Session

func init() {
	session = New()
}

func New() *mgo.Session {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
	}
	return session
}

func Close() {
	defer session.Close()
}

func GetCollection() *mgo.Collection {
	return session.DB("test").C("testcol")
}

func Find() {
	// log.Println(session)
	coll := session.DB("test").C("testcol")
	// var result interface{}

	result := Row{}
	err := coll.Find(nil).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	var result2 []Row
	err = coll.Find(nil).All(&result2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result2)

}

func Insert() {
	coll := session.DB("test").C("testcol")
	row := &Row{Name: "golang"}
	err := coll.Insert(row)
	if err != nil {
		log.Fatal(err)
	}

	row2 := map[string]string{"name": "chenjinle"}
	err = coll.Insert(&row2)
	if err != nil {
		log.Fatal(err)
	}

	err = coll.Insert(bson.M{"name": "wangwu"})
	if err != nil {
		log.Fatal(err)
	}
}

func Update() {
	coll := GetCollection()
	err := coll.Update(bson.M{"_id": bson.ObjectIdHex("5fd2f001a4b6dcaeb7d4be3a")}, bson.M{"name": "bbbb"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(bson.NewObjectId().Hex())
}

func Delete() {
	coll := GetCollection()
	err := coll.RemoveId(bson.ObjectIdHex("5fd2f12ea4b6dcaeb7d4be4e"))
	if err == mgo.ErrNotFound {
		log.Println("=========", err, "==========")
	} else if err != nil {
		log.Println(err)
	}
	log.Println("continue ... ")

	err = coll.Remove(bson.M{"name": "golang"})
	if err != nil {
		log.Println(err)
	}
}
