package pubsub

import (
	"sync"
)

type fn func([]string)

// Config represents a configuration file.
type Router struct {
	tbl *map[string][]fn
}

var instance *Router
var once sync.Once
var r_tbl map[string][]fn

// New creates a new Config object.
func GetInstance() *Router {
	once.Do(func() {
		r_tbl = make(map[string][]fn)
		instance = &Router{&r_tbl}
	})
	return instance
}

// Subscribe to a message.
func (r *Router) Subscribe(key string, f fn) {
	var fl []fn
	fl = r_tbl[key]
	r_tbl[key] = append(fl, f)
}

// Publish Message
func (r *Router) Publish(key string, data []string) {
	var ok bool
	var fl []fn
	fl, ok = r_tbl[key]
	if ok {
		//call each fn inside fl
		for _, f := range fl {
			f(data)
		}
	}
}
