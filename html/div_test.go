package html

import (
	"log"
	"strings"
	"testing"
)

func TestDiv(t *testing.T) {
	var out strings.Builder
	one := Div(WithID("first-child"), WithText("Hello, World"))
	two := Div(WithID("second-child"), WithText("Second, Hello World!"))
	three := Div(WithID("third-child"), WithText("Third, Hello World"))
	elm := Div(WithClassName("my-class"), WithID("my-id"), WithChildren(one, two, three))
	if err := elm.Render(&out); err != nil {
		log.Fatal(err)
	}

	if len(out.String()) == 0 {
		t.Fail()
	}
	t.Log(out.String())
}
