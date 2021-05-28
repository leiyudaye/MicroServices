/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-28 21:20:14
 * @LastEditors: lly
 * @LastEditTime: 2021-05-28 23:13:32
 */

package main

import (
	"context"
	"errors"
	"fmt"
	"pb/user"

	"github.com/garyburd/redigo/redis"
)

const (
	gRedisAddr         = "159.75.212.214:6380"
	gUserInfoPrefix    = "userinfo_user_"
	gUserSessionPrefix = "userinfo_session"
)

type UserInfoSvr struct {
}

// 注册用户
func (u UserInfoSvr) RegisterUser(ctx context.Context, req *user.RegisterUserReq) (*user.RegisterUserRsp, error) {
	redisfd, err := redis.Dial("tcp", gRedisAddr)
	if err != nil {
		fmt.Println("redis dial failed, err=%v", err)
		return nil, errors.New("redis dial failed")
	}
	defer redisfd.Close()

	// 自增用户ID
	userID, err := redis.Int64(redisfd.Do("INCR", gUserSessionPrefix))
	if err != nil {
		fmt.Println("user session incrby failed. err=%v", err)
		return nil, errors.New("redis incrby failed")
	}

	// 存储用户信息
	_, err = redisfd.Do("HMSet", fmt.Sprintf("%s%v", gUserInfoPrefix, userID),
		"Name", req.User.Name,
		"Gender", req.User.Gender.Number(),
		"Age", req.User.Age,
		"Addr", req.User.Addr,
		"Phone", req.User.Phone)
	if err != nil {
		fmt.Println("HMSet failed, err=%v", err)
		return nil, errors.New("HMSet failed")
	}

	// 返回信息
	rsp := &user.RegisterUserRsp{
		Code:   0,
		Msg:    "success",
		UserID: userID,
	}
	return rsp, nil
}

// 查询用户信息
func (u UserInfoSvr) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (*user.GetUserInfoRsp, error) {

	fmt.Println(req)
	rsp := new(user.GetUserInfoRsp)
	info := &user.UserInfo{Name: "lly"}
	rsp.User = info
	return rsp, nil

	/*
		redisfd, err := redis.Dial("tcp", gRedisAddr)
		rsp := new(user.GetUserInfoRsp)
		if err != nil {
			fmt.Println("redis dial failed, err=%v", err)
			return nil, errors.New("redis dial failed")
		}
		defer redisfd.Close()

		// 取出用户信息
		userInfo, err := redis.Strings(redisfd.Do("HGetALL", gUserInfoPrefix+strconv.FormatInt(req.UserID, 10)))
		if err != nil {
			fmt.Println("HMSet failed, err=%v", err)
			return nil, errors.New("HMSet failed")
		}
		fmt.Println(userInfo)
		info := &user.UserInfo{}
		for i := 0; i < len(userInfo); i += 2 {
			if userInfo[i] == "Name" {
				info.Name = userInfo[i+1]
			}
			if userInfo[i] == "Addr" {
				info.Addr = userInfo[i+1]
			}
			if userInfo[i] == "Phone" {
				info.Phone = userInfo[i+1]
			}
			if userInfo[i] == "Gender" {
				gender, _ := strconv.ParseInt(userInfo[i+1], 10, 64)
				info.Gender = user.Gender(gender)
			}
			if userInfo[i] == "Age" {
				age, _ := strconv.ParseInt(userInfo[i+1], 10, 64)
				info.Age = int32(age)
			}
		}
		rsp.User = info
		return rsp, nil
	*/
}
