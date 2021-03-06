package redis

import (
	"sync"
	"time"
)

var NewsCachePool *sync.Pool

func init() {
	NewsCachePool = &sync.Pool{
		New: func() interface{} {
			return NewSimpleCache(NewStringOperation(), WithExpire(15),
				SERILIZER_JSON,
				NewCrossPolicy("^/news/\\d{1,5}$", time.Second * 30))
		},
	}
}

func NewsCache() *SimpleCache {
	return NewsCachePool.Get().(*SimpleCache)
}

func ReleaseNewsCache(cache *SimpleCache)  {
	NewsCachePool.Put(cache)
}