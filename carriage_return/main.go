package main

import (
	"fmt"
	"time"
)

func main() {
	for range time.Tick(100 * time.Microsecond) {
		fmt.Printf("\r %s", time.Now())
	}
}
