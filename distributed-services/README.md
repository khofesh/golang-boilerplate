# Distributed services with Go

## chapter 1 - test the API

run the server

```shell
cd proglog

go run cmd/server/main.go
```

test the endpoint

```shell
curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzEK"}}'

curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzIK"}}'

curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzMK"}}'

curl -X GET localhost:8080 -d '{"offset": 0}'

curl -X GET localhost:8080 -d '{"offset": 1}'

curl -X GET localhost:8080 -d '{"offset": 2}'
```

## chapter 2

https://github.com/protocolbuffers/protobuf-go

get protobuf

```shell
go get google.golang.org/protobuf/...@v1.28.0
```

## chapter 3

get testify

```shell
go get github.com/stretchr/testify
```

gommap

```shell
sudo dnf install bzr

export GOVCS=*:all

go get -u github.com/tysonmote/gommap
```

https://command-not-found.com/bzr.bzr

run test

```shell
make test
```
