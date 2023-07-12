package dulogv2

import (
	"fmt"
	"testing"
)

func TestValidJson(t *testing.T) {
	m := `{"a":"b "}`
	if validJson(m) {
		fmt.Println("validJson(m) is true")
	} else {
		fmt.Println("validJson(m) is false")
	}
}
