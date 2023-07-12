package ducachecenter

import (
	"fmt"
	"testing"
)

func TestAndGetCustomerInfoByCid(t *testing.T) {
	choose := &CacheChoose{
		CidCustomerInfoAnd:       true,
		CidCustomerInfoIos:       true,
		PkgProductInfoAnd:        true,
		PkgProductInfoIos:        true,
		PkgCustomerInfoAnd:       true,
		PkgCustomerInfoIos:       true,
		AndCidToIosCid:           true,
		AppidProductInfoH5:       true,
		AppidProductInfoApplet:   true,
		PkgCustomerInfoDirectIos: true,
	}
	//全部模块的缓存都需要Init函数的第一个参数填nil
	Init(choose, "172.17.0.130:8185")
	res1 := AndGetCustomerInfoByCid(238)
	fmt.Println(res1)
	res2 := IosGetCustomerInfoByCid(6)
	fmt.Println(res2)
	res3 := AndGetProductInfoByPkg("com.yek.lafaso")
	fmt.Println(res3)
	res4 := IosGetProductInfoByPkg("com.moji.MojiWeather")
	fmt.Println(res4)
	res5 := AndGetCustomerInfoByPkg("com.yek.lafaso")
	fmt.Println(res5)
	res6 := IosGetCustomerInfoByPkg("com.moji.MojiWeather")
	fmt.Println(res6)
	res7 := IosGetIosCidByAndCid(238)
	fmt.Println(res7)
	res8 := AppletGetInfoByAppid("wx21c7506e98a2fe75")
	fmt.Println(res8)
	res9 := H5GetInfoByAppid("h5U6fo1h")
	fmt.Println(res9)
	res10 := IosGetCustomerInfoByPkgDirect("com.moji.MojiWeather")
	fmt.Println(res10)
}
