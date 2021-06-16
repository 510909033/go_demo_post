package my_etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"log"

	"testing"
	"time"
)

func getCli() (*clientv3.Client, error) {
	log.SetFlags(log.Lshortfile)
	return clientv3.New(clientv3.Config{
		// 集群列表
		//Endpoints:   []string{IP},
		Endpoints: []string{IP, IP1, IP2},
		//Endpoints:   []string{IP2},
		DialTimeout: 5 * time.Second,
	})
}

func getCliV2() *clientv3.Client {
	log.SetFlags(log.Lshortfile)
	client, err := clientv3.New(clientv3.Config{
		// 集群列表
		//Endpoints:   []string{IP},
		Endpoints: []string{IP, IP1, IP2},
		//Endpoints:   []string{IP2},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return client
}

//监听 name
//etcdctl put name test_watch
//etcdctl put name test_watch
func TestEtcdWatch(t *testing.T) {
	name := "test_watch"
	log.Println("name=", name)
	cli, err := getCli()
	if err != nil {
		t.Error(err)
	}

	defer func() {
		log.Println("defer")
		err := cli.Close()
		log.Println("defer, err=", err)

	}()

	// 监听值
	go func() {
		watch := cli.Watch(ctx, name)
		for {
			res := <-watch
			t.Logf("name发生改变, res=%+v", res)
			log.Println(res.Events[0].Type)             //PUT
			log.Println(string(res.Events[0].Kv.Value)) //v3
			time.Sleep(time.Second)
		}
	}()

	select {}
}

/*
etcdctl --endpoints=$ENDPOINTS put web1 value1
etcdctl --endpoints=$ENDPOINTS put web2 value2
etcdctl --endpoints=$ENDPOINTS put web3 value3

etcdctl --endpoints=$ENDPOINTS get web --prefix
*/
func TestEtcdPrefix(t *testing.T) {
	name := "web"
	log.Println("前缀测试， 前缀=", name)
	cli, err := getCli()
	if err != nil {
		t.Error(err)
	}

	defer func() {
		log.Println("defer")
		err := cli.Close()
		log.Println("defer, err=", err)

	}()

	response, err := cli.Get(ctx, "web", clientv3.WithPrefix())
	if err != nil {
		log.Printf("err=%+v", err)
		return
	}
	for k, v := range response.Kvs {
		log.Printf("k=%d, key=%s, val=%s", k, v.Key, v.Value)
	}

}

func TestEtcdTransaction(t *testing.T) {
	log.Println("事务测试")
	cli, err := getCli()
	if err != nil {
		t.Error(err)
	}

	defer func() {
		log.Println("defer")
		err := cli.Close()
		log.Println("defer, err=", err)

	}()

	//txn := cli.Txn(ctx)

	//txn.If()

}

/*

etcdctl --endpoints=$ENDPOINTS lease -h
COMMANDS:
	grant		Creates leases
	keep-alive	Keeps leases alive (renew)
	list		List all active leases
	revoke		Revokes leases
	timetolive	Get lease information
*/
func TestEtcdLease_create(t *testing.T) {
	cli := getCliV2()
	defer func() {
		cli.Close()
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	_ = cancel
	// minimum lease TTL is 5-second
	ttl := 5
	resp, err := cli.Create(ctx, int64(ttl))
	if err != nil {
		log.Printf("resp=%+v", resp)
		log.Fatal(err)
	}

	// after 5 seconds, the key 'foo' will be removed
	_, err = cli.Put(context.TODO(), "foo", "bar", clientv3.WithLease(clientv3.LeaseID(resp.ID)))
	if err != nil {
		log.Fatal(err)
	}
}
