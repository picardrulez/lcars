package lcars

import (
	"math/rand"
	"strings"
	"time"
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
		var thiscolor string
		if menucolor == "random" {
			thiscolor = pickColor()
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

func pickColor() string {
	colors := []string{"pale-canary", "danub", "hopbush"}
	rand.Seed(time.Now().Unix())
	randnum := rand.Intn(3-1) + 1
	return colors[randnum]
}
