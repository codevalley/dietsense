package main

import (
	"dietsense/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.SetupRoutes(router)
	router.Run(":8080") // listens and serves on 0.0.0.0:8080
}
