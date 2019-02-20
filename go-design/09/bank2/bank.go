package bank

var (
	sema = make(chan struct{}, 1) // 用来保护 balance 的信号量
	balance int
)

func Deposit(amount int)  {
	sema <- struct{}{} // 获取令牌
	balance = balance + amount
	<-sema // 释放令牌
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b

}