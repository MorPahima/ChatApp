package cmd

import (
	"ChatApp/cmd/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config          Config
}

func NewServer(config Config) *Server {
	return &Server{config: config}
}

func (s *Server) setupServer() *gin.Engine {
	r := gin.Default()
	r.GET("/health", handlers.Health())
	r.GET("/", handlers.Noam())
	//r.POST("/reports/export/pdf/*reportId", handlers.ReportExportHandler(s.reportService, "pdf"))
	return r
}

func (s *Server) Start() error {
	srv := s.setupServer()
	return srv.Run(fmt.Sprintf(":%d", s.config.ServerPort))
}
