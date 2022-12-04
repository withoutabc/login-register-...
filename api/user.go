package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"goproject1/model"
	"goproject1/service"
	"goproject1/util"
	"log"
	"strconv"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//有效输入
	if username == " " || password == "" {
		util.RespParamErr(c)
		return
	}
	//检索不到
	u, err := service.SearchUserByUsername(username)
	log.Printf("search user error:%v", err)
	if err != nil && err != sql.ErrNoRows {
		util.RespInternalErr(c)
		return
	}
	//用户名是否存在
	if u.Username != "" {
		util.NormErr(c, 300, "账户已存在")
		return
	}
	//创建用户
	err = service.CreateUser(model.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//有效输入
	if username == " " || password == "" {
		util.RespParamErr(c)
		return
	}
	//检索不到用户处理
	u, err := service.SearchUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	//密码错误
	if u.Password != password {
		util.NormErr(c, 300, "密码错误")
		return
	}
	util.RespOK(c)
	//设置cookie
	c.SetCookie("name", username, 3600, "/", "localhost", false, true)
}
func ChangePassword(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	newPassword := c.Query("newPassword")
	//有效输入
	if username == "" || password == "" || newPassword == "" {
		util.RespParamErr(c)
		return
	}
	//检索不到用户处理
	u, err := service.SearchUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	//判断密码是否正确
	if u.Password != password {
		util.NormErr(c, 200, "密码错误")
		return
	}
	//修改密码
	err = service.ChangePassword(newPassword, username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
func DeleteUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//有效输入
	if username == "" || password == "" {
		util.RespParamErr(c)
		return
	}
	//检索不到用户处理
	u, err := service.SearchUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	//判断密码是否正确
	if u.Password != password {
		util.NormErr(c, 200, "密码错误")
		return
	}
	//注销账户
	err = service.DeleteUser(username)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func GetMessage(c *gin.Context) {

}
func SendMessage(c *gin.Context) {
	var m model.Message
	sendUid := c.Query("send_uid")
	if sendUid == "" {
		util.RespParamErr(c)
		return
	}
	RecUID := c.Query("rec_uid")
	m.SendUID, _ = strconv.ParseInt(sendUid, 10, 64)
	m.RecUID, _ = strconv.ParseInt(RecUID, 10, 64)
	m.Detail = c.Query("detail")
	err := service.SendMessage(m)
	if err != nil {
		log.Printf("search user error:%v", err)
		util.NormErr(c, 300, "发送失败")
		return
	}
	util.RespOK(c)
	return
}
func ChangeMessage(c *gin.Context) {

}
func DeleteMessage(c *gin.Context) {

}
