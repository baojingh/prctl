package prsys

import (
	"bufio"
	"os"
	"os/user"
	"strings"
)

/*
$ cat /etc/os-release
PRETTY_NAME="Ubuntu 22.04 LTS"
NAME="Ubuntu"
VERSION_ID="22.04"
VERSION="22.04 (Jammy Jellyfish)"
VERSION_CODENAME=jammy
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=jammy
*/

type OSInfo struct {
	Distribution string
	Version      string
}

func GetOs() (OSInfo, error) {
	osInfo := OSInfo{}

	// default path in linux
	path := "/etc/os-release"
	file, err := os.Open(path)
	if err != nil {
		return osInfo, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "=")
		if len(fields) >= 2 && fields[0] == "ID" {
			dis := strings.TrimSpace(fields[1])
			osInfo.Distribution = dis
		}
		if len(fields) >= 2 && fields[0] == "VERSION_CODENAME" {
			version := strings.TrimSpace(fields[1])
			osInfo.Version = version
		}
	}
	if scanner.Err(); err != nil {
		return osInfo, err
	}
	return osInfo, nil
}

func CurrentUserPath() string {
	currUser, _ := user.Current()
	userPath := currUser.HomeDir
	return userPath
}

func GetGoInfo(key string) string {
	val := os.Getenv(key)
	return val
}
