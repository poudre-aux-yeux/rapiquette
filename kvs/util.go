package kvs

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// Ping will return nil if everything is OK or an error
func Ping() error {
	conn := Pool.Get()
	defer conn.Close()

	if _, err := redis.String(conn.Do("PING")); err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

// Get some value by its key
func Get(key string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

// Set some value by its key
func Set(key string, value []byte) error {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

// Exists will return true if an item exists
func Exists(key string) (bool, error) {
	conn := Pool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

// Delete will remove an item by its key
func Delete(key string) error {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

// GetKeys will return all the keys
func GetKeys(pattern string) ([]string, error) {
	conn := Pool.Get()
	defer conn.Close()

	iter := 0
	keys := []string{}
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}

// Incr will ... increment a counter ? tbd
func Incr(counterKey string) (int, error) {
	conn := Pool.Get()
	defer conn.Close()

	return redis.Int(conn.Do("INCR", counterKey))
}
