package demo

import "sync"

// singleton Define
type singleton struct{}

// 饿汉式
var hungryIns *singleton = &singleton{}

// GetHungryIns is used to return singleton instance
func GetHungryIns() *singleton {
	return hungryIns
}

// 懒汉式
var lazyIns *singleton

// GetLazyIns is used to return singleton instance
func GetLazyIns() *singleton {
	if lazyIns == nil {
		lazyIns = &singleton{}
	}
	return lazyIns
}

// 双重检测
var doubleIns *singleton
var mu sync.Mutex

// GetDoubleIns is used to return singleton instance
func GetDoubleIns() *singleton {
	if doubleIns == nil {
		mu.Lock()
		defer mu.Unlock()
		if doubleIns == nil {
			doubleIns = &singleton{}
		}
	}
	return doubleIns
}

// sync.Once实现
var onceIns *singleton
var once sync.Once

// GetOnceIns is used to return singleton instance
func GetOnceIns() *singleton {
	once.Do(func() {
		onceIns = &singleton{}
	})
	return onceIns
}
