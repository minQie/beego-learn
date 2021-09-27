package utils

import (
	"sync"
)

/*
如果接口对应的业务具有重量（消耗资源、耗时）、要求幂等，那就应该给接口上锁（且要求实时响应）
又如果业务操作是针对每一个用户、或是特定的场景，就应该使用当前方法

1.为指定的接口绑定一个 Limiter 实例
l := NewLimiter()

2.具体的接口方法中
if ok := l.Lock(key); !ok {
	// 相应的响应提示
}
defer l.Unlock(key)
*/

type limiter sync.Map

func NewLimiter() Limiter {
	return new(limiter)
}

type Limiter interface {
	Lock(key interface{}) bool
	Unlock(key interface{}) bool
}

func (l *limiter) Lock(key interface{}) bool {
	_, ok := ((*sync.Map)(l)).LoadOrStore(key, struct{}{})
	return !ok
}

func (l *limiter) Unlock(key interface{}) bool {
	_, ok := ((*sync.Map)(l)).LoadAndDelete(key)
	return ok
}
