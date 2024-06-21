package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Forma de trabalhar com json sem ter que criar struct
func main() {
	var p fastjson.Parser
	jsonData := `{"user": {"name": "john", "age": 30}}`

	value, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	user := value.GetObject("user")

	fmt.Printf("User name: %s\n", user.Get("name"))
	fmt.Printf("User age: %s\n", user.Get("age"))

	userJSON := value.GetObject("user").String()

	var user2 User
	if err := json.Unmarshal([]byte(userJSON), &user2); err != nil {
		panic(err)
	}

	fmt.Printf(user2.Name, user2.Age)
}

//todo testAR RODAR POIS NN FUNIONOU
