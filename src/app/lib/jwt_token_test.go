package lib

import (
	"testing"
)

func TestGetUserToken(t *testing.T) {
	tokenString, err := GetUserToken(1, 10000)
	if err != nil {
		t.Fatal(err)
	}

	token, err := ParseUserToken(tokenString)
	if err != nil {
		t.Fatal(err)
	}

	println("%v", token)
}
