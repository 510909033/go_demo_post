package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	//	dir := `/data/wangbaotian/git`

	os.Chdir("e:/a")

	matchList, err := filepath.Glob("*")
	_ = err

	fmt.Println(matchList)

}
