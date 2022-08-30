package html

import "io"

type Renderer interface {
	Node
	Render(w io.Writer) error
}
