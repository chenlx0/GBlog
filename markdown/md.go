package markdown

import (
	"strings"
	"time"
)

// Article represent each article in the blog
type Article struct {
	Title       string
	Date        time.Time
	Categories  []string
	Description string
	Content     string
	HTML        string
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
		}
		if strings.HasPrefix(line, "description: ") {
			res.Description = strings.TrimPrefix(line, "description: ")
		}
		if strings.HasPrefix(line, "date: ") {
			date := strings.TrimPrefix(line, "date: ")
			res.Date, _ = time.Parse("2006-01-02", date)
		}
	}

	return res
}
