package action

import (
	"golangexample/consts"
	state "golangexample/model"
	"golangexample/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateHandler(state *state.State) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if state.GetState() == consts.StateEmpty || state.GetState() == consts.StateOk {
			state.SetState(consts.StateUpdating)

			go func() {
				service.Update()
				state.SetState(consts.StateOk)
			}()

			ctx.JSON(http.StatusOK, gin.H{
				"result": true,
				"info":   "",
				"code":   http.StatusOK,
			})

			return
		}

		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"result": false,
			"info":   "Service unavailable",
			"code":   http.StatusServiceUnavailable,
		})
	}
}
