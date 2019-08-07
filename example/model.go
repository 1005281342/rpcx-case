package example

import "context"

type C int

type A struct {
	A1 int
	A2 int
}

type B struct {
	B int
}

func (t *C) Mul(ctx context.Context, a *A, b *B) error {
	b.B = a.A1 * a.A2
	return nil
}
