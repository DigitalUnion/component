package duratelimit

type Resource struct {
	Name          string `json:"name" yaml:"name"`                       // 资源名称
	RatePerSecond int    `json:"rate_per_second" yaml:"rate_per_second"` // 每秒限制的请求个数（秒，分钟，小时三选一，多选默认秒）
	RatePerMinute int    `json:"rate_per_minute" yaml:"rate_per_minute"` // 每分钟限制的请求个数
	RatePerHour   int    `json:"rate_per_hour" yaml:"rate_per_hour"`     // 每小时限制的请求个数
	Allow         int    `json:"allow" yaml:"allow"`                     // 每次去 redis 拿到的令牌个数（大于1切不得小于 qps 的 %0.5）
}

func (r Resource) verify() (int, error) {
	rateTp := perSecond
	if len(r.Name) == 0 {
		return rateTp, nameEmpty
	}
	if r.RatePerSecond == 0 && r.RatePerMinute == 0 && r.RatePerHour == 0 {
		return rateTp, rateEmpty
	}
	var rate int
	if r.RatePerHour != 0 {
		rateTp = perHour
		rate = r.RatePerHour
	}
	if r.RatePerMinute != 0 {
		rateTp = perMinute
		rate = r.RatePerMinute
	}
	if r.RatePerSecond != 0 {
		rateTp = perSecond
		rate = r.RatePerSecond
	}
	if r.Allow == 0 {
		return rateTp, allowEmpty
	}
	if r.Allow*200 < rate && rate > 200 {
		return rateTp, allowExceed
	}
	return rateTp, nil
}
