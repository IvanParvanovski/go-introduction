package main

type Func func(key string, done <-chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	// Run the actual function and store the result
	e.res.value, e.res.err = f(key, done)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result, done <-chan struct{}) {
	select {
	case <-e.ready:
		response <- e.res
	case <-done:
		// Aborted, do not send result
	}
}

type request struct {
	key      string
	response chan<- result
	done     <-chan struct{}
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result, 1) // buffered to prevent goroutine leaks
	memo.requests <- request{key, response, done}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// Create new entry and call function
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go func() {
				e.call(f, req.key, req.done)
				// If cancelled, remove from cache
				select {
				case <-e.ready:
					// success, keep it
				case <-req.done:
					// cancelled, delete from cache
					delete(cache, req.key)
				}
			}()
		}
		// Deliver the result (if not already cancelled)
		go e.deliver(req.response, req.done)
	}
}