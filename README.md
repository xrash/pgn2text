# pgn2text

Transform PGN into nice narration text.

# Installation

`go install github.com/xrash/pgn2text/cmd/pgn2text`

# Usage

```
$ cat ./game.pgn
```

```
[Event "Live Chess"]
[Site "Chess.com"]
[Date "2017.11.12"]
[White "wardaddy20"]
[Black "b-d-l"]
[Result "0-1"]
[ECO "D00"]
[WhiteElo "2100"]
[BlackElo "2139"]
[TimeControl "180+2"]
[EndTime "12:26:00 PST"]
[Termination "b-d-l won by checkmate"]
[CurrentPosition "r3k3/pp3pp1/2p1pn2/3p2pr/2PP2K1/4PP2/PP1NBbP1/R2Q1R2 w q - 2 21"]

1.d4 {[%clk 0:03:01]} d5 {[%clk 0:02:59]} 2.Bf4 {[%clk 0:03:02]} Bf5 {[%clk 0:02:58]} 3.e3 {[%clk 0:03:01]} e6 {[%clk 0:02:59]} 4.Nf3 {[%clk 0:02:58]} Be7 {[%clk 0:02:58]} 5.Be2 {[%clk 0:02:54]} Nf6 {[%clk 0:02:56]} 6.O-O {[%clk 0:02:50]} Nbd7 {[%clk 0:02:55]} 7.Nh4 {[%clk 0:02:49]} Be4 {[%clk 0:02:55]} 8.Nd2 {[%clk 0:02:47]} c6 {[%clk 0:02:50]} 9.f3 {[%clk 0:02:46]} Bg6 {[%clk 0:02:49]} 10.Nxg6 {[%clk 0:02:43]} hxg6 {[%clk 0:02:49]} 11.Bg3 {[%clk 0:02:41]} Nh5 {[%clk 0:02:50]} 12.Bf2 {[%clk 0:02:39]} Qc7 {[%clk 0:02:50]} 13.c4 {[%clk 0:02:36]} Qxh2+ {[%clk 0:01:43]} 14.Kxh2 {[%clk 0:02:31]} Ng3+ {[%clk 0:01:44]} 15.Kxg3 {[%clk 0:02:32]} Bh4+ {[%clk 0:01:45]} 16.Kh3 {[%clk 0:02:16]} Bxf2+ {[%clk 0:01:41]} 17.Kg4 {[%clk 0:02:17]} Rh4+ {[%clk 0:01:24]} 18.Kg5 {[%clk 0:02:15]} Rh5+ {[%clk 0:01:22]} 19.Kf4 {[%clk 0:02:14]} g5+ {[%clk 0:01:02]} 20.Kg4 {[%clk 0:02:13]} Nf6# {[%clk 0:01:02]}  0-1
```

```
$ cat ./game.pgn | pgn2text
```

```
d4, d5, bishop to f4, bishop to f5, e3, e6, knight to f3, bishop to e7, bishop to e2, knight to f6, castles kingside, knight to d7, knight to h4, bishop to e4, knight to d2, c6, f3, bishop to g6, knight takes on g6, pawn takes on g6, bishop to g3, knight to h5, bishop to f2, queen to c7, c4, queen takes on h2 check, king takes on h2, knight to g3 check, king takes on g3, bishop to h4 check, king to h3, bishop takes on f2 check, king to g4, rook to h4 check, king to g5, rook to h5 check, king to f4, g5 check, king to g4, knight to f6 checkmate, black wins
```
