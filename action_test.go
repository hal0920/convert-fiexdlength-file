package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestDoConvert(t *testing.T) {

	inStream := new(bytes.Buffer)
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)

	stdCLI = &CLI{
		inputStream: inStream,
		outStream:   outStream,
		errStream:   errStream,
	}

	wd, _ := os.Getwd()

	type pattern struct {
		confPath  string
		stdinStr  string
		flags     string
		expectMsg string
	}

	//OK
	func() {
		patterns := []pattern{
			{"/testdata", "", "--layout example1 testdata/test_ok.dat", "1,22,333,4444\n"},
			{"/testdata", "1223334444", "--layout example1", "1,22,333,4444\n"},
		}

		for _, p := range patterns {
			reset := setTestEnv("HOME", wd+p.confPath)
			defer reset()

			inStream.WriteString(p.stdinStr)

			err := doConvert(makeCliContext("doConvertTest", p.flags))
			if err != nil {
				t.Errorf("Output error=%s", err)
			}
			if outStream.String() != p.expectMsg {
				t.Errorf("Outout=%s, but Want=%s", outStream.String(), p.expectMsg)
			}
			inStream.Reset()
			outStream.Reset()
		}
	}()

	// NG config file not exist
	func() {

		patterns := []pattern{
			{"/testdata_NG", "", "", "load config error"},
			{"/testdata", "", "--layout", "layout option is Required"},
			{"/testdata", "", "--layout ngtest", " is not defined filelayout"},
			{"/testdata", "", "--layout example1 ngFile", "no such file or directory"},
			{"/testdata", "", "--layout example1 ngFile1 ngfile2", "Invalid number of argument"},
		}

		for _, p := range patterns {
			reset := setTestEnv("HOME", wd+p.confPath)
			defer reset()

			inStream.WriteString(p.stdinStr)

			err := doConvert(makeCliContext("doConvertTest", p.flags))
			if !strings.Contains(err.Error(), p.expectMsg) {
				t.Errorf("Output error=%s", err)
			}
		}

	}()

}

func TestConvCSV(t *testing.T) {

	type pattern struct {
		inStr  string
		layout []int
		expect string
	}

	// OK pattern
	func() {
		patterns := []pattern{
			{"", []int{1, 2}, ""},
			{"1223334444", []int{1, 2, 3, 4}, "1,22,333,4444"},
			{"abbcccdddd", []int{1, 2, 3, 4}, "a,bb,ccc,dddd"},
			{"1a22bb333ccc", []int{1, 1, 2, 2, 3, 3}, "1,a,22,bb,333,ccc"},
			{"あかかさささ", []int{1, 2, 3}, "あ,かか,さささ"},
			{"りんご　01100円", []int{4, 2, 4}, "りんご　,01,100円"},
		}

		for _, p := range patterns {
			act := convCSV(p.inStr, p.layout)
			if p.expect != act {
				t.Fatal("fatal TestConvCSV OK pattern")
			}
		}
	}()

}
