package kvs

import (
	"fmt"
	"strings"

	"github.com/garyburd/redigo/redis"
)

// Ping will return nil if everything is OK or an error
func (rd *Redis) Ping() error {
	conn := rd.Pool.Get()
	defer conn.Close()

	if _, err := redis.String(conn.Do("PING")); err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

// Get some value by its key
func (rd *Redis) Get(key string) ([]byte, error) {
	conn := rd.Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

// Set some value by its key
func (rd *Redis) Set(key string, value []byte) error {
	conn := rd.Pool.Get()
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

// AddToSet will add keys to a set
func (rd *Redis) AddToSet(set string, keys ...interface{}) error {
	conn := rd.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SADD", set, keys)

	if err != nil {
		return fmt.Errorf("error adding keys to set %s: %v", set, err)
	}

	return err
}

// Exists will return true if an item exists
func (rd *Redis) Exists(key string) (bool, error) {
	conn := rd.Pool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", "["+key+"]"))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

// ExistsInSet will return true if an item is present in a set
func (rd *Redis) ExistsInSet(set, key string) (bool, error) {
	conn := rd.Pool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("SISMEMBER", set, "["+key+"]"))
	if err != nil {
		return ok, fmt.Errorf("error checking the key %s exists in set %s: %v", key, set, err)
	}
	return ok, err
}

// Delete will remove an item by its key
func (rd *Redis) Delete(key string) error {
	conn := rd.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

// GetKeys will return all the keys
func (rd *Redis) GetKeys(pattern string) ([]string, error) {
	conn := rd.Pool.Get()
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

// GetSetMembers will return all the keys in a set
func (rd *Redis) GetSetMembers(set string) ([]string, error) {
	conn := rd.Pool.Get()
	defer conn.Close()

	arr, err := redis.Values(conn.Do("SMEMBERS", set))
	if err != nil {
		return nil, fmt.Errorf("error retrieving the keys from the set '%v': %v", set, err)
	}

	keys, err := redis.Strings(arr, nil)

	// The keys are surrounded by [ ], we need to trim them
	for i, k := range keys {
		keys[i] = trimBrackets(k)
	}

	return keys, err
}

// Incr will ... increment a counter ? tbd
func (rd *Redis) Incr(counterKey string) (int, error) {
	conn := rd.Pool.Get()
	defer conn.Close()

	return redis.Int(conn.Do("INCR", counterKey))
}

func trimBrackets(s string) string {
	trimmed := strings.Trim(s, "[")
	trimmed = strings.Trim(trimmed, "]")
	return trimmed
}
