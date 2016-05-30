package main

import (
	//"errors"
	"fmt"
	"reflect"
	//"github.com/graphql-go/graphql"
)

type Test map[string]string

type Ts struct {
	a string `json:"species=gopher;color=blue" ibbd:"test2"`
	b string `json:"b" ibbd:"test"`
}

func main() {
	t := Test{
		"a": "test",
		"b": "test",
	}

	for k, v := range t {
		fmt.Printf("\nkey = %+v, val = %+v ", k, v)
	}

	ts := &Ts{
		a: "tes",
		b: "tes",
	}
	fmt.Printf("\n%+v  ", ts)
	s := reflect.TypeOf(ts).Elem()
	for i := 0; i < s.NumField(); i++ {
		fmt.Println("====")
		//fmt.Println(s.Field(i).Tag)                            //将tag输出出来
		fmt.Printf("\njson: %s  ", s.Field(i).Tag.Get("json")) //将tag输出出来
		fmt.Printf("\nibbd: %s  ", s.Field(i).Tag.Get("ibbd")) //将tag输出出来
		fmt.Println(s.Field(i).Type)                           //将tag输出出来
	}

}
