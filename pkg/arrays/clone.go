package arrays

/*
# Clone

the Clone method returns a shallow copy of the array elements
*/
func (a Array[T]) Clone() Array[T] {
	return Array[T]{
		elements: a.elements,
		lastPos:  a.lastPos,
	}
}
