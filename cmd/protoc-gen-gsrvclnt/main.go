// protoc-gen-gsrvclnt is a protoc plugin to generate wrapper of grpc server implementation
// as client.
package main

import (
	"github.com/fardream/gsrvclnt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	options := protogen.Options{}
	options.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range plugin.Files {
			if !f.Generate {
				continue
			}

			g := plugin.NewGeneratedFile(f.GeneratedFilenamePrefix+"_gsrvclnt.pb.go", f.GoImportPath)
			g.P("package ", f.GoPackageName)
			g.P()
			gsrvclnt.GenForFile(g, f)
		}
		return nil
	})
}
