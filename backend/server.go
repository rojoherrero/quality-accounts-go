package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/go-nats"
)

type Server struct {
	engine *gin.Engine
	api    *api
}

func InitServer(db *sqlx.DB, nc *nats.Conn) *Server {
	a := &Server{
		api:    newApi(db, nc),
		engine: gin.Default(),
	}
	a.engine.Use(CORSMiddleware())
	a.createRoutes()
	return a

}

func (a *Server) createRoutes() {
	baseRoute := a.engine.Group("/api")
	roleRoutes := baseRoute.Group("/role")
	{
		roleRoutes.POST("", a.api.role.Save)
		roleRoutes.PUT("", a.api.role.Update)
		roleRoutes.GET("", a.api.role.Paginate)
		roleRoutes.DELETE("", a.api.role.Delete)
	}
	departmentRoutes := baseRoute.Group("/department")
	{
		departmentRoutes.POST("", a.api.department.Save)
		departmentRoutes.PUT("", a.api.department.Update)
		departmentRoutes.GET("", a.api.department.Paginate)
		departmentRoutes.DELETE("", a.api.department.Delete)
	}
	userRoutes := baseRoute.Group("/user")
	{
		userRoutes.POST("", a.api.user.Save)
		userRoutes.PUT("", a.api.user.Update)
		userRoutes.GET("", a.api.user.Paginate)
		userRoutes.DELETE("", a.api.user.Delete)
	}

}


func (a *Server) Run(port string) {
	listeningPort := fmt.Sprintf(":%s", port)
	a.engine.Run(listeningPort)
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
