package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
)

func PredecesorInOrder(bt *binarytree.BinarySearchTree[int], k int) (int, error) {
	// Implementar

	nodo := bt.GetRoot()
	var predecesor *binarytree.BinaryNode[int]

	for nodo != nil {
		if k == nodo.GetData() {
			// Caso: clave encontrada y tiene hijo izquierdo
			if nodo.GetLeft() != nil {
				max := maximo(nodo.GetLeft())
				return max.GetData(), nil
			}
			break // clave encontrada pero sin hijo izquierdo
		} else if k > nodo.GetData() {
			// Posible predecesor, seguimos a la derecha
			predecesor = nodo
			nodo = nodo.GetRight()
		} else {
			// Menor o igual, vamos a la izquierda
			nodo = nodo.GetLeft()
		}
	}

	if predecesor != nil {
		return predecesor.GetData(), nil
	}

	return 0, errors.New("No hay predecesores menores que el mínimo")
}

func maximo(n *binarytree.BinaryNode[int]) *binarytree.BinaryNode[int] {
	for n.GetRight() != nil {
		n = n.GetRight()
	}
	return n
}
