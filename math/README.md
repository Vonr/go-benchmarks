# Math Benchmark Results

```sh
$ go test -bench=.

Floor Tests 1: true 2: true 3: true
Mod Tests 1: true 2: true 3: true
Max Tests 1: true 2: true 3: true
goos: linux
goarch: amd64
pkg: fastmath
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkStdMod-8          62706             18841 ns/op
BenchmarkFastMod-8        390474              2741 ns/op
BenchmarkStdFloor-8     279545673                4.259 ns/op
BenchmarkFastFloor-8    286258305                4.288 ns/op
BenchmarkStdMax-8       100000000               11.15 ns/op
BenchmarkFastMax-8      962329552                1.236 ns/op
BenchmarkStdAbs-8       953286294                1.267 ns/op
BenchmarkFastAbs-8      953832920                1.255 ns/op
PASS
ok      fastmath        10.870s
```
