# Trig Benchmark Results

```sh
$ go test -bench=.
```

```
goos: linux
goarch: amd64
pkg: trig
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkStdSine-8      1000000000               0.1322 ns/op
BenchmarkFastSine-8     1000000000               0.03907 ns/op
BenchmarkStdCos-8       1000000000               0.1465 ns/op
BenchmarkFastCos-8      1000000000               0.04561 ns/op
PASS
ok      trig    3.724s
```

```sh
$ go test
```

```
FastSin Max Error: 0.0009201468244426358
FastCos Max Error: 0.0009201468244426358
PASS
ok      trig    2.607s
```
