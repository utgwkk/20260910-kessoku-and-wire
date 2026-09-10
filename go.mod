module github.com/utgwkk/20260910-kessoku-and-wire

go 1.26

require (
	github.com/google/wire v0.7.0
	github.com/mazrean/kessoku v1.1.1
)

require (
	github.com/alecthomas/kong v1.15.0 // indirect
	github.com/google/subcommands v1.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.36.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/tools v0.45.0 // indirect
)

tool (
	github.com/google/wire/cmd/wire
	github.com/mazrean/kessoku/cmd/kessoku
)
