package main


import (
	"fmt"
	"makesite/files"
)

func main() {
	fmt.Println(files.ReadFile("first-post.txt"))
}
