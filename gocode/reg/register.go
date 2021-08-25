package reg

import (
	monitor12 "go_demo_post/aa_githubcode/my_tdigest-opts2/monitor1"
	"go_demo_post/aa_githubcode/my_tdigest/monitor1"
	"go_demo_post/gocode/my_gob/gob_interface"
	"go_demo_post/gocode/my_gob/gob_interface1"
	"go_demo_post/gocode/my_gob/gob_true_interface"
	"go_demo_post/gocode/my_gob/gob_true_interface2"
	"go_demo_post/gocode/my_json"
	"go_demo_post/gocode/my_p99"
	"go_demo_post/gocode/my_regexp"
	"go_demo_post/gocode/mystruct/struct_pointer"
	"go_demo_post/gocode/mystruct/struct_pointer2"
	"log"
	"reflect"
	"runtime"
	"sync"
)

type IndexFuncMap map[int]string

var mu sync.Mutex
var Func = make(map[string]func())
var IndexFunc = make(IndexFuncMap)
var index = 0

func init() {
	MyRegister(gob_interface.MyGobInterface)
	MyRegister(gob_interface1.MyGobInterface)
	MyRegister(gob_true_interface.MyGobInterface)
	MyRegister(gob_true_interface2.MyGobInterface)
	MyRegister(struct_pointer.MyStructPointer)
	MyRegister(struct_pointer2.MyStructPointer2)
	MyRegister(monitor1.DemoMonitor1)
	MyRegister(monitor12.DemoMonitorOpts2)
	MyRegister(my_p99.DemoMyP99)
	MyRegister(my_regexp.DemoMyRegexp)
	MyRegister(my_json.MyJson)
}

func MyRegister(fn func()) {
	mu.Lock()
	defer mu.Unlock()

	val := reflect.ValueOf(fn)
	name := runtime.FuncForPC(val.Pointer()).Name()
	log.Println(name)
	Func[name] = fn
	IndexFunc[index] = name
	index++
}

func (fn IndexFuncMap) Dump() {
	for k, v := range fn {
		log.Println(k, v)
	}
}
