package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

type options struct {
	size     int64
	dest     string
	helpFlag bool
	args     []string
}

func helpMessage(prog string) string {
	return fmt.Sprintf(`%s [OPTIONS]
OPTIONS
    -s, --size <SIZE>     specifies the data size for generation. Default is 1024.
    -d, --dest <DEST>     specifies the destination file. Default is stdout.
    -h, --help            prints this message.`, filepath.Base(prog))
}

func (opts *options) Destination() (io.WriteCloser, error) {
	if opts.dest == "" {
		return os.Stdout, nil
	}
	return os.Create(opts.dest)
}

func parseOptions(args []string) (*options, error) {
	opts := &options{}
	flags := flag.NewFlagSet("randomizer", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.Int64VarP(&opts.size, "size", "s", 1024, "specifies the data size for generation")
	flags.StringVarP(&opts.dest, "dest", "d", "", "specifies the destination file")
	flags.BoolVarP(&opts.helpFlag, "help", "h", false, "prints this message")
	flags.Parse(args)
	opts.args = flags.Args()
	return opts, nil
}

func printHelpIfNeeeded(name string, opts *options) bool {
	if opts.helpFlag {
		fmt.Println(helpMessage(name))
		return true
	}
	return false
}
