package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

// Forma de trabalhar com json sem ter que criar struct
func main() {
	var p fastjson.Parser
	jsonData := `{"foo": "bar", "nun": 123, "bool": true, "arr": [1, 2, 3]}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("foo=%s\n", v.GetStringBytes("foo"))
	fmt.Printf("nun=%d\n", v.GetInt("nun"))
	fmt.Printf("bool=%v\n", v.GetStringBytes("bool"))
	fmt.Printf("arr=%s\n", v.GetStringBytes("foo"))

	a := v.GetArray("arr")
	for i, value := range a {
		fmt.Printf("Index %d: %s\n", i, value)
	}

}
