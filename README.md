# unicode-forms
Output command line input into various unicode normalization forms.

# Usage

Simply run the Go program for an example output.

```
$ go run main.go
[NOTE] No arguments given, using example data.
[NOTE]To read stdin pass '--' as the argument
Raw Input: Nord̩n Näs Åmål Västra Mêléé - 가구
i: N  o  r  d    ̩  n     N  a    ̈  s     A    ̊  m  a    ̊  l     V    ä  s  t  r  a     M    ê  l    é  e    ́     -         가      구
u:4E 6F 72 64 0329 6E 20 4E 61 0308 73 20 41 030A 6D 61 030A 6C 20 56 00E4 73 74 72 61 20 4D 00EA 6C 00E9 65 0301 20 2D 20 00AC00 00AD6C
x:4E 6F 72 64 CCA9 6E 20 4E 61 CC88 73 20 41 CC8A 6D 61 CC8A 6C 20 56 C3A4 73 74 72 61 20 4D C3AA 6C C3A9 65 CC81 20 2D 20 EAB080 EAB5AC

NFC: Nord̩n Näs Åmål Västra Mêléé - 가구
i: N  o  r  d    ̩  n     N    ä  s       Å  m    å  l     V    ä  s  t  r  a     M    ê  l    é    é     -         가      구
u:4E 6F 72 64 0329 6E 20 4E 00E4 73 20 00C5 6D 00E5 6C 20 56 00E4 73 74 72 61 20 4D 00EA 6C 00E9 00E9 20 2D 20 00AC00 00AD6C
x:4E 6F 72 64 CCA9 6E 20 4E C3A4 73 20 C385 6D C3A5 6C 20 56 C3A4 73 74 72 61 20 4D C3AA 6C C3A9 C3A9 20 2D 20 EAB080 EAB5AC

NFKC: Nord̩n Näs Åmål Västra Mêléé - 가구
i: N  o  r  d    ̩  n     N    ä  s       Å  m    å  l     V    ä  s  t  r  a     M    ê  l    é    é     -         가      구
u:4E 6F 72 64 0329 6E 20 4E 00E4 73 20 00C5 6D 00E5 6C 20 56 00E4 73 74 72 61 20 4D 00EA 6C 00E9 00E9 20 2D 20 00AC00 00AD6C
x:4E 6F 72 64 CCA9 6E 20 4E C3A4 73 20 C385 6D C3A5 6C 20 56 C3A4 73 74 72 61 20 4D C3AA 6C C3A9 C3A9 20 2D 20 EAB080 EAB5AC

NFD: Nord̩n Näs Åmål Västra Mêléé - 가구
i: N  o  r  d    ̩  n     N  a    ̈  s     A    ̊  m  a    ̊  l     V  a    ̈  s  t  r  a     M  e    ̂  l  e    ́  e    ́     -         ᄀ      ᅡ      ᄀ      ᅮ
u:4E 6F 72 64 0329 6E 20 4E 61 0308 73 20 41 030A 6D 61 030A 6C 20 56 61 0308 73 74 72 61 20 4D 65 0302 6C 65 0301 65 0301 20 2D 20 001100 001161 001100 00116E
x:4E 6F 72 64 CCA9 6E 20 4E 61 CC88 73 20 41 CC8A 6D 61 CC8A 6C 20 56 61 CC88 73 74 72 61 20 4D 65 CC82 6C 65 CC81 65 CC81 20 2D 20 E18480 E185A1 E18480 E185AE

NFKD: Nord̩n Näs Åmål Västra Mêléé - 가구
i: N  o  r  d    ̩  n     N  a    ̈  s     A    ̊  m  a    ̊  l     V  a    ̈  s  t  r  a     M  e    ̂  l  e    ́  e    ́     -         ᄀ      ᅡ      ᄀ      ᅮ
u:4E 6F 72 64 0329 6E 20 4E 61 0308 73 20 41 030A 6D 61 030A 6C 20 56 61 0308 73 74 72 61 20 4D 65 0302 6C 65 0301 65 0301 20 2D 20 001100 001161 001100 00116E
x:4E 6F 72 64 CCA9 6E 20 4E 61 CC88 73 20 41 CC8A 6D 61 CC8A 6C 20 56 61 CC88 73 74 72 61 20 4D 65 CC82 6C 65 CC81 65 CC81 20 2D 20 E18480 E185A1 E18480 E185AE

```

A string may be passed as an argument, or "--" may be passed to read STDIN.

