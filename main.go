package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()
	router.StaticFile("/", "index.html")
	router.Static("/static", "./static")
	router.Run(":3000")
}
