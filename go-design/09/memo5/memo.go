
package memo
// Func 用于缓存的函数类型
type Func func(key string) (interface{}, error)
// result 是调用 Func 的返回结果
type result struct {
	value interface{}
	err error
}

type entry struct {
	res result
	ready chan struct{} // 当res 准备好后关闭该通道
}
// request 是一条请求消息， key需要用Func来调用
type request struct {
	key string
	response chan<- result // 客户端需要单个 result
}

type Memo struct {requests chan request}
// New 返回 f 的函数记忆，客户端之后需要调用 Close
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests  <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// 对这个 key 的第一次请求
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e;
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// 执行函数
	e.res.value, e.res.err = f(key)
	// 通知数据已准备完毕
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// 等该数据准备完毕
	<-e.ready
	// 向客户端发送结果
	response <- e.res
}