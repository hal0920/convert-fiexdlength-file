package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestLoadConfigToml(t *testing.T) {
	wd, _ := os.Getwd()

	// OK
	func() {
		reset := setTestEnv("HOME", wd+"/testdata")
		defer reset()

		l1 := layout{Length: []int{1, 2, 3, 4}}
		l2 := layout{Length: []int{4, 3, 2, 1}}

		expect := tomlConfig{
			Layouts: map[string]layout{
				"example1": l1,
				"example2": l2,
			},
		}

		act, err := loadConfigToml()
		if err != nil {
			t.Fatal("unexpected error occurred error:", err)
		}
		if !reflect.DeepEqual(act, expect) {
			t.Errorf("Output=%v, want %v", act, expect)
		}
	}()

	// NG
	func() {
		reset := setTestEnv("HOME", wd+"/testdata_NG")
		defer reset()

		_, err := loadConfigToml()
		if err == nil {
			t.Fatal("error occurred error:")
		}

		if !strings.Contains(err.Error(), "load config error") {
			t.Errorf("Output=%s", err)
		}
	}()

}
