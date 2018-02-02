package kvs

import (
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

// NewRedisPool inits the redis connection and returns a Pool
func NewRedisPool() *redis.Pool {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = ":6379"
	}
	defer cleanupHook()
	return newPool(host)
}

func newPool(host string) *redis.Pool {
	dial := func() (redis.Conn, error) {
		return redis.Dial("tcp", host)
	}

	testOnBorrow := func(c redis.Conn, t time.Time) error {
		_, err := c.Do("PING")
		return err
	}

	return &redis.Pool{
		MaxIdle:      3,
		IdleTimeout:  4 * time.Minute,
		Dial:         dial,
		TestOnBorrow: testOnBorrow,
	}
}
