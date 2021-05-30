package main

import (
	"AynaAPI/api/provider/susudm"
)

func main() {
	//fmt.Println(vfile.CalcFileMD5("main.go"))
	//fmt.Println(vfile.CalcFileMD5("README.md"))
	//fmt.Println(susudm.GetInfo("63251","acg","11"))
	//fmt.Println(susudm.GetPlayData("63251","1"))
	susudm.Search("刀剑神域", 0)
}
