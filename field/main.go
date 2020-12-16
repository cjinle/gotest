package main

import (
	"log"
	"reflect"
)

type MyStruct struct {
	Field1 int
	Field2 string
	Field3 float64
}

func (ms *MyStruct) UpdateField(field string, v interface{}) bool {
	log.Println("update field", field, v)
	switch field {
	case "Field1":
		ms.Field1 = v.(int)
	case "Field2":
		ms.Field2 = v.(string)
	case "Field3":
		ms.Field3 = v.(float64)
	}
	return true
}

func (ms *MyStruct) UpdateField2(field string, v interface{}) bool {
	log.Println("update field2", field, v)
	ptr := reflect.ValueOf(ms)
	s := ptr.Elem()
	if s.Kind() != reflect.Struct {
		return false
	}
	f := s.FieldByName(field)
	if !f.IsValid() {
		return false
	}
	f.Type()
	val := reflect.Indirect(reflect.ValueOf(ms))
	if val.CanAddr() && val.FieldByName(field).CanAddr() {

	}
	return false
}

func main() {
	ms := &MyStruct{1, "a", 2.0}
	ms.UpdateField("Field1", 2)
	log.Println(ms)
	ms.UpdateField("Field2", "b")
	log.Println(ms)
	ms.UpdateField("Field3", 3.0)
	log.Println(ms)

	val := reflect.Indirect(reflect.ValueOf(ms))

	log.Println(val.FieldByName("Field2").Type().String())
}
