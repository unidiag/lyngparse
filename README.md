# lyngparse

`go get github.com/unidiag/lyngparse`

Usage:
`spew.Dump(lyngparse.LyngSat("/Astra-4A.html"))`

Output:
```
([]string) (len=16 cap=16) {
 (string) (len=19) "4.8:S:11727:H:27500",
 (string) (len=20) "4.8:S2:11747:V:30000",
 (string) (len=20) "4.8:S2:11766:H:30000",
 (string) (len=20) "4.8:S2:11785:V:27500",
 (string) (len=20) "4.8:S2:11938:V:27500",
 (string) (len=20) "4.8:S2:11958:H:27500",
 (string) (len=19) "4.8:S:11996:H:27500",
 (string) (len=20) "4.8:S2:12073:H:30000",
 (string) (len=19) "4.8:S:12130:V:27500",
 (string) (len=20) "4.8:S2:12207:V:30000",
 (string) (len=20) "4.8:S2:12226:H:27500",
 (string) (len=19) "4.8:S:12284:V:27500",
 (string) (len=19) "4.8:S:12380:H:27500",
 (string) (len=19) "4.8:S:12418:H:27500",
 (string) (len=20) "4.8:S2:12437:V:30000",
 (string) (len=20) "4.8:S2:12565:V:30000"
}
```
