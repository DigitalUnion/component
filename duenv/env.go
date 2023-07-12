package duenv

import "runtime"

func GetEnv(key string) (string, error) {
	if runtime.GOOS == "windows" {
		return GetWindowsEnv(key)
	} else {
		return GetLinuxEnv(key), nil
	}
}
