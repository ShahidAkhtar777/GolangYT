package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang::")
	content := "Hardwork and patience makes together makes fortune."

	files, err := os.Create("./myfile.txt")
	checkNilError(err)

	length, err := io.WriteString(files, content)
	checkNilError(err)

	fmt.Println("Total length of the content is:", length)
	defer files.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Content in the file is:\n", string(databyte))
}

func checkNilError(err error){
	if err!=nil{
		panic(err)
	}
}
