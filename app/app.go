package app

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chenlx0/GBlog/config"

	"github.com/chenlx0/GBlog/markdown"
)

var articleCache markdown.ArticleList
var aboutPage *markdown.Article
var globalConf *config.Conf

func init() {
	articleCache, _ = markdown.LoadArticles("articles/")
	aboutContent, _ := ioutil.ReadFile("custom_page/about.md")
	aboutPage = markdown.RetrieveArticle(string(aboutContent))
}

func Run() {
	http.HandleFunc("/articles/", article)
	http.HandleFunc("/about", about)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
