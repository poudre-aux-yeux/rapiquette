package tennis

import (
	"encoding/json"
	"fmt"
)

// Create adds an item to the key-value store
func (c *Client) Create(item interface{}, id, set string) error {
	data, err := json.Marshal(item)

	if err != nil {
		return fmt.Errorf("impossible to marshal: %v", err)
	}

	if err = c.redis.Set(id, data); err != nil {
		return fmt.Errorf("could not set the data: %v", err)
	}

	if err = c.redis.AddToSet(set, id); err != nil {
		return fmt.Errorf("couldn't add the key %v to set %v: %v", id, set, err)
	}

	return nil
}

// KeyExists checks if the key exists in the key-value store
func (c *Client) KeyExists(key string) (bool, error) {
	return c.redis.Exists(key)
}

// KeyExistsInSet checks if the key exists in the set
func (c *Client) KeyExistsInSet(set, key string) (bool, error) {
	return c.redis.ExistsInSet(set, key)
}
