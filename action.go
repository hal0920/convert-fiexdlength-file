package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func doConvert(c *cli.Context) error {

	cfg, err := loadConfigToml()
	if err != nil {
		return err
	}

	ltype := c.String("layout")
	if ltype == "" {
		return errors.New("layout option is Required")
	}

	if lay, ok := cfg.Layouts[ltype]; ok {
		switch c.Args().Len() {
		case 0:
			scanner := bufio.NewScanner(stdCLI.inputStream)
			for scanner.Scan() {
				fmt.Fprintln(stdCLI.outStream, convCSV(scanner.Text(), lay.Length))
			}
		case 1:
			in, err := os.Open(c.Args().First())
			if err != nil {
				return err
			}
			defer in.Close()
			scanner := bufio.NewScanner(in)
			for scanner.Scan() {
				fmt.Fprintln(stdCLI.outStream, convCSV(scanner.Text(), lay.Length))
			}
		default:
			return errors.New("error : Invalid number of argument")
		}
	} else {
		return errors.New(ltype + " is not defined filelayout")
	}
	return nil
}

func convCSV(str string, layout []int) (ret string) {
	const camma = ","

	if len(str) == 0 {
		return ""
	}

	from := 0
	for _, digit := range layout {
		ret += string([]rune(str)[from:from+digit]) + camma
		from += digit
	}

	return strings.TrimRight(ret, camma)
}
