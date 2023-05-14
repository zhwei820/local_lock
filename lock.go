package local_lock

import "sync"

type KeyedMutex struct {
	mutexes sync.Map // Zero value is empty and ready for use
}

// https://stackoverflow.com/questions/64564781/golang-lock-per-value
func (m *KeyedMutex) Lock(key string) func() {
	value, _ := m.mutexes.LoadOrStore(key, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()

	return func() { mtx.Unlock() }
}

var km *KeyedMutex

func init() {
	km = &KeyedMutex{}
}

func Lock(key string) func() {
	return km.Lock(key)
}
