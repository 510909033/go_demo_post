package my_race

import (
	"log"
	"time"
)

func DemoSliceRace() {
	list := make([]int, 100)
	ch := make(chan struct{}, 10)

	i := 0
	for {
		i++
		if i > 99 {
			i = 0
		}
		ch <- struct{}{}
		go func(i int) {
			//i = 10
			defer func() {
				<-ch
			}()
			list[i] = i
			log.Printf("i=%d", i)
			time.Sleep(time.Second)
		}(i)

	}
}
