package shell

import "os/exec"

func DoShellCmd(cmd string, params ...string) (string, error) {
	command := exec.Command(cmd, params...)
	out, err := command.CombinedOutput()
	return string(out), err
}

func IsCmdExist(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
