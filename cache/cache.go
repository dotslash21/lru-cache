package cache

import "fmt"

type node[T comparable] struct {
	Value T
	Left  *node[T]
	Right *node[T]
}

type queue[T comparable] struct {
	Head     *node[T]
	Tail     *node[T]
	Length   int
	Capacity int
}

type hashTable[T comparable] map[T]*node[T]

func initQueue[T comparable](capacity int) queue[T] {
	head := &node[T]{}
	tail := &node[T]{}

	head.Right = tail
	tail.Left = head

	return queue[T]{Head: head, Tail: tail, Length: 0, Capacity: capacity}
}

type Cache[T comparable] struct {
	Queue     queue[T]
	HashTable hashTable[T]
}

func InitCache[T comparable](capacity int) Cache[T] {
	return Cache[T]{Queue: initQueue[T](capacity), HashTable: hashTable[T]{}}
}

func (cache *Cache[T]) Get(item T) {
	if node, ok := cache.HashTable[item]; ok {
		cache.Remove(node)
	}

	node := &node[T]{Value: item}
	cache.Add(node)
	cache.HashTable[item] = node
}

func (cache *Cache[T]) Remove(node *node[T]) {
	node.Left.Right = node.Right
	node.Right.Left = node.Left
	node.Right = nil
	node.Left = nil

	delete(cache.HashTable, node.Value)

	cache.Queue.Length--
}

func (cache *Cache[T]) Add(node *node[T]) {
	if cache.Queue.Length == cache.Queue.Capacity {
		cache.Remove(cache.Queue.Tail.Left)
	}

	node.Right = cache.Queue.Head.Right
	node.Left = cache.Queue.Head
	cache.Queue.Head.Right.Left = node
	cache.Queue.Head.Right = node

	cache.HashTable[node.Value] = node

	cache.Queue.Length++
}

func (cache Cache[T]) String() string {
	return fmt.Sprintf("Cache: %v", cache.Queue)
}

func (queue queue[T]) String() string {
	node := queue.Head.Right
	str := ""

	for node != queue.Tail {
		str += fmt.Sprintf("{%v} ", node.Value)
		node = node.Right
	}

	return str
}
