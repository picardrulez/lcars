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
	for _, k := range menu.Items {
		newitem := strings.Split(k, "|")
		content = content + `
		<a href="`
		content = content + newitem[0]
		content = content + `" style="text-decoration:none" class="lcars-element button lcars-` + menucolor + `-bg">`
		content = content + newitem[1]
		content = content + `</a>`
	}
	content = content + `
	</div>
	</div>
	`
	return content
}
