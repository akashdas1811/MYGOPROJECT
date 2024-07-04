package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance(wg *sync.WaitGroup) *single {
	defer wg.Done()
	if singleInstance == nil {

		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go getInstance(&wg)
	}
	wg.Wait()

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
	time.Sleep(1 * time.Second)
}

//----USING MUTEX LOCK------//

var lock = &sync.Mutex{}

type single1 struct {
}

var singleInstance1 *single1

func getInstance1() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
