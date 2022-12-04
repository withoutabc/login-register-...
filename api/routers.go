package api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	u := r.Group("/user")
	{
		u.POST("/register", Register)
		u.POST("/login", Login)
		u.PUT("/changePassword", ChangePassword)
		u.DELETE("/deleteUser", DeleteUser)
	}
	m := r.Group("/message")
	{
		//m.Use(middleware.Auth())
		m.GET("/message", GetMessage)
		m.POST("/message", SendMessage)
		m.PUT("/message", ChangeMessage)
		m.DELETE("/message", DeleteMessage)
	}
	r.Run()
}
