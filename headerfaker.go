package main

import (
	"fmt"
	"github.com/aURORA-JC/headerfaker/config"
	"github.com/aURORA-JC/headerfaker/data"
	"github.com/aURORA-JC/headerfaker/database"
	"github.com/aURORA-JC/headerfaker/router"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

const version string = "0.2.0-dev"

func main() {
	color.Green("HeaderFaker v%s is starting...", version)
	// Read config
	configError := config.InitConfig("./config.ini")
	if configError != nil {
		return
	}

	// Load test data
	color.Green("Loading data...")
	data.ReadData()

	// Connect Database
	dbError := database.InitDataBase(config.Config.MySQLConfig)
	if dbError != nil {
		return
	}

	// Set gin Mode
	if config.Config.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	// Set routers
	color.Green("Setting routers...")
	r := router.SetRouters()

	// Start program
	address := fmt.Sprint("0.0.0.0", ":", config.Config.Port)
	color.Green(fmt.Sprint("Starting http service at ", address, "..."))
	err := r.Run(address)
	if err != nil {
		color.Red("Failed! Address may be occupied!")
		return
	}
}
