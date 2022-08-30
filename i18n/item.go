package i18n

type Item struct {
	items map[string]string
}

func (i Item) GetItem(item string) (string, bool) {
	v, ok := i.items[item]
	return v, ok
}

func NewItem() *Item {
	return &Item{items: map[string]string{}}
}
