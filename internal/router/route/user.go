package route

import (
	"blog/global"
	"blog/internal/controller"
	"blog/internal/middleware"
	"blog/internal/repo"
	"blog/internal/repo/dao"
	"blog/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(router *gin.RouterGroup) {
	publicRouter := router.Group("/user")
	privateRouter := router.Group("/user")
	privateRouter.Use(middleware.Auth())

	db := global.Db
	svc := service.NewUserService(
		repo.NewUserRepo(dao.NewUserDao(db)),
	)

	ctl := controller.NewUserController(svc)
	{
		publicRouter.POST("/login", ctl.Login)
		publicRouter.POST("/register", ctl.Register)

		privateRouter.GET("/info", ctl.GetUserInfo)
	}
}
