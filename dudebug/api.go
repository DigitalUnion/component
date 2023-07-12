package dudebug

import (
	"log"
	"strings"
)

var DebugMode = true

func init() {
	env, err := GetEnv("DU_ENV")
	if err != nil {
		log.Println("get DU_ENV error:", err.Error())
		return
	}
	if strings.Contains(env, "PROD") {
		DebugMode = false
		log.Println("PROD mode,DebugMode disabled")
	} else {
		log.Println("TEST mode,DebugMode enabled")
	}
}
