/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/sabilimaulana/studybuddy/cmd"
	"github.com/sabilimaulana/studybuddy/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
