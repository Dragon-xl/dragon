package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model //继承
	Name       string
	Age        int
}

var DB *gorm.DB
var err error

func main() {
	dsn := "root:12345678@tcp(192.168.101.45:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("open mysql err:", err)
		return
	}

	//err = db.AutoMigrate(&User{}) //创建表
	err = DB.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("auto migrate err:", err)
		return
	}
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("sqlDB err:", err)
		return
	}
	sqlDB.SetMaxIdleConns(10)  //初始数目
	sqlDB.SetMaxOpenConns(100) //最大数目
	defer sqlDB.Close()
	SelectData()
	//InsertDate()
	//UpdateData()
	//DeleteData()

}
func InsertDate() {
	DB.Create(&User{
		Name: "123",
		Age:  9,
	})
}
func SelectData() {
	user := make([]User, 2)

	DB.Select("name,age").Where("name=?", "123").Unscoped().Find(&user)
	fmt.Println(user)
}
func UpdateData() {
	err = DB.Model(new(User)).Where("name=?", "123").
		Update("name", "xulong").Error
	if err != nil {
		fmt.Println("update err:", err)
		return
	}

}
func DeleteData() {
	err = DB.Where("name=?", "123").Delete(&User{}).Error
	if err != nil {
		fmt.Println("delete err:", err)
		return
	}
}
