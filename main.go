package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type IndexPage struct {
	PageTitle string
}

func main() {

	router := gin.Default()

	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")

	router.GET("/", indexHandler)
	router.GET("/blog", blogIndexHandler)
	router.GET("/blog/:post", blogPostHandler)
	router.GET("/about", aboutHandler)

	router.Run(":3000")

}

func indexHandler(c *gin.Context) {
	index := IndexPage{PageTitle: "Dave J Kane - Personal Website"}
	c.HTML(http.StatusOK, "index.tmpl", index)
}
