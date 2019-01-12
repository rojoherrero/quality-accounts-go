package server

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/go-nats"
	"github.com/rs/zerolog"
)

type Server struct {
	engine *gin.Engine
	api    *api
}

func NewServer(db *sqlx.DB, nc *nats.Conn, log zerolog.Logger) *Server {
	a := &Server{
		api:    newApi(db, nc, log),
		engine: configEngine(log),
	}
	a.createRoutes()
	return a

}

func configEngine(log zerolog.Logger) *gin.Engine {
	ginLogger := logger.SetLogger(logger.Config{
		Logger: &log,
		UTC:    true,
	})
	engine := gin.New()
	engine.Use(ginLogger, corsMiddleware)
	return engine
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
	a.engine.Run(port)
}

var corsMiddleware = func(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
