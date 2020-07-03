
package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"html/template"
)

type content struct {
	Content string

}

func main() {
	save()
}


func save() {
	filename := flag.String("file", "test3.txt", "Name of your file(.txt) you want to convert into a template(.html)")
	directory := flag.String("dir", "", "All you files(.txt) are contained in this directory")
	flag.Parse()
	writeTemplateFromFile(*filename)
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
