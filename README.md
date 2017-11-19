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
  //Create an string array containing your menu items.  each aray item should contain the url followed by the link name delimited  by a pipe "|"
  mymenuitems := []string("/settings|settings", "/about|about", "/logout|logout")
  //Create an lcars.Menu struct containing your array of menu items
  mymenu := lcars.Menu{Items: mymenuitems}
  //create a variable containing a string of the html you want in the main area of the page
  content := `
    <h1> Here's a header</h1>
    and here's some text
   `
   //run the lcars.MakePage function including the content and menu struct as attributes
   t, createPage := lcars.MakePage(content, mymenu)
   //Execute the page
   t.Execute(w, createPage)
```
