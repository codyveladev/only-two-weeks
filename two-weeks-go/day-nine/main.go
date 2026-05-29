package main

import (
	"fmt"
	"math/rand"
	"sync"

	//"sync"
	"time"

	"github.com/codyveladev/day-nine/models"
)

func main() {
	var wg sync.WaitGroup
	c := models.NewCache()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", rand.Intn(5))
			switch rand.Intn(3) {
			case 0:
				c.Set(key, fmt.Sprintf("value-%d", id), time.Second)
			case 1:
				c.Get(key)
			case 2:
				c.Delete(key)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("done")
}
