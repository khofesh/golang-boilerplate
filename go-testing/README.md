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

golangci-lint

https://golangci-lint.run/usage/install/

```shell
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2

golangci-lint --version
```

lint

```shell
golangci-lint run
```

lint result

```shell
[fahmad@ryzen go-testing]$ golangci-lint run
linting/lint.go:5:6: `checkFlag` is unused (deadcode)
func checkFlag(flag bool) bool {
     ^
linting/lint.go:13:6: `errChecking` is unused (deadcode)
func errChecking() (int, error) {
     ^
linting/lint.go:6:5: S1002: should omit comparison to bool constant, can be simplified to `flag` (gosimple)
	if flag == true {
	   ^
linting/lint.go:14:2: ineffectual assignment to a (ineffassign)
	a := 1
	^
```

formatting

```shell
gofmt -s -w .
```

testMain

```shell
go test yamltohtml/yamltohtml_test.go -v
```

testify mock

```shell
go test service/service_test.go -v
```
