package main

import (
	"fmt"
	"log"
	"reflect"
)

type MyStruct struct {
	Field1 int
	Field2 string
	Field3 float64
}

var MyFieldType = map[string]reflect.Kind{
	"Field1": reflect.Int,
	"Field2": reflect.String,
	"Field3": reflect.Float64,
}

var MyFieldType2 = map[string]interface{}{
	"Field1": FormatInt,
	"Field2": FormatString,
	"Field3": FormatFloat,
}

func FormatInt(oldVal, newVal interface{}) (int, error) {
	ov := reflect.ValueOf(oldVal)
	nv := reflect.ValueOf(newVal)
	if ov.Kind() != nv.Kind() {
		return oldVal.(int), fmt.Errorf("oldValue: %v(%T) newValue: %v(%T) type error", oldVal, newVal)
	}
	ov.SetInt(nv.Int())
	return newVal.(int), nil
}

func FormatString(oldVal, newVal interface{}) {}
func FormatFloat(oldVal, newVal interface{})  {}

func (ms *MyStruct) UpdateField(field string, v interface{}) bool {
	log.Println("update field", field, v)
	if rtype, ok := MyFieldType[field]; !ok || reflect.TypeOf(v).Kind() != rtype {
		return false
	}

	switch field {
	case "Field1":
		ms.Field1 = v.(int)
	case "Field2":
		ms.Field2 = v.(string)
	case "Field3":
		ms.Field3 = v.(float64)
		// val := reflect.ValueOf(v)

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
	err := ms.UpdateField("Field1", 2)
	log.Println(ms, err)
	err = ms.UpdateField("Field2", 2)
	log.Println(ms, err)
	err = ms.UpdateField("Field3", 3.0)
	log.Println(ms, err)

}
