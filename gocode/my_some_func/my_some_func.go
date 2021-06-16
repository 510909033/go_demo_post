package main

import (
	"fmt"
	"net"
)

func main() {
	CNAME("114.114.114.114")
	CNAME("www.baidu.com")
	CNAME("api.babytree.com")
}

func CNAME(src string) (dst string, err error) {
	//114.114.114.114
	dst, err = net.LookupCNAME(src)
	fmt.Printf("%#v\n %#v\n", dst, err)
	net.ListenUnix()
	return
}

/*
time.Now 得到的当前时间的时区跟电脑的当前时区一样。
time.Parse 把时间字符串转换为Time，时区是UTC时区。
不管Time变量存储的是什么时区，其Unix()方法返回的都是距离UTC时间：1970年1月1日0点0分0秒的秒数。
Unix()返回的秒数可以是负数，如果时间小于1970-01-01 00:00:00的话。
Zone方法可以获得变量的时区和时区与UTC的偏移秒数，应该支持夏令时和冬令时。
time.LoadLocation可以根据时区名创建时区Location，所有的时区名字可以在$GOROOT/lib/time/zoneinfo.zip文件中找到，解压zoneinfo.zip可以得到一堆目录和文件，我们只需要目录和文件的名字，时区名是目录名+文件名，比如"Asia/Shanghai"。中国时区名只有"Asia/Shanghai"和"Asia/Chongqing"，而没有"Asia/Beijing"。
time.ParseInLocation可以根据时间字符串和指定时区转换Time。

作者：云上听风
链接：https://www.jianshu.com/p/f809b06144f7
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
