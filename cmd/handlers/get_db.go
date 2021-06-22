package handlers

import (
	"ChatApp/internal/data"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetDB(mongoRepo data.MongoRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dbData, err := mongoRepo.GetAll()
		if err != nil {
			logrus.Error(err)
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": dbData,
		})
	}
}
