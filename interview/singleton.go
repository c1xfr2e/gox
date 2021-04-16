package interview

import "sync"

var (
	mtx  sync.Mutex
	inst interface{}
)

// GetInstance returns the singleton instance.
func GetInstance() interface{} {
	if inst != nil {
		return inst
	}
	mtx.Lock()
	defer mtx.Unlock()

	// double check
	if inst != nil {
		return inst
	}

	inst = struct{}{}
	return inst
}
