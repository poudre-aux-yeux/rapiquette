package kvs

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Redis : Redis instance
type Redis struct {
	Pool *redis.Pool
}

// NewRedis instantiates a new RedisPool
func NewRedis(host string) *Redis {
	pool := newPool(host)

	r := Redis{
		Pool: pool,
	}

	r.cleanupHook()
	return &r
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

// Intercept the exit signal and close the pool properly before exiting
func (rd *Redis) cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		rd.Pool.Close()
		os.Exit(0)
	}()
}
