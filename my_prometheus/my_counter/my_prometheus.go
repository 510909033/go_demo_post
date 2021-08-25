package my_counter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

/*

下列查询将生成在 5 分钟间隔内每秒钟处理的任务数。
sum by (type) (rate(worker_jobs_processed_total[5m]))

查询当前系统中，访问量前10的HTTP地址：
topk(10, http_requests_total)

promhttp_metric_handler_requests_total - promhttp_metric_handler_requests_total offset 1h

假设我们要查过去一小时内每5分钟有多少次请求可以这么写。
increase(promhttp_metric_handler_requests_total[5m])

线上的每一个点就是当前时间上所有code和handler对应指标的总和
sum(increase(promhttp_metric_handler_requests_total[5m]))

当前时间上该code和handler对应指标的总和
sum(increase(promhttp_metric_handler_requests_total[5m])) by (code)
sum(increase(worker_jobs_processed_total[5m])) by (code)

sum(increase(worker_jobs_processed_total[5m])) / (300)   等于  sum(rate(worker_jobs_processed_total[5m]))



*/

type Job struct {
	Type string
}

var (
	totalCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "worker",
			Subsystem: "jobs",
			Name:      "processed_total",
			Help:      "Total number of jobs processed by the workers",
		},
		// We will want to monitor the worker ID that processed the
		// job, and the type of job that was processed
		//这个内容必须提前定义好，
		//如果 WithLabelValues 传的数量不匹配 ，panic
		// 如果 With 传的labels不匹配，panic
		[]string{"method", "code"},
	)
)

func init() {
	//...
	// register with the prometheus collector
	prometheus.MustRegister(totalCounterVec)
	//...
}

func MyPrometheus() {

	jobChannel := make(chan *Job, 10000)

	go startWorker("woker1", jobChannel)
	go createJobs(jobChannel)

	log.Println("starting server")
	server()
	//time.Sleep(time.Second)
	//log.Println("started server")
}

func server() {
	// create a new mux server
	server := http.NewServeMux()
	// register a new handler for the /metrics endpoint
	server.Handle("/metrics", promhttp.Handler())
	// start an http server using the mux server
	log.Fatalln(http.ListenAndServe(":9002", server))
}

func createJobs(jobChannel chan<- *Job) {
	go methodDemo()

	types := []string{"type1", "type2"}
	for {
		for _, v := range types {
			jobChannel <- &Job{Type: v}
			countInc()
		}
		time.Sleep(time.Second)
	}
}

func startWorker(workerID string, jobs <-chan *Job) {
	for {
		select {
		case job := <-jobs:
			totalCounterVec.WithLabelValues(workerID, job.Type).Inc()
		}
	}
}

func countInc() {
	nano := time.Now().UnixNano()
	i := int(nano % 9)
	switch {
	case i < 2:
		totalCounterVec.WithLabelValues("get", "404").Inc()
	case i < 5:
		totalCounterVec.WithLabelValues("get", "200").Inc()
	case i < 7:
		totalCounterVec.WithLabelValues("post", "200").Inc()
	case i < 8:
		totalCounterVec.WithLabelValues("post", "500").Inc()
	}
	totalCounterVec.With(prometheus.Labels{"code": "404", "method": "put"}).Add(42)
	return
}

func methodDemo() {
	//20毫秒  1秒钟50个
	ticker := time.NewTicker(time.Millisecond * 20)
	for {
		select {
		case <-ticker.C:
			totalCounterVec.With(prometheus.Labels{"code": "302", "method": "demo"}).Add(1)
		}

	}
}
