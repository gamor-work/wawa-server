package middleware

import (
	"gin/global"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/sirupsen/logrus"
	"time"
)

func ZookeeperRegister() {
	hosts := []string{global.ZookeeperAddr}
	conn, _, err := zk.Connect(hosts, time.Second*5)

	if err != nil {
		// todo
		logrus.Errorf("Connect %s", err.Error())
	}

	global.ZookeeperConn = conn

	defer conn.Close()

}
