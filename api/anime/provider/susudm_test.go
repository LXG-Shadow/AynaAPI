package provider

import "testing"

func TestSusuDm_Search(t *testing.T) {
	provider := _newSusuDm()
	provider.Search("刀剑神域")
}
