package files

import (
  "io/ioutil"
  "fmt"
  "os"
  "path/filepath"
  "strings"
  "path"
  "runtime"

)

func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func WriteToFile(file string, template []byte) os.FileInfo {
	fileName := fmt.Sprintf("%s.html", strings.TrimSuffix(file, filepath.Ext(file)))

	// err := ioutil.WriteFile(fileName, template, 0666)
	err := ioutil.WriteFile(fileName, template, 0644)

	if err != nil {
		panic(err)
	}

	fileInfo, err := os.Stat(fileName)

	if err != nil {
		panic(err)
	}

	return fileInfo
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
