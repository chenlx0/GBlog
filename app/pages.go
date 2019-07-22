package app

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/chenlx0/GBlog/markdown"
)

const (
	HEADER = "templates/header.html"
	HOME   = "templates/home.html"
	ITEM   = "templates/item.html"
	FOOTER = "templates/footer.html"
	ABOUT  = "templates/about.html"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(HOME, HEADER, FOOTER)
	if err != nil {
		w.Write([]byte("Parse template error"))
		w.WriteHeader(500)
		return
	}

	data := struct {
		Title       string
		ArticleList markdown.ArticleList
	}{
		Title:       "俺的博客",
		ArticleList: articleCache,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.Write([]byte("Execute template error"))
		w.WriteHeader(500)
		return
	}
}

func article(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(ITEM, HEADER, FOOTER)
	if err != nil {
		w.Write([]byte("Parse template error"))
		w.WriteHeader(500)
		return
	}

	paths := strings.Split(r.URL.Path, "/")
	articleTitle := paths[len(paths)-1]
	for _, v := range articleCache {
		if v.Title == articleTitle {
			tmpl.Execute(w, v)
			return
		}
	}

	w.WriteHeader(404)
	w.Write([]byte("Sorry, page is not found"))
}

func about(w http.ResponseWriter, r *http.Request) {

}
