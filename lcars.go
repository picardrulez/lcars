package lcars

import (
	"text/template"
)

func MakePage(content string) (*template.Template, Page) {
	t := template.New("template")
	createPage := Page{Content: content}
	pageTemplate := pageBuilder()
	t, _ = t.Parse(pageTemplate)
	return t, createPage
}
