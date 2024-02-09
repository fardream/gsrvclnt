// gsrvclnt generates go code to use gRPC's server implementation as a client.
//
// For service named Svc, a function NewSvcSrvClient(server SvcServer) will be generated,
// which will wrap a SvcServer and returns a SvcClient.
//
// A protoc plugin is at cmd/protoc-gen-gsrvclnt, which can be used to generate
// the necessary go code files. this plugin needs to be used together with protoc-gen-go-grpc and
// protoc-gen-go, and takes the same parameters - such as
//
//	protoc --proto_path=. --go_out=. --go_opt=paths=source_relative \
//	    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
//	    --gsrvclnt_out=. --gsrvclnt_opt=paths=source_relative \
//	    /path/to/proto/files
//
// The plugin can be installed by
//
//	go install github.com/fardream/gsrvclnt/cmd/protoc-gen-gsrvclnt@latest
package gsrvclnt
