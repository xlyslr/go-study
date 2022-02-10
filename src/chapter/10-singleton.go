package chapter

import "sync"

type singleton struct {
}

var instance *singleton

var once sync.Once

func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
