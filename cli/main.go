package main

import (
	"CreditCard/cli/cmd"
	"fmt"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Errorf("failed to run command. Error: %s", err.Error())
	}
}
