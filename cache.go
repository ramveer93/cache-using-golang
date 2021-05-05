package main

import (
	"container/list"
	"errors"
	"fmt"
)

type Cache struct {
	capacity   int
	linkedList *list.List
	cacheMap   map[string]*list.Element
}

type CacheStruct struct {
	key   string
	value string
}

/**
*Initialize the cache with capacity
 */
func New(capacity int) Cache {
	cache := Cache{
		capacity:   capacity,
		linkedList: list.New(),
		cacheMap:   make(map[string]*list.Element),
	}
	return cache
}
func Cost(m map[string]*list.Element, capacity int) int {
	if len(m) >= capacity {
		return 2
	} else {
		return 0
	}
}

func (cache Cache) Add(key string, value string, costFunc func(m map[string]*list.Element, capacity int) int) bool {
	//see if map has key
	//if yes then remove from list and put in front
	totalCost := 0
	elem, found := cache.cacheMap[key]
	if found {
		elem.Value = value
		cache.linkedList.MoveBefore(elem, cache.linkedList.Front()) //move elem to front of link list
		return true
	} else {
		if len(cache.cacheMap) >= cache.capacity {
			//this is the eviction function
			totalCost = costFunc(cache.cacheMap, cache.capacity)
			fmt.Println("Cost : ", totalCost)
			el := cache.linkedList.Back()
			cache.linkedList.Remove(el)
			newKey := el.Value.(CacheStruct).key
			delete(cache.cacheMap, newKey)
		}
		//push in the front of queue
		entry := CacheStruct{key: key, value: value}
		newItem := cache.linkedList.PushFront(entry)
		cache.cacheMap[key] = newItem
		return true
	}

}

func (cache Cache) Get(key string) (string, error) {
	//check if key exists
	elem, found := cache.cacheMap[key]
	if found {
		cache.linkedList.MoveBefore(elem, cache.linkedList.Front()) //move elem to front of link list
		val := elem.Value.(CacheStruct).value
		return val, nil
	} else {
		return "", errors.New("no key exists in map")
	}
}

func (cache Cache) Update(key string, value string) bool {
	//if key exists then
	//update its value and move that to front
	elem, found := cache.cacheMap[key]
	if found {
		elem.Value = CacheStruct{key: key, value: value}
		cache.linkedList.MoveBefore(elem, cache.linkedList.Front()) //move elem to front of link list
		return true
	} else {
		return false
	}
}

func (cache Cache) Evict(key string) bool {
	//evict policy is that
	//it will remove last item from
	elem, found := cache.cacheMap[key]
	if found {
		cache.linkedList.Remove(elem)
		delete(cache.cacheMap, key)
		return true
	} else {
		return false
	}

}

func (cache Cache) ShowContent() {
	for e := cache.linkedList.Front(); e != nil; e = e.Next() {
		k := e.Value.(CacheStruct).key
		v := cache.cacheMap[k].Value.(CacheStruct).value
		fmt.Println(" k and v", k, v)
	}
}
