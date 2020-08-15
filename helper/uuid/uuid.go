package uuid

import (
	"sync"

	u "github.com/google/uuid"
)

func NewUUID() string {
	var mu sync.Mutex
	mu.Lock()
	res := u.New().String()
	mu.Unlock()
	return res
}
