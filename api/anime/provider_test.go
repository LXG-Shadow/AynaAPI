package anime

import (
	"AynaAPI/api/core"
	"fmt"
	"testing"
)

func TestGetAnimeProviderList(t *testing.T) {
	fmt.Println(Providers.GetProviderList())
}

func TestGetAnimeProvider(t *testing.T) {
	fmt.Println(Providers.GetProvider("a"))
	fmt.Println(Providers.GetAnimeProvider("a"))
}

type name struct {
}

func (n *name) Search(keyword string) core.ApiResource {
	fmt.Println(123)
	return core.ApiResource{}
}

func TestProviderMeta(t *testing.T) {
	p := core.ProviderMeta{Name: "aa", Url: "https://大家好我超市你"}
	s := p.Dump()
	fmt.Printf("%s", s)
	pnew := core.ProviderMeta{}
	err := pnew.Load(s)
	fmt.Println(err)
	fmt.Println(pnew)
}
