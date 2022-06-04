https://tutorialedge.net/courses/go-testing-bible/

run test

```shell
go test ./... -v
```

coverage

```shell
go test ./... --cover

go test ./... --coverprofile=coverage.out

go tool cover -html=coverage.out
```

benchmark

```shell
go test ./... -run=Benchmark -bench=.
```

benchmark result

```shell
goos: linux
goarch: amd64
pkg: go-test/calculator
cpu: AMD Ryzen 5 1600 Six-Core Processor
BenchmarkCalculateIsArmstrong-12    	15675831	        75.24 ns/op
PASS
ok  	go-test/calculator	1.262s
PASS
ok  	go-test/yamltohtml	0.004s
```
