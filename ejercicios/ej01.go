package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
)

func SecondLargestElement(bt *binarytree.BinarySearchTree[int]) (int, error) {
	// Implementar
	if bt == nil {
		return 0, errors.New("El árbol no puede ser nil")
	}
	return findSecLargeElem(bt.GetRoot())
}

func findSecLargeElem(n *binarytree.BinaryNode[int]) (int, error) {
	if n == nil {
		return 0, errors.New("No hay valores")
	}

	right := n.GetRight()
	if right == nil {
		if left := n.GetLeft(); left != nil {
			for left.GetRight() != nil {
				left = left.GetRight()
			}
			return left.GetData(), nil
		}
		return 0, errors.New("No hay más de un elemento")
	}

	if right.GetRight() == nil {
		if rLeft := right.GetLeft(); rLeft != nil {
			for rLeft.GetRight() != nil {
				rLeft = rLeft.GetRight()
			}
			return rLeft.GetData(), nil
		}
		return n.GetData(), nil
	}

	return findSecLargeElem(right)
}
