package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func War() {

	start := time.Now()
	evilNinjas := []string{"Surya", "Sai", "Ritesh", "Keerthi"}
	defer func() {
		fmt.Println(time.Since(start))
	}()

	// wg.Add(4)
	for _, name := range evilNinjas {
		go Attack(name)

	}
	//  wg.Wait()
	time.Sleep(time.Second * 1)
}

func Attack(name string) {
	fmt.Println("Attack By ", name)
	// wg.Done()
	time.Sleep(time.Second)
}
