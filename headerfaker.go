package main

import (
	"github.com/aURORA-JC/headerfaker/data"
	"github.com/aURORA-JC/headerfaker/router"
	"github.com/gin-gonic/gin"
)

func main() {
	data.ReadData()
	gin.DisableConsoleColor()

	r := router.SetRouter()

	err := r.Run("0.0.0.0:80")
	if err != nil {
		return
	}

}
