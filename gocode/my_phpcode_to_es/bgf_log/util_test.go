package bgf_log

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestStringSliceContains(t *testing.T) {
	find := "abc"
	s := []string{
		"zzz",
		find,
		"bcd",
		"def",
	}
	expect := true
	if get := stringSliceContains(s, find, nil); get != expect {
		t.Errorf("expect val is %t, but get %v", expect, get)
	}
}

func TestFileExists(t *testing.T) {
	file, err := ioutil.TempFile("", "tempfile")
	if err != nil {
		t.Fatalf("Failed to create temp file, err: %v", err)
	}
	defer os.Remove(file.Name())
	t.Logf("temp file name: %v", file.Name())
	isExist := fileExists(file.Name())
	if !isExist {
		t.Errorf("file %s exists, but got not.", file.Name())
	}
}

func TestStack(t *testing.T) {
	b := stack(1)
	if !strings.Contains(string(b), "util_test.go") {
		t.Errorf("Failed to stack.")
	}
}

func TestGetDeployRootPath(t *testing.T) {
	p := getDeployRootPath(true)
	if !strings.Contains(p, "go-build") {
		t.Error("Failed to getDeployRootPath")
	}
}
