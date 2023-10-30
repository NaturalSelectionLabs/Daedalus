package types

type Response[T any] struct {
	msg  string
	data T
}
