package duratelimit

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
	"sync"
	"time"
)

const (
	perSecond = iota
	perMinute
	perHour
)

type DuRateLimit struct {
	limiter     *redis_rate.Limiter
	limit       redis_rate.Limit
	name        string
	allowN      int
	accessN     int
	lastReqTime time.Time
	retryAfter  time.Duration
	mu          *sync.Mutex
}

func (a *DuRateLimit) Allow() bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.accessN > 0 {
		a.accessN -= 1
		return true
	}
	n := time.Now()
	if n.Sub(a.lastReqTime) > a.retryAfter {
		res, err := a.limiter.AllowN(context.Background(), a.name, a.limit, a.allowN)
		if err != nil {
			fmt.Println(err.Error())
		}
		a.lastReqTime = n
		if res.RetryAfter != -1 {
			a.retryAfter = res.RetryAfter
		}
		if res.Allowed > 0 {
			a.accessN += res.Allowed - 1
			return true
		} else {
			return false
		}
	}
	return false
}

func NewDuRateLimit(client *redis.Client, cfg Resource) *DuRateLimit {
	limiter := redis_rate.NewLimiter(client)
	rateTp, err := cfg.verify()
	if err != nil {
		panic(err)
	}
	rateLimit := &DuRateLimit{
		limiter:     limiter,
		name:        cfg.Name,
		allowN:      cfg.Allow,
		accessN:     0,
		lastReqTime: time.Now(),
		retryAfter:  0,
		mu:          new(sync.Mutex),
	}
	switch rateTp {
	case perSecond:
		rateLimit.limit = redis_rate.PerSecond(cfg.RatePerSecond)
	case perMinute:
		rateLimit.limit = redis_rate.PerMinute(cfg.RatePerMinute)
	case perHour:
		rateLimit.limit = redis_rate.PerHour(cfg.RatePerSecond)
	}
	return rateLimit
}
