package vdom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffElementTrees_empty(t *testing.T) {
	// Given ... empty old and new trees
	old := Element{}
	new := Element{}

	// When ...
	result := diffElementTrees(&old, &new)

	// Then ... result patch is empty
	assert.Equal(t, 0, len(result.Patch))
}

func TestDiffElementTrees_attributes(t *testing.T) {
	// Given ... new attribute
	old := Element{Path: []int{}}
	new := Element{Path: []int{}, Attrs: map[string]interface{}{"width": "100%"}}

	// When ...
	result := diffElementTrees(&old, &new)

	// Then ... result patch has single patch of new tree
	assert.Equal(t, 1, len(result.Patch))
	assert.Equal(t, Patch{Type: AttrSet, Path: []int{}, Attr: Attr{Name: "width", Value: "100%"}}, result.Patch[0])

	// Given ... changed attribute
	old = Element{Path: []int{}, Attrs: map[string]interface{}{"width": "80%"}}
	new = Element{Path: []int{}, Attrs: map[string]interface{}{"width": "100%"}}

	// When ...
	result = diffElementTrees(&old, &new)

	// Then ... result patch has single patch of new tree
	assert.Equal(t, 1, len(result.Patch))
	assert.Equal(t, Patch{Type: AttrSet, Path: []int{}, Attr: Attr{Name: "width", Value: "100%"}}, result.Patch[0])

	// Given ... multiple changed attributes
	old = Element{Path: []int{}, Attrs: map[string]interface{}{"width": "80%"}}
	new = Element{Path: []int{}, Attrs: map[string]interface{}{"width": "100%", "height": "50%"}}

	// When ...
	result = diffElementTrees(&old, &new)

	// Then ... result patch has single patch of new tree
	assert.Equal(t, 2, len(result.Patch))
	assert.Equal(t, Patch{Type: AttrSet, Path: []int{}, Attr: Attr{Name: "width", Value: "100%"}}, result.Patch[0])
	assert.Equal(t, Patch{Type: AttrSet, Path: []int{}, Attr: Attr{Name: "height", Value: "50%"}}, result.Patch[1])

	// Given ... removed attribute
	old = Element{Path: []int{}, Attrs: map[string]interface{}{"width": "100%"}}
	new = Element{Path: []int{}}

	// When ...
	result = diffElementTrees(&old, &new)

	// Then ... result patch has single patch of new tree
	assert.Equal(t, 1, len(result.Patch))
	assert.Equal(t, Patch{Type: AttrRemove, Path: []int{}, Attr: Attr{Name: "width"}}, result.Patch[0])
}

func TestDiffElementTrees_elements(t *testing.T) {
	// Given ... identical children
	old := Element{Path: []int{}, Children: []Element{Element{Path: []int{0}, Name: "div"}}}
	new := Element{Path: []int{}, Children: []Element{Element{Path: []int{0}, Name: "div"}}}

	// When ...
	result := diffElementTrees(&old, &new)

	// Then ... result patch is empty
	assert.Equal(t, 0, len(result.Patch))

	// Given ... new children
	old = Element{Path: []int{}}
	new = Element{Path: []int{}, Children: []Element{Element{Path: []int{0}, Name: "div"}}}

	// When ...
	result = diffElementTrees(&old, &new)

	// Then ... full DOM patch
	assert.Equal(t, 1, len(result.Patch))

	// Given ... removed children
	old = Element{Path: []int{}, Children: []Element{Element{Path: []int{0}, Name: "div"}}}
	new = Element{Path: []int{}}

	// When ...
	result = diffElementTrees(&old, &new)

	// Then ... full DOM patch
	assert.Equal(t, 1, len(result.Patch))

	// Todo - implemente changed children
	// // Given ... changed children
	// old = Element{Path: []int{}, Children: []Element{Element{Path: []int{0}, Name: "div"}}}
	// new = Element{Path: []int{}, Children: []Element{Element{Path: []int{0}, Name: "a"}}}
	//
	// // When ...
	// result = diffElementTrees(&old, &new)
	//
	// // Then ... full DOM patch
	// assert.Equal(t, 1, len(result.Patch))
}

func TestDiffElementTrees_tree(t *testing.T) {
	// Given ... changed attribute in div
	old := Element{Path: []int{}, Children: []Element{Element{Path: []int{0}, Name: "div", Attrs: map[string]interface{}{"width": "80%"}}}}
	new := Element{Path: []int{}, Children: []Element{Element{Path: []int{0}, Name: "div", Attrs: map[string]interface{}{"width": "100%"}}}}

	// When ...
	result := diffElementTrees(&old, &new)

	// Then ...
	assert.Equal(t, 1, len(result.Patch))
	assert.Equal(t, Patch{Type: AttrSet, Path: []int{0}, Attr: Attr{Name: "width", Value: "100%"}}, result.Patch[0])
}
