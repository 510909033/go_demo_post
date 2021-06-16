package multi_groutinue2

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func demo1() {
	dataCh := make(chan int, 0)
	overCh := make(chan int, 1)

	go func() {
		defer func() {
			overCh <- 1
		}()
		close(dataCh)
		log.Println(dataCh == nil)
		select {
		case <-dataCh:
			log.Println("recive dataCh")
		case <-overCh:
		default:
			log.Println("default")
		}
	}()

	select {
	case val := <-overCh:
		log.Println("overCh:", val)
	}
}

func demo2() {
	//timer := time.NewTimer(time.Second * 2)
	timer := time.NewTimer(time.Second * -1)

	i := 4
	_ = i
	for {
		select {
		case <-timer.C:
			log.Println("trigger timer")
			//default:
			//	if !timer.Stop() {
			//		select {
			//		case <-timer.C:
			//			log.Println("select timer.C")
			//		default:
			//			log.Println("select default")
			//		}
			//	}
			//
			//	timer.Reset(time.Second * time.Duration(i))
			//	log.Println("reset ", i)
			//	i--
			//	time.Sleep(time.Second * 2)
		}
	}

	return

	ch1 := make(chan int, 4)
	ch1 <- 30
	ch1 <- 40
	ch1 <- 50
	ch2 := ch1
	ch2 <- 100
	log.Println(len(ch1), ch1)
	log.Println(ch2)
	for v := range ch2 {
		log.Println(v)
	}

}

func DemoMultiGroutinue2() {
	demo2()
	//demo1()
	return

	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the moderator goroutine shown below.
	// Its reveivers are all senders and receivers of dataCh.
	toStop := make(chan string, 1)
	// the channel toStop is used to notify the moderator
	// to close the additional signal channel (stopCh).
	// Its senders are any senders and receivers of dataCh.
	// Its reveiver is the moderator goroutine shown below.

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop // part of the trick used to notify the moderator
		// to close the additional signal channel.
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				if value == 0 {
					// here, a trick is used to notify the moderator
					// to close the additional signal channel.
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// the first select here is to try to exit the
				// goroutine as early as possible.
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// same as senders, the first select here is to
				// try to exit the goroutine as early as possible.
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == MaxRandomNumber-1 {
						// the same trick is used to notify the moderator
						// to close the additional signal channel.
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}
