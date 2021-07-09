package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in> ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func sloppy_copy() {
	var (
		reader FooReader
		writer FooWriter
	)

	input := make([]byte, 4096)

	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}

func sexy_copy() {
	var (
		reader FooReader
		writer FooWriter
	)

	fmt.Println("here1")
	success, err := io.Copy(&writer, &reader)
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("unable to read/write data")
	}

	return
}

func main() {
	sexy_copy()
}
