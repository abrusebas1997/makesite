package main


import (
	"fmt"
	"makesite/files"
	"html/template"
	"bytes"
	"os"
	"log"
	"path/filepath"

)



func main() {
	RenderTemplate()
}
type data struct {
	Content template.HTML
}

func print(text string) {
	fmt.Println(text)
}

var templatesDir string = files.RootDir() + "/generated_templates"
var templatePath string = files.RootDir() + "template.tmpl"

var file string = "first-post.txt"


func RenderTemplate() {
	post_data := files.ReadFile(file)
	contentType := data{Content: template.HTML(post_data)}


	tmpl, err := template.ParseFiles("template.tmpl")

	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer

	if err := tmpl.ExecuteTemplate(&buf, "template.tmpl", contentType); err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
		if err = os.Mkdir(templatesDir, 0777); err != nil {
			log.Printf("Error occured when creating generated_templates directory. Error: %v", err)
		}
	}


	fileInfo := files.WriteToFile(templatesDir+"/"+filepath.Base(file), buf.Bytes())


	fmt.Println(fileInfo)
	//
	// 	var buf bytes.Buffer
	//
	// 	if err := tmpl.ExecuteTemplate(&buf, "template.tmpl", contentType); err != nil {
	// 			log.Fatal(err)
	// 	}
	//
	// 		tmpl = buf.Bytes()
	//
	// 		if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
	// 			if err = os.Mkdir(templatesDir, 0777); err != nil {
	// 				log.Printf("Error occured when creating generated_templates directory. Error: %v", err)
	// 			}
	// 		}
	//
	// 		fmt.Println(temp)

}
