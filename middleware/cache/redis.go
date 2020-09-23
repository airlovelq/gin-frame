package cache

import (
	"github.com/garyburd/redigo/redis"
)

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "192.168.100.103:10203")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}

type redisCache struct {
	Pool *redis.Pool
}

var RedisPool = redisCache{Pool: nil}

func init() {
	RedisPool.Pool = newRedisPool()
}

func (cache *redisCache) Set(key string, value interface{}, options ...CacheOptions) error {
	co := CacheOption{Ex: false, ExTime: 0}
	for _, option := range options {
		option(&co)
	}
	conn := cache.Pool.Get()
	defer conn.Close()
	var err error
	if co.Ex {
		_, err = conn.Do("set", key, value, "EX", co.ExTime)
	} else {
		_, err = conn.Do("set", key, value)
	}
	return err
}

func (cache *redisCache) Get(key string) (interface{}, error) {
	conn := cache.Pool.Get()
	defer conn.Close()
	res, err := conn.Do("get", key)
	return res, err
}

func (cache *redisCache) Delete(key string) error {
	conn := cache.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("del", key)
	return err
}

func (cache *redisCache) String(obj interface{}) (string, error) {
	return redis.String(obj, nil)
}
