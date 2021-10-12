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

func TestProvider(t *testing.T) {

}
