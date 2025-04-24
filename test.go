package main

import (
	"fmt"
	"time"
)

func main() {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Now())
	fmt.Println(time.Now().In(loc))
	fmt.Println(loc)
	fmt.Println(loc.String())
}