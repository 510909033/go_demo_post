package struct_pointer

import "log"

type User struct {
	Id         int
	Address    Address
	AddPointer *Address
}

type Address struct {
	Name         string
	LocationInfo Location
}

type Location struct {
	PoiId int
}

func (u User) SetId1(id int) {
	log.Printf("SetId1 %p", &u)
	log.Printf("SetId1 Address %p", &(u.Address))
	log.Printf("SetId1 *Address %p", u.AddPointer)
	log.Printf("SetId1 *Address.Location %p", &(u.AddPointer.LocationInfo))
	u.Id = id
	u.AddPointer.Name = "haha"
	u.AddPointer.LocationInfo.PoiId = 666
}

func (u *User) SetId2(id int) {
	log.Printf("SetId2 %p", u)
	log.Printf("SetId1 Address %p", &(u.Address))
	log.Printf("SetId1 *Address %p", u.AddPointer)
	log.Printf("SetId1 *Address.Location %p", &(u.AddPointer.LocationInfo))
	u.Id = id
}

func MyStructPointer() {
	var u = User{Id: 100, AddPointer: &Address{}}
	log.Printf("user %p", &u)
	log.Printf("user Address %p", &(u.Address))
	log.Printf("user * Address%p", u.AddPointer)
	log.Printf("SetId1 *Address.Location %p", &(u.AddPointer.LocationInfo))

	u.SetId1(5)

	u.SetId2(555)

	log.Println("address name", u.AddPointer.Name)
	log.Println("location name", u.AddPointer.LocationInfo.PoiId)

	u2 := u
	log.Printf("%#v", u)
	log.Printf("%#v", u2)
}
