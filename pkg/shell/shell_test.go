package shell

import (
	"fmt"
	"testing"
)

func TestShellCmd(t *testing.T) {
	out, err := DoShellCmd("ls", "-l", "-a")
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}
	fmt.Printf("success %s", string(out))
}

func TestIsCmdExist(t *testing.T) {
	res := IsCmdExist("ls")
	fmt.Println(res)
}
