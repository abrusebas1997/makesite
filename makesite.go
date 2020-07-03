package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"html/template"
	"fmt"
	"github.com/labstack/gommon/color"
)

type content struct {
	Content string

}

func main() {
	color.Println(color.Red("Start making your own websites!"))
	save()
}

// function that reads the file
func readFile(file string) string {
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}
// function that writes the file
func writeFile(fileName string) *os.File {
	fileName = strings.Split(fileName, ".")[0] + ".html"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	return file
}
// function that converts the .txt file to a html template
func writeTemplateFromFile(file string) {
	var fileData content
	fileData.Content = readFile(file)

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	File := writeFile(file)

	err := t.Execute(File, fileData)
	if err != nil {
		panic(err)
	}
}

func save() {

	filename := flag.String("file", "test3.txt", "Name of your file(.txt) you want to convert into a template(.html)")
	directory := flag.String("dir", "", "All you files(.txt) are contained in this directory")
	flag.Parse()
	// adding count of pages
	numberofPages := 0

	if *directory != "" {
		files, err := ioutil.ReadDir(*directory)
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			fname := f.Name()
			// making sure if the filename has .txt in it
			if strings.Contains(fname, ".") {
				fmt.Println("There is a file(.txt) in your directory", fname)
				writeTemplateFromFile((fname))
				numberofPages += 1
			}
		}

	} else {
		writeTemplateFromFile(*filename)
		numberofPages += 1
	}

	print(numberofPages)

}
