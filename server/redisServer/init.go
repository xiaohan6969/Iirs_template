package redisServer

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
	"../../config"
)

var (
	ADDRESS      = config.Config.Get("redis.ip").(string) + config.Config.Get("redis.port").(string)
	PASSWORD      = config.Config.Get("redis.password").(string)
	pool   *redis.Pool
	RedisC redis.Conn
)

func init() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ADDRESS, redis.DialPassword(PASSWORD))
			if nil != err {
				return nil, err
			}
			return c, nil
		},
	}
	RedisC = pool.Get()
}

//Find
func RInputKeyGetValue(redis_key string, library_number int) (interface{}, error) {
	var (
		R = RedisC
	)
	defer func() {
		if err := R.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	err := R.Send("select", library_number)
	if err != nil {
		return nil, err
	}
	return redis.String(R.Do("get", redis_key))
}

//insert
func RInsertKeyAndValue(redis_key, value string, library_number int) interface{} {
	var (
		R   = RedisC
		err error
	)
	defer func() {
		if err := R.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	err = R.Send("select", library_number)
	if err != nil {
		return err
	}
	_, err = R.Do("set", redis_key, value)
	return err
}

// 检车 key 是否存在
func RCheckKey(redis_key string, library_number int) bool {
	var (
		R   = RedisC
		bo  bool
		err error
	)
	defer func() {
		if err := R.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	_, err = R.Do("select", library_number)
	if err != nil {
		return false
	}
	bo, err = redis.Bool(R.Do("EXISTS", redis_key))
	if err != nil {
		return false
	}
	return bo
}
