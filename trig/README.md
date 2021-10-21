# Trig Benchmark Results

```sh
$ go test -bench=.

FastSin Max Error: 0.0009201468244426358
FastCos Max Error: 0.0009201468244426358
goos: linux
goarch: amd64
pkg: trig
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkStdSine-8      1000000000               0.1332 ns/op
BenchmarkFastSine-8     1000000000               0.04140 ns/op
BenchmarkStdCos-8       1000000000               0.1490 ns/op
BenchmarkFastCos-8      1000000000               0.04632 ns/op
PASS
ok      trig    6.921s
```
