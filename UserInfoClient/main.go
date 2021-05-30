/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-28 22:44:51
 * @LastEditors: lly
 * @LastEditTime: 2021-05-31 01:16:31
 */

package main

import (
	"context"
	"fmt"
	"log"
	"pb/user"
	// "github.com/micro/go-micro/v2"
)

const (
	Address = "127.0.0.1:8899"
)

func main() {

	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	client := user.NewUserInfoSvrService("user.UserInfo", service.Client())
	rsp, err := client.GetUserInfo(context.Background(), &user.GetUserInfoReq{UserID: 1})
	if err != nil {
		log.Fatal("")
	}
	fmt.Println(rsp)

	/*
		grpc 使用
		// 连接
		conn, err := grpc.Dial(Address, grpc.WithDefaultCallOptions(grpc.CallContentSubtype(codec.JSON{}.Name())))
		if err != nil {
			//grpclog.Fatalln(err)
			return
		}
		defer conn.Close()


		// // 注册用户
		regsReq := &user.RegisterUserReq{User: &user.UserInfo{
			Name:   "雷雨",
			Gender: 1,
			Age:    18,
			Addr:   "深圳",
			Phone:  "17685244412",
		}}
		regsRsp, err := c.RegisterUser(context.Background(), regsReq)
		if err != nil {
			grpclog.Fatalln(err)
		}
		fmt.Printf("regsRsp=%+v\n", regsRsp)

		// // 查询用户
		// err = conn.Invoke(context.Background(), "/user.UserInfoSvr/GetUserInfo", req, out)
		// if err != nil {
		// 	fmt.Printf("err=%v\n", err)
		// 	return
		// }
		// fmt.Printf("rsp=%+v\n", out)
	*/
}
