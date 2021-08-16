package reg

import (
	"go_demo_post/aa_githubcode/my_tdigest/monitor1"
	"go_demo_post/gocode/my_gob/gob_interface"
	"go_demo_post/gocode/my_gob/gob_interface1"
	"go_demo_post/gocode/my_gob/gob_true_interface"
	"go_demo_post/gocode/my_gob/gob_true_interface2"
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
