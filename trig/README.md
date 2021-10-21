# Trig Benchmark Results

```sh
$ go test -bench=.

FastSin/Cos Max Error: 0.0009201468244426358
FastAtan2 Max Error: 3.4323974483774222e-06
FastAtan 3.4323973303607147e-06
goos: linux
goarch: amd64
pkg: trig
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkStdSine-8     	98542628	        12.31 ns/op
BenchmarkFastSine-8    	362214169	         3.483 ns/op
BenchmarkStdCos-8      	86710432	        13.47 ns/op
BenchmarkFastCos-8     	319870407	         3.785 ns/op
BenchmarkStdAtan2-8    	72719572	        16.06 ns/op
BenchmarkFastAtan2-8   	89731520	        13.07 ns/op
BenchmarkStdAtan-8     	120159442	         9.685 ns/op
BenchmarkFastAtan-8    	88352276	        12.84 ns/op
PASS
ok  	trig	16.331s
```
