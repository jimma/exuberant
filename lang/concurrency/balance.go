package concurrency

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"
)

type balanceType struct {
	AccountNo string
	Balance   int
	Mutex     sync.Mutex
}

func init() {
	flag.Parse()
}

func NewBalance(accountNo string, balance int) balanceType {
	var mutex = sync.Mutex{}
	myBalance := balanceType{accountNo, balance, mutex}
	return myBalance
}

func (b *balanceType) InitBalance(amount int) int {
	glog.V(2).Infoln("init balance\n")
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	b.Balance = amount
	fmt.Println(b.Balance)
	return b.Balance
}

func (b *balanceType) Income(amount int) int {
	glog.V(2).Infoln("Balance Income\n")
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	time.Sleep(700)
	b.Balance = b.Balance + amount
	return b.Balance
}

func (b *balanceType) Expense(amount int) int {
	glog.V(2).Infoln("Balance expense\n")
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	time.Sleep(300)
	b.Balance = b.Balance - amount
	return b.Balance
}

func (b *balanceType) Get() int {
	return b.Balance
}
