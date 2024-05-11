package prsys

import (
	"fmt"
	"testing"
)

func TestGetOs(t *testing.T) {
	info, err := GetOs()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(info.Distribution, info.Version)
}

func TestCurrentUserPath(t *testing.T) {
	curr := CurrentUserPath()
	fmt.Println(curr)
}

func TestGetGoInfo(t *testing.T) {
	res := GetGoInfo("GOMODCACHE2")
	fmt.Println(res)
}
