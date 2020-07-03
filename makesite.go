package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"
	"os"
	"github.com/labstack/gommon/color"
)

type create struct {
	Content string
}

func readFile(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(fileContents))
	return string(fileContents)
}
