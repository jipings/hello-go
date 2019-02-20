package memo

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex // 保护cache
	cache map[string]result
}

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	// 主调 goroutine 会分两次获取锁， 第一次用于查询，第二次用于在查询无返回结果时进行更新
	// 在这两次之间，其他 gorutine 也可以使用缓存
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}