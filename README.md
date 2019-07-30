# grpc-client-pool

grpc-client-pool's scenes are mainly about grpc-sidecar and grpc-proxy !

## desc

Each grpc request contains PATH in http2 header, the path is fullMethod.

**what is fullMethod ?**

```
/api.v1.election.CandidateSvc/Register
```

**what is serviceName ?**

```
/api.v1.election.CandidateSvc
```

**what is methodName ?**

```
/Register
```

## process

1. init grpc client pool
2. get grpc-client by grpc fullMethod/serviceName
3. use gprc invoke request/response.
