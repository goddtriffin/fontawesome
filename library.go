package fontawesome

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
)

var styleMap = map[string]string{
	"fab": "brands",
	"fal": "light",
	"far": "regular",
	"fas": "solid",
}

// Library is a container for Font Awesome icons.
type Library struct {
	Path  string
	icons map[string]Icon
}

// New returns a new Font Awesome Library loaded with data from the given path.
func New(path string) (fa *Library, err error) {
	var library Library

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return &library, err
	}

	if err := json.Unmarshal(dat, &library.icons); err != nil {
		return &library, err
	}

	library.Path = path
	return &library, nil
}

// Icon returns the Font Awesome icon with the given name.
func (fa *Library) Icon(name string) (Icon, error) {
	icon, ok := fa.icons[name]
	if !ok {
		return Icon{}, fmt.Errorf("Font Awesome icon doesn't exist with name: '%v'", name)
	}

	return icon, nil
}

// SVG returns the given icon as a raw SVG element.
func (fa *Library) SVG(prefix, name string) (template.HTML, error) {
	icon, err := fa.Icon(name)
	if err != nil {
		return template.HTML(""), err
	}

	style, ok := styleMap[prefix]
	if !ok {
		return template.HTML(""), fmt.Errorf("No such Font Awesome icon style: '%v'", prefix)
	}

	svg, ok := icon.SVG[style]
	if !ok {
		return template.HTML(""), fmt.Errorf("Font Awesome '%v' icon doesn't have style: '%v'", name, prefix)
	}

	return template.HTML(svg.Raw), nil
}
