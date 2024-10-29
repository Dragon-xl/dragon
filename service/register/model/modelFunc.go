package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

// 声明全局连接池句柄
var RedisPool redis.Pool

func InitRedis() {
	RedisPool = redis.Pool{
		MaxIdle:         20,
		MaxActive:       20,
		MaxConnLifetime: 300 * time.Second,
		IdleTimeout:     30 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "192.168.101.45:6379")
		},
	}
}
func CheckCaptchaCode(uuid, code string) bool {
	conn := RedisPool.Get()
	defer conn.Close()
	codeStr, err := redis.String(conn.Do("get", uuid))
	if err != nil {
		return false
	}
	return codeStr == code
}

// 去redis存短信验证码
func SaveSmsCode(phone, code string) error {
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("setex", phone, 180, code)
	if err != nil {
		return err
	}
	return nil
}

//校验短信验证码

func CheckSmsCode(phone, code string) bool {
	conn := RedisPool.Get()
	defer conn.Close()
	codeStr, err := redis.String(conn.Do("get", phone))
	if err != nil {
		fmt.Println("service CheckSmsCode failed err:", err)
		return false
	}
	return codeStr == code
}

// 注册用户信息 存入MYSQL数据库
func SaveRegisterUser(phone, passwd string) error {

	//使用md5对密码加密
	m5 := md5.New()
	m5.Write([]byte(passwd))
	hash := hex.EncodeToString(m5.Sum(nil))
	//写入数据库
	err := DB.Create(&User{
		Name:          phone,
		Password_hash: hash,
		Mobile:        phone,
	}).Error
	return err
}
