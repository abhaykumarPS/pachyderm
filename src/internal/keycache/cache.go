package keycache

import (
	"sync/atomic"
	"time"

	"github.com/gogo/protobuf/proto"
	logrus "github.com/sirupsen/logrus"

	"github.com/pachyderm/pachyderm/v2/src/internal/backoff"
	col "github.com/pachyderm/pachyderm/v2/src/internal/collection"
	"github.com/pachyderm/pachyderm/v2/src/internal/watch"
)

// Cache watches a key in etcd and caches the value in an atomic value
// This is useful for frequently read but infrequently updated values
type Cache struct {
	readOnly     col.ReadOnlyCollection
	defaultValue proto.Message
	key          string
	value        *atomic.Value
}

// NewCache returns a cache for the given key in the etcd collection
func NewCache(readOnly col.ReadOnlyCollection, key string, defaultValue proto.Message) *Cache {
	value := &atomic.Value{}
	value.Store(defaultValue)
	return &Cache{
		readOnly:     readOnly,
		value:        value,
		key:          key,
		defaultValue: defaultValue,
	}
}

// Watch should be called in a goroutine to start the watcher
func (c *Cache) Watch() {
	backoff.RetryNotify(func() error {
		return c.readOnly.WatchOneF(c.key, func(ev *watch.Event) error {
			switch ev.Type {
			case watch.EventPut:
				val := proto.Clone(c.defaultValue)
				if err := proto.Unmarshal(ev.Value, val); err != nil {
					return err
				}
				c.value.Store(val)
			case watch.EventDelete:
				c.value.Store(c.defaultValue)
			}
			return nil
		})
	}, backoff.NewInfiniteBackOff(), func(err error, d time.Duration) error {
		logrus.Printf("error from watcher for %v: %v; retrying in %v", c.key, err, d)
		return nil
	})
}

// Load retrieves the current cached value
func (c *Cache) Load() interface{} {
	return proto.Clone(c.value.Load().(proto.Message))
}