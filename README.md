# cache-using-golang
# How to setup
```git
git clone https://github.com/ramveer93/cache-using-golang.git
```
- install go version 1.16.3
- import the project in visual studio code
- run go test using 
```go
go test -v
```

## About test cases
The functionality of cache is being testing by using test cases , there are test cases for :
- Initialize the cache with capacity
- test the Add and Get operations in cache 
- test the cache when cache reaches its capacity, to find out the eviction 
- test when an entry get refreshed, or a Get call happen to that entry, the entry should come in front of the list to make it recently used
- test the update functionality of cache 
- test the evict functionality of cache 
- test concurrent , we will add 10 entries into cache in chunks of 5 - 5 entries by subscribing to a chennal, then we will try to collect these 10 entries from channel and store them into an array, sorting the array in asc order and comparing if any of the entry is missing will validate the syn operations on cache 

## Usages
There are 4 method this cache has , below is the detailed explanation about the usages of these:

### Add(key, value , costFunc)
This method will add an entry to cache 
#### Inputs:
- key: string type , a key to be added in cache , same will be used to retrieve the data 
- value: string type, value of the key 
- costFunc: fun which will print the cost of adding this entry to map , calculation of cost is sum of number of entries needs to be removed and put front when add happen. The cost of adding an entry if the cache reaches its capacity is constant(2), because we just need to perform two steps: one remove entry from back and add entry to front, both operations will take constant time as we are using list to store there entries.
### Evict(key)
This method will remove the key and corrosponding value from cache 
#### Inputs:
- key: string type, key to be removed
### Get(key)
This method will retrieve corrosponding value from cache , it will return value,error, if the key is not present in cache then error will be non nill else error will be nil and value will be returned.
#### Inputs:
- key: string type
### Update(key, value)
This method will update the value which corrosponds to key , if the key is not present in cache , it will do nothing and return false , else update value and return true.
#### Inputs:
- key: string type
- value: string type

## About the code 
We use list (https://golang.org/pkg/container/list/) to have add and remove operations in O(1) and map(https://tour.golang.org/moretypes/19) to make get operations in O(1)
Also for sync the library uses https://golang.org/pkg/sync/ to make method mutual exclusive.

## Limitations:
- Only supports string type key and values

# License 
MIT License 2021

