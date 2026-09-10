package pkg_a

import (
	"github.com/google/wire"
	"github.com/mazrean/kessoku"
	"github.com/utgwkk/20260910-kessoku-and-wire/pkg_b"
)

type A struct{}

func NewA(
	B *pkg_b.B,
) *A {
	return &A{}
}

var WireSet = wire.NewSet(NewA)

var KessokuSet = kessoku.Set(kessoku.Provide(NewA))
