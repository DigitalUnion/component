package load_jobs

import (
	"encoding/json"
	"git.du.com/cloud/du_component/ducachecenter/consts"
	"git.du.com/cloud/du_component/ducachecenter/local_cache"
	"git.du.com/cloud/du_component/ducachecenter/tools"
	"git.du.com/cloud/du_component/ducachecenter/types"
	"github.com/coocood/freecache"
	"strconv"
)

func IosCidCustomerInfo(msg []byte, cache *freecache.Cache) error {
	var tmp []types.CustomerIos
	err := json.Unmarshal(msg, &tmp)
	if err != nil {
		return err
	}
	for _, v := range tmp {
		key := tools.JointKV([]byte(consts.IosQueryCustomerInfoByCid), []byte(strconv.Itoa(v.Id)))
		c, err := json.Marshal(v)
		if err != nil {
			return err
		}
		err = local_cache.SetCacheValueNoLock(key, c, cache)
		if err != nil {
			return err
		}
	}
	return nil
}
