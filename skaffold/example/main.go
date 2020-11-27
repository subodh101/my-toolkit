package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello yo!")

		time.Sleep(time.Second * 1)
	}
}
