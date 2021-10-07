package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	// 기본 만료 시간이 5분, 30초마다 만료된 항목 제거하는 캐시 생성
	c := cache.New(5 * time.Minute, 30 * time.Second)

	c.Set("keyExample", "valueExample", cache.DefaultExpiration)
	// cache.DefaultExpiration 값은 0s

	v, found := c.Get("keyExample")
	if found {
		fmt.Printf("key: keyExample, value: %s\n", v)
	}
}
