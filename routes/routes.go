package routes

import (
	"project-root/providers"

	"github.com/gin-gonic/gin"

	auth "project-root/modules/auth/routes"
	books "project-root/modules/books/routes"
	ex "project-root/modules/examples/routes"
	users "project-root/modules/users/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.Engine, p *providers.Providers) {
	api := r.Group("api/v1")

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ex.RegisterRoutes(api, p.Examples)
	users.RegisterRoutes(api, p.Users)
	books.RegisterRoutes(api, p.Books)
	auth.RegisterRoutes(api, p.Auth)
}
