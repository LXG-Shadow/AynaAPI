package models

import (
	"testing"
)

func TestAuthUser(t *testing.T) {
	var tests = []struct {
		un   string
		pwd  string
		want bool
	}{
		{"admin", "ayapiadmin", true},
		{"a", "b", false},
	}
	for _, test := range tests {
		if ok, _ := AuthUser(test.un, test.pwd); ok != test.want {
			t.Errorf("AuthUser %s/%s -> %t should be %t", test.un, test.pwd, ok, test.want)
		}
	}
}

func TestIsUserNameExists(t *testing.T) {
	var tests = []struct {
		un   string
		want bool
	}{
		{"admin", true},
		{"a", false},
	}
	for _, test := range tests {
		if ok, _ := IsUserNameExists(test.un); ok != test.want {
			t.Errorf("IsUserNameExists %s -> %t should be %t", test.un, ok, test.want)
		}
	}
}

func TestRegisterUser(t *testing.T) {
	user, err := RegisterUser("testa", "testb")
	t.Log(user, err)
}

func TestGetUser(t *testing.T) {
	t.Log(GetUser("admin"))
	t.Log(GetUser("a"))
	t.Log(GetUser("admin1"))
}

func TestPermission(t *testing.T) {
	u, _ := GetUser("admin")
	p, _ := GetPermissionByUser(u)
	t.Log(p)
}
