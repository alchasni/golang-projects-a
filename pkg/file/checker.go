package file

import "os"

func CheckExistence(fPath string) (exist bool) {
	info, err := os.Stat(fPath)
	if err != nil || info.IsDir() {
		return false
	}

	return true
}
