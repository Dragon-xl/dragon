package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

/* 用户 table_name = user */
type User struct {
	ID            uint          `gorm:"primaryKey"`        //用户编号
	Name          string        `gorm:"size:32;unique"`    //用户名
	Password_hash string        `gorm:"size:128" `         //用户密码加密的
	Mobile        string        `gorm:"size:11;unique" `   //手机号
	Real_name     string        `gorm:"size:32" `          //真实姓名  实名认证
	Id_card       string        `gorm:"size:20" `          //身份证号  实名认证
	Avatar_url    string        `gorm:"size:256" `         //用户头像路径       通过fastdfs进行图片存储
	Houses        []*House      `gorm:"foreignKey:UserId"` //用户发布的房屋信息  一个人多套房
	Orders        []*OrderHouse `gorm:"foreignKey:UserId"` //用户下的订单       一个人多次订单
}

/* 房屋信息 table_name = house */
type House struct {
	gorm.Model                    //房屋编号
	UserId          uint          //房屋主人的用户编号  与用户进行关联
	AreaId          uint          //归属地的区域编号   和地区表进行关联
	Title           string        `gorm:"size:64" `                 //房屋标题
	Address         string        `gorm:"size:512"`                 //地址
	Room_count      int           `gorm:"default:1" `               //房间数目
	Acreage         int           `gorm:"default:0" json:"acreage"` //房屋总面积
	Price           int           `json:"price"`
	Unit            string        `gorm:"size:32;default:''" json:"unit"`               //房屋单元,如 几室几厅
	Capacity        int           `gorm:"default:1" json:"capacity"`                    //房屋容纳的总人数
	Beds            string        `gorm:"size:64;default:''" json:"beds"`               //房屋床铺的配置
	Deposit         int           `gorm:"default:0" json:"deposit"`                     //押金
	Min_days        int           `gorm:"default:1" json:"min_days"`                    //最少入住的天数
	Max_days        int           `gorm:"default:0" json:"max_days"`                    //最多入住的天数 0表示不限制
	Order_count     int           `gorm:"default:0" json:"order_count"`                 //预定完成的该房屋的订单数
	Index_image_url string        `gorm:"size:256;default:''" json:"index_image_url"`   //房屋主图片路径
	Facilities      []*Facility   `gorm:"many2many:house_facilities" json:"facilities"` //房屋设施   与设施表进行关联
	Images          []*HouseImage `gorm:"foreignKey:HouseId" json:"img_urls"`           //房屋的图片   除主要图片之外的其他图片地址
	Orders          []*OrderHouse `gorm:"foreignKey:HouseId" json:"orders"`             //房屋的订单    与房屋表进行管理
}

/* 区域信息 table_name = area */ //区域信息是需要我们手动添加到数据库中的
type Area struct {
	Id     uint     `gorm:"primaryKey" json:"aid"`           //区域编号     1    2
	Name   string   `gorm:"size:32" json:"aname"`            //区域名字     昌平 海淀
	Houses []*House `gorm:"foreignKey:AreaId" json:"houses"` //区域所有的房屋   与房屋表进行关联
}

/* 设施信息 table_name = "facility"*/ //设施信息 需要我们提前手动添加的
type Facility struct {
	Id     uint     `gorm:"primaryKey" json:"fid"`      //设施编号
	Name   string   `gorm:"size:32"`                    //设施名字
	Houses []*House `gorm:"many2many:house_facilities"` //都有哪些房屋有此设施  与房屋表进行关联的
}

/* 房屋图片 table_name = "house_image"*/
type HouseImage struct {
	Id      uint   `gorm:"primaryKey" json:"house_image_id"` //图片id
	Url     string `gorm:"size:256" json:"url"`              //图片url     存放我们房屋的图片
	HouseId uint   `json:"house_id"`                         //图片所属房屋编号
}

/* 订单 table_name = order */
type OrderHouse struct {
	gorm.Model            //订单编号
	UserId      uint      `json:"user_id"`       //下单的用户编号   //与用户表进行关联
	HouseId     uint      `json:"house_id"`      //预定的房间编号   //与房屋信息进行关联
	Begin_date  time.Time `gorm:"type:datetime"` //预定的起始时间
	End_date    time.Time `gorm:"type:datetime"` //预定的结束时间
	Days        int       //预定总天数
	House_price int       //房屋的单价
	Amount      int       //订单总金额
	Status      string    `gorm:"default:'WAIT_ACCEPT'"` //订单状态
	Comment     string    `gorm:"size:512"`              //订单评论
	Credit      bool      //表示个人征信情况 true表示良好
}

var DB *gorm.DB

func InitDB() error {
	dsn := "root:12345678@tcp(192.168.101.45:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	DB = db
	if err != nil {
		fmt.Println("connect mysql failed:", err)
		return err
	}
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	err = DB.AutoMigrate(
		&User{},       // 1. 先创建用户表
		&Area{},       // 2. 再创建区域表
		&Facility{},   // 3. 再创建设施表
		&House{},      // 4. 再创建房屋表（依赖于用户表和区域表）
		&HouseImage{}, // 5. 再创建房屋图片表（依赖于房屋表）
		&OrderHouse{}, // 6. 最后创建订单表（依赖于用户表和房屋表）
	)
	if err != nil {
		fmt.Println("auto migrate failed:", err)
		return err
	}
	return nil
}
