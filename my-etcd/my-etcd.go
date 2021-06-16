package my_etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"log"
	"os"
	"strconv"
	"strings"

	"time"
)

var errlog = log.New(os.Stdout, " [ERROR] ", log.Lshortfile)
var infolog = log.New(os.Stdout, " [INFO] ", log.Lshortfile)

var (
	ctx = context.TODO()
	//IP  = "192.168.6.5:2379"
	//IP  = "192.168.6.5:2379"
	//IP1 = "192.168.6.5:22379"
	//IP2 = "192.168.6.5:32379"
	IP  = "172.20.10.40:2379"
	IP1 = "172.20.10.40:22379"
	IP2 = "172.20.10.40:32379"
)

func NewClient() *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		// 集群列表
		//Endpoints:   []string{IP},
		Endpoints: []string{IP, IP1, IP2},
		//Endpoints:   []string{IP2},
		DialTimeout: 51 * time.Second,
	})
	if err != nil {
		errlog.Println(err)
		return nil
	}
	return cli
	//defer cli.Close()
}

func Monitor(cli *clientv3.Client) {

}

func Watch(ctx context.Context, cli *clientv3.Client, key string) {
	infolog.Println("开始 Watch")
	watch := cli.Watch(ctx, key)
	for {
		res := <-watch
		//t.Logf("name发生改变, watchRes=%#v", res)
		infolog.Printf("Watch method, name发生改变, res=%+v,", res)
		time.Sleep(time.Second)
	}
}

func ModifyKey(ctx context.Context, cli *clientv3.Client, key string, val string) {
	cli.Put(ctx, key, val)
	//cli.Sync(ctx)
	//if resp, err := cli.Put(ctx, "name", "pibigstar", clientv3.WithPrevKV()); err != nil {
	//if resp, err := cli.Put(ctx, "name", "pibigstar", clientv3.WithPrevKV()); err != nil {
	//	t.Error(err)
	//} else {
	//	t.Log("旧值: ", string(resp.PrevKv.Value))
	//}
}

func AutoSetDirSubVal(ctx context.Context, cli *clientv3.Client, dir string) {
	if strings.LastIndex(dir, "/") != 0 {
		dir = dir + "/"
	}

	for {
		time.Sleep(time.Second * 5)
		key := dir + strconv.FormatUint(uint64(time.Now().Unix()), 10)
		putResponse, err := cli.Put(ctx, key, time.Now().String())
		infolog.Printf("key=%s", key)
		if err != nil {
			errlog.Printf("cli.Put err=%+v\n", err)
			return
		}
		infolog.Printf("putResponse=%+v", putResponse)

		getResponse, err := cli.Get(ctx, key)
		if err != nil {
			errlog.Printf("cli.Get err=%+v\n", err)
			return
		}
		infolog.Printf("key=%s, getResponse=%+v", key, getResponse)
	}

}

func DemoMyEtcd() {
	cli := NewClient()
	key := "h5"
	_ = key
	go Watch(ctx, cli, key)
	go func() {
		for {
			ModifyKey(ctx, cli, key, time.Now().String())
			time.Sleep(time.Second)
		}
	}()

	//go AutoSetDirSubVal(ctx, cli, "mydir1")

	select {}

}
