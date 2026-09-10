package main

import (
	"github.com/mazrean/kessoku"
	"github.com/utgwkk/20260910-kessoku-and-wire/pkg_a"
	"github.com/utgwkk/20260910-kessoku-and-wire/pkg_b"
)

//go:generate go tool kessoku $GOFILE

var _ = kessoku.Inject[*pkg_a.A](
	"NewAKessoku",
	pkg_a.KessokuSet,
	pkg_b.KessokuSet,
)
