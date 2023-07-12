/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/11/17 11:59
 */

package dudecoder

import (
	"fmt"
	"os"
	"testing"
)

// addr := "http://172.17.129.33:10000/api/field/v1/list"
// idaa: `{"page": 1,"limit": 99999,"channel_tp": 2,"message_type_id": 5}`
// daa: `{"page": 1,"limit": 99999,"channel_tp": 1,"message_type_id": 2}`
func TestGetFieldListDAA(t *testing.T) {
	bs, err := GetFieldListFromApi("http://172.17.129.33:10000/api/field/v1/list", 1, 2)
	fmt.Println(err)
	fmt.Println(string(bs))
	os.WriteFile("./fieldDAA.json", bs, os.FileMode(0644))
}
func TestGetFieldListDNA(t *testing.T) {
	bs, err := GetFieldListFromApi("http://172.17.129.33:10000/api/field/v1/list", 1, 1)
	fmt.Println(err)
	fmt.Println(string(bs))
	os.WriteFile("./fieldDNA.json", bs, os.FileMode(0644))
}
func TestGetFieldListIDAA(t *testing.T) {
	bs, err := GetFieldListFromApi("http://172.17.129.33:10000/api/field/v1/list", 2, 5)
	fmt.Println(err)
	fmt.Println(string(bs))
	os.WriteFile("./fieldIDAA.json", bs, os.FileMode(0644))
}
func TestGetFieldListIDNA(t *testing.T) {
	bs, err := GetFieldListFromApi("http://172.17.129.33:10000/api/field/v1/list", 2, 4)
	fmt.Println(err)
	fmt.Println(string(bs))
	os.WriteFile("./fieldIDNA.json", bs, os.FileMode(0644))
}
func TestGetFieldListDAI(t *testing.T) {
	bs, err := GetFieldListFromApi("http://172.17.129.33:10000/api/field/v1/list", 1, 3)
	fmt.Println(err)
	os.WriteFile("./fieldDAI.json", bs, os.FileMode(0644))
}

func TestGetFieldListApplet(t *testing.T) {
	bs, err := GetFieldListFromApi("http://172.17.129.33:10000/api/field/v1/list", 3, 6)
	fmt.Println(err)
	os.WriteFile("./fieldAppletDaa.json", bs, os.FileMode(0644))
}
