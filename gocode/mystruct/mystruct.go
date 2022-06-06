package mystruct

import (
	"log"
)

func main() {

	//c1 := &mya.AController{}
	//fmt.Println(c1.GetName())

}

type user struct {
	Id int
}

var list = []user{
	{Id: 1},
	{Id: 2},
	{Id: 3},
}

func DemoRangeTest() {

	var u = &user{Id: 111}
	u1 := u
	u1.Id = 222
	log.Println(*u, *u1)
	return

	for k, v := range list {
		log.Printf("%p", &v)
		v.Id++
		list[k] = v
		v.Id++
	}
	log.Printf("%#v", list)

	for _, v := range list {
		log.Printf("%p", &v)
	}

}
