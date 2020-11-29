package main

import (
	"Go-000/Week02/service"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := service.UserService{}
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		input = 0
	}
	_, err = s.GetUser(input)
	fmt.Printf("%+v\r\n", err)
}
