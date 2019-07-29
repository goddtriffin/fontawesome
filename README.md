# Golang Font Awesome

An open-source Go(lang) library for server side rendering Font Awesome icons.

## Setup

1) Download the [Font Awesome for Desktop](https://fontawesome.com/how-to-use/on-the-desktop/setup/getting-started) package.
2) Once downloaded, find the `icons.json` file in the `metadata/` sub-directory (full path: `fontawesome-free-X.X.X-desktop/metadata/icons.json`).
3) Copy the `icons.json` file into your repository. (e.g. `repo/static/`)

## How to use

```go
package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/MagnusFrater/fontawesome"
)

func main() {
	// initialize a Font Awesome library by loading the `icons.json` file
	fa, err := fontawesome.New("static/icons.json")
	if err != nil {
		log.Fatalln(err)
	}

	// add the library's `SVG` function to your template FuncMap
	funcMap := template.FuncMap{
		"fontawesome": fa.SVG,
	}

	// call the 'fontawesome' FuncMap function in your templates
	// it takes two parameters:
	// 1) icon prefix (e.g. "fab" = brands, "fal" = light, "far" = regular, "fas" = solid)
	// 2) icon name
	const exampleTemplate = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Go(lang) SSR Font Awesome library</title>
    <meta name="description" content="Go(lang) SSR Font Awesome library">
	<meta name="author" content="Todd Griffin">

    <style>svg{width:5em;height:5em;}</style>
  </head>
  <body>
    {{fontawesome "fas" "home"}}
    {{fontawesome "fas" "hamburger"}}
    {{fontawesome "fas" "dice"}}
  </body>
</html>
`

	// Create a template, add the function map, and parse the text.
	tmpl, err := template.New("fontawesomeTest").Funcs(funcMap).Parse(exampleTemplate)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// create a simple http route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// run the example server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```
