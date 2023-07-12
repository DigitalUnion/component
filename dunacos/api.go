package dunacos

import (
	"os/exec"
	"runtime"
)

func GetEnv(key string) (string, error) {
	if runtime.GOOS == "windows" {
		return GetWindowsEnv(key)
	} else {
		val := GetLinuxEnv(key)
		if val == "" {
			return execLinux("echo $" + key)
		} else {
			return val, nil
		}
	}
}
func execLinux(c string) (string, error) {
	cmd := exec.Command(c)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
