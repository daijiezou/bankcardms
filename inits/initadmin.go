package inits

import (
	"BankCardMS/internal/data/do"
	"BankCardMS/internal/data/mysql"
	"BankCardMS/internal/pkg/glog"
	"BankCardMS/internal/service/user"
	"time"
)

func createAdmin() {
	now := time.Now().UnixMilli()
	pwd := "123456"
	salt, newpwd := user.MakePwd(pwd)
	adminUser := do.User{
		UserId:      1,
		UserName:    "xuxueqin",
		Password:    newpwd,
		Salt:        salt,
		DisplayName: "许学勤",
		Phone:       "",
		Email:       "",
		UpdateTime:  now,
		CreateTime:  now,
	}
	_, err := mysql.MySQL().Insert(adminUser)
	if err != nil {
		glog.Fatalf("init admin failed,err: %v", err)
	}
}

func InitAdmin() {
	has, err := mysql.MySQL().Table("user").Where("user_id = 1").Get(new(do.User))
	if err != nil {
		glog.Fatalf("init admin failed,err: %v", err)
	}
	if !has {
		createAdmin()
	}
}