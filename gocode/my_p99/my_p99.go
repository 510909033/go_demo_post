package my_p99

import (
	"log"
	"time"
)

const BUCKET_SIZE = 135

type P99Log interface {
}

type p99Monitor struct {
	Name        string
	observeChan chan int64
	Bucket      []int // 0- 0秒， 999表示999毫秒
	Total       int64
}

func NewP99() *p99Monitor {
	p99m := &p99Monitor{
		observeChan: make(chan int64, 100),
		Bucket:      make([]int, BUCKET_SIZE),
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				//todo log
				log.Println(err)
			}
		}()
		p99m.monitor()
	}()
	return p99m
}

//val 毫秒
func (m *p99Monitor) Observe(val int64) {
	m.observeChan <- val
}

func getIndex(val int64) int {
	if val < 1 {
		return 0
	}
	if val < 1000 { // 99个桶
		return int(val / 10) //[0,99]
	}
	if val < 2000 { // 20个桶
		//50毫秒一个桶
		return int((val-1000)/50) + 100 //[100, 119]
	}
	if val < 5000 { // 15个桶
		return int((val-2000)/200) + 120 //[120, 134]
	}
	return 135
}

func convertKToMilis(k int) int {
	if k == 0 {
		return 0
	}
	if k < 100 {
		return k * 10
	}
	if k < 120 {
		return (k-100)*50 + 1000
	}
	if k < 125 {
		return (k-120)*200 + 2000
	}
	return 5000
}

func (m *p99Monitor) monitor() {
	ticker := time.NewTicker(time.Millisecond * 1000)
	debugTicker := time.NewTicker(time.Millisecond * 1000)

	prevTime := time.Now()
	for {
		select {
		case <-debugTicker.C:
		case <-ticker.C:
			if time.Now().Sub(prevTime).Seconds() >= 3 {
				//计算
				p95Cnt := int(0.95 * float64(m.Total))
				p99Cnt := int(0.99 * float64(m.Total))

				var p95, p99 = -1, -1

				cnt := 0
				for k, v := range m.Bucket {
					cnt += v
					if cnt >= p95Cnt && p95 < 0 {
						p95 = k
					}
					if cnt >= p99Cnt && p99 < 0 {
						p99 = k
					}
				}
				p99 = convertKToMilis(p99)
				p95 = convertKToMilis(p95)

				log.Printf("total=%d, p95Cnt=%d, p99=%d", m.Total, p95, p99)
				log.Println(m.Bucket)

				//将当前p99写入日志
				//log.Printf("p95Cnt=%.8f, p99=%.08f", m.digest.Quantile(0.95), m.digest.Quantile(0.99))
				//time.Sleep(time.Millisecond * 300) //todo
				//m.digest = tdigest.New()
				prevTime = time.Now()
				m.Bucket = make([]int, BUCKET_SIZE)
				m.Total = 0
			}
		case val := <-m.observeChan:
			m.Bucket[getIndex(val)]++
			m.Total++
		}
	}

}
