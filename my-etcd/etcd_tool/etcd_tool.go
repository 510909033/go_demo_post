// 生成etcd的配置文件和其它配置的工具
//
// 省去每次都要编写的问题
package etcd_tool

import (
	"fmt"
	"log"
	"strings"
)

//生成goreman的配置文件
//
//可以将输出保存为文件
//
//goreman -f procfile start
func CreateProcfile() {
	type config struct {
		ProcessName              string
		Name                     string
		ListenClientUrls         string
		AdvertiseClientUrls      string
		ListenPeerUrls           string
		InitialAdvertisePeerUrls string
		InitialClusterToken      string
		InitialCluster           string
		InitialClusterState      string
	}

	templateDemo := `etcd1: etcd 
--name infra1 
--listen-client-urls http://0.0.0.0:2379 
--advertise-client-urls http://0.0.0.0:2379 
--listen-peer-urls http://0.0.0.0:12380 
--initial-advertise-peer-urls http://0.0.0.0:12380 
--initial-cluster-token etcd-cluster-1 
--initial-cluster 'infra1=http://0.0.0.0:12380,infra2=http://0.0.0.0:22380,infra3=http://0.0.0.0:32380' 
--initial-cluster-state new 
--enable-pprof`
	_ = templateDemo

	template := `%s: etcd 
--name %s 
--listen-client-urls %s 
--advertise-client-urls %s 
--listen-peer-urls %s 
--initial-advertise-peer-urls %s 
--initial-cluster-token %s 
--initial-cluster '%s' 
--initial-cluster-state %s 
--enable-pprof`

	port := 2379
	port2 := 12380
	ip := "0.0.0.0"
	configList := make([]config, 0)
	max := 10
	initialClusterList := make([]string, max)

	for i := 0; i < max; i++ {
		initialClusterList[i] = fmt.Sprintf("infra-%d=http://%s:%d", i, ip, port2+i)
	}
	for i := 0; i < max; i++ {
		c := config{
			ProcessName:              fmt.Sprintf("etcd-%d", i),
			Name:                     fmt.Sprintf("infra-%d", i),
			ListenClientUrls:         fmt.Sprintf("http://%s:%d", ip, port+i),
			AdvertiseClientUrls:      fmt.Sprintf("http://%s:%d", ip, port+i),
			ListenPeerUrls:           fmt.Sprintf("http://%s:%d", ip, port2+i),
			InitialAdvertisePeerUrls: fmt.Sprintf("http://%s:%d", ip, port2+i),
			InitialClusterToken:      "etcd-cluster-1",
			InitialCluster:           strings.Join(initialClusterList, ","),
			InitialClusterState:      "new",
		}
		configList = append(configList, c)
	}

	//生成字符串
	var builder strings.Builder
	for _, c := range configList {
		tem := fmt.Sprintf(template,
			c.ProcessName,
			c.Name,
			c.ListenClientUrls,
			c.AdvertiseClientUrls,
			c.ListenPeerUrls,
			c.InitialAdvertisePeerUrls,
			c.InitialClusterToken,
			c.InitialCluster,
			c.InitialClusterState)
		builder.WriteString(strings.Replace(tem, "\n", " ", -1))
		builder.WriteString("\n")
	}
	log.Println("procfile结果")
	log.Println(builder.String())
}
