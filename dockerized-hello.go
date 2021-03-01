package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	hostname, _ := os.Hostname()
	fmt.Printf("I am %s running on %s/%s.", hostname, runtime.GOOS, runtime.GOARCH)
}
