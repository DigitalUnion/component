package dudownload

import (
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func Init(cfg *Config) {
	if cfg == nil {
		log.Println("No files to download")
		return
	}
	l := len(cfg.Files)
	log.Printf("Found [%d] files to download\n", l)
	for k, e := range cfg.Files {
		log.Printf("Download [%s] from [%s],save to [%s]\n", k, e.Url, e.Target)
		err := Download(e.Url, e.Target)
		if err != nil {
			log.Printf("Download file: [%s] from [%s] error:%s\n", k, e.Url, err.Error())
		}
	}
}
func Download(url string, output string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return ioutil.WriteFile(output, data, 0644)
}
func ReadLines(url string, lineHandler func([]byte)) (int64, int64, error) {
	log.Printf("Read File: [%s]\n", url)
	if lineHandler == nil {
		return 0, 0, errors.New("lineHandler not defined!")
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("File: [%s] not found\n", url)
		return 0, 0, errors.New(resp.Status)
	}

	log.Printf("File Size: [%d B]\n", resp.ContentLength)
	var c int64 = 0
	br := bufio.NewReader(resp.Body)
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			break
		}
		lineHandler(line)
		c++
	}
	log.Printf("File [%s] handle finish! Lines: [%d], Size: [%d B]\n", url, c, resp.ContentLength)
	return resp.ContentLength, c, nil
}
