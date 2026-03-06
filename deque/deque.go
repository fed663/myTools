package deque

type Deque[T any] struct {
	data []T
}

// добавить в начало
func (d *Deque[T]) PushFront(val T) {
	d.data = append([]T{val}, d.data...)
}

// добавить в конец
func (d *Deque[T]) PushBack(val T) {
	d.data = append(d.data, val)
}

// взять из начала
func (d *Deque[T]) PopFront() (T, bool) {
	var nol T
	if len(d.data) == 0 {
		return nol, false
	}

	val := d.data[0]
	d.data = d.data[1:]
	return val, true
}

// взять из конца
func (d *Deque[T]) PopBack() (T, bool) {
	var nol T
	if len(d.data) == 0 {
		return nol, false
	}

	val := d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]
	return val, true
}
