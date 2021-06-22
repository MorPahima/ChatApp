package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Noam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Just test to deploy with github")
	}
}
