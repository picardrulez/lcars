package lcars

import (
	"strings"
)

type Menu struct {
	Items []string
}

func makeMenu(menu Menu) string {
	content := `
	<div class="lcars-row">
		<div class="lcars-column">
	`
	for _, k := range menu.Items {
		newitem := strings.Split(k, ":")
		content = content + `
		<a href="`
		content = content + newitem[0]
		content = content + `" class=lcars-element button lcars-blue-bell-bg">`
		content = content + newitem[1]
		content = content + `</a>`
	}
	content = content + `
	</div>
	</div>
	`
	return content
}
