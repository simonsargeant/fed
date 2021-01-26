package main

import (
	"fmt"
	"os"

	"github.com/simonsargeant/fed/internal/command"
)

func main() {
	root := command.NewRoot()

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
