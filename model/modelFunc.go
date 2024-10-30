package model

import (
	"crypto/md5"
	"encoding/hex"
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
			conn, err := redis.Dial("tcp", "192.168.101.45:6379")
			if err != nil {
				return nil, err
			}
			if _, err = conn.Do("AUTH", "123"); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
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

// 处理登录 根据手机号 密码 获取用户名
func LoginRetName(phone, password string) (string, error) {
	m5 := md5.New()
	m5.Write([]byte(password))
	pwd_hash := hex.EncodeToString(m5.Sum(nil))
	var user User
	err := DB.Select("name").Where("mobile = ?", phone).Where("password_hash = ?", pwd_hash).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

// 获取用户信息
func GetUserInfo(userName string) (User, error) {
	var user User
	err := DB.Where("name = ?", userName).First(&user).Error
	return user, err
}

// 更新用户名
func UpdateUserName(oldname, newname string) error {
	var user User
	err := DB.Model(&user).Where("name = ?", oldname).Update("name", newname).Error
	return err

}
