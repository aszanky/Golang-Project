package memcache

import (
	"github.com/bluele/gcache"
)

type Memcache interface {
	Set(key interface{}, value interface{}) error
	Update(key interface{}, value interface{}) error
	Delete(key interface{}) bool
	Get(key interface{}) (interface{}, error)
	GetALL() map[interface{}]interface{}
}

//memcache is memory struct cache
//here I am using cache for data store
type memcache struct {
	gcache gcache.Cache
}

//New instantiation
func New() Memcache {
	client := gcache.New(256).LRU().Build()
	return &memcache{
		gcache: client,
	}
}

//Set for set the key/value data
func (m *memcache) Set(key interface{}, value interface{}) error {
	return m.gcache.Set(key, value)
}

//Update for updating data in cache
func (m *memcache) Update(key interface{}, value interface{}) error {
	var err error
	checkKey := m.gcache.Has(key)
	if checkKey {
		err = m.gcache.Set(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

//Delete for deleting data
func (m *memcache) Delete(key interface{}) bool {
	return m.gcache.Remove(key)
}

//Get for getting only specific key
func (m *memcache) Get(key interface{}) (interface{}, error) {
	return m.gcache.Get(key)
}

//GetALL for getting all data in memory cache
func (m *memcache) GetALL() map[interface{}]interface{} {
	checkExpired := false
	return m.gcache.GetALL(checkExpired)
}
