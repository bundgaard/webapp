package i18n

type Category struct {
	categories map[string]Item
}

func (c *Category) GetCategory(category string) (Item, bool) {
	v, ok := c.categories[category]
	return v, ok
}

func NewCategory() *Category {
	return &Category{categories: map[string]Item{}}
}
