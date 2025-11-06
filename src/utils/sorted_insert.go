package utils

import "errors"

// SortedInsert inserts an item into a sorted slice while maintaining sort order.
// If the slice is at maxSize and the new item would be inserted, the last item is removed and returned.
// Comparator should return: negative if a < b, 0 if a == b, positive if a > b.
// Returns the removed item if the slice was full, or nil if no item was removed.
func SortedInsert[T any](items *[]T, add T, maxSize int, comparator func(a, b T) int) (*T, error) {
	if maxSize <= 0 {
		return nil, errors.New("MAX_SIZE_ZERO")
	}
	// this is an invariant because the interface cannot return multiple removed items if items.length exceeds maxSize
	if len(*items) > maxSize {
		return nil, errors.New("ITEMS_SIZE")
	}

	// short circuit first item add
	if len(*items) == 0 {
		*items = append(*items, add)
		return nil, nil
	}

	isFull := len(*items) == maxSize
	// short circuit if full and the additional item does not come before the last item
	if isFull && comparator((*items)[len(*items)-1], add) <= 0 {
		result := add
		return &result, nil
	}

	lo := 0
	hi := len(*items)

	for lo < hi {
		mid := (lo + hi) >> 1
		if comparator((*items)[mid], add) <= 0 {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	// Insert at position lo
	*items = append(*items, add)
	copy((*items)[lo+1:], (*items)[lo:len(*items)-1])
	(*items)[lo] = add

	if isFull {
		// Remove the last element and return it
		removed := (*items)[len(*items)-1]
		*items = (*items)[:len(*items)-1]
		return &removed, nil
	}

	return nil, nil
}
