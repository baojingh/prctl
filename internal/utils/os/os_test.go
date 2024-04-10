package os

import (
	"fmt"
	"testing"
)

func TestGetOs(t *testing.T) {
	release, version, err := GetOs()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(release, version)
}
