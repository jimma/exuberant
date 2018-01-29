package main

import (
	"fmt"
	"sync"

	"github.com/golang/glog"
	"github.com/jimma/exuberant/lang/concurrency"
)

func main() {

	//
	var wg sync.WaitGroup
	myBalance := concurrency.NewBalance("0001", 0)

	myBalance.InitBalance(30)
	fmt.Printf("init balance is : %d \n", myBalance.Get())
	//the equailent thing with "java thread.join"
	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			myBalance.Expense(10)
		}()
		go func() {
			defer wg.Done()
			myBalance.Income(10)
		}()

	}

	wg.Wait()
	fmt.Printf("Current balance is : %d \n", myBalance.Get())

	myBalance2 := concurrency.NewBalance("0002", 0)
	myBalance2.InitBalance(100)
	fmt.Printf("init Balance2 is : %d \n", myBalance2.Get())
	done := make(chan bool)
	for i := 0; i < 50; i++ {
		go func() {
			myBalance2.Expense(10)
			done <- true
		}()
		go func() {
			myBalance2.Income(10)
			done <- true
		}()

	}
	for j := 0; j < 100; j++ {
		<-done
	}
	fmt.Printf("Current Balancer2 is : %d \n", myBalance2.Get())
	glog.Flush()
}
