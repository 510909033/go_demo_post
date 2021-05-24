package my_bucket_v1

import (
	"fmt"
	"github.com/juju/ratelimit"
	_ "github.com/juju/ratelimit"
	"log"
	"time"
)

func DebugMyBucketV1() {
	demo2()
	return

	var fillInterval = time.Millisecond * 10
	var capacity = 100
	var tokenBucket = make(chan struct{}, capacity)

	fillToken := func() {
		ticker := time.NewTicker(fillInterval)
		for {
			select {
			case <-ticker.C:
				select {
				case tokenBucket <- struct{}{}:
				default:
				}
				fmt.Println("current token cnt:", len(tokenBucket), time.Now())
			}
		}
	}

	go fillToken()
	time.Sleep(time.Hour)
}

func demo2() {
	//	github.com/juju/ratelimit

	//bucket := ratelimit.NewBucket(time.Second, 10)
	//ratelimit.NewBucketWithRateAndClock()
	//clock:= ratelimit.Clock()
	var clock ratelimit.Clock
	bucket := ratelimit.NewBucketWithQuantumAndClock(time.Second, 10, 1, clock)
	//bucket.Available()

	go func() {
		for {
			fmt.Println("Available-1", bucket.Available())
			fmt.Println("Available-2", bucket.Available())
			//fmt.Println("bucket.TakeAvailable", bucket.TakeAvailable(900))
			time.Sleep(time.Second)
		}
	}()

	for {
		//bucket.Wait(2)
		//duration := bucket.Take(71)
		//log.Println("duration = ", duration)

		//log.Println(bucket.TakeAvailable(70))

		//如果不足，不做操作
		log.Println(bucket.WaitMaxDuration(9, time.Second))

		log.Println(time.Now())
		time.Sleep(time.Second * 5)
		//break
	}

}
