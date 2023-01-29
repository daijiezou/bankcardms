package user

import (
	"BankCardMS/internal/data/do"
	"BankCardMS/internal/data/mysql"
	"BankCardMS/internal/pkg/gerr"
	"BankCardMS/internal/pkg/glog"
	"BankCardMS/internal/pkg/jwt"
	"BankCardMS/internal/pkg/response"
	"BankCardMS/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {
	type LoginReq struct {
		UserName string `json:"user_name" form:"user_name" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
	type LoginRspData struct {
		Token       string `json:"token"`
		UserId      int64  `json:"user_id"`
		UserName    string `json:"user_name"`
		DisplayName string `json:"display_name"`
		PhoneNum    string `json:"phone"`
		Email       string `json:"email"`
	}
	req := new(LoginReq)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}
	user := new(do.User)
	has, err := mysql.MySQL().Where("user_name = ?", req.UserName).Get(user)
	if err != nil {
		response.ErrorCode(c, gerr.ErrCodeDbError)
		return
	}
	if !has {
		response.ErrorCode(c, gerr.ErrCodeInvalidPassword)
		return
	}
	if !CheckPwd(user.Salt, req.Password, user.Password) {
		response.ErrorCode(c, gerr.ErrCodeInvalidPassword)
		return
	}
	accessToken, err := jwt.GenToken(user.UserId, req.UserName)
	if err != nil {
		glog.Warnf("jwt GenToken error", glog.String("msg", err.Error()))
		response.ErrorCode(c, gerr.ErrCodeServerBusy)
		return
	}
	result := LoginRspData{
		Token:       accessToken,
		UserId:      user.UserId,
		UserName:    user.UserName,
		DisplayName: user.DisplayName,
		PhoneNum:    user.Phone,
		Email:       user.Email,
	}
	response.Success(c, result)
}

func EditPwd(c *gin.Context) {
	type EditUserPasswordReq struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password" binding:"required,min=6,max=25"`
	}
	req := new(EditUserPasswordReq)
	if err := c.ShouldBind(req); err != nil {
		glog.Warnf("req params check failed:%v,req params:%+v", err, req)
		response.ErrorCode(c, gerr.ErrCodeWrongParam)
		return
	}
	userId, err := utils.GetCurrentUserID(c)
	if err != nil {
		glog.Warnf("get current UserID failed,err:%v", err)
		response.ErrorCode(c, gerr.ErrCodeUnauthorized)
		return
	}
	user := new(do.User)
	has, err := mysql.MySQL().Where("user_id = ?", userId).Get(user)
	if err != nil {
		glog.Errorf("db error:%v", err)
		response.ErrorCode(c, gerr.ErrCodeDbError)
		return
	}
	if !has {
		response.ErrorCode(c, gerr.ErrCodeInvalidPassword)
		return
	}
	if !CheckPwd(user.Salt, req.OldPassword, user.Password) {
		response.ErrorCode(c, gerr.ErrCodeInvalidPassword)
		return
	}
	user.Salt, user.Password = MakePwd(req.NewPassword)
	user.UpdateTime = time.Now().UnixMilli()
	_, err = mysql.MySQL().Cols("salt", "update_time", "password").Where("user_id = ?", userId).Update(user)
	if err != nil {
		glog.Errorf("db error:%v", err)
		response.ErrorCode(c, gerr.ErrCodeDbError)
		return
	}
	response.Success(c, nil)
	return
}
