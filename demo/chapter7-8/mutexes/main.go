package main 

// import (
// 	"time"
// 	"sync"	
// )

// type safeCounter struct {
// 	counts map[string]int 
// 	mux *sync.Mutex
// }

// func (sc safeCounter) inc(key string) {
// 	sc.mux.Lock()
// 	defer sc.mux.Unlock()
// 	sc.slowIncrement(key)
// }

// func (sc safeCounter) val(key string) int {
// 	return sc.counts[key]
// }

// // func slowIncrement(key) {
// // }
