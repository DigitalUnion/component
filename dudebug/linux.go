package dudebug

import "os"

func GetLinuxEnv(key string) string {
	return os.Getenv(key)
}
