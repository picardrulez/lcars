# lcars
A Go library that builds html pages styled similarly to the LCARS user interface from the 24th Century.  The LCARS CSS template is based off of https://github.com/joernweissenborn/lcars

## Example usage:

```
package main
import (
  "github.com/picardrulez/lcars"
  "net/http"
 )
 
 func main() {
  http.HandleFunc("/", exampleHandler)
  http.ListenAndServe(":8000", nil)
 }
 
func exampleHandler(w http.ResponseWriter, r *http.Request) {
  mymenuitems := []string("/settings|settings", "/about|about", "/logout|logout")
  mymenu := lcars.Menu{Items: mymenuitems}
  content := `
    <h1> Here's a header</h1>
    and here's some text
   `
   t, createPage := lcars.MakePage(content, mymenu)
   t.Execute(w, createPage)
```
