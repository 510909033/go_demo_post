package myredistest

import (
	"fmt"
	"sync"
	"time"
)

func Set() {
	key := fmt.Sprintf("myredistest_set_%d", time.Now().UnixNano())
	value := 1
	expiration := time.Duration(time.Second * 30)
	statusCmd := client.Set(key, value, expiration)
	if statusCmd.Err() != nil {
		fmt.Printf("set err, %s\n", statusCmd.Err().Error())
	}
	//
	_ = statusCmd
}

func SetMulti(max int) {
	t := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < max; i++ {
		wg.Add(1)
		func() {
			defer wg.Done()
			Set()

		}()
	}

	wg.Wait()
	fmt.Printf("SetMulti time:%s\n", time.Since(t))
}

var incrSyncM sync.Map
var w sync.WaitGroup

func GetId() {
	key := fmt.Sprintf("%d", time.Now().Unix())
	for i := 0; i < 3; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			cmd := client.Incr(key)
			if cmd.Err() != nil {
				fmt.Println(cmd.Err())
				return
			}
			id := cmd.Val()
			v, ok := incrSyncM.Load(id)
			//			v, ok = incrSyncM.Load(id)
			if ok {
				panic(v)
			}
			//			fmt.Println(id)
			incrSyncM.Store(id, id)
		}()
		//		time.Sleep(time.Millisecond * 10)
	}

	w.Wait()
	incrSyncM.Range(func(key, val interface{}) bool {
		fmt.Println(key, val)
		return true
	})
	v, ok := incrSyncM.Load(int64(1))
	fmt.Println(v, ok)
}

func AddData(val int) {
	key := fmt.Sprintf("cache_key_%d", val)
	fields := make(map[string]interface{})
	fields["status"] = 1
	fields["update_ts"] = time.Now().UnixNano()
	fields["value"] = val
	statusCmd := client.HMSet(key, fields)
	fmt.Printf("statusCmd-1=%#v\n", statusCmd)

	var saveToDbRet = false
	if statusCmd.Err() == nil {
		//save to db
		saveToDbRet = true
	}

	if saveToDbRet {
		statusCmd := client.HSet(key, "status", 2)
		fmt.Printf("statusCmd-2=%#v\n", statusCmd)
	}

	fmt.Println(client.HGetAll(key).Val())
}

//
//func GetData(val int) {
//	key:= fmt.Sprintf("cache_key_%d", val)
//
//	ret:=client.HGetAll(key)
//
//	if ret.Val()["status"] == "1" {
//		//无效的缓存，因为存储到db失败了， 或者存储db成功，更新redis的status失败了
//		fmt.Println("status=1 ")
//		//分布式锁
//		//lock
//		//从db里取数据，从库可能没有， 再从主库取
//		for {
//			success:= client.SetNX()
//			if !success {
//				time.Sleep(time.Millisecond*100)
//				//再从缓存里取
//				var a获取不到  = false
//				if !a获取不到 {
//					return //缓存的数据
//				} else {
//
//				}
//			}
//		}
//	} else if ret.Val()["status"] == "2" {
//		//ok
//	}
//}

/*
	//set cache
	//hset status 1进行中， 2已完成缓存设置
	//hset update_ts 操作的时间
	//save db

	//update cache ok

	//get from redis
	//分布式lock and get from db

*/
