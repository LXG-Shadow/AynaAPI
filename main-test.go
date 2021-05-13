package main

import (
	"AynaAPI/api/provider"
	"fmt"
)

func main() {
	var pvdr provider.ApiProvider
	fmt.Println(pvdr, pvdr == nil)
}
