package anime

import (
	"fmt"
	"testing"
)

func TestGetAnimeProviderList(t *testing.T) {
	fmt.Println(GetAnimeProviderList())
}

func TestGetAnimeProvider(t *testing.T) {
	fmt.Println(GetAnimeProvider("a"))
}
