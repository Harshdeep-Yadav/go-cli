package main

import (
	"fmt"

	"github.com/harshdeep/ghcli/cmd"
)

func main(){
	fmt.Println("In the main function")
	// Execute the root command
	cmd.Execute()
}