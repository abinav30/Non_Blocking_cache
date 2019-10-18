package Non_Blocking_cache

import (
	"sync"
)
type Func func(string)(interface{},error)
type result struct{
	value interface{}
	err error
}
type entry struct{
	res result
	ready chan struct{}
}
type Nbc struct{
	f Func
	mu sync.Mutex
	cache map[string]*entry
}
//Returns New map
func New(f Func) *Nbc{
	return &Nbc{f:f,cache:make(map[string]*entry)}
}
//Gets the Value from the cache or stores it if no valu is already present
//Concurrency safe
func (nbc *Nbc) Get(key string)(value interface{},err error){
	nbc.mu.Lock()
	e := nbc.cache[key]
	if e == nil{
		//Creates a reference to an entry and stores teh pointer in the map
		e = &entry{ready:make(chan struct{})}
		nbc.cache[key]=e
		nbc.mu.Unlock()
		e.res.value, e.res.err = nbc.f(key)
		close(e.ready)
	}else{
		nbc.mu.Unlock()
		//This operation blocks until the ready channel is closed
		<-e.ready
	}
	return e.res.value,e.res.err
}