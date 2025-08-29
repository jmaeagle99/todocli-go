/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/jmaeagle99/todocli/cli"
)

func main() {
	err := cli.NewCommand().Execute()
	if err != nil {
		os.Exit(1)
	}
}
