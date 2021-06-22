package handlers

import (
	"ChatApp/internal/data"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Set(mongoRepo data.MongoRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body interface{}
		if err := ctx.BindJSON(&body); err != nil {
			logrus.Error(err)
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		err := mongoRepo.Set(body)
		if err != nil {
			logrus.Error(err)
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"insertStatus": "OK",
		})
	}
}
