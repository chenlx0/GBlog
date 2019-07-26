package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chenlx0/GBlog/config"

	"github.com/chenlx0/GBlog/markdown"
)

var articleCache markdown.ArticleList
var aboutPage *markdown.Article
var globalConf *config.Config

func init() {
	var err error
	articleCache, err = markdown.LoadArticles("articles/")
	if err != nil {
		panic(err)
	}

	aboutContent, err := ioutil.ReadFile("custom_page/about.md")
	if err != nil {
		panic(err)
	}

	aboutPage = markdown.RetrieveArticle(string(aboutContent))

	globalConf, err = config.FromYamlFile("conf.yaml")
	if err != nil {
		panic(err)
	}
}

func Run() {
	http.HandleFunc("/articles/", article)
	http.HandleFunc("/about", about)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", globalConf.Server.Host, globalConf.Server.Port), nil))
}
