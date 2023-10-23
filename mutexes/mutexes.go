package main

// lock access to protected resources so that only 1 goroutine can access the resource at a time
// built in implementation of a mutex is sync.Mutex and has methods Lock() and Unlock()
// func protected() {
//     mux.Lock()
//     defer mux.Unlock()
//     // the rest of the function is protected
//     // any other calls to mux.Lock() will block
// }

// type safeCounter struct {
// 	counts map[string]int
// 	mux    *sync.Mutex
// }

// func (sc safeCounter) inc(key string) {
// 	sc.mux.Lock()
// 	defer sc.mux.Unlock()
// 	sc.slowIncrement(key)
// }

// func (sc safeCounter) val(key string) {
//     sc.mux.Lock()
//     defer sc.mux.Unlock()
//     return sc.counts[key]
// }

func main() {
}
