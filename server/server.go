package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	nats "github.com/nats-io/go-nats"
	"github.com/rs/zerolog"
)

// Server hides gin engine
type Server struct {
	engine *gin.Engine
}

// New creates a new server
func New(db *sqlx.DB, nc *nats.Conn, log zerolog.Logger) *Server {
	s := &Server{engine: newEngine(log)}
	s.createRoutes(newAPI(db, nc, log))
	return s
}

func newEngine(log zerolog.Logger) *gin.Engine {
	ginLogger := logger.SetLogger(logger.Config{
		Logger: &log,
		UTC:    true,
	})
	engine := gin.New()
	engine.Use(ginLogger, cors.Default())
	return engine
}

func (s *Server) createRoutes(api *api) {
	baseRoute := s.engine.Group("/api")
	roleRoutes := baseRoute.Group("/role")
	{
		roleRoutes.POST("", api.role.Save)
		roleRoutes.PUT("", api.role.Update)
		roleRoutes.GET("", api.role.Paginate)
		roleRoutes.DELETE("", api.role.Delete)
	}
	departmentRoutes := baseRoute.Group("/department")
	{
		departmentRoutes.POST("", api.department.Save)
		departmentRoutes.PUT("", api.department.Update)
		departmentRoutes.GET("", api.department.Paginate)
		departmentRoutes.DELETE("", api.department.Delete)
	}
	userRoutes := baseRoute.Group("/user")
	{
		userRoutes.POST("", api.user.Save)
		userRoutes.PUT("", api.user.Update)
		userRoutes.GET("", api.user.Paginate)
		userRoutes.DELETE("", api.user.Delete)
	}
}

// Run launch starts the server
func (s *Server) Run(port string) {
	s.engine.Run(port)
}

