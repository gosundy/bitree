## BitMap
Use Examples:
#### Set
add 7 to bitmap
```golang
    bm := NewBitMap(8)
    _=bm.Set(7)
```
#### SetN
```golang
    bm := NewBitMap(8)
    isSedBefore,_:=bm.SetN(7)
    fmt.Println(isSedBefore)
```
#### ResetN
```golang
    bm := NewBitMap(8)
    isNotSedBefore,_:=bm.ResetN(7)
    fmt.Println(isNotSedBefore)
```
#### get
get 7 position result, if data is one, return true, else return false
```golang
    bm := NewBitMap(8)
    isSed,_:=bm.Get(7)
    fmt.Println(isSed)
```

Benchmark Result:
```shell script
goos: darwin
goarch: amd64
pkg: github.com/gosundy/bitmap
BenchmarkSetOneMillion
BenchmarkSetOneMillion-8         	     406	   2917147 ns/op
BenchmarkGetOneMillion
BenchmarkGetOneMillion-8         	     417	   2815003 ns/op
BenchmarkResetSetNOneMillion
BenchmarkResetSetNOneMillion-8   	     418	   2820011 ns/op
```