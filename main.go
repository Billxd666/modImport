package main

import (
	"fmt"

	"github.com/Billxd666/modImport.git/greetings"
)

func main() {
	message := greetings.Hello("Billy")
	fmt.Println(message)
}
