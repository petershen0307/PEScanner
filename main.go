package main

import (
	"fmt"

	"github.com/petershen0307/PEScanner/cmd"
)

func main() {
	config := cmd.Parse()
	fmt.Println(config)
}
