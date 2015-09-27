package gopath

import (
	"simonwaldherr.de/go/golibs/file"
	"strings"
	"testing"
)

var archlist = map[string]bool{
	"386":   true,
	"amd64": true,
	"arm":   true,
}

var oslist = map[string]bool{
	"darwin":  true,
	"linux":   true,
	"windows": true,
}

func Test_Compiler(t *testing.T) {
	if v := Compiler(); v != "gc" {
		t.Fatalf("Compiler should be \"gc\" but is: %v", v)
	}
}

func Test_GOARCH(t *testing.T) {
	v := GOARCH()
	if _, ok := archlist[v]; !ok {
		t.Fatalf("GOARCH should be \"386\", \"amd64\" or \"arm\" but is: %v", v)
	}
}

func Test_GOOS(t *testing.T) {
	v := GOOS()
	if _, ok := oslist[v]; !ok {
		t.Fatalf("GOOS should be \"darwin\", \"linux\" or \"windows\" but is: %v", v)
	}
}

func Test_GOROOT(t *testing.T) {
	if v := GOROOT(); !file.IsDir(v) {
		t.Fatalf("GOROOT should be a directory but is: %v", v)
	}
}

func Test_GOPATH(t *testing.T) {
	if v := GOPATH(); !file.IsDir(v) {
		t.Fatalf("GOPATH should be a directory but is: %v", v)
	}
}

func Test_Compiled(t *testing.T) {
	if v := Compiled(); v == true {
		t.Fatalf("Compiled should be false but is: %v", v)
	}
}

func Test_Path(t *testing.T) {
	if v := Path(); !strings.HasSuffix(v, "gopath_test.go") {
		t.Fatalf("Path should end with \"gopath_test.go\" but is: %v", v)
	}
}

func Test_Name(t *testing.T) {
	if v := Name(); !strings.HasSuffix(v, "gopath_test.go") {
		t.Fatalf("Name should end with \"gopath_test.go\" but is: %v", v)
	}
}

func Test_Dir(t *testing.T) {
	if v := Dir(); !file.IsDir(v) {
		t.Fatalf("Dir should be a directory but is: %v", v)
	}
}

func Test_GetFileType(t *testing.T) {
	if v := GetFileType(Name()); v != 2 {
		t.Fatalf("GetFileType should be \"2\" but is: %v", v)
	}
}
