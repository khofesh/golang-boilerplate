# Production ready REST API

## taskfile

https://taskfile.dev/#/installation

```shell
task build

task test

task lint

task run
```

## golangci-lint

https://golangci-lint.run/

```shell
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2

golantci-lint --version
```
