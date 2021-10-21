# Trig Benchmark Results

```sh
$ go test -bench=.

FastSin/Cos Max Error: 0.2722665829677713
FastAtan2 Max Error: 3.4323974483774222e-06
FastAtan 3.4323973303607147e-06
goos: linux
goarch: amd64
pkg: trig
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkStdSine-8      95230574                12.36 ns/op
BenchmarkFastSine-8     480795496                2.537 ns/op
BenchmarkStdCos-8       88366902                13.82 ns/op
BenchmarkFastCos-8      444905314                2.610 ns/op
BenchmarkStdAtan2-8     74973613                15.94 ns/op
BenchmarkFastAtan2-8    74182797                13.48 ns/op
BenchmarkStdAtan-8      120804951               10.10 ns/op
BenchmarkFastAtan-8     94576054                12.54 ns/op
PASS
ok      trig    15.972s
```
