package beleine

import (
	"testing"
	"fmt"
)

func  TestStyle(t *testing.T)  {
	style := NewStyle()
	item := NewStyleItem("h1")
	item.AddAttr("color", "white")
	item.AddAttr("text-align", "center")
	style.AddItem(item)

	if style.Render() != fmt.Sprintf(`
<style>

h1 {color: white;text-align: center;}

</style>
`) {t.Fail()}

	item2 := NewStyleItem("body")
	item2.AddAttr("font-size", "12")

	style.AddItem(item2)

	if style.Render() != fmt.Sprintf(`
<style>

h1 {color: white;text-align: center;}

body {font-size: 12;}

</style>
`) {t.Fail()}

}
