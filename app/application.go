package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication for handling http request
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}