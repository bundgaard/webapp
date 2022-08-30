package html

type Option func(component *Component) error

func WithID(id string) Option {
	return func(c *Component) error {
		c.ID = id
		return nil
	}
}
func WithChildren(nodes ...Node) Option {
	return func(c *Component) error {
		for _, child := range nodes {
			c.Children = append(c.Children, child)
		}
		return nil
	}
}
func WithText(textNode string) Option {
	return func(c *Component) error {
		c.Children = append(c.Children, TextNode{text: textNode})
		return nil
	}
}

func WithClassName(class string) Option {
	return func(c *Component) error {
		c.ClassName = class
		return nil
	}
}
