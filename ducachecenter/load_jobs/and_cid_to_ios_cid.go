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

// AndCidToIosCid 建立and和ios的cid互查的缓存
func AndCidToIosCid(msg []byte, cache *freecache.Cache) error {
	var tmp []types.IdMap
	err := json.Unmarshal(msg, &tmp)
	if err != nil {
		return err
	}
	for _, v := range tmp {
		key := tools.JointKV([]byte(consts.IosQueryCidByAndCid), []byte(strconv.Itoa(v.AndId)))
		c, err := json.Marshal(v)
		if err != nil {
			return err
		}
		err = local_cache.SetCacheValueNoLock(key, c, cache)
		if err != nil {
			return err
		}
		key = tools.JointKV([]byte(consts.AndQueryCidByIosCid), []byte(strconv.Itoa(v.IosId)))
		err = local_cache.SetCacheValueNoLock(key, c, cache)
		if err != nil {
			return err
		}
	}
	return nil
}
