package main

import (
	// "flag"
	"fmt"
	"io/ioutil"
	"os"
	"html/template"
)

type Data struct {
	Content string
}

func writeFile() {
	bytesToWrite := []byte("filename")
	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}


	return string(fileContents)
}

func writeTemplate() {
	var fileData Data
	fileData.Content = readFile()

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err := t.Execute(os.Stdout, fileData)
	if err != nil {
		panic(err)
	}
}

func renderTemplate() string{
	var filename string
	bytesToWrite := []byte(filename)
	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
	return fmt.Sprintf("%s file created", err)
}
func main() {
	// flags
	writeTemplate()
}
