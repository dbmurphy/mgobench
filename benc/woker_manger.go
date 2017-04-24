package main

// BufferPool is used to recive task
type BufferPool struct {
	T Task
}

func NewBufferPool(bufferSize int) (pool chan BufferPool) {
	pool = make(chan BufferPool, bufferSize)
	return

}
