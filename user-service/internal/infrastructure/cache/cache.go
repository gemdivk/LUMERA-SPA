package cache

import (
	"github.com/gemdivk/LUMERA-SPA/user-service/internal/domain"
	"sync"
)

type UserCache struct {
	data  map[string]*domain.User
	mutex sync.RWMutex
}

func NewUserCache() *UserCache {
	return &UserCache{
		data: make(map[string]*domain.User),
	}
}

func (c *UserCache) Get(id string) (*domain.User, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	user, found := c.data[id]
	return user, found
}

func (c *UserCache) Set(user *domain.User) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[user.ID] = user
}

func (c *UserCache) Delete(id string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, id)
}

func (c *UserCache) LoadInitial(users []*domain.User) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for _, u := range users {
		c.data[u.ID] = u
	}
}
