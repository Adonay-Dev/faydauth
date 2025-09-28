package store

import "sync"

type MemoryCache struct {
	mu    sync.RWMutex
	cache map[string]string
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{cache: make(map[string]string)}
}

func (m *MemoryCache) Save(token, userID string, ttl time.Duration) error {
	m.mu.Lock()
	m.cache[token] = userID
	m.mu.Unlock()
	return nil
}

func (m *MemoryCache) Get(token string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	uid, ok := m.cache[token]
	return uid, ok
}

func (m *MemoryCache) Delete(token string) error {
	m.mu.Lock()
	delete(m.cache, token)
	m.mu.Unlock()
	return nil
}
