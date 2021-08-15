package main

import (
	"github.com/aURORA-JC/headerfaker/data"
	"github.com/aURORA-JC/headerfaker/router"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load test data
	color.Green("Loading data...")
	data.ReadData()

	// Close colorful console
	gin.DisableConsoleColor()

	// Set routers
	color.Green("Setting routers...")
	r := router.SetRouters()

	// Start program
	color.Green("Starting http service...")
	err := r.Run("0.0.0.0:11451")
	if err != nil {
		color.Red("Fail! Address may be occupied!")
		return
	}
}
