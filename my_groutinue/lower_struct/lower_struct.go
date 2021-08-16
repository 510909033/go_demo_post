package lower_struct

import (
	"fmt"
	"log"
)

type lowerStruct struct {
	Id int64
}

func NewLowerStruct() *lowerStruct {
	return &lowerStruct{}
}

func (service *lowerStruct) GetId() {
	log.Println(service.Id)
}

type Lower2 struct {
	Id int64
}

func (service *Lower2) GetId() {
	log.Printf("low2.p=%p", service)
	log.Println(service.Id)
}
func (service Lower2) GetId2() {
	
	log.Printf("low2.p=%p", &service)
	log.Println(service.Id)

	fmt.Printf("%.2f", 22.20)
}
