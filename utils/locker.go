package utils

import (
	"sync"
)

/*
如果接口对应的业务具有重量（消耗资源、耗时）、要求幂等，那就应该给接口上锁（且要求等待阻塞）
又如果业务操作是针对每一个用户、或是特定的场景，就应该使用当前方法

1.为指定的接口绑定一个 Locker 实例
l := NewLocker()

2.具体的接口方法中
l.Lock(key)
defer l.Unlock(key)
*/

type locker sync.Map

func NewLocker() Locker {
	return new(locker)
}

type Locker interface {
	Lock(key interface{})
	Unlock(key interface{})
}

func (l *locker) Lock(key interface{}) {
	c, _ := ((*sync.Map)(l)).LoadOrStore(key, make(chan struct{}, 1))
	c.(chan struct{}) <- struct{}{}
}

func (l *locker) Unlock(key interface{}) {
	c, ok := ((*sync.Map)(l)).Load(key)
	if !ok {
		panic("unlock none")
	}
	<-c.(chan struct{})
}
