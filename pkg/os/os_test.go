package os

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
