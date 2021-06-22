package cmd

import (
	"ChatApp/cmd/handlers"
	"ChatApp/internal/data"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config          Config
	mongoRepo data.MongoRepository
}

func NewServer(config Config, mongoRepo data.MongoRepository) *Server {
	return &Server{config: config, mongoRepo: mongoRepo}
}


func (s *Server) setupServer() *gin.Engine {
	r := gin.Default()
	r.GET("/health", handlers.Health())
	r.GET("/", handlers.Noam())
	r.GET("/get-full-db", handlers.GetDB(s.mongoRepo))
	r.POST("/insert-user", handlers.Set(s.mongoRepo))
	//r.POST("/reports/export/pdf/*reportId", handlers.ReportExportHandler(s.reportService, "pdf"))
	return r
}

func (s *Server) Start() error {
	srv := s.setupServer()
	return srv.Run(fmt.Sprintf(":%d", s.config.ServerPort))
}
