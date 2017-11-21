package lcars

import (
	"strings"
)

type Menu struct {
	Items []string
}

func makeMenu(menu Menu, settings Settings) string {
	menucolor := settings.MenuColor
	content := `
	<div class="lcars-row">
		<div class="lcars-column">
	`
	for i, k := range menu.Items {
		var thiscolor string
		if menucolor == "random" {
			thiscolor = pickColor(i)
		} else {
			thiscolor = menucolor
		}
		newitem := strings.Split(k, "|")
		content = content + `
		<a href="`
		content = content + newitem[0]
		content = content + `" style="text-decoration:none" class="lcars-element button lcars-` + thiscolor + `-bg">`
		content = content + newitem[1]
		content = content + `</a>`
	}
	content = content + `
	</div>
	</div>
	`
	return content
}

func pickColor(iteration int) string {
	colors := []string{"pale-canary", "danub", "hopbush"}
	numpicked := false
	for numpicked == false {
		if iteration <= len(colors) {
			numpicked = true
		} else {
			for iteration <= len(colors) {
				iteration = iteration - len(colors)
				if iteration <= len(colors) {
					numpicked = true
				}
			}
		}
	}
	return colors[iteration]
}
