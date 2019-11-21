# unicode-forms
Output command line input into various unicode normalization forms.

# Usage

Simply run the Go program for an example output.

```
$ go run main.go
[NOTE] No arguments given, using example data.
[NOTE] To read stdin pass '--' as the argument
Raw Input: Nord̩n 4² Näs Åmål Västra Mêléé 가구
i: N  o  r  d    ̩  n     4    ²     N  a    ̈  s     A    ̊  m  a    ̊  l     V    ä  s  t  r  a     M    ê  l    é  e    ́         가      구
u:4E 6F 72 64  329 6E 20 34   B2 20 4E 61  308 73 20 41  30A 6D 61  30A 6C 20 56   E4 73 74 72 61 20 4D   EA 6C   E9 65  301 20   AC00   AD6C
x:4E 6F 72 64 CCA9 6E 20 34 C2B2 20 4E 61 CC88 73 20 41 CC8A 6D 61 CC8A 6C 20 56 C3A4 73 74 72 61 20 4D C3AA 6C C3A9 65 CC81 20 EAB080 EAB5AC

NFC: Nord̩n 4² Näs Åmål Västra Mêléé 가구
i: N  o  r  d    ̩  n     4    ²     N    ä  s       Å  m    å  l     V    ä  s  t  r  a     M    ê  l    é    é         가      구
u:4E 6F 72 64  329 6E 20 34   B2 20 4E   E4 73 20   C5 6D   E5 6C 20 56   E4 73 74 72 61 20 4D   EA 6C   E9   E9 20   AC00   AD6C
x:4E 6F 72 64 CCA9 6E 20 34 C2B2 20 4E C3A4 73 20 C385 6D C3A5 6C 20 56 C3A4 73 74 72 61 20 4D C3AA 6C C3A9 C3A9 20 EAB080 EAB5AC

NFKC: Nord̩n 42 Näs Åmål Västra Mêléé 가구
i: N  o  r  d    ̩  n     4  2     N    ä  s       Å  m    å  l     V    ä  s  t  r  a     M    ê  l    é    é         가      구
u:4E 6F 72 64  329 6E 20 34 32 20 4E   E4 73 20   C5 6D   E5 6C 20 56   E4 73 74 72 61 20 4D   EA 6C   E9   E9 20   AC00   AD6C
x:4E 6F 72 64 CCA9 6E 20 34 32 20 4E C3A4 73 20 C385 6D C3A5 6C 20 56 C3A4 73 74 72 61 20 4D C3AA 6C C3A9 C3A9 20 EAB080 EAB5AC

NFD: Nord̩n 4² Näs Åmål Västra Mêléé 가구
i: N  o  r  d    ̩  n     4    ²     N  a    ̈  s     A    ̊  m  a    ̊  l     V  a    ̈  s  t  r  a     M  e    ̂  l  e    ́  e    ́         ᄀ      ᅡ      ᄀ      ᅮ
u:4E 6F 72 64  329 6E 20 34   B2 20 4E 61  308 73 20 41  30A 6D 61  30A 6C 20 56 61  308 73 74 72 61 20 4D 65  302 6C 65  301 65  301 20   1100   1161   1100   116E
x:4E 6F 72 64 CCA9 6E 20 34 C2B2 20 4E 61 CC88 73 20 41 CC8A 6D 61 CC8A 6C 20 56 61 CC88 73 74 72 61 20 4D 65 CC82 6C 65 CC81 65 CC81 20 E18480 E185A1 E18480 E185AE

NFKD: Nord̩n 42 Näs Åmål Västra Mêléé 가구
i: N  o  r  d    ̩  n     4  2     N  a    ̈  s     A    ̊  m  a    ̊  l     V  a    ̈  s  t  r  a     M  e    ̂  l  e    ́  e    ́         ᄀ      ᅡ      ᄀ      ᅮ
u:4E 6F 72 64  329 6E 20 34 32 20 4E 61  308 73 20 41  30A 6D 61  30A 6C 20 56 61  308 73 74 72 61 20 4D 65  302 6C 65  301 65  301 20   1100   1161   1100   116E
x:4E 6F 72 64 CCA9 6E 20 34 32 20 4E 61 CC88 73 20 41 CC8A 6D 61 CC8A 6C 20 56 61 CC88 73 74 72 61 20 4D 65 CC82 6C 65 CC81 65 CC81 20 E18480 E185A1 E18480 E185AE
```

A string may be passed as an argument, or "--" may be passed to read STDIN.

A compiled binary can be created in the usual Go way:
```
$ go build -o unicode_forms main.go
```

An example of viewing a file line-by-line:
```
$ cat main.go | ./unicode_forms -- | less
```

