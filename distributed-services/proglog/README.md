## TLS

```shell
mkdir ~/.proglog
make gencert
make test
```

## observability

```shell
cd internal/server
go test -v -debug=true
```

result

```shell
[fahmad@ryzen server]$ go test -v -debug=true
=== RUN   TestServer
=== RUN   TestServer/produce/consume_a_message_to/from_the_log_succeeeds
    server_test.go:60: metrics log file: /tmp/metrics-779215398.log
    server_test.go:60: traces log file: /tmp/traces-443977549.log

```

## raft

turns out the following line is still needed, otherwise the test will fail.

```shell
go mod edit -replace github.com/hashicorp/raft-boltdb=github.com/travisjeffery/raft-boltdb@v1.0.0
```

## resolver.go and resolver_test.go

`Endpoint` is deprecated, so in `resolver.go` I replaced with

```go
r.resolverConn, err = grpc.Dial(target.URL.Host, dialOpts...)
if err != nil {
	return nil, err
}
```

and in `resolver_test.go`

```go
_, err = r.Build(
	resolver.Target{
		URL: url.URL{
			Host: l.Addr().String(),
		},
	},
	conn,
	opts,
)
require.NoError(t, err)
```

I'm using the `Host` param
