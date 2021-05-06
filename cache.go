package main

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

type Cache struct {
	capacity   int                      //capacity of cache, only full eviction will happen
	linkedList *list.List               //list to make add,delete operations O(1)
	cacheMap   map[string]*list.Element //map to store key and corrosponding list as value
}

type CacheStruct struct {
	key   string
	value string
}

var lock sync.Mutex //sync lock to make methods mutually exclusive

/**
*Initialize the cache with capacity
 */
func New(capacity int) Cache {
	lock.Lock()
	defer lock.Unlock()
	cache := Cache{
		capacity:   capacity,
		linkedList: list.New(),
		cacheMap:   make(map[string]*list.Element),
	}
	return cache
}

/**
This method will calculate the cost of eviction
this cost will be constant, because every time when cache is full
and an entry needs to be added , least recent used entry (back)
will be removed and new entry will be added in front
so the const is just constant
*/
func Cost(m map[string]*list.Element, capacity int) int {
	lock.Lock()
	defer lock.Unlock()

	if len(m) >= capacity {
		return 2
	} else {
		return 0
	}
}

/**
This method will add a new key , value to cache
Below can be the scenerios:
1. When the key already exists: Make this key recently used by
moving it to front
2. When the key is not present in cache: if cache has reached
its capacity, we will evict the element from back and put the
entry in front
*/
func (cache Cache) Add(key string, value string, costFunc func(m map[string]*list.Element, capacity int) int) bool {
	lock.Lock()
	defer lock.Unlock()
	totalCost := 0
	elem, found := cache.cacheMap[key]
	if found {
		elem.Value = value
		cache.linkedList.MoveBefore(elem, cache.linkedList.Front())
		return true
	} else {
		if len(cache.cacheMap) >= cache.capacity {
			totalCost = costFunc(cache.cacheMap, cache.capacity)
			fmt.Println("Cost : ", totalCost)
			el := cache.linkedList.Back()
			cache.linkedList.Remove(el)
			newKey := el.Value.(CacheStruct).key
			delete(cache.cacheMap, newKey)
		}
		entry := CacheStruct{key: key, value: value}
		newItem := cache.linkedList.PushFront(entry)
		cache.cacheMap[key] = newItem
		return true
	}

}

/**
This method will get the value for a key from cache
If the key exists in cache then move it to front to make
it recently used and return it.
If key is not present then throw error
*/
func (cache Cache) Get(key string) (string, error) {
	lock.Lock()
	defer lock.Unlock()
	elem, found := cache.cacheMap[key]
	if found {
		cache.linkedList.MoveBefore(elem, cache.linkedList.Front()) //move elem to front of link list
		val := elem.Value.(CacheStruct).value
		return val, nil
	} else {
		return "", errors.New("no key exists in map")
	}
}

/**
This method will perform an update to value of existing key
If the key present in cache then update its value to input
value and make it recently used and return bool true
If the key is not present in cache then return bool false
*/
func (cache Cache) Update(key string, value string) bool {
	lock.Lock()
	defer lock.Unlock()
	elem, found := cache.cacheMap[key]
	if found {
		elem.Value = CacheStruct{key: key, value: value}
		cache.linkedList.MoveBefore(elem, cache.linkedList.Front()) //move elem to front of link list
		return true
	} else {
		return false
	}
}

/**
The method will remove the key and corrosponing value from cache
Remove the key ,value from cache only if the key is present in cache
else return bool false
*/
func (cache Cache) Evict(key string) bool {
	lock.Lock()
	defer lock.Unlock()
	elem, found := cache.cacheMap[key]
	if found {
		cache.linkedList.Remove(elem)
		delete(cache.cacheMap, key)
		return true
	} else {
		return false
	}

}

/**
This is utility method to print the cache content
at any time for testing
*/
func (cache Cache) ShowContent() {
	lock.Lock()
	defer lock.Unlock()
	for e := cache.linkedList.Front(); e != nil; e = e.Next() {
		k := e.Value.(CacheStruct).key
		v := cache.cacheMap[k].Value.(CacheStruct).value
		fmt.Println(" k and v", k, v)
	}
}
