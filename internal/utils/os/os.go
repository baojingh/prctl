package os

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

func GetOs() (string, string, error) {
	return "ubuntu", "jammy", nil
	// // default path in linux
	// path := "/etc/os-release"
	// file, err := os.Open(path)
	// if err != nil {
	// 	return "", err
	// }

	// osMap := make(map[string]string)

	// defer file.Close()
	// scanner := bufio.NewScanner(file)
	//
	//	for scanner.Scan() {
	//		line := scanner.Text()
	//		fields := strings.Split(line, "=")
	//		osMap[]
	//		if len(fields) >= 2 && fields[0] == "ID" {
	//			dis := strings.TrimSpace(fields[1])
	//			return dis, nil
	//		}
	//	}
	//
	//	if scanner.Err(); err != nil {
	//		return "", err
	//	}
	//
	// return "", fmt.Errorf("Cannot find Linux distribution")
}
