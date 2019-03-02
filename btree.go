package btree

const maxItems = 31 // use an odd number

type item struct {
	key   string
	value interface{}
}

// the node must satisfy the constraint that len(children) == len(items) + 1 unless len(children) == 0
type node struct {
	items    [maxItems]item
	children [maxItems + 1]*node
}

// BTree represents the B tree data structure
type BTree struct {
	height int
	length int
	root   *node
}

// Set inserts a key-value pair
func (b *BTree) Set(key string, value interface{}) {
	if b.root == nil {
		b.root = new(node)
		b.root.items[0] = item{key, value}
		b.length = 1
	}
}

func (n *node) set(key string, value interface{}, height int) (previous interface{}, replaced bool) {
	if height == 0 {
		for j := len(n.items); j > 0; j-- {
			n.items[j] = n.items[j-1] //create new item on node with value of previous item
		}
		n.items[0] = item{key, value}
		return nil, false
	}
	return
}
