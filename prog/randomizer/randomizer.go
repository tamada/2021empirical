package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// 返り値に名前を付けてもOK！
func copyData(reader io.Reader, writer io.Writer, size int64) (statusCode int) {
	writtenSize, err := io.CopyN(writer, reader, size)
	if err != nil {
		fmt.Println(err.Error())
		return 2
	}
	if writtenSize != size {
		fmt.Printf("written data is not equals to %d (written %d bytes)\n", size, writtenSize)
		return 3
	}
	return 0
}

func perform(opts *options) int {
	writer, err := opts.Destination()
	if err != nil {
		return 1
	}
	defer writer.Close()
	return copyData(rand.Reader, writer, opts.size)
}

func goMain(args []string) int {
	opts, err := parseOptions(args)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	if printHelpIfNeeeded(args[0], opts) {
		return 0
	}
	return perform(opts)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
