package main

import (
	//"github.com/russross/blackfriday"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {

	//fmt.Print(BlogPosts)
	d3 := BlogPosts[1]
	fmt.Println("Title:")
	fmt.Println(d3.Title)
	fmt.Println("Subitle:")
	fmt.Println(d3.Subtitle)
	fmt.Println("Date:")
	fmt.Println(d3.Date)
	fmt.Println("BodyString:")
	fmt.Print(d3.BodyString)
	fmt.Println("")
	router := gin.Default()
	router.StaticFile("/", "index.html")
	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")
	router.GET("/blog", blogIndexHandler)
	router.GET("blog/:post", blogPostHandler)
	router.Run(":3000")
}
