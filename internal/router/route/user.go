package route

import (
	"github.com/gin-gonic/gin"
	"template/global"
	"template/internal/controller"
	"template/internal/middleware"
	"template/internal/repo"
	"template/internal/repo/dao"
	"template/internal/service"
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
