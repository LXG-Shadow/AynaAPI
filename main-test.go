package main

import (
	"AynaAPI/utils/vfile"
	"fmt"
)

func main() {
	fmt.Println(vfile.CalcFileMD5("main.go"))
	fmt.Println(vfile.CalcFileMD5("README.md"))
}
