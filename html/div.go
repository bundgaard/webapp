package html

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type Component struct {
	ID        string
	ClassName string
	Children  []Node
}

type TextNode struct {
	text string
}

func (tn TextNode) getID() string {
	return ""
}
func (tn TextNode) elementNode() {}

type div struct {
	*Component
}

func (d *div) getID() string {
	return d.ID
}

func (d *div) elementNode() {
	// So we can have a bunch of nodes []Node
}

func Div(opts ...Option) Renderer {
	d := &div{
		&Component{Children: []Node{}},
	}
	for _, opt := range opts {
		if err := opt(d.Component); err != nil {
			log.Println(err)
		}
	}
	return d
}

func (d *div) String() string {
	var out strings.Builder
	out.WriteString("<div")
	if d.ID != "" {
		out.WriteString(fmt.Sprintf(" id=%q", d.ID))
	}
	if d.ClassName != "" {
		out.WriteString(fmt.Sprintf(" class=%q", d.ClassName))
	}

	out.WriteString(">")
	for idx := range d.Children {
		child := d.Children[idx]

		out.WriteString("\n")
		out.WriteString("\t")
		switch v := child.(type) {
		case TextNode:
			out.WriteString("\t")
			out.WriteString(v.text)
		case *div:
			out.WriteString(v.String())
		default:
			log.Printf("not added for this type: %T\n", v)
		}
	}
	out.WriteString("\n")
	out.WriteString("</div>")
	return out.String()
}
func (d *div) Render(w io.Writer) error {

	_, err := fmt.Fprintf(w, "%s", d.String())
	if err != nil {
		return err
	}
	return nil
}
