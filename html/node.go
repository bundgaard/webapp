package html

type Node interface {
	elementNode()
	getID() string
}
