package superctx

import "context"

// key - Defines a phantom type key based on type T
type key[T any] struct{}

// WithValue - Attaches a value of type T to the specified context.
// The newly created context will have the values attached.
func WithValue[T any](ctx context.Context, val T) context.Context {
	return context.WithValue(ctx, key[T]{}, val)
}

// Value - Gets a value of type T from the specified context.
// If the value does not exist, a zero value of type T is returned.
func Value[T any](ctx context.Context) (T, bool) {
	val, ok := ctx.Value(key[T]{}).(T)
	return val, ok
}
