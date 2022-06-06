package main

import (
	"go_demo_post/gocode/reg"
	_ "go_demo_post/gocode/reg"
	"net/http"
	"os"
	"strconv"

	//"go_demo_post/my_prometheus/monitor/demo3"
	"log"
	_ "net/http/pprof"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	reg.IndexFunc.Dump()

	//mystruct.DemoRangeTest()
	//return

	go func() {
		log.Println("pprof port = ", 8005)
		log.Println(http.ListenAndServe("0.0.0.0:8005", nil))
	}()

	indexStr := ""
	if len(os.Args) > 1 {
		indexStr = os.Args[1]
	}
	index, _ := strconv.ParseInt(indexStr, 10, 64)
	if index > 0 {
		reg.Func[reg.IndexFunc[int(index)]]()
	}
	//gob_true_interface2.MyGobInterface()
	//log.Println()

	//DemoFushi()
	//my_etcd.DemoMyEtcd()
	//my_groutinue_cpu_only_one.DemoMyGroutinueCpuOnlyOne()
	//go my_guage.MyGauge()
	//demo3.GetUserInfo()
	//go my_summary.MySummary()
	//my_counter.MyPrometheus()
	return

	//my_groutinue.DemoGroutinue()
	//multi_groutinue.DemoMultiGroutinue()
	//multi_groutinue2.DemoMultiGroutinue2()
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
