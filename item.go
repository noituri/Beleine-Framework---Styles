package beleine

import "fmt"

type StyleItem struct {
	name string
	attr []Attr
}

type Attr struct {
	name  string
	value string
}

func NewStyleItem(name string) StyleItem {
	return StyleItem{name: name}
}

func (si *StyleItem) AddAttr(name, value string) {
	si.attr = append(si.attr, Attr{name: name, value: value})
}

func (si *StyleItem) Render() string {
	attr := ""
	for _, v := range si.attr {
		attr += fmt.Sprintf(`%s: %s;`, v.name, v.value)
	}

	return fmt.Sprintf(`
%s {%s}
`, si.name, attr)
}
