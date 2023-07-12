package ducachecenter

import (
	"encoding/json"
	"fmt"
	"git.du.com/cloud/du_component/ducachecenter/consts"
	"git.du.com/cloud/du_component/ducachecenter/load_jobs"
	"git.du.com/cloud/du_component/ducachecenter/local_cache"
	"git.du.com/cloud/du_component/ducachecenter/tools"
)

func LoadInfo() {
	tmpCache := local_cache.GetNewCache()
	if LocalCacheChoose.CidCustomerInfoAnd {
		//更新安卓的cid_customer_info
		tmp, err := tools.HttpPost(Addr+consts.AndQueryCustomerInfoByCidUrl, nil, nil)
		if err != nil {
			fmt.Println("LoadInfo:HttpPost err:", err.Error())
		}
		err = load_jobs.AndCidCustomerInfo(tmp, tmpCache)
		if err != nil {
			fmt.Println("LoadInfo:AndCidCustomerInfo err:", err.Error())
		}
	}
	if LocalCacheChoose.PkgProductInfoAnd {
		//更新安卓的pkg_product_info
		tmp, err := tools.HttpPost(Addr+consts.AndQueryProductInfoByPkgUrl, nil, nil)
		if err != nil {
			fmt.Println("LoadInfo:HttpPost err:", err.Error())
		}
		err = load_jobs.AndPkgProductInfo(tmp, tmpCache)
		if err != nil {
			fmt.Println("LoadInfo:AndPkgProductInfo err:", err.Error())
		}
	}

	if LocalCacheChoose.CidCustomerInfoIos {
		//更新IOS的cid_customer_info
		tmp, err := tools.HttpPost(Addr+consts.IosQueryCustomerInfoByCidUrl, nil, nil)
		if err != nil {
			fmt.Println("LoadInfo:HttpPost err:", err.Error())
		}
		err = load_jobs.IosCidCustomerInfo(tmp, tmpCache)
		if err != nil {
			fmt.Println("LoadInfo:IosCidCustomerInfo err:", err.Error())
		}
	}
	if LocalCacheChoose.PkgProductInfoIos {
		//更新IOS的pkg_product_info
		tmp, err := tools.HttpPost(Addr+consts.IosQueryProductInfoByPkgUrl, nil, nil)
		if err != nil {
			fmt.Println("LoadInfo:HttpPost err:", err.Error())
		}
		err = load_jobs.IosPkgProductInfo(tmp, tmpCache)
		if err != nil {
			fmt.Println("LoadInfo:IosPkgProductInfo err:", err.Error())
		}
	}
	if LocalCacheChoose.AndCidToIosCid || LocalCacheChoose.IosCidToAndCid {
		//更新安卓cid到ios的cid的映射和ios的cid到安卓的cid的映射
		tmp, err := tools.HttpPost(Addr+consts.IosQueryCidByAndCidUrl, nil, nil)
		if err != nil {
			fmt.Println("LoadInfo:HttpPost err:", err.Error())
		}
		err = load_jobs.AndCidToIosCid(tmp, tmpCache)
		if err != nil {
			fmt.Println("LoadInfo:AndCidToIosCid err:", err.Error())
		}
	}
	if LocalCacheChoose.PkgCustomerInfoDirectIos {
		//更新安卓pkg到ios的customer的直接映射
		tmp, err := tools.HttpPost(Addr+consts.IosQueryCustomerInfoByPkgUrl, nil, nil)
		if err != nil {
			fmt.Println("LoadInfo:HttpPost err:", err.Error())
		}
		err = load_jobs.IosPkgCustomerInfo(tmp, tmpCache)
		if err != nil {
			fmt.Println("LoadInfo:IosPkgCustomerInfo err:", err.Error())
		}
	}
	if LocalCacheChoose.AppidProductInfoApplet {
		//更新小程序的appid到产品信息的映射
		tmp, err := tools.HttpPost(Addr+consts.AppletQueryProductInfoByAppidUrl, nil, nil)
		if err != nil {
			fmt.Println("LoadInfo:HttpPost err:", err.Error())
		}
		err = load_jobs.AppidProductInfoApplet(tmp, tmpCache)
	}
	if LocalCacheChoose.AppidProductInfoH5 {
		//更新H5的appid到产品信息的映射
		tmp, err := tools.HttpPost(Addr+consts.H5QueryProductInfoByAppIdUrl, nil, nil)
		if err != nil {
			fmt.Println("LoadInfo:HttpPost err:", err.Error())
		}
		err = load_jobs.AppidProductInfoH5(tmp, tmpCache)
	}

	if LocalCacheChoose.CidActivateIdfaInfoFront {
		//更新前端客户产品激活信息的映射
		body, _ := json.Marshal(CidActivateIdfaInfoFrontReq{ProductValue: "product_idx_idfa_padding"})
		tmp, err := tools.HttpPost(Addr+consts.FrontQueryCustomerInfoByCidUrl, nil, body)
		if err != nil {
			fmt.Println("LoadInfo:HttpPost err:", err.Error())
		} else {
			err = load_jobs.CidActivateInfoFront("product_idx_idfa_padding", tmp, tmpCache)
		}
	}

	local_cache.UpdateCache(tmpCache)
}

type CidActivateIdfaInfoFrontReq struct {
	ProductValue string `json:"product_value"`
}
