package btree

type item struct {
	key   string
	value interface{}
}

type node struct {
	items    []item
	children []node
}

// BTree represent the B tree data structure
type BTree struct {
	height int
	length int
	root   *node
}
