package discovery

import "github.com/coreos/etcd/clientv3"

//监控
type MonitorService struct {
}

func d1() {
	clientv3.Client{}.MoveLeader()
}
