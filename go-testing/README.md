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
