package lock

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type DistributedLock struct {
	Client  *redis.Client
	Key     string
	Value   string
	Timeout time.Duration
}

func New(client *redis.Client, key, value string, timeout time.Duration) *DistributedLock {
	return &DistributedLock{
		Client:  client,
		Key:     key,
		Value:   value,
		Timeout: timeout,
	}
}

func (d *DistributedLock) Acquire(ctx context.Context) (bool, error) {
	result, err := d.Client.SetNX(ctx, d.Key, d.Value, d.Timeout).Result()
	if err != nil {
		return false, err
	}
	return result, nil
}

func (d *DistributedLock) Release(ctx context.Context) error {
	val, err := d.Client.Get(ctx, d.Key).Result()
	if err != nil {
		return err
	}
	if val != d.Value {
		return fmt.Errorf("distributed lock value is not equal with origin value")
	}
	_, err = d.Client.Del(ctx, d.Key).Result()
	return err
}
