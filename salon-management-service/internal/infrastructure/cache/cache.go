package cache

import (
	"sync"

	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/entity"
)

type MemoryCache struct {
	procedures  []*entity.Procedure
	specialists []*entity.Specialist
	mu          sync.RWMutex
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{}
}

func (c *MemoryCache) SetProcedures(p []*entity.Procedure) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.procedures = p
}

func (c *MemoryCache) GetProcedures() ([]*entity.Procedure, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.procedures == nil {
		return nil, false
	}
	return c.procedures, true
}

func (c *MemoryCache) SetSpecialists(s []*entity.Specialist) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.specialists = s
}

func (c *MemoryCache) GetSpecialists() ([]*entity.Specialist, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.specialists == nil {
		return nil, false
	}
	return c.specialists, true
}
