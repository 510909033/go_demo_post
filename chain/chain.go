package chain

import (
	"log"
	"time"
)

type MyHandler interface {
	Handle(content string)
	next(content string)
}

type IsJsonHandler struct {
	handler MyHandler
	Name    string
}

func (i *IsJsonHandler) Handle(content string) {
	//做一些处理
	log.Println("name=", i.Name)
	log.Println("是json格式，继续处理")
	i.next(content)
}

func (i *IsJsonHandler) next(content string) {
	if i.handler == nil {
		return
	}
	i.handler.Handle(content)
}

type SaveFileHandler struct {
	handler MyHandler
}

func (i *SaveFileHandler) Handle(content string) {
	//做一些处理
	log.Println("已保存问文件，继续处理")
	time.Sleep(time.Second)
	i.next(content)
}

func (i *SaveFileHandler) next(content string) {
	if i.handler == nil {
		return
	}
	i.handler.Handle(content)
}

type SendImHandler struct {
	handler MyHandler
}

func (i *SendImHandler) Handle(content string) {
	//做一些处理
	log.Println("发消息成功，即系")
	//log.Fatalln("发消息失败，停止处理")
	i.next(content)
}

func (i *SendImHandler) next(content string) {
	if i.handler == nil {
		return
	}
	i.handler.Handle(content)
}

func DemoChain() {
	var handler MyHandler
	var handler1 MyHandler
	log.Printf("hander的地址为 %p", handler)
	jsonHandler := &IsJsonHandler{}
	saveFileHandler := &SaveFileHandler{}
	sendImHandler := &SendImHandler{}

	jsonHandler.handler = saveFileHandler
	saveFileHandler.handler = sendImHandler
	sendImHandler.handler = jsonHandler

	handler = jsonHandler
	handler1 = jsonHandler
	handler1 = nil
	_ = handler1

	log.Printf("hander的地址为 %p", handler)
	log.Printf("jsonHandler的地址为 %p", handler)
	handler.Handle("嘿嘿")
}
