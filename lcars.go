package lcars

import (
	"text/template"
)

type Settings struct {
	Title string
	Color string
	Menu  bool
}

func MakePage(content string, mymenu Menu, mysettings Settings) (*template.Template, Page) {
	t := template.New("template")
	createPage := Page{Content: content}
	pageTemplate := pageBuilder(mymenu, mysettings)
	t, _ = t.Parse(pageTemplate)
	return t, createPage
}
