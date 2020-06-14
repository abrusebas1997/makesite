package files

import (
  "io/ioutil"
  
)

func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(data)
}
