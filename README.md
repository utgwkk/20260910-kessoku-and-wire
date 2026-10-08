# 20260910-kessoku-and-wire

```
% go generate .
time=2026-09-10T19:29:07.123+09:00 level=INFO msg="Generating dependency injection code" files=[kessoku.go]
time=2026-09-10T19:29:07.163+09:00 level=WARN msg="parseInjectCall failed" callExpr="&{Fun:0x1f4e8ed873e0 Lparen:263062 Args:[0x1f4e8ed87410 0x1f4e8edb4348 0x1f4e8edb4360] Ellipsis:0 Rparen:263118}" error="parse provider argument: Set call expression from another package is not supported: KessokuSet"
wire: github.com/utgwkk/20260910-kessoku-and-wire: wrote /Users/utgwkk/ghq/github.com/utgwkk/20260910-kessoku-and-wire/wire_gen.go
```

## Verification with `feat/external-package-import` (mazrean/kessoku#203)

```
% go get github.com/mazrean/kessoku@507b39a   # HEAD of feat/external-package-import
% go generate .
time=2026-10-08T13:22:13.049Z level=INFO msg="Generating dependency injection code" files=[kessoku.go]
time=2026-10-08T13:22:13.117Z level=INFO msg="Found inject directives" file=kessoku.go count=1
wire: github.com/utgwkk/20260910-kessoku-and-wire: wrote /home/user/20260910-kessoku-and-wire/wire_gen.go
```

`kessoku_band.go` is generated successfully.
