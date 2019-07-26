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
	path := r.URL.Path
	if path != "/" && path != "" {
		w.WriteHeader(404)
		w.Write([]byte("Page not found"))
		return
	}

	tmpl, err := template.ParseFiles(HOME, HEADER, FOOTER)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Parse template error"))
		return
	}

	data := struct {
		Title       string
		ArticleList markdown.ArticleList
	}{
		Title:       globalConf.Blog.Title,
		ArticleList: articleCache,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Execute template error"))
		return
	}
}

func article(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(ITEM, HEADER, FOOTER)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Parse template error"))
		return
	}

	paths := strings.Split(r.URL.Path, "/")
	articleID := paths[len(paths)-1]
	for _, v := range articleCache {
		if v.ID == articleID {
			tmpl.Execute(w, v)
			return
		}
	}

	w.WriteHeader(404)
	w.Write([]byte("Sorry, page is not found"))
}

func about(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(ITEM, HEADER, FOOTER)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Parse template error"))
		return
	}

	err = tmpl.Execute(w, *aboutPage)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Execute template error"))
	}
}
