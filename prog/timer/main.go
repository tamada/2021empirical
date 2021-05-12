package main

import (
	"fmt"
	"os"
	"time"
)

func handle(opts *options) int {
	return opts.commands.Execute()
}

func perform(opts *options) int {
	select {
	case <-time.After(opts.time.Duration()):
		return handle(opts)
	}
}

func goMain(args []string) int {
	opts, err := ParseArgs(args)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	if opts.help {
		fmt.Println(helpMessage(args[0]))
		return 0
	}
	return perform(opts)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
