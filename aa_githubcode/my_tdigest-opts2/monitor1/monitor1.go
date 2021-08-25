package monitor1

import (
	tdigest "go_demo_post/aa_githubcode/my_tdigest-opts2"
	"log"
	"time"
)

type Monitor struct {
	digest      *tdigest.TDigest
	observeChan chan float64
}

func NewMonitor() *Monitor {
	monitor := &Monitor{
		digest:      tdigest.New(),
		observeChan: make(chan float64, 1000),
	}
	go monitor.monitor()
	return monitor
}

func (m *Monitor) Observe(val float64) {
	m.observeChan <- val
}

func (m *Monitor) monitor() {
	ticker := time.NewTicker(time.Millisecond * 100)
	debugTicker := time.NewTicker(time.Millisecond * 1000)

	prevTime := time.Now()
	_ = prevTime
	for {
		select {
		case <-debugTicker.C:
			log.Println(m.digest.Dump())
		case t := <-ticker.C:
			if time.Now().Sub(prevTime).Seconds() >= 300 {
				//将当前p99写入日志
				log.Println(t)
				log.Printf("p95=%.8f, p99=%.08f", m.digest.Quantile(0.95), m.digest.Quantile(0.99))
				time.Sleep(time.Millisecond * 300) //todo
				m.digest = tdigest.New()
				prevTime = time.Now()
			}
		case val := <-m.observeChan:
			m.digest.Add(val, 1)
		}
	}

}
