package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"math"
	"math/big"
	"net"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile)

	u := uuid.NewV4()
	log.Println(u)
	consumer := fmt.Sprintf("amq.ctag-%s", base64.StdEncoding.EncodeToString(u.Bytes()))
	log.Println(consumer)
	hash := md5.New()
	hash.Write(u.Bytes())
	fmt.Sprintf("%x", hash.Sum(nil))

	log.Println(1 / 3)
	var a1 = float64(1) / 3
	var a3 = float32(1) / 3
	a2 := a1 + 0.002
	log.Println("float32", a3)
	log.Printf("%.3f", float64(1)/3)
	log.Println(a2, a2 == 0.3353333333333333)
	log.Println(a2, a2 == 0.3353333333333332)
	log.Println(a2, a2 == 0.3353333333333334)
	log.Printf("%.30f", float64(1)/3)
	log.Printf("%.31f", float64(1)/3)
	return
	fmt.Println(math.Trunc(9.815*1e2+0.5) * 1e-2)              //9.82
	fmt.Println(math.Trunc(9.825*1e2+0.5) * 1e-2)              //9.83
	fmt.Println(math.Trunc(9.835*1e2+0.5) * 1e-2)              //9.84
	fmt.Println(math.Trunc(9.845*1e2+0.5) * 1e-2)              //9.85
	fmt.Println(math.Trunc(3.3*1e2+0.5) * 1e-2)                //3.3000000000000003
	fmt.Println(math.Trunc(3.3000000000000003*1e2+0.5) * 1e-2) //3.3000000000000003

	var f float32 = 3.1415926
	var f1 float32 = 3.1415925
	var a = float32(25.0)
	var b = float64(25.0)
	log.Println(a, b)
	log.Println(float64(a) == b)
	log.Println(f, f1, f == f1)

	float1 := big.NewFloat(float64(f))
	float2 := big.NewFloat(float64(f1))
	log.Println(float1, float2)
	log.Println(float1.String(), float2.String(), float1 == float2)

	float3 := big.NewFloat(float64(25))
	log.Println("float3", float3, int(float64(10.5/0.5)*100), float64(25)*100 == 2500)

	log.Println(f)

	// float32 转 float64
	log.Printf("%v\n", float64(f)) // 输出：3.141592502593994，6位后的小数精度是错误的

	// float64 转 float32
	var f2 float64 = 3.141592653589793
	log.Printf("%v\n", float32(f2)) // 输出：3.1415927，6位后的小数精度是错误的
	//64.5   64.499999

	return
	demo1()
}
func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func demo1() {

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IP("0.0.0.0"),
		Port: 23232,
		Zone: "",
	})
	errPanic(err)
	go func() {
		for {
			conn, err := listener.Accept()
			errPanic(err)
			log.Println(conn)
		}
	}()

	time.Sleep(time.Second)

	//net.DialTCP("tcp", "", "")

}
