package main

import (
	"fmt"
	"git.du.com/cloud/du_component/dugoroutingpool"
	"strconv"
)

func main() {
	//同步池
	syncPool := dugoroutingpool.NewSyncPool(10, syncHandle)
	defer syncPool.Close()
	//异步池
	asyncPool := dugoroutingpool.NewAsyncPool(10, func(line interface{}) {
		res := syncPool.Process(line)
		fmt.Println("input:", line, "output:", res.(string))
	})
	defer asyncPool.Close()
	for i := 0; i < 100; i++ {
		asyncPool.SendLine(strconv.Itoa(i))
	}
	asyncPool.Wait()

}

func syncHandle(data interface{}) interface{} {
	return data.(string) + "_hello"
}
