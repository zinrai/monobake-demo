package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("hello worker")
		time.Sleep(time.Second)
	}
}
