package main

import (
	"container/list"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func main() {

}

func TestCacheInit(t *testing.T) {
	cache := New(3)
	assert.Equal(t, 3, cache.capacity)
}
func TestCacheAddGet(t *testing.T) {
	cache := New(3)
	assert.Equal(t, 3, cache.capacity)

	costFunc := func(m map[string]*list.Element, capacity int) int {
		if len(m) >= capacity {
			return 2
		} else {
			return 0
		}
	}

	isRajAdded := cache.Add("Rajasthan", "Jaipur", costFunc)
	assert.Equal(t, isRajAdded, true)

	isUpAdded := cache.Add("UP", "Lucknow", costFunc)
	assert.Equal(t, isUpAdded, true)

	isKrAdded := cache.Add("Karnataka", "Bangalore", costFunc)
	assert.Equal(t, isKrAdded, true)

	rJCap, _ := cache.Get("Rajasthan")
	assert.Equal(t, rJCap, "Jaipur")

	uPCap, _ := cache.Get("UP")
	assert.Equal(t, uPCap, "Lucknow")

	kRCap, _ := cache.Get("Karnataka")
	assert.Equal(t, kRCap, "Bangalore")

}

func TestCacheWhenCapacityFull(t *testing.T) {
	cache := New(3)
	assert.Equal(t, 3, cache.capacity)

	costFunc := func(m map[string]*list.Element, capacity int) int {
		if len(m) >= capacity {
			return 2
		} else {
			return 0
		}
	}

	isRajAdded := cache.Add("Rajasthan", "Jaipur", costFunc)
	assert.Equal(t, isRajAdded, true)

	isUpAdded := cache.Add("UP", "Lucknow", costFunc)
	assert.Equal(t, isUpAdded, true)

	isKrAdded := cache.Add("Karnataka", "Bangalore", costFunc)
	assert.Equal(t, isKrAdded, true)

	rJcap, _ := cache.Get("Rajasthan")
	assert.Equal(t, rJcap, "Jaipur")

	uPcap, _ := cache.Get("UP")
	assert.Equal(t, uPcap, "Lucknow")

	kRcap, _ := cache.Get("Karnataka")
	assert.Equal(t, kRcap, "Bangalore")

	isMhAdded := cache.Add("Maharastra", "Mumbai", costFunc)
	assert.Equal(t, isMhAdded, true)

	mhCap, _ := cache.Get("Maharastra")
	assert.Equal(t, mhCap, "Mumbai")

	_, e := cache.Get("Rajasthan")
	if e != nil {
		assert.Equal(t, true, true)
	}
}

func TestWhenEntryGetRefreshed(t *testing.T) {
	cache := New(3)
	assert.Equal(t, 3, cache.capacity)

	costFunc := func(m map[string]*list.Element, capacity int) int {
		if len(m) >= capacity {
			return 2
		} else {
			return 0
		}
	}

	isRajAdded := cache.Add("Rajasthan", "Jaipur", costFunc)
	assert.Equal(t, isRajAdded, true)

	isUpAdded := cache.Add("UP", "Lucknow", costFunc)
	assert.Equal(t, isUpAdded, true)

	isKrAdded := cache.Add("Karnataka", "Bangalore", costFunc)
	assert.Equal(t, isKrAdded, true)

	k, _ := cache.Get("Rajasthan")
	assert.Equal(t, k, "Jaipur")

	isMhAdded := cache.Add("Maharastra", "Mumbai", costFunc)
	assert.Equal(t, isMhAdded, true)

	k1, _ := cache.Get("UP")
	if k1 == "" {
		assert.Equal(t, true, true)
	} else {
		assert.Equal(t, false, true)
	}

}

func TestUpdate(t *testing.T) {
	cache := New(3)
	assert.Equal(t, 3, cache.capacity)

	costFunc := func(m map[string]*list.Element, capacity int) int {
		if len(m) >= capacity {
			return 2
		} else {
			return 0
		}
	}

	isRajAdded := cache.Add("Rajasthan", "Jaipur", costFunc)
	assert.Equal(t, isRajAdded, true)

	isUpAdded := cache.Add("UP", "Lucknow", costFunc)
	assert.Equal(t, isUpAdded, true)

	isKrAdded := cache.Add("Karnataka", "Bangalore", costFunc)
	assert.Equal(t, isKrAdded, true)

	status := cache.Update("Rajasthan", "New Jaipur")
	assert.Equal(t, true, status)

	isMhAdded := cache.Add("Maharastra", "Mumbai", costFunc)
	assert.Equal(t, isMhAdded, true)

	k1, _ := cache.Get("UP")
	if k1 == "" {
		assert.Equal(t, true, true)
	} else {
		assert.Equal(t, false, true)
	}
	k2, _ := cache.Get("Rajasthan")
	assert.Equal(t, k2, "New Jaipur")

	falseStatus := cache.Update("Test", "Testing")
	assert.Equal(t, falseStatus, false)
}

func TestEvict(t *testing.T) {
	cache := New(3)
	assert.Equal(t, 3, cache.capacity)

	costFunc := func(m map[string]*list.Element, capacity int) int {
		if len(m) >= capacity {
			return 2
		} else {
			return 0
		}
	}

	isRajAdded := cache.Add("Rajasthan", "Jaipur", costFunc)
	assert.Equal(t, isRajAdded, true)

	isUpAdded := cache.Add("UP", "Lucknow", costFunc)
	assert.Equal(t, isUpAdded, true)

	isKrAdded := cache.Add("Karnataka", "Bangalore", costFunc)
	assert.Equal(t, isKrAdded, true)

	status := cache.Evict("Karnataka")
	assert.Equal(t, status, true)

	_, e := cache.Get("Karnataka")
	if e != nil {
		assert.Equal(t, true, true)
	} else {
		assert.Equal(t, true, false)
	}
}

func TestConsurrent(t *testing.T) {
	cache := New(3)
	assert.Equal(t, 3, cache.capacity)
	ch := make(chan string)
	const iterations = 1000
	var a [iterations]string

	costFunc := func(m map[string]*list.Element, capacity int) int {
		if len(m) >= capacity {
			return 2
		} else {
			return 0
		}
	}

	//first add 5 values to map
	go func() {
		for i := 0; i < iterations/2; i++ {
			cache.Add(strconv.Itoa(i), strconv.Itoa(i), costFunc)
			val, _ := cache.Get(strconv.Itoa(i))
			ch <- val
		}
	}()

	//second add 5 values to map
	go func() {
		for i := iterations / 2; i < iterations; i++ {
			cache.Add(strconv.Itoa(i), strconv.Itoa(i), costFunc)
			val, _ := cache.Get(strconv.Itoa(i))
			ch <- val
		}
	}()

	count := 0
	for elem := range ch {
		a[count] = elem
		count++
		if count == iterations {
			break
		}
	}
	sort.Strings(a[0:iterations])

	if len(a) != iterations {
		t.Error("Expecting 1000 elements.")
	}

	for i := 0; i < iterations; i++ {
		if strconv.Itoa(i) != a[i] {
			t.Error("missing value", i)
		}
	}

}
