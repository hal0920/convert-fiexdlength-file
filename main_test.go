package main

import (
	"reflect"
	"testing"
)

func TestReadLengthListByToml(t *testing.T) {

	type pattern struct {
		tomlPath string
		ltype    string
		expect   []int
	}

	// OK pattern
	func() {
		patterns := []pattern{
			{"testdata/config.toml", "example1", []int{1, 2, 3, 4}},
		}

		for _, p := range patterns {

			act, err := readLengthListByToml(p.tomlPath, p.ltype)
			if err != nil {
				t.Fatal("toml is not exists")
			}

			if !reflect.DeepEqual(act, p.expect) {
				t.Fatal("fatal success test")
			}
		}

	}()

	// NG pattern
	func() {
		patterns := []pattern{
			{"testdata/fatal.toml", "example1", []int{1, 2, 3, 4}},
			{"testdata/config.toml", "exampleF", []int{1, 2, 3, 4}},
			{"testdata/fatal.toml", "exampleF", []int{1, 2, 3, 4}},
		}

		for _, p := range patterns {

			_, err := readLengthListByToml(p.tomlPath, p.ltype)
			if err == nil {
				t.Fatal("fatal error test")
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
			{"abbcccdddd", []int{1, 2, 3, 4}, "a,bb,ccc,dddd"},
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
