/*
 * Simple caching library with expiration capabilities
 *     Copyright (c) 2012, Radu Ioan Fericean
 *                   2013-2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE.txt
 */

package cache2go

import (
	"sync"
)

var (
	//而new返回的是指向类型的指针
	//make也是用于内存分配的，但是和new不同，它只用于chan、map以及切片的内存创建，
	//而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，
	//所以就没有必要返回他们的指针了。
	cache = make(map[string]*CacheTable)
	mutex sync.RWMutex
)

// Cache returns the existing cache table with given name or creates a new one
// if the table does not exist yet.
// 根据table name 查询 或 新建 一个CacheTable
func Cache(table string) *CacheTable {
	mutex.RLock()
	//返回两个值
	t, ok := cache[table]
	mutex.RUnlock()

	if !ok {
		mutex.Lock()
		t, ok = cache[table]
		// Double check whether the table exists or not.
		if !ok {
			t = &CacheTable{
				name:  table,
				items: make(map[interface{}]*CacheItem),
			}
			cache[table] = t
		}
		mutex.Unlock()
	}

	return t
}
