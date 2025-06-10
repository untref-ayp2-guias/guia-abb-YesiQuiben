package ejercicios

import (
	"testing"

	"untref-ayp2/guia-abb/binarytree"

	"github.com/stretchr/testify/assert"
)

/* Construct the following tree
          15
        /    \
       /      \
      10       20
     / \      /  \
    /   \    /    \
   8    12  16    25
*/

func TestSecLargeElem(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()
	bst.Insert(15)
	bst.Insert(10)
	bst.Insert(20)
	bst.Insert(8)
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(25)

	p, _ := SecondLargestElement(bst)
	assert.Equal(t, 20, p)
}
func TestSecLargeElemIs25(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()
	bst.Insert(15)
	bst.Insert(10)
	bst.Insert(20)
	bst.Insert(8)
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(25)
	bst.Insert(30)
	p, _ := SecondLargestElement(bst)
	assert.Equal(t, 25, p)
}

func TestSecLargeElemTwoNodesIs5(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()
	bst.Insert(10)
	bst.Insert(5)
	p, _ := SecondLargestElement(bst)
	assert.Equal(t, 5, p)
}

func TestSecLargeElemRightDeepIs20(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(20)
	bst.Insert(30)
	p, _ := SecondLargestElement(bst)
	assert.Equal(t, 20, p)
}

func TestSecLargeElemLeftSubtreeOfMaxIs19(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(20)
	bst.Insert(18)
	bst.Insert(19)
	p, _ := SecondLargestElement(bst)
	assert.Equal(t, 19, p)
}

func TestNilSecLargeElem(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()
	_, err := SecondLargestElement(bst)
	assert.EqualError(t, err, "No hay valores")
}
