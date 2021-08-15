package router

import (
	"github.com/aURORA-JC/headerfaker/controller"
	"github.com/aURORA-JC/headerfaker/middleware"
	"github.com/gin-gonic/gin"
)

// SetRouters Set HeaderFaker Routers
func SetRouters() *gin.Engine {
	r := gin.Default()

	// Use Middlewares
	r.Use(middleware.ServerType())

	// Load HTML template files
	r.LoadHTMLGlob("template/**/*")

	// Index router
	r.GET("/", controller.IndexHandler)

	// Tests routers
	tests := r.Group("/test")
	{
		tests.GET("/getresponses", controller.GetResponsesHandler)

		tests.GET("/easypassword", controller.EasyPasswordGETHandler)
		tests.POST("/easypassword", controller.EasyPasswordPOSTHandler)

		tests.GET("/getorpost", controller.GetHandler)
		tests.POST("/getorpost", controller.PostHandler)

		tests.GET("/fakerefer", controller.FakeReferHandler)

		tests.GET("/fakeagent", controller.FakeAgentHandler)
	}

	// Keys Checker Router
	r.POST("/check", controller.CheckHandler)

	return r
}
