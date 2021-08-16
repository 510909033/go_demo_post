package gob_interface1

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Getter interface {
	Get() string
}

type Foo string

func (f Foo) Get() string {
	return string(f)
}

func MyGobInterface() {
	// init and register
	buf := bytes.NewBuffer(nil)

	//todo 不加这句的话 dec.Decode会报错
	gob.Register(new(Foo))

	// create a getter of Foo
	g := Getter(Foo("haha"))
	//g := Foo("haha")

	log.Println("g=", g)
	log.Println("&g=", &g)

	// encode
	enc := gob.NewEncoder(buf)
	encoderr := enc.Encode(&g)

	log.Println("encode encoderr = ", encoderr)

	// decode
	dec := gob.NewDecoder(buf)
	//var gg Foo //panic: gob: attempt to decode into a non-pointer
	//var gg = new(Foo)
	var gg Getter
	if err := dec.Decode(&gg); err != nil {
		panic(err)
	}
	log.Println("decode result = ", gg, gg.Get())
}
