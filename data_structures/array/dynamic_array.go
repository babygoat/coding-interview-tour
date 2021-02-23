package array

// Practice the implementing an automatically resizing array
// Since golang comes with builtin slice, this practice tries not to utilize the builtin functionality
// Instead, try to maintain the array status (capacity, length) myself.
// Ref: https://github.com/jwasham/coding-interview-university#arrays

import "errors"

var ErrArrIndexOutOfBound = errors.New("array index out of bound")

// DynamicArray defines the status(capacity, length) to be tracked of and data reference of the dynamic array structure.
type DynamicArray struct {
	cap  int
	len  int
	data []int
}

// New creates a DynamicArray object
func New(capacity int) *DynamicArray {
	return &DynamicArray{
		cap:  capacity,
		data: make([]int, capacity),
	}
}

// Push appends the value from the end and grows the capacity to its double if current array is full
// Time Complexity: amortized O(1)
func (a *DynamicArray) Push(v int) {
	if a.len == a.cap {
		a.resize(a.cap * 2)
	}
	a.data[a.len] = v
	a.len++
}

// InvalidValue defines the invalid value if accessing the value of array occurs error
// For simplicity, the values to be stored in this practice are assumed non-negative.
const InvalidValue = -1

// IndexAt returns the value of the array index reference
// Time Complexity: O(1)
func (a *DynamicArray) IndexAt(index int) (int, error) {
	if index < 0 || index >= a.len {
		return InvalidValue, ErrArrIndexOutOfBound
	}
	return a.data[index], nil
}

// resize grows or shrinks the DynamicArray as required
// Time complexity: O(n)
func (a *DynamicArray) resize(grow int) {
	old := a.data
	a.data = make([]int, grow)
	a.cap = grow
	for ind, val := range old {
		a.data[ind] = val
	}
}

// Insert insert the value at the specific index and shift the elements after the index if required
// Time complexity: O(n)
// Special case: O(1), if insert at the end
func (a *DynamicArray) Insert(index int, value int) error {
	if index < 0 || index > a.len {
		return ErrArrIndexOutOfBound
	}
	if a.len == a.cap {
		a.resize(a.cap * 2)
	}

	// if insert value in the middle of array, shift the elements after index
	if index < a.len {
		copy(a.data[index+1:], a.data[index:])
	}
	a.data[index] = value
	a.len++

	return nil
}

// EraseAt removes the data at the specific index and shift the elements to the left after the index if required
// Time complexity: O(n)
// Special case: O(1), if delete at the end
func (a *DynamicArray) EraseAt(index int) error {
	if index < 0 || index >= a.len {
		return ErrArrIndexOutOfBound
	}

	copy(a.data[index:], a.data[index+1:])
	a.len--

	// if the length of the array is 1/4 of current capacity, resize the array to half of the space
	if a.len < a.cap/4 {
		a.resize(a.cap / 2)
	}
	return nil
}

// Remove removes the value in the array and the array remains unchanged if the value is not in the array
// Time complexity: O(n^2) worst case
//                  O(n) if value is not in the array
func (a *DynamicArray) Remove(val int) {
	for index := 0; index < a.len; {
		got, _ := a.IndexAt(index)
		// If the element is found, erase the element at the index
		if got == val {
			a.EraseAt(index)
			continue
		}
		// If the element does not found in current index, shift to the right
		index++
	}
}

// Size returns the number of elements in the array
func (a *DynamicArray) Size() int {
	return a.len
}

//Capacity returns the allocated space of the array
func (a *DynamicArray) Capacity() int {
	return a.cap
}
