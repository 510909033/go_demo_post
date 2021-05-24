package my_middleware

import (
	"log"
	"net/http"
	"time"
)

func DemoMyMiddleware() {
	go func() {
		time.Sleep(time.Second)
		client()
	}()

	//http.HandleFunc("/", textMiddleware(timeMiddleware(http.HandlerFunc(hello))))
	//http.HandleFunc("/", textMiddleware(timeMiddleware(http.HandlerFunc(hello))))
	//http.HandleFunc("/", timeMiddleware(hello))
	//http.HandleFunc("/", textMiddleware(timeMiddleware(hello)))
	http.HandleFunc("/", textMiddleware(timeMiddleware(recoverMiddleware(hello))))
	err := http.ListenAndServe(":8080", nil)

	textMiddleware(timeMiddleware(recoverMiddleware(hello)))
	log.Println(err)
}

func hello(wr http.ResponseWriter, r *http.Request) {
	log.Println("hello")
	wr.Write([]byte("hello"))
	time.Sleep(time.Millisecond * 700)

}
