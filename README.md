# gsrvclnt

`gsrvclnt` wraps golang gRPC's server implementation as client. For example, for a service named `Svc`, a function `NewSvcSrvClient(server SvcServer) SvcClient` will be generated, which wraps a `SvcServer` and returns a `SvcClient`.

A protoc plugin is provided at [cmd/protoc-gen-gsrvclnt](./cmd/protoc-gen-gsrvclnt). The plugin must be used together with `protoc-gen-go-grpc` and `protoc-gen-go`.

```shell
 protoc --proto_path=. --go_out=. --go_opt=paths=source_relative \
	    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	    --gsrvclnt_out=. --gsrvclnt_opt=paths=source_relative \
        /path/to/proto/files
```

The plugin can be installed by

```shell
go install github.com/fardream/gsrvclnt/cmd/protoc-gen-gsrvclnt@latest
```
