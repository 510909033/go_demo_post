package my_middleware

import (
	"fmt"
	_ "github.com/davecgh/go-spew/spew"
	"net/http"
)

//var fmt  spew
func client() {
	fmt.Println(http.Get("http://localhost:8080/hello"))
	//spew.Dump(http.Get("http://localhost:8080/hello"))

}
