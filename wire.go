//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/utgwkk/20260910-kessoku-and-wire/pkg_a"
	"github.com/utgwkk/20260910-kessoku-and-wire/pkg_b"
)

//go:generate go tool wire $GOFILE

func NewAWire() *pkg_a.A {
	wire.Build(
		pkg_b.WireSet,
		pkg_a.WireSet,
	)
	return nil
}
