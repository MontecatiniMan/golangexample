package main

import (
	"golangexample/action"
	"golangexample/consts"
	state "golangexample/model"
	"golangexample/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	state := state.NewState(consts.StateEmpty)
	router := gin.Default()

	repository.Init()

	router.GET("/update", action.UpdateHandler(state))
	router.GET("/state", action.StateHandler(state))
	router.GET("/get_names", action.NamesHandler())
	router.Run()
}
