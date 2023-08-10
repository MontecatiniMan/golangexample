package action

import (
	state "golangexample/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StateHandler(state *state.State) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"result": true,
			"info":   state.GetState(),
		})
	}
}
