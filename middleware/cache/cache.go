package cache

import (
	"scoremanager/utils"
)

type CacheOp interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, options ...CacheOptions) error
	Delete(key string) error
	String(obj interface{}) (string, error)
}

type CacheOption struct {
	Ex     bool
	ExTime int
}
type CacheOptions func(*CacheOption)

func WithEx(seconds int) CacheOptions {
	return func(co *CacheOption) {
		co.Ex = true
		co.ExTime = seconds
	}
}

func NewCacheOp() CacheOp {
	cacheType := utils.GetEnvDefault("CACHE_TYPE", "redis")
	var cacheOp CacheOp
	switch cacheType {
	case "redis":
		cacheOp = &RedisPool
	}
	return cacheOp
}
