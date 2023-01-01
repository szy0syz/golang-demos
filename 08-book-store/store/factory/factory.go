package factory

import (
	"bookstore/store"
	"fmt"
	"sync"
)

var (
	providersMu sync.RWMutex
	providers   = make(map[string]store.Store)
)

func Register(name string, p store.Store) {
	providersMu.Lock()
	defer providersMu.Unlock()

	if p == nil {
		panic("store: Register provider is nil")
	}

	if _, dup := providers[name]; dup {
		panic("store: Register called twice for provider " + name)
	}

	providers[name] = p
}

func New(providerName string) (store.Store, error) {
	// 上读锁
	providersMu.RLock()

	// ---- 意味着这句话是有数据并行风险的
	// ---- 啥意思呢？就是不能要 "脏读" 的意思，不能做无赖
	p, ok := providers[providerName]

	// 解开读锁
	providersMu.RUnlock()

	if !ok {
		// 噢哟，原来这样也能生产一个error出去
		return nil, fmt.Errorf("store: unkonow provider %s", providerName)
	}

	return p, nil
}
