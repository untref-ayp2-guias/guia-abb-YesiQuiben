package ejercicios

import (
	"untref-ayp2/guia-abb/binarytree"
)

func IsBst(bt *binarytree.BinaryTree[int]) bool {
	// Implementar
	return isBstByNode(bt.GetRoot(), nil, nil)
}

func isBstByNode(n *binarytree.BinaryNode[int], min *int, max *int) bool {
	// Implementar
	if n == nil {
		return true
	}

	val := n.GetData()
	// Verificamos si el valor actual está fuera de los límites
	if min != nil && val <= *min {
		return false
	}
	if max != nil && val >= *max {
		return false
	}

	// Validamos recursivamente los hijos con límites actualizados
	return isBstByNode(n.GetLeft(), min, &val) &&
		isBstByNode(n.GetRight(), &val, max)
}
