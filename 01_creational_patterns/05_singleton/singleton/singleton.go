package singleton

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = new(Singleton)
		fmt.Println("creating instance")
	})
	return instance
}
