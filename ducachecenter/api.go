package ducachecenter

import (
	"encoding/json"
	"git.du.com/cloud/du_component/ducachecenter/consts"
	"git.du.com/cloud/du_component/ducachecenter/local_cache"
	"git.du.com/cloud/du_component/ducachecenter/tools"
	"git.du.com/cloud/du_component/ducachecenter/types"
	"github.com/robfig/cron"
	"strconv"
)

var LocalCacheChoose *CacheChoose
var Addr string
var crontab = cron.New()

type CacheChoose struct {
	//安卓
	CidCustomerInfoAnd bool // 安卓的cid查询用户信息
	PkgProductInfoAnd  bool // 安卓的pkgName查询产品信息
	PkgCustomerInfoAnd bool // 安卓的pkgName查询用户信息
	AndCidToIosCid     bool // 安卓的cid查询IOS的cid
	IosCidToAndCid     bool // IOS的cid查询安卓的cid
	//IOS
	CidCustomerInfoIos       bool // IOS的cid查询用户信息
	PkgProductInfoIos        bool // IOS的pkgName查询产品信息
	PkgCustomerInfoIos       bool // IOS的pkgName查询用户信息
	PkgCustomerInfoDirectIos bool //IOS的pkgName查询用户信息(直接查询，延迟低，推荐)
	//小程序
	AppidProductInfoApplet bool // 小程序的appid查询产品信息
	//H5
	AppidProductInfoH5 bool // H5的appid查询产品信息
	//Front
	CidActivateIdfaInfoFront bool // 前端的cid查询idfa补全产品激活信息
}

func Init(choose *CacheChoose, addr string) {
	if choose == nil {
		choose = &CacheChoose{
			CidCustomerInfoAnd:       true,
			PkgProductInfoAnd:        true,
			PkgCustomerInfoAnd:       true,
			AndCidToIosCid:           true,
			CidCustomerInfoIos:       true,
			PkgProductInfoIos:        true,
			PkgCustomerInfoIos:       true,
			PkgCustomerInfoDirectIos: true,
			AppidProductInfoApplet:   true,
			AppidProductInfoH5:       true,
			CidActivateIdfaInfoFront: true,
			IosCidToAndCid:           true,
		}
	}
	LocalCacheChoose = choose
	if LocalCacheChoose.PkgCustomerInfoAnd {
		LocalCacheChoose.CidCustomerInfoAnd = true
		LocalCacheChoose.PkgProductInfoAnd = true
	}
	if LocalCacheChoose.PkgCustomerInfoIos {
		LocalCacheChoose.CidCustomerInfoIos = true
		LocalCacheChoose.PkgProductInfoIos = true
	}
	if LocalCacheChoose.IosCidToAndCid || LocalCacheChoose.AndCidToIosCid {
		LocalCacheChoose.IosCidToAndCid = true
		LocalCacheChoose.AndCidToIosCid = true
	}
	Addr = addr
	LoadInfo()
	crontab.AddFunc("0 */10 * * * *", LoadInfo)
	crontab.Start()
}

func AndGetCustomerInfoByCid(cid int) *types.CustomerAnd {
	key := tools.JointKV([]byte(consts.AndQueryCustomerInfoByCid), []byte(strconv.Itoa(cid)))
	value := local_cache.GetLocalCacheValue(key)
	if len(value) == 0 {
		return nil
	}
	var tmp types.CustomerAnd
	err := json.Unmarshal(value, &tmp)
	if err != nil {
		return nil
	}
	return &tmp
}

func IosGetCustomerInfoByCid(cid int) *types.CustomerIos {
	key := tools.JointKV([]byte(consts.IosQueryCustomerInfoByCid), []byte(strconv.Itoa(cid)))
	value := local_cache.GetLocalCacheValue(key)
	if len(value) == 0 {
		return nil
	}
	var tmp types.CustomerIos
	err := json.Unmarshal(value, &tmp)
	if err != nil {
		return nil
	}
	return &tmp
}

func AndGetProductInfoByPkg(pkg string) *types.ProductAnd {
	key := tools.JointKV([]byte(consts.AndQueryProductInfoByPkg), []byte(pkg))
	value := local_cache.GetLocalCacheValue(key)
	if len(value) == 0 {
		return nil
	}
	var tmp types.ProductAnd
	err := json.Unmarshal(value, &tmp)
	if err != nil {
		return nil
	}
	return &tmp
}

func IosGetProductInfoByPkg(pkg string) *types.ProductIos {
	key := tools.JointKV([]byte(consts.IosQueryProductInfoByPkg), []byte(pkg))
	value := local_cache.GetLocalCacheValue(key)
	if len(value) == 0 {
		return nil
	}
	var tmp types.ProductIos
	err := json.Unmarshal(value, &tmp)
	if err != nil {
		return nil
	}
	return &tmp
}

func AndGetCustomerInfoByPkg(pkg string) *types.CustomerAnd {
	product := AndGetProductInfoByPkg(pkg)
	if product == nil {
		return nil
	}
	return AndGetCustomerInfoByCid(product.Cid)
}

func IosGetCustomerInfoByPkg(pkg string) *types.CustomerIos {
	product := IosGetProductInfoByPkg(pkg)
	if product == nil {
		return nil
	}
	return IosGetCustomerInfoByCid(product.Cid)
}

func IosGetIosCidByAndCid(cid int) *types.IdMap {
	key := tools.JointKV([]byte(consts.IosQueryCidByAndCid), []byte(strconv.Itoa(cid)))
	value := local_cache.GetLocalCacheValue(key)
	if len(value) == 0 {
		return nil
	}
	var tmp types.IdMap
	err := json.Unmarshal(value, &tmp)
	if err != nil {
		return nil
	}
	return &tmp
}

func AndGetAndCidByIosCid(cid int) *types.IdMap {
	key := tools.JointKV([]byte(consts.AndQueryCidByIosCid), []byte(strconv.Itoa(cid)))
	value := local_cache.GetLocalCacheValue(key)
	if len(value) == 0 {
		return nil
	}
	var tmp types.IdMap
	err := json.Unmarshal(value, &tmp)
	if err != nil {
		return nil
	}
	return &tmp
}

func IosGetCustomerInfoByPkgDirect(pkg string) *types.CustomerInfoByPkg {
	key := tools.JointKV([]byte(consts.IosQueryCustomerInfoByPkg), []byte(pkg))
	value := local_cache.GetLocalCacheValue(key)
	if len(value) == 0 {
		return nil
	}
	var tmp types.CustomerInfoByPkg
	err := json.Unmarshal(value, &tmp)
	if err != nil {
		return nil
	}
	return &tmp

}

func AppletGetInfoByAppid(appid string) *types.ProductApplet {
	key := tools.JointKV([]byte(consts.AppletQueryProductInfoByAppid), []byte(appid))
	value := local_cache.GetLocalCacheValue(key)
	if len(value) == 0 {
		return nil
	}
	var tmp types.ProductApplet
	err := json.Unmarshal(value, &tmp)
	if err != nil {
		return nil
	}
	return &tmp
}

func H5GetInfoByAppid(appid string) *types.ProductH5 {
	key := tools.JointKV([]byte(consts.H5QueryProductInfoByAppId), []byte(appid))
	value := local_cache.GetLocalCacheValue(key)
	if len(value) == 0 {
		return nil
	}
	var tmp types.ProductH5
	err := json.Unmarshal(value, &tmp)
	if err != nil {
		return nil
	}
	return &tmp
}

func FrontGetActivateInfoByCid(cid int) *types.FrontCustomerActivateCacheCenter {
	key := tools.JointKV([]byte(consts.FrontQueryCustomerInfoByCid+"&"+"product_idx_idfa_padding"), []byte(strconv.Itoa(cid)))
	value := local_cache.GetLocalCacheValue(key)
	if len(value) == 0 {
		return nil
	}
	var tmp types.FrontCustomerActivateCacheCenter
	err := json.Unmarshal(value, &tmp)
	if err != nil {
		return nil
	}
	return &tmp
}
