# Math Benchmark Results

```sh
$ go test -bench=.

Floor Tests 1: true 2: true 3: true
goos: linux
goarch: amd64
pkg: fastmath
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkStdMod-8         132207              9001 ns/op
BenchmarkFastMod-8        758137              1331 ns/op
BenchmarkStdFloor-8     886775138                1.336 ns/op
BenchmarkFastFloor-8    896988750                1.331 ns/op
PASS
ok      fastmath        4.968s
```
