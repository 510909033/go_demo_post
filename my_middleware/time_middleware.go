package my_middleware

import (
	"log"
	"net/http"
	"time"
)

func timeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	log.Println("timeMiddleware")
	t := time.Now()
	log.Println("before time")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next.ServeHTTP(w, r)
		log.Println("time:", time.Now().Sub(t))
		log.Println("after time")
	})

}

func textMiddleware(next http.HandlerFunc) http.HandlerFunc {
	log.Println("textMiddleware")
	log.Println("before text ")
	defer func() {
		log.Println("after text")
	}()
	time.Sleep(time.Second)
	//panic("test panic")

	return next
}

func recoverMiddleware(next http.HandlerFunc) http.HandlerFunc {
	log.Println("recoverMiddleware")
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover err=%+v", err)
		}
	}()

	return next
}

type MyHandler interface {
	Handler()
	Next() MyHandler
}

type a int
type b int

func (b b) Handler() {
	panic("implement me")
}

func (b b) Next() MyHandler {
	panic("implement me")
}

func (a a) Next() MyHandler {
	panic("implement me")
}

func (a a) Handler() {

}

func demo1() {

	var a1 a
	a1.Next()

	//value := atomic.Value{}
	//value.Store()

	//reflect.StringHeader{}

}
