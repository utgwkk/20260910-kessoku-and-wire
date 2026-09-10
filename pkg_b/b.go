package pkg_b

import (
	"github.com/google/wire"
	"github.com/mazrean/kessoku"
)

type B struct{}

func NewB() *B {
	return &B{}
}

var WireSet = wire.NewSet(NewB)

var KessokuSet = kessoku.Set(kessoku.Provide(NewB))
