
package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"github.com/labstack/gommon/color"
)

type FileCont struct {
	Title   string
	Content string
}

func generator() {
	fileContents, err := ioutil.ReadFile("third-post.txt")

	fmt.Println(string(fileContents))
	if err != nil {
		panic(err)
	}

	content := FileCont{
		Title:   "first-post-w",
		Content: string(fileContents),
	}
	t := template.Must(template.ParseFiles("template.tmpl"))

	file, err := os.Create("third-post.html")
	t.Execute(os.Stdout, content)
	t.Execute(file, content)
	file_flag := flag.String("latest-post.html", "defaultValue", ".txt")
	flag.Parse()
	fmt.Println("file:", *file_flag)
}

func main() {
	color.Println(color.Green("Welcome to makesite!"))
	generator()
}
