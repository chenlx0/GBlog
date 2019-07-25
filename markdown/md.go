package markdown

import (
	"io/ioutil"
	"sort"
	"strings"
	"time"

	"gopkg.in/russross/blackfriday.v2"
)

// Article represent each article in the blog
type Article struct {
	Title       string
	ID          string
	Date        time.Time
	Categories  []string
	Description string
	Content     string
	HTML        string
}

type ArticleList []*Article

func (l ArticleList) Len() int {
	return len(l)
}

func (l ArticleList) Less(i, j int) bool {
	return l[i].Date.Unix() > l[j].Date.Unix()
}

func (l ArticleList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func md2html(text string) string {
	content := []byte(text)
	res := blackfriday.Run(content)
	return string(res)
}

// RetrieveArticle initialize a markdown format article
func RetrieveArticle(text string) *Article {
	// basic information start with '---'
	if !strings.HasPrefix(text, "---") {
		return nil
	}

	res := &Article{}

	// basic information end with '---'
	basicInfoBytes := make([]byte, 0)
	for i := 3; i < len([]byte(text))-3; i++ {
		x := text[i]
		if x == '-' && text[i+1] == '-' && text[i+2] == '-' {
			// left are content
			res.Content = text[i+3:]
			break
		}
		basicInfoBytes = append(basicInfoBytes, x)
	}

	// split into lines
	basicInfoSlice := strings.Split(string(basicInfoBytes), "\n")

	for _, line := range basicInfoSlice {
		if strings.HasPrefix(line, "title: ") {
			res.Title = strings.TrimPrefix(line, "title: ")
		} else if strings.HasPrefix(line, "description: ") {
			res.Description = strings.TrimPrefix(line, "description: ")
		} else if strings.HasPrefix(line, "date: ") {
			date := strings.TrimPrefix(line, "date: ")
			res.Date, _ = time.Parse("2006-01-02", date)
		} else if strings.HasPrefix(line, "categories") {
			categoryString := strings.TrimPrefix(line, "categories: ")
			res.Categories = strings.Split(categoryString, ",")
		}
	}

	res.HTML = md2html(res.Content)

	return res
}

// LoadArticles load all articles from specified dir
func LoadArticles(dirname string) (ArticleList, error) {
	res := make(ArticleList, 0)

	files, _ := ioutil.ReadDir(dirname)
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".md") {
			continue
		}
		content, err := ioutil.ReadFile(dirname + f.Name())
		if err != nil {
			return nil, err
		}
		article := RetrieveArticle(string(content))
		article.ID = f.Name()
		if article != nil {
			res = append(res, article)
		}
	}

	// sort articles by date
	sort.Sort(res)

	return res, nil
}
