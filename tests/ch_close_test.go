package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChanClose(t *testing.T) {
	ch := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		ch <- struct{}{}
		fmt.Println("over")
	}()
	go func() {
		defer wg.Done()
		ch <- struct{}{}
		fmt.Println("over")
	}()
	time.Sleep(time.Second)
	close(ch)
	wg.Wait()
}
