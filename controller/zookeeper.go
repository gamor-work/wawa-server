package controller

var (
	PATH  = "/wawa"
	DATA1 = "zoujun"
)

//
///*
//	zookeeper 服务注册
//*/
//func ZookeeperRegisterHandler(c *gin.Context) {
//	err = global.ZookeeperConn.Create(PATH, []byte(DATA1))
//	if err != nil {
//		log.Fatal("创建 " + PATH + " 失败")
//	} else {
//		log.Println("创建 " + PATH + " 成功")
//	}
//}
