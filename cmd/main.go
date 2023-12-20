package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/UtkarshM-hub/Hex/internal/application/api"
	"github.com/UtkarshM-hub/Hex/internal/application/core"
)

func main() {
	t:=time.Now()
	// initialize one wait group
	var wg sync.WaitGroup

	// getting instance of arithmetic struct
	Arithmetic := core.New()

	// dependency injection
	API := api.New(Arithmetic)

	// create a channel to receive results
	resultChannel := make(chan int32, 10)

	// run a for loop from i to 10000, call GetAddition, and send result to channel
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val, err := API.GetAddition(int32(i), int32(i))
			if err != nil {
				log.Fatal(err)
			}
			resultChannel <- val
		}(i)
	}

	// close the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// receive and print results from the channel
	for result := range resultChannel {
		fmt.Println(result)
	}
	fmt.Println("Time: ",time.Since(t))
}
