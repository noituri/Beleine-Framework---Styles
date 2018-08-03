package beleine

import (
	"github.com/kubastick/BeleineFramework"
	"fmt"
)

type Style struct {
	id 	  string
	items []StyleItem
}

func NewStyle() Style {
	return Style{id: beleine.GetGlobalID()}
}

func (s *Style) AddItem(item StyleItem) {
	s.items = append(s.items, item)
}

func (s *Style) Render() string {
	items := ""

	for _, v := range s.items {
		items += v.Render()
	}

	return fmt.Sprintf(`
<style>
%s
</style>
`, items)
}