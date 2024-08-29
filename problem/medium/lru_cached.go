package medium

// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

// Implement the LRUCache class:

// LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
// int get(int key) Return the value of the key if the key exists, otherwise return -1.
// void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
// The functions get and put must each run in O(1) average time complexity.

// Input
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// Output
// [null, null, null, 1, null, -1, null, -1, 3, 4]

type LRUCache struct {
	mapCache    map[int]int
	top, bottom *Node
	capacity    int
}

type Node struct {
	key, value int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	top := &Node{}
	bottom := &Node{}

	bottom.next = top
	top.prev = bottom

	return LRUCache{
		mapCache: make(map[int]int, capacity),
		top:      top,
		bottom:   bottom,
		capacity: capacity,
	}

}

func (this *LRUCache) MoveNodeToTop(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = this.top.prev
	this.top.prev.next = node
	this.top.prev = node
	node.next = this.top
}

func (this *LRUCache) AddNodeToTop(node *Node) {
	node.prev = this.top.prev
	node.next = this.top
	node.prev.next = node
	this.top.prev = node
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.mapCache[key]; ok {
		this.MoveNodeToTop(&Node{key: key, value: val})
		return val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.mapCache[key]; ok {
		this.mapCache[key] = value
		this.MoveNodeToTop(&Node{key: key, value: value})
		return
	}

	if len(this.mapCache) == this.capacity {
		delete(this.mapCache, this.bottom.next.key)
		this.bottom.next = this.bottom.next.next
		this.bottom.next.prev = this.bottom
	}
	this.mapCache[key] = value
	this.AddNodeToTop(&Node{key: key, value: value})
}
