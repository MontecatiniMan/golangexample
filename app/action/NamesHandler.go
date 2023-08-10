package action

import (
	"golangexample/component"
	"golangexample/entity"
	"golangexample/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NamesHandler() gin.HandlerFunc {
	conn := component.DbConnection()

	return func(ctx *gin.Context) {
		result := []entity.SdnEntity{}

		if ctx.Query("type") == "strong" {
			result = repository.SearchStrong(conn, ctx.Query("name"))
		} else {
			result = repository.SearchWeak(conn, ctx.Query("name"))
		}

		ctx.JSON(http.StatusOK, result)
	}
}
