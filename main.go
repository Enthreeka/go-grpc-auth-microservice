package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	//b := make([]byte, 32)
	//
	//s := rand.NewSource(time.Now().Unix())
	//r := rand.New(s)
	//
	//_, err := r.Read(b)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%x", string(b))

	var s struct {
		Data    string
		Time    time.Time
		Seasson string
	}
	s.Data = "f2342"
	s.Time = time.Now()

	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Name

		fmt.Println("field - ", field)
		fmt.Println("fieldName - ", fieldName)
	}
}
