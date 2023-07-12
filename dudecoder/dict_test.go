package dudecoder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestDAA(t *testing.T) {
	bs, err := ioutil.ReadFile("./fieldDAA.json")
	parseFields(bs, DAA)
	log.Println(len(infoMap))
	originData, err := ioutil.ReadFile("./daa1.txt")
	res, err := DecodeDAA(originData)
	log.Println(err)
	j, _ := json.Marshal(res)
	fmt.Println(string(j))
}
func TestDNA(t *testing.T) {
	bs, err := ioutil.ReadFile("./fieldDNA.json")
	parseFields(bs, DNA)
	log.Println(len(infoMap))
	originData, err := ioutil.ReadFile("./dna.txt")
	res, err := DecodeDNA(originData)
	log.Println(err)
	j, _ := json.Marshal(res)
	fmt.Println(string(j))
}
func TestIDNA(t *testing.T) {
	bs, err := ioutil.ReadFile("./fieldIDNA.json")
	parseFields(bs, IDNA)
	log.Println(len(infoMap))
	originData, err := ioutil.ReadFile("./idna.txt")
	res, err := DecodeIDNA(originData)
	log.Println(err)
	j, _ := json.Marshal(res)
	fmt.Println(string(j))
}

func TestApplet(t *testing.T) {
	bs, err := ioutil.ReadFile("./fieldAppletDaa.json")
	parseFields(bs, Applet)
	log.Println(len(infoMap))
	originData, err := ioutil.ReadFile("./applet.txt")
	res, err := DecodeApplet(originData)
	log.Println(err)
	j, _ := json.Marshal(res)
	fmt.Println(string(j))
}
