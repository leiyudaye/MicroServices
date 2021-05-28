/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-28 22:44:51
 * @LastEditors: lly
 * @LastEditTime: 2021-05-28 22:47:25
 */

package main

import (
	"context"
	"fmt"
	"pb/user"

	"github.com/johanbrandhorst/grpc-json-example/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	Address = "127.0.0.1:8899"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithDefaultCallOptions(grpc.CallContentSubtype(codec.JSON{}.Name())))
	if err != nil {
		//grpclog.Fatalln(err)
		return
	}
	defer conn.Close()

	// // 初始化客户端
	c := user.NewUserInfoSvrClient(conn)

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
}
