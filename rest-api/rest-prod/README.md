## CRUD

post comment

```shell
curl --location --request POST 'localhost:8080/api/v1/comment' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "slug": "hello",
        "body": "body",
        "author": "me"
    }'
```

get comment

```shell
curl --location --request GET 'localhost:8080/api/v1/comment/96c43014-9208-4901-926e-6a7f82843597'
```

update comment

```shell
curl --location --request PUT 'localhost:8080/api/v1/comment/96c43014-9208-4901-926e-6a7f82843597' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "slug": "testing-put",
        "body": "body",
        "author": "me"
    }'
```

delete comment

```shell
curl --location --request DELETE 'localhost:8080/api/v1/comment/96c43014-9208-4901-926e-6a7f82843597'
```

auth header

```shell
curl --location --request POST 'localhost:8080/api/v1/comment' -v \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.7fKsho4l39TKxWFjDzo-zAdNI5wFMn2-wqBCDfEuYA4' \
    --data-raw '{
        "slug": "hello",
        "body": "body",
        "author": "me"
    }'

```

integration testing

```shell
go test -tags=integration ./... -v

```
