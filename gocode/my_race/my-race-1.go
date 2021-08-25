package my_race

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	value int
	mtx   *sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{0, &sync.Mutex{}}
}

func (c *Counter) inc() {
	c.mtx.Lock()
	c.value++
	c.mtx.Unlock()
}

func (c Counter) get() int {
	c.mtx.Lock()
	res := c.value
	c.mtx.Unlock()
	return res
}

type User struct {
	mu sync.Mutex
}

func (u *User) String() string {
	return "user string"
}

func main() {
	var wg sync.WaitGroup
	counter := NewCounter()
	max := 100
	max = 1
	wg.Add(max)

	// consumer
	go func() {
		for i := 0; i < max; i++ {
			value := counter.get()
			fmt.Printf("counter value = %d\n", value)
			wg.Done()
		}
	}()
	// producer
	go func() {
		for i := 0; i < max; i++ {
			counter.inc()
		}
	}()

	wg.Wait()

	u := &User{}
	go func() {
		for {
			u.mu.Lock()
			fmt.Sprintf("111-user=%+v, %+v, %s\n", u, u, u)
			u.mu.Unlock()
		}
	}()
	go func() {
		for {
			u.mu.Lock()
			fmt.Sprintf("222-user=%+v, %+v, %s\n", u, u, u)
			u.mu.Unlock()
		}
	}()

	go func() {
		for {
			u.mu.Lock()
			fmt.Sprintf("222-user=%+v, %+v, %s\n", u, u, u)
			u.mu.Unlock()
		}
	}()
	time.Sleep(time.Second * 3)
}
