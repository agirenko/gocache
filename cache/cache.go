package cache

type Cc struct {
	cMap map[string]interface{}
}

func New() *Cc {
	return &Cc{
		cMap: make(map[string]interface{}),
	}
}

func (cache *Cc) Set(key string, value interface{}) {
	cache.cMap[key] = value
}

func (cache Cc) Get(key string) interface{} {
	return cache.cMap[key]
}

func (cache *Cc) Delete(key string) bool {
	_, exists := cache.cMap[key]
	if exists {
		delete(cache.cMap, key)
		return true
	}
	return false
}
