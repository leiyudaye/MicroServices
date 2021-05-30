/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-28 21:19:01
 * @LastEditors: lly
 * @LastEditTime: 2021-05-31 00:57:03
 */

package main

import (
	"log"

	"github.com/micro/micro/v3/service"

	pb "pb/user"
)

const (
	gEtcdAddress = "127.0.0.1:2379"
)

func main() {

	// micro使用
	// etcdReg := etcd.NewRegistry(registry.Addrs(gEtcdAddress))
	svr := service.New(
		service.Name("user.UserInfo"), //可以通过这个名字调用到此服务
		service.Address(":8810"),
		// micro.Registry(etcdReg),
	)

	svr.Init()
	pb.RegisterUserInfoSvrHandler(svr.Server(), new(UserInfoSvrImpl))
	if err := svr.Run(); err != nil {
		log.Fatal(err)
	}
}
