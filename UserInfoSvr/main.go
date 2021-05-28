/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-28 21:19:01
 * @LastEditors: lly
 * @LastEditTime: 2021-05-28 22:50:16
 */

package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"pb/user"

	_ "github.com/johanbrandhorst/grpc-json-example/codec"
	"google.golang.org/grpc/grpclog"
)

const (
	Address = "127.0.0.1:8899"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	info := user.UserInfo{Name: "lileiyu", Age: 13}
	fmt.Println(info.String())

	// 注册HelloService
	user.RegisterUserInfoSvrServer(s, UserInfoSvr{})
	fmt.Println("Listen on " + Address)
	s.Serve(listen)
}
