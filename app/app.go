package app

import (
	"log"
	"net/http"

	"github.com/chenlx0/GBlog/config"

	"github.com/chenlx0/GBlog/markdown"
)

var articleCache markdown.ArticleList
var globalConf *config.Conf

func init() {
	articleCache, _ = markdown.LoadArticles("articles/")
}

func Run() {
	http.HandleFunc("/articles/", article)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
