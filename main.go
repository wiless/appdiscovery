package main

import (
	"fmt"
	"go/build"
	"log"

	"github.com/kr/pretty"
)

func ParsePakage() {
	log.Print("Starting ..  ")
	var ctx build.Context
	ctx = build.Default
	// pkg, err := build.ImportDir("github.com/wiless/singlecell", build.FindOnly)
	pkg, err := ctx.Import("github.com/wiless/singlecell", "", build.AllowBinary)
	if err != nil {
		fmt.Printf("The Error %v", err)
		return
	}

	pretty.Printf("The pkg info is % #v", pkg)
	pretty.Printf("\nIs it Command % #v\n", pkg.IsCommand())
}
