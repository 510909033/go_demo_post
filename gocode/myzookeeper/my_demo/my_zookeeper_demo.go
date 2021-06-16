package my_demo

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"os"
	"time"
)

var (
	hosts       = []string{"192.168.6.3:2181"}
	path        = "/demo22"
	flags int32 = zk.FlagSequence
	data        = []byte("zk data 001")
	acls        = zk.WorldACL(zk.PermAll)
)

var conn *zk.Conn
var err error

func Exception(err error) {
	if err != nil {
		panic(err)
	}
}
func Start() {
	//defer conn.Close()
	_conn()

	log.Printf("conn.SessionID()=%d\n", conn.SessionID())

	//s,err:=conn.Create(path,data,zk.FlagSequence,acls)
	//Exception(err)
	//log.Printf("create res=%s\n",s)

	//testRace()
	getChild()
	time.Sleep(time.Second)

	//for {
	//	_,stat,ch,err:=conn.GetW("/demo/d1")
	//	fmt.Println(len(ch),err)
	//	evCh :=<-ch
	//	log.Printf("evCh=%#v\n", evCh, stat)
	//}
}

func getChild() {
	list, stats, err := conn.Children("/me/1/11")
	fmt.Println(list, stats, err)

}

func testRace() {
	ch := make(chan int, 10)
	var i uint64
	for {
		ch <- 1
		go func() {
			defer func() {
				<-ch
			}()
			_, err := conn.Create(path, data, zk.FlagEphemeral, acls)
			if err != nil {
				//
				return
			}
			Exception(err)
			_, stats, err := conn.Get(path)
			_ = stats
			Exception(err)
			defer func() {
				//fmt.Println("delete")
				err := conn.Delete(path, stats.Version)
				Exception(err)
			}()

			i++
			if i%100 == 0 {
				fmt.Println(i)
			}
		}()

		//break
	}
}

func _conn() {
	// 创建监听的option，用于初始化zk
	eventCallbackOption := zk.WithEventCallback(callback)

	zk.WithLogger(log.New(os.Stdout, "zi_", log.Lshortfile))
	zk.WithLogInfo(true)
	//eventCallbackOption = nil
	// 连接zk
	conn, _, err = zk.Connect(hosts, time.Second*30, eventCallbackOption)

	if err != nil {
		panic(err)
		return
	}
}

func callback(event zk.Event) {
	// zk.EventNodeCreated
	// zk.EventNodeDeleted
	fmt.Println("--Start--###########################")
	fmt.Println("path: ", event.Path)
	fmt.Println("type: ", event.Type.String())
	fmt.Println("state: ", event.State.String())
	fmt.Println("---------------------------\n\n")
}
