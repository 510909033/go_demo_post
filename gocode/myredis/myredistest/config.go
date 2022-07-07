package myredistest

import (
	"github.com/go-redis/redis"
	"log"
	"sync/atomic"
	"time"
)

/*

密码错误：
	WRONGPASS invalid username-password pair or user is disabled.


*/

//const ADDR = "192.168.6.3:6379"
const (
	//ADDR = "172.17.0.7:6379"
	ADDR = "172.20.10.40:36383"
	//ADDR     = "172.20.10.40:6379"
	PASSWORD = "bitnami"
	//PASSWORD = "bitnami_fail"
)

var addrs = []string{
	"172.20.10.40:36381",
	"172.20.10.40:36382",
	"172.20.10.40:36384",
}

var client *redis.ClusterClient
var debugA = int64(0) //用于调试 OnConnect

func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	log.SetFlags(log.Lshortfile)
	log.Println("addr", ADDR)
	//var opt = &redis.Options{
	//	Network:            "tcp",
	//	Addr:               ADDR,
	//	Dialer:             nil,
	//	OnConnect:          nil,
	//	Password:           PASSWORD,
	//	DB:                 0,
	//	MaxRetries:         0,
	//	MinRetryBackoff:    0,
	//	MaxRetryBackoff:    0,
	//	DialTimeout:        0,
	//	ReadTimeout:        0,
	//	WriteTimeout:       0,
	//	PoolSize:           0,
	//	MinIdleConns:       0,
	//	MaxConnAge:         0,
	//	PoolTimeout:        0,
	//	IdleTimeout:        0,
	//	IdleCheckFrequency: 0,
	//	TLSConfig:          nil,
	//}
	//_ = opt
	//client = redis.NewClient(opt)

	client = GetClusterClient()
	//client1 = GetClusterClient()
}

func getOpts1() *redis.ClusterOptions {
	return &redis.ClusterOptions{
		Addrs: addrs,
		//MaxRedirects:       0,
		//ReadOnly:           false,
		//RouteByLatency:     false,
		//RouteRandomly:      false,
		//ClusterSlots:       nil,
		//OnNewNode:          nil,
		OnConnect: func(conn *redis.Conn) error {
			old := atomic.AddInt64(&debugA, 1)
			if old == 54 {
				log.Println("debug OnConnect ")
				panic(old)
			}
			log.Printf("OnConnect, conn=%+v", conn)
			return nil
		},
		Password: PASSWORD,
		//MaxRetries:         0,
		//MinRetryBackoff:    0,
		//MaxRetryBackoff:    0,
		DialTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		//PoolSize:           0,
		//MinIdleConns:       0,
		//MaxConnAge:         0,
		//PoolTimeout:        time.Second,
		IdleTimeout: time.Second * 60,
		//IdleCheckFrequency: 0,
		//TLSConfig:          nil,
	}
}

func getOpts2() *redis.ClusterOptions {
	return &redis.ClusterOptions{
		Addrs: addrs,
		//MaxRedirects:       0,
		//ReadOnly:           false,
		//RouteByLatency:     false,
		//RouteRandomly:      false,
		//ClusterSlots:       nil,
		OnNewNode: func(newClient *redis.Client) {
			log.Printf("OnNewNode, %+v \n", newClient.String())
		},
		OnConnect: func(conn *redis.Conn) error {
			old := atomic.AddInt64(&debugA, 1)
			if old == 84 {
				log.Println("debug OnConnect ")
				//panic(old)
			}
			//log.Printf("OnConnect, conn=%+v", conn)
			return nil
		},
		Password: PASSWORD,
		//MaxRetries:         0,
		//MinRetryBackoff:    0,
		//MaxRetryBackoff:    0,
		DialTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,

		//NODE
		PoolSize:     20,
		MinIdleConns: 5,
		MaxConnAge:   time.Second * 200,
		PoolTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
		//IdleCheckFrequency: 0,
		//TLSConfig:          nil,
	}
}

func GetClusterClient() *redis.ClusterClient {

	opts := getOpts2()

	client := redis.NewClusterClient(opts)
	ping := client.Ping()
	//WRONGPASS invalid username-password pair or user is disabled.
	if ping.Err() != nil {
		panic("ping报错， err=" + ping.Err().Error())
	}
	log.Println(ping.Args())

	log.Printf("opts=%#v\n", opts)

	clusterNodes := client.ClusterNodes()

	log.Println(clusterNodes)

	//	usage := client.MemoryUsage("a")
	//	log.Println(usage)
	//
	//	log.Println(client.ClientGetName())
	//	log.Println(client.ClientID())
	//	log.Println(client.ClientList())
	//	log.Println(client.ClusterSlots())
	//	log.Println(client.Time())
	//	log.Println(client.DBSize())
	//	log.Println(client.Scan(0, "", 5))
	//	script := `local name=redis.call("get", KEYS[1])
	//local greet=ARGV[1]
	//local result=greet.." "..name
	//return result`
	//	log.Println(client.ScriptLoad(script))
	//	log.Println(client.Eval(script, []string{"b"}, "haha"))
	//	log.Println(client.ClusterSlots())
	//	log.Println(client.ClusterKeySlot("aaa"))

	go monitorRedisInfo()

	return client
}

func NewClientConn() <-chan int {
	ch := make(chan int, 0)
	for i := 0; i < 10; i++ {
		GetClusterClient()
	}
	return ch
}

func monitorRedisInfo() {
	callback := func() {
		defer func() {
			recover()
		}()

		//clusterInfo := client.ClusterInfo()
		//
		//debug(clusterInfo)
		//log.Println(clusterInfo.String())
		//
		//info := client.Info("all")
		//debug(info)
		//log.Println(info.String())

		log.Printf("PoolStats=%+v", client.PoolStats())

	}
	_ = callback

	ticker := time.NewTicker(time.Second)
	log.Print("monitorRedisInfo\n")
	for {
		select {
		case <-ticker.C:
			//todo
			callback()
		}
	}

}
