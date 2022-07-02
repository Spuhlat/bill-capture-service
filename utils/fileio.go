package utils

import (
	"fmt"
	"os"
	"io/ioutil"
)

func ReadFromFile(fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return []byte("Error")
	}
	return data
}

func WriteToFile(fileName, content string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := content
	_, err = fmt.Fprintln(f, newLine)

	if err != nil {
		fmt.Println(err)
		f.Close()
		return 
	}
}