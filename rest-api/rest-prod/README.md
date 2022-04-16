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
