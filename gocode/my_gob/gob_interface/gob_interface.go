package gob_interface

import (
	"bytes"
	"encoding/gob"
)

type Getter interface {
	Get() string
}

type Foo struct {
	Bar string
}

func (f Foo) Get() string {
	return f.Bar
}

func MyGobInterface() {
	// init and register
	buf := bytes.NewBuffer(nil)

	//todo 不加这句的话 dec.Decode会报错 
	//gob.Register(Foo{})

	// create a getter of Foo
	g := Getter(Foo{"wazzup"})

	// encode
	enc := gob.NewEncoder(buf)
	enc.Encode(&g)

	// decode
	dec := gob.NewDecoder(buf)
	var gg Getter
	if err := dec.Decode(&gg); err != nil {
		panic(err)
	}
}
