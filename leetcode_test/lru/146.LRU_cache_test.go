package lru

import (
	"container/list"
	"fmt"
	"testing"
)

/**
No.146 LRU 缓存机制
描述：
	运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制 。
	实现 LRUCache 类：
		LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
		int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
		void put(int key, int value)如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
	你是否可以在 O(1) 时间复杂度内完成这两种操作？
思路：
    1、
		双线链表：节点存储key+value，可以保证顺序，且增、删的时间复杂度为O(1)
		hash表：key为查询键，value为链表节点地址，读的时间复杂度为O(1)
		通过双向链表 + hash的方式，让顺序结构并且增删改为O(1)
		读：查询hash表，有则返回值并且将数据插到链表头，时间复杂度O(1)
		写：查询hash表，有则修改值并且将数据插到链表头；没有则新增链表节点排在头部，增加hash表，同时判断如果超过限制长度cap，删除链表尾部且删除hash表，时间复杂度O(1)
时间：
    1、2021/9/20
*/

type LRUCache struct {
	L *list.List
	M map[int]*list.Element

	Cap int
}

type node struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{L: list.New(), M: make(map[int]*list.Element, 10), Cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	// 从map中获取，没有则返回-1
	val, ok := this.M[key]
	if !ok {
		fmt.Println(-1)
		return -1
	}
	// 有则返回值，调整至链表头部
	this.L.MoveToFront(val)
	fmt.Println(val.Value.(*node).value)
	return val.Value.(*node).value
}

func (this *LRUCache) Put(key int, value int) {
	val, ok := this.M[key]
	// 如果不存在，则新增到头部
	if !ok {
		val = this.L.PushFront(&node{key: key, value: value})
		// 修改hash值
		this.M[key] = val
		// **淘汰规则：如果超过长度，淘汰末尾，即淘汰内存中时间最长的且没有访问的
		if this.Cap < this.L.Len() {
			last := this.L.Back()
			delete(this.M, last.Value.(*node).key)
			this.L.Remove(last)
		}
		return
	}
	// 存在则修改链表值，并且调整到头部
	val.Value.(*node).value = value
	this.L.MoveToFront(val)
}

func TestLRUCache(t *testing.T) {
	obj := Constructor(2)

	obj.Put(1, 0)
	obj.Put(2, 2)
	obj.Get(1)
	obj.Put(3, 3) // [3,3 1,0]
	obj.Get(2)
	obj.Put(4, 4) // [4,4 3.3]
	obj.Get(1)    // 0
	obj.Get(3)
	obj.Get(4)
}
