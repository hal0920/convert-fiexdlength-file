package main

import (
	"flag"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func setTestEnv(key, val string) func() {
	preVal := os.Getenv(key)
	os.Setenv(key, val)
	return func() {
		os.Setenv(key, preVal)
	}
}

func makeCliContext(name, argv string) *cli.Context {
	app := newApp()

	set := flag.NewFlagSet(name, flag.ContinueOnError)
	for _, f := range app.Flags {
		f.Apply(set)
	}
	set.Parse(strings.Split(argv, " "))

	return cli.NewContext(app, set, nil)
}
