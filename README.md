# Go(lang) SSR Font Awesome library

An open-source [Go(lang)](https://golang.org/) library for server-side-rendering [Font Awesome](https://fontawesome.com/) icons.

## Setup

1. `go get -u github.com/MagnusFrater/fontawesome`
2. Download the [Font Awesome for Desktop](https://fontawesome.com/how-to-use/on-the-desktop/setup/getting-started) package.
3. Once downloaded, find `metadata/icons.json` (full path: `fontawesome-free-X.X.X-desktop/metadata/icons.json`).
4. Copy the `icons.json` file into your repository (e.g. `repo/static/`).

## Example

1. `git clone https://github.com/MagnusFrater/fontawesome`
2. `cd examples/example-server/`
3. `go run example-server.go`
4. Visit `localhost:8080` in your favourite browser!

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

	// create a FuncMap and create a new mapping for the new Font Awesome library's `SVG` function
	funcMap := template.FuncMap{
		"fontawesome": fa.SVG,
	}

	// call the 'fontawesome' mapped function in your template(s); it takes two parameters:
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
    {{fontawesome "fas" "american-sign-language-interpreting"}}
  </body>
</html>
`
	// create a template, add the `funcMap`, and parse the `exampleTemplate`
	tmpl, err := template.New("fontawesomeTest").Funcs(funcMap).Parse(exampleTemplate)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// create a simple http route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// run the example server
	log.Println("Listening on localhost:8080!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Result

![Example Server Result](/assets/example-server-result.png?raw=true "Example Server Result")

## Credits

- [Todd Griffin](https://github.com/MagnusFrater) - Author

## License

[MIT](/LICENSE)
