package main

import (
	"fmt"
	"io"
	"os"
)

func showUsage(statusCode int) int {
	fmt.Println("copy <FROM> <TO>")
	return statusCode
}

func copyImpl(reader io.Reader, writer io.Writer) int {
	io.Copy(writer, reader)
	return 0
}

func perform(from, to string) int {
	reader, err := os.Open(from)
	if err != nil {
		return 2
	}
	defer reader.Close()
	writer, err := os.Create(to)
	if err != nil {
		return 3
	}
	defer writer.Close()
	return copyImpl(reader, writer)
}

func goMain(args []string) int {
	if len(args) != 3 {
		return showUsage(1)
	}
	return perform(args[1], args[2])
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
