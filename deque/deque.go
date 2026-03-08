package deque

import "fmt"

const defaultCapacity = 8

type Deque[T any] struct {
	data []T
	head int
	tail int
	size int
}

// конструктор 2 в 1
func NewDeque[T any](capacity ...int) *Deque[T] {
	if len(capacity) > 1 {
		panic("NewDeque takes up to one argument.")
	}

	newCap := defaultCapacity
	if len(capacity) > 0 && capacity[0] > 0 {
		newCap = capacity[0]
	}
	return &Deque[T]{data: make([]T, newCap)}
}

// блок инициализации внутреннего среза
func (d *Deque[T]) initBlock() {
	if len(d.data) == 0 {
		d.data = make([]T, defaultCapacity)
	}
}

// проверка пустой структуры
func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

// функция увелечения вложенного среза
func (d *Deque[T]) grow() {
	if d.size < len(d.data) {
		return
	}
	newData := make([]T, len(d.data)*2)
	for i := 0; i < d.size; i++ {
		newData[i] = d.data[(d.head+i)%len(d.data)]
	}

	d.data = newData
	d.head = 0
	d.tail = d.size
}

func (d *Deque[T]) Len() int {
	return d.size
}

func (d *Deque[T]) Values() []T {
	res := make([]T, d.size)
	for i := 0; i < d.size; i++ {
		res[i] = d.data[(d.head+i)%len(d.data)]
	}
	return res
}

func (d *Deque[T]) String() string {
	return fmt.Sprint(d.Values())
}

// добавить в начало
func (d *Deque[T]) PushFront(val T) {
	d.initBlock()
	d.grow()

	d.head = (d.head - 1 + len(d.data)) % len(d.data)
	d.data[d.head] = val
	d.size++
}

// добавить в конец
func (d *Deque[T]) PushBack(val T) {
	d.initBlock()
	d.grow()

	d.data[d.tail] = val
	d.tail = (d.tail + 1) % len(d.data)
	d.size++
}

// взять из начала
func (d *Deque[T]) PopFront() (T, bool) {
	var nul T
	if d.size == 0 {
		return nul, false
	}

	val := d.data[d.head]
	d.data[d.head] = nul
	d.head = (d.head + 1) % len(d.data)
	d.size--

	return val, true
}

// взять из конца
func (d *Deque[T]) PopBack() (T, bool) {
	var nul T
	if d.size == 0 {
		return nul, false
	}

	d.tail = (d.tail - 1 + len(d.data)) % len(d.data)
	val := d.data[d.tail]
	d.data[d.tail] = nul
	d.size--

	return val, true
}
