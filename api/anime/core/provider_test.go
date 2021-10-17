package core

import (
	"AynaAPI/api/core"
	"fmt"
	"testing"
)

type name struct {
}

func (n *name) Search(keyword string) core.ApiResource {
	fmt.Println(123)
	return core.ApiResource{}
}

func TestProviderMeta(t *testing.T) {
	p := ProviderMeta{Name: "aa", Url: "https://大家好我超市你"}
	s := p.Dump()
	fmt.Printf("%s", s)
	pnew := ProviderMeta{}
	err := pnew.Load(s)
	fmt.Println(err)
	fmt.Println(pnew)
}
