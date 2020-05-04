package main

import (
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const version = "0.1.0"

// CLI ...
type CLI struct {
	inputStream io.Reader
	outStream   io.Writer
	errStream   io.Writer
}

var stdCLI *CLI

func main() {

	stdCLI = &CLI{
		inputStream: os.Stdin,
		outStream:   os.Stdout,
		errStream:   os.Stderr,
	}

	err := newApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func newApp() *cli.App {

	app := cli.NewApp()

	app.Name = "cvfv"
	app.Usage = "Convert Fixed-length file into variable-length file"
	app.Version = version
	app.Action = doConvert
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "layout",
			Aliases: []string{"l"},
			Usage:   "Fixed-length file layout `FILELAYOUT`",
		},
	}

	return app
}
