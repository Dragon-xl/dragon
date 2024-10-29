package model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 存储验证码到redis

func SaveCaptchaCode(code, uuid string) error {
	conn, err := redis.Dial("tcp", "192.168.101.45:6379")
	if err != nil {
		fmt.Println("SaveCaptchaCode redis dial err:", err)
		return err
	}
	defer conn.Close()
	_, err = redis.String(conn.Do("setex", uuid, 300, code)) //五分钟有效时间
	if err != nil {
		fmt.Println("SaveCaptchaCode redis set err:", err)
		return err
	}

	return nil
}
