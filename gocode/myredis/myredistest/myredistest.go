package myredistest

import (
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
	"runtime"
	"strconv"
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

func LuaDel() {
	filename := "/home/wangbaotian/go_demo_post/gocode/myredis/lua/del.lua"
	bytes, e := ioutil.ReadFile(filename)
	errPanic(e)

	key := "lua_del_test"
	client.Set(key, "val", time.Minute)
	eval := client.Eval(string(bytes), []string{key}, []interface{}{"val"}...)
	log.Printf("%+v", eval.Err())
	log.Printf("%+v", eval.Val())
}

func debug(cmder redis.Cmder) {
	ret := map[string]interface{}{
		"Err()-redis.nil": cmder.Err() == redis.Nil,
		"Err()":           cmder.Err(),
		"Name()":          cmder.Name(),
		//"String()":        cmder.String(),
		"Args()": cmder.Args(),
		//"FullName()":      cmder.FullName(),
	}
	log.Printf("ret=%+v\n", ret)
	if cmder.Err() != nil {
		log.Printf("ERROR: %+v", cmder.Err())
		//panic(cmder.Err())
	}
}

func Big() {
	key := "{sadd}_test2"

	size := 120
	val := make([]interface{}, size)
	for i := 0; i < size; i++ {
		val[i] = i
	}
	intCmd := client.SAdd(key, val...)
	debug(intCmd)

	val = make([]interface{}, size)
	for i := 0; i < size; i++ {
		val[i] = i + size
	}
	client.SAdd(key, val...)
	log.Println("sadd success")

	go func() {
		log.Println("smember start")
		t := time.Now()
		members := client.SMembers(key)
		strings, _ := members.Result()
		log.Println("smember", len(strings), time.Since(t).Milliseconds())
	}()

	runtime.Gosched()
	go func() {
		for i := 0; i < 20; i++ {
			t := time.Now()
			client.Get("{sadd}_ha" + strconv.Itoa(i))
			log.Println(i, time.Since(t).Milliseconds())
			//runtime.Gosched()
		}
	}()

	time.Sleep(time.Second * 3)

}

func GetId() {

	Big()
	return

	key := fmt.Sprintf("%d", time.Now().UnixNano())
	log.Println("key=", key)

	for i := 0; i < 3; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			cmd := client.Incr(key)
			log.Println(cmd.Result())
			if cmd.Err() != nil {
				log.Println(cmd.Err())
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
		log.Println(key, val)
		return true
	})
	v, ok := incrSyncM.Load(int64(1))
	log.Println(v, ok)

	//ch := NewClientConn()
	//<-ch
}

func LuaSomeMethod() {
	filename := "/home/wangbaotian/go_demo_post/gocode/myredis/lua/some_method.lua"

	key := "LuaSomeMethod"
	client.Set(key, "LuaSomeMethod", time.Minute)

	bytes, e := ioutil.ReadFile(filename)
	errPanic(e)

	stringCmd := client.ScriptLoad(string(bytes))
	errPanic(stringCmd.Err())
	log.Println(stringCmd)

	hashids, e := stringCmd.Result()
	errPanic(e)
	log.Println("hashids", hashids)

	time.Sleep(time.Second)
	exists := client.ScriptExists(hashids)
	log.Println(exists)

	cmd := client.EvalSha(hashids, []string{key}, []interface{}{"mmm"}...)
	log.Println(cmd)
	log.Printf("%#v", cmd.Val())

}

func LuaSomeMethod2() {
	filename := "/home/wangbaotian/go_demo_post/gocode/myredis/lua/some_method2.lua"

	key := "LuaSomeMethod"
	client.Set(key, "LuaSomeMethod", time.Minute)

	bytes, e := ioutil.ReadFile(filename)
	errPanic(e)

	cmd := client.Eval(string(bytes), []string{"{k}1", "{k}2", "{k}3", "{k}4"}, []interface{}{"a1", "a2", "a3", "a4"}...)
	log.Printf("%#v", cmd)
	log.Printf("%#v", cmd.Val())

}

func Mset() {
	var list []interface{}

	for i := 0; i < 10; i++ {
		//mset mset_0 val0 mset_1 val1 mset_2 val2 mset_3 val3 mset_4 val4 mset_5 val5 mset_6 val6 mset_7 val7 mset_8 val8 mset_9 val9:
		// CROSSSLOT Keys in request don't hash to the same slot
		//list = append(list, "mset_"+strconv.Itoa(i), "val"+strconv.Itoa(i))

		list = append(list, "{mset}_"+strconv.Itoa(i), "val"+strconv.Itoa(i))
	}

	cmd := client.MSet(list...)
	cmd1 := client.MSetNX(list...)
	log.Println(cmd)
	log.Println(cmd1)
}

func Mget() {
	var list []string

	for i := 0; i < 10; i++ {
		list = append(list, "{mset}_"+strconv.Itoa(i))
		//list = append(list, "{mset}_"+strconv.Itoa(i), "val"+strconv.Itoa(i))
	}

	cmd := client.MGet(list...)
	log.Println(cmd)
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

func TestRedisShutdown() {
	key := "TestRedisShutdown"
	for {
		statusCmd := client.Set(key, "1", time.Minute)
		debug(statusCmd)
		//time.Sleep(time.Second * 3)
		time.Sleep(time.Millisecond)
	}
}
func TestRedisSlaveDown() {

	client.ClusterFailover()

	//key := "TestRedisShutdown"
	//for {
	//	statusCmd := client.Set(key, "1", time.Minute)
	//	debug(statusCmd)
	//	//time.Sleep(time.Second * 3)
	//	time.Sleep(time.Millisecond * 200)
	//}
}
