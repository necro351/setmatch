Setmatch
========

Setmatch is a calculator that prints out all the valid Sets in the card game
Set.

Setmatch is written in Go. You can build the program with `go build` and try
running the sample input `sets.txt`.

```
go build
cat sets.txt|./setmatch
```

Cards are entered by typing in four-character code words. one per line. Each
character maps to a property of the card. Every card in Set has four
properties, and each property has three valid values:

```
count: 1, 2, 3
color: (r)ed, (p)urple, (g)reen
shade: (p)en, (s)olid, (h)ollow
shape: (w)ave, (p)ill, (d)iamond
```

You can find sample valid code words in sets.txt. For example `2ppw` stands for
"two purple, penned-in, waves" which is one of the cards in Set.
