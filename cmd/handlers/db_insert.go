package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Set(s string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//templates, err := tmplSrv.ListTemplates()
		//if err != nil {
		//	logrus.Error(err.Error())
		//	ctx.String(http.StatusInternalServerError, err.Error())
		//	return
		//}

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"list": "banana",
		})
	}
}
