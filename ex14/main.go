package main

import (
	"fmt"
	"reflect"
)

// type assertion
func AssertType(i interface{}) {
	if _, ok := i.(string); ok {
		fmt.Printf("это string\n")
	} else if _, ok := i.(int); ok {
		fmt.Printf("это integer\n")
	} else if _, ok := i.(bool); ok {
		fmt.Printf("это boolian\n")
	} else if _, ok := i.(chan int); ok {
		fmt.Printf("это channel\n")
	} else {
		fmt.Printf("это неизвестный тип\n")
	}
}

// пакет reflect
func ReflectType(i interface{}) {
	fmt.Printf("это %v\n", reflect.TypeOf(i))
}

func main() {
	set := map[string]interface{}{
		"x": 0,
		"y": "",
		"z": false,
		"w": make(chan interface{}),
		"q": make(chan int),
	}

	func(fs ...func(interface{})) {
		for _, fn := range fs {
			for k, v := range set {
				fmt.Printf("%s - ", k)
				fn(v)
			}
			fmt.Println("-------------")
		}
	}(AssertType, ReflectType)

}
