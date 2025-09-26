package cache

import (
	"encoding/json"
	"time"
)

// GetOrSetCache checks Redis first, otherwise calls fetchFn and stores result
func GetOrSetCache[T any](key string, ttl time.Duration, target *T, fetchFn func() (*T, error)) error {
	// 1. Try to get from cache only if Redis is connected
	if IsConnected() {
		val, err := Rdb.Get(Ctx, key).Result()
		if err == nil {
			// Found in Redis, unmarshal
			if unmarshalErr := json.Unmarshal([]byte(val), target); unmarshalErr == nil {
				return nil
			}
			// if unmarshal fails, we'll fall back to fetchFn
		}
	}

	// 2. Fetch from DB or other source
	freshData, fetchErr := fetchFn()
	if fetchErr != nil {
		return fetchErr
	}

	// Assign result to target
	*target = *freshData

	// 3. Store in Redis only if connected
	if IsConnected() {
		if jsonData, marshalErr := json.Marshal(freshData); marshalErr == nil {
			Rdb.Set(Ctx, key, jsonData, ttl)
		}
	}

	return nil
}
