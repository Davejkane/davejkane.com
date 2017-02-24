package main

import (
	"bufio"
	"github.com/russross/blackfriday"
	"gopkg.in/gin-gonic/gin.v1"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
)

type BlogPost struct {
	PageTitle  string
	Slug       string
	Title      string
	Subtitle   string
	Date       string
	BodyString string
	BodyHTML   template.HTML
}

type BlogIndex struct {
	BlogPostList BlogPostList
	PageTitle    string
}

type BlogPostList []BlogPost

func (slice BlogPostList) Len() int {
	return len(slice)
}

func (slice BlogPostList) Less(i, j int) bool {
	return slice[j].Date < slice[i].Date
}

func (slice BlogPostList) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

var BlogPosts []BlogPost

func init() {
	BlogPosts = parseAllBlogPosts("blog")
}

//Takes a markdown file and puts it into a struct to be used to render the blog post into the blogpost template.
func parseBlogPost(filename string) BlogPost {
	var bp BlogPost

	//The slug is the file name minus the .md extension, and will be used as the url for the blog post after /blog/
	bp.Slug = filename[:len(filename)-3]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//Open a scanner on the file and read the first line
	scanner.Scan()
	//The first line in the markdown file has to be the title of the blog post. Hard coded.
	bp.Title = scanner.Text()
	scanner.Scan()
	//After scanning the next line, parse it into the subtitle, again hardcoded.
	bp.Subtitle = scanner.Text()
	scanner.Scan()
	//Same for the date. The data ***MUST BE IN FORMAT*** YYYY-MM-DD!!!!
	bp.Date = scanner.Text()
	//Scan the rest of the lines to the end of the file for the blog post body.
	for scanner.Scan() {
		bp.BodyString += scanner.Text() + "\n"
	}
	bp.BodyHTML = template.HTML(blackfriday.MarkdownBasic([]byte(bp.BodyString)))
	bp.PageTitle = "Dave J Kane " + bp.Title
	return bp
}

//Takes a folder and parses all of the files using parseBlogPost(). Folder should only contain .md files.
//It returns a map of url slugs to BlogPost structs to be stored in memory on startup for quick rendering.
func parseAllBlogPosts(directory string) BlogPostList {
	var blogPosts BlogPostList
	filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
		bp := parseBlogPost(path)
		blogPosts = append(blogPosts, bp)
		return nil
	})
	sort.Sort(blogPosts)
	return blogPosts
}

func blogIndexHandler(c *gin.Context) {
	bi := BlogIndex{PageTitle: "Dave J Kane - Blog", BlogPostList: BlogPosts}
	c.HTML(http.StatusOK, "blogindex.tmpl", bi)
}

func blogPostHandler(c *gin.Context) {
	filename := "blog/" + c.Param("post") + ".md"
	bp := parseBlogPost(filename)
	c.HTML(http.StatusOK, "blogpost.tmpl", bp)
}
