package writer

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/agonzalezro/go-polo/parser"
)

type Site struct {
	Config          map[string]string // TODO
	Pages, Articles []parser.ParsedFile
	OutputPath      string

	TemporalParsedFile parser.ParsedFile // TODO: This is not the best option, but it does the work
}

func (site Site) writeIndex() {
	indexFile := fmt.Sprintf("%s/index.html", site.OutputPath)

	template := template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))

	file, err := os.Create(indexFile)
	if err != nil {
		log.Panic(err)
	}
	template.ExecuteTemplate(file, "base", site)
}

func (site Site) writeArticle(article parser.ParsedFile) {
	filePath := fmt.Sprintf("%s/%s.html", site.OutputPath, article.Metadata["slug"])

	template := template.Must(template.ParseFiles("templates/article.html", "templates/base.html"))

	file, err := os.Create(filePath)
	if err != nil {
		log.Panic(err)
	}
	site.TemporalParsedFile = article
	template.ExecuteTemplate(file, "base", site)
}

func (site Site) writePage(page parser.ParsedFile) {
	// First of all create the pages/ folder if it doesn't exist
	pagesPath := fmt.Sprintf("%s/pages", site.OutputPath)
	if _, err := os.Stat(pagesPath); os.IsNotExist(err) {
		os.Mkdir(pagesPath, 0777)
	}

	filePath := fmt.Sprintf("%s/%s.html", pagesPath, page.Metadata["slug"])

	template := template.Must(template.ParseFiles("templates/page.html", "templates/base.html"))

	file, err := os.Create(filePath)
	if err != nil {
		log.Panic(err)
	}
	site.TemporalParsedFile = page
	template.ExecuteTemplate(file, "base", site)
}

func (site Site) WriteSite() {
	site.writeIndex()

	for _, article := range site.Articles {
		site.writeArticle(article)
	}

	for _, page := range site.Pages {
		site.writePage(page)
	}
}
