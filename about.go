package main

import (
	"bufio"
	"github.com/russross/blackfriday"
	"gopkg.in/gin-gonic/gin.v1"
	"html/template"
	"log"
	"net/http"
	"os"
)

type About struct {
	PageTitle string
	AboutHTML template.HTML
}

var about About

func init() {
	about.PageTitle = "Dave J Kane - About"
	var filestring string
	file, err := os.Open("about/about.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		filestring += scanner.Text() + "\n"
	}
	about.AboutHTML = template.HTML(blackfriday.MarkdownBasic([]byte(filestring)))
}

func aboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "about.tmpl", about)
}
