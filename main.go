package main

type heap struct {
	xs []int
}

func (h *heap) Len() int {
	// a heap is 1 indexed
	return len(h.xs) - 1
}

// insert to end
// then bubble up to find min
func (h *heap) Insert(x int) {
	// Note that parenthesis are required to dereference the heap pointer in order to save the return value of append()
	(*h).xs = append(h.xs, x)
	h.bubbleUp(h.Len())
}

func (h *heap) bubbleUp(k int) {
	p, ok := parent(k)
	if !ok {
		return // case k is root
	}
	if h.xs[p] > h.xs[k] {
		h.xs[k], h.xs[p] = h.xs[p], h.xs[k]
		h.bubbleUp(p)
	}
}

func parent(k int) (int, bool) {
	if k == 1 {
		return 0, false
	}
	return k / 2, true
}

// get index of left child of node at index k
func left(k int) int {
	return 2 * k
}

// get index of right child of node at index k
func right(k int) int {
	return 2*k + 1
}

// pop min off of top
// top is 1 indexed
// replace with last element
// bubble down to maintain order
func (h *heap) PopMin() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}

	min := h.xs[1]           // 1 indexed
	h.xs[1] = h.xs[h.Len()]  // min is now the last element
	(*h).xs = h.xs[:h.Len()] // save
	h.bubbleDown(1)
	return min, true

}

func (h *heap) bubbleDown(index int) {
	min := index
	leftOfIndex := left(index)

	// find index of minimum value, between k, and k's left and k's right child
	// 0 - left of k
	// 1 - right of k (2k+1)
	// if min is not reset, k is in fact min
	for i := 0; i < 2; i++ {
		if (leftOfIndex + i) <= h.Len() {
			if h.xs[min] > h.xs[leftOfIndex+i] {
				min = leftOfIndex + i
			}
		}
	}

	// swap if min has changed
	if min != index {
		h.xs[index], h.xs[min] = h.xs[min], h.xs[index]
		h.bubbleDown(min)
	}
}
