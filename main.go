package main

import (
	my_etcd "go_demo_post/my-etcd"
	"go_demo_post/my_groutinue/multi_groutinue2"
	"log"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	//DemoFushi()
	my_etcd.DemoMyEtcd()
	return

	//my_groutinue.DemoGroutinue()
	//multi_groutinue.DemoMultiGroutinue()
	multi_groutinue2.DemoMultiGroutinue2()
	//my_pointer.DebugPointer()
	//return
	//modifyTitle()
	//my_defer.DemoDefer()
	//demo1rpc.DemoRpc1()
	//demo_rpc_json.DemoRpcJson()
	//demo_rpc_http.DemoRpcHttp()

	//my_bucket_v1.DebugMyBucketV1()

	//my_parser.DemoMyParser()
	//my_middleware.DemoMyMiddleware()

	//my_gc.DemoGc()
	//demokuaishou()
	return

}
