package lcars

import (
	"text/template"
)

func MakePage(content string, mymenu Menu) (*template.Template, Page) {
	t := template.New("template")
	createPage := Page{Content: content}
	pageTemplate := pageBuilder(mymenu)
	t, _ = t.Parse(pageTemplate)
	return t, createPage
}
