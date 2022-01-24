package routes

import (
	utils "bytedance/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters() *gin.Engine{
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	return r
}

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")

	//成员管理
	g.POST("/member/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"msg" : "ok",
		})
	})
	g.GET("/member")
	g.GET("/member/list")
	g.POST("/member/update")
	g.POST("/member/delete")

	//登录
	g.POST("/auth/login")
	g.POST("/auth/logout")
	g.GET("/course/who_am_i")

	//排课
	g.POST("/course/create")
	g.GET("/course/get")

	g.POST("/teacher/bind_course")
	g.POST("/teacher/unbind_course")
	g.GET("/teacher/get_course")
	g.POST("/course/schedule")

	//抢课
	g.POST("/student/book_course")
	g.GET("/student/course")

}