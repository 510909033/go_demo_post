package main

import (
	"fmt"
	"net/http"
	"os"
)

//http://31t29j8955.zicp.vip/hello
/*

http://192.168.6.3/go_upload/api/demo/demo
curl 'http://192.168.6.3/go_upload'
nginx -s reload -c /home/nginxWebUI/nginx.conf


*/
func SayHello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hhhhh")
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":54322", nil)

	os.Setenv("ucm_remote", "192.168.6.3:9602")
}
