package tree

import (
	"fmt"
	"io"

	_ "github.com/jinzhu/gorm"
)

type Tree struct {
	Value    interface{} `json:"name"`
	Children []*Tree     `json:"children"`
}

func NewTree(v interface{}) *Tree {
	return &Tree{
		Children: []*Tree{},
		Value:    v,
	}
}

func (t *Tree) String() string {
	return fmt.Sprint(t.Value)
}

func (t *Tree) PrettyPrint(w io.Writer, prefix string) {
	var inner func(int, *Tree)
	inner = func(depth int, child *Tree) {
		for i := 0; i < depth; i++ {
			io.WriteString(w, prefix)
		}
		io.WriteString(w, child.String()+"\n") // you should really observe the return value here.
		for _, grandchild := range child.Children {
			inner(depth+1, grandchild)
		}
	}
	inner(0, t)
}

func (t *Tree) AddChild(child interface{}) {
	switch c := child.(type) {
	case *Tree:
		t.Children = append(t.Children, c)
	default:
		t.Children = append(t.Children, NewTree(c))
	}
}
