package main

import (
	"fmt"
	"lru-cache/cache"
)

func main() {
	c := cache.InitCache[string](3)

	for index, word := range []string{"hello", "world", "foo", "hello", "bar"} {
		c.Get(word)
		fmt.Printf("index: %v: %v\n", index, c)
	}
}
