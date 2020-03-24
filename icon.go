package fontawesome

// Icon is a Font Awesome icon.
type Icon struct {
	Changes   []string `json:"changes"`
	Ligatures []string `json:"ligatures"`
	Search    struct {
		Terms []string `json:"terms"`
	} `json:"search"`
	Styles  []string       `json:"styles"`
	Unicode string         `json:"unicode"`
	Label   string         `json:"label"`
	Voted   bool           `json:"voted"`
	SVG     map[string]SVG `json:"svg"`
	Free    []string       `json:"free"`
}

// SVG is a Font Awesome icon's SVG data.
type SVG struct {
	LastModified int      `json:"last_modified"`
	Raw          string   `json:"raw"`
	ViewBox      []string `json:"viewBox"`
	Width        int      `json:"width"`
	Height       int      `json:"height"`
	Path         string   `json:"path"`
}

// TemplateIcon is how a Font Awesome icon is defined in a template.Template.
type TemplateIcon struct {
	Name   string `json:"name"`
	Prefix string `json:"prefix"`
}
