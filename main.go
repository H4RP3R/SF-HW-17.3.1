// Напишите программу, аналогичную той, что мы только что написали, однако она
// должна использовать уже не 1000 горутин, а только 10.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	step            int64 = 1
	endCounterValue int64 = 1000
	gNum            int   = 10
)

func main() {

	var counter int64 = 0
	var wg sync.WaitGroup

	increment := func() {
		defer wg.Done()
		for {
			current := atomic.LoadInt64(&counter)
			if current >= endCounterValue {
				return
			}
			atomic.CompareAndSwapInt64(&counter, current, current+step)
		}
	}

	for i := 0; i < gNum; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Println(counter)
}
