package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main()  {
	gin.SetMode(gin.ReleaseMode)

	// Load Env
	var port, adminPath string
	_ = godotenv.Load()

	port = os.Getenv("PORT")
	if len(port) == 0 {
		port = "9527"
	}
	adminPath = os.Getenv("ADMIN_PATH")
	if len(adminPath) == 0 {
		adminPath = "/admin"
	}

	// Run server
	router := gin.Default()
	router.LoadHTMLFiles("dist/index.html")
	router.GET(adminPath, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.Static("/em", "./dist/em")

	_ = router.Run(":" + port)
}
