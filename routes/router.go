package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ginblog/api/v1"
	"github.com/ginblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		// 用户模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		router.POST("category/add", v1.AddCategory)
		router.GET("category", v1.GetCategory)
		router.PUT("category/:id", v1.EditCategory)
		router.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块的路由接口
		router.POST("article/add", v1.AddArticle)
		router.GET("article", v1.GetArticle)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.PUT("article/:id", v1.EditArticle)
		router.DELETE("article/:id", v1.DeleteArticle)
	}
	r.Run(utils.HttpPort)
}
