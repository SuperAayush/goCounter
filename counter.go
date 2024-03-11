package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type PageData struct {
	Results bool
	Words   int
	Letters int
	Text    string
}

func main() {
	req := gin.Default()
	req.LoadHTMLGlob("ui/html/*")

	req.Static("/ui", "./ui")
	req.GET("/", getRootHandler)
	req.POST("/", postRootHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	req.Run(":" + port)
}

func getRootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func postRootHandler(c *gin.Context) {
	inputText := c.PostForm("inputText")
	wordsCount := len(strings.Fields(inputText))
	letterCount := len(strings.ReplaceAll(inputText, " ", ""))

	data := PageData{
		Results: true,
		Words:   wordsCount,
		Letters: letterCount,
		Text:    inputText,
	}

	c.HTML(http.StatusOK, "index.html", data)
}
