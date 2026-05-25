package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

func readAndWrite(i ReadWriter) {
	i.Read()
	i.Write()
}

type FileReadWriter struct{}

func (r FileReadWriter) Read() {
	fmt.Println("Read")
}

func (r FileReadWriter) Write() {
	fmt.Println("Write")
}

func main() {
	fileReadeWriter := FileReadWriter{}

	readAndWrite(fileReadeWriter)
}
