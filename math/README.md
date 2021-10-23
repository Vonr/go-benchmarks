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
BenchmarkStdMod-8         128118              8387 ns/op
BenchmarkFastMod-8        857655              1294 ns/op
BenchmarkStdFloor-8     929981188                1.299 ns/op
BenchmarkFastFloor-8    924504414                1.288 ns/op
BenchmarkStdMax-8       88428781                11.42 ns/op
BenchmarkFastMax-8      869061790                1.429 ns/op
BenchmarkStdAbs-8       836941008                1.425 ns/op
BenchmarkFastAbs-8      858260697                1.432 ns/op
PASS
ok      fastmath        11.034s
```
