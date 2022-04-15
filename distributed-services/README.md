# Distributed services with Go

## chapter 1 - test the API

```shell
curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzEK"}}'

curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzIK"}}'

curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzMK"}}'

curl -X GET localhost:8080 -d '{"offset": 0}'

curl -X GET localhost:8080 -d '{"offset": 1}'

curl -X GET localhost:8080 -d '{"offset": 2}'
```

## chapter 2

get protobuf

```shell
go get google.golang.org/protobuf/...@v1.28.0
```
