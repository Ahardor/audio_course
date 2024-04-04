package cacher

import (
	"iotvisual/processor/domain/melody"
	"sync"
	"time"
)

type MelodyCache struct {
	store                 sync.Map
	cleanapInterval       time.Duration
	defaultExpirationTime time.Duration
}

type Option func(*MelodyCache)

func WithCleanupInterval(interval time.Duration) Option {
	return func(mc *MelodyCache) { mc.cleanapInterval = interval }
}

func WithExpirationTime(expiration time.Duration) Option {
	return func(mc *MelodyCache) { mc.defaultExpirationTime = expiration }
}

func New(opts ...Option) *MelodyCache {
	mc := &MelodyCache{}

	for i := range opts {
		opts[i](mc)
	}

	return mc
}

// MelodyItem - обертка мелодии для хранения в кэше
type melodyItem struct {
	// Мелодия.
	melody melody.Melody
	// Время создания в Unixtime.
	createdAt int64
	// Время истечения кеша.
	expiresAt int64
}

func (c *MelodyCache) Load(key melody.ID) (melody.Melody, bool) {
	v, ok := c.store.Load(key)
	if !ok {
		return melody.Melody{}, false
	}
	item, _ := v.(melodyItem)
	if item.expiresAt < time.Now().UnixNano() {
		return melody.Melody{}, false
	}
	return item.melody, true
}

func (c *MelodyCache) Store(key melody.ID, value melody.Melody) {
	now := time.Now()
	item := melodyItem{
		melody:    value,
		createdAt: now.UnixNano(),
		expiresAt: now.Add(c.defaultExpirationTime).UnixNano(),
	}
	c.store.Store(key, item)
}

func (c *MelodyCache) Delete(key melody.ID) {
	c.store.Delete(key)
}

func (c *MelodyCache) Cleanup() {
	go func() {
		for {
			<-time.After(c.cleanapInterval)
			c.store.Range(func(key, value any) bool {
				item, _ := value.(melodyItem)
				if item.expiresAt < time.Now().UnixNano() {
					c.store.Delete(key)
				}
				return true
			})
		}
	}()
}
