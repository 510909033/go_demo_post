package gob_true_interface2

import (
	"bytes"
	"encoding/gob"
	"log"
)

type User struct {
	Data interface{}
}

func MyGobInterface() {

	var g = User{}

	g.Data = []interface{}{
		1, 2, "222", "333", map[string]interface{}{
			"haha": "1",
		},
	}

	// init and register
	buf := bytes.NewBuffer(nil)

	//todo 不加这句的话 dec.Decode会报错
	gob.Register([]interface{}{})
	//gob.Register([]interface{}{})
	//gob.Register(map[string]interface{}{})

	// create a getter of Foo

	// encode
	enc := gob.NewEncoder(buf)
	encoderr := enc.Encode(&g)
	log.Println("encode err=", encoderr)

	// decode
	dec := gob.NewDecoder(buf)
	//var gg []interface{}
	var gg User
	if err := dec.Decode(&gg); err != nil {
		panic(err)
	}

	log.Printf("gg=%#v\n", gg)
}
