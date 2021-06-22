package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Noam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Do you like it Baby? :D")
	}
}
