package main

import (
	"fmt"
	"time"
)

const Version = "1.0.0"

func main() {
	fmt.Printf("Starting application version %s\n", Version)
	
	for {
		fmt.Printf("Hello World - Version: %s\n", Version)
		time.Sleep(5 * time.Second)
	}
}
