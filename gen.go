package gsrvclnt

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

var (
	contextContext    = protogen.GoIdent{GoName: "Context", GoImportPath: "context"}
	contextCancelFunc = protogen.GoIdent{GoName: "CancelFunc", GoImportPath: "context"}
	contextWithCancel = protogen.GoIdent{GoName: "WithCancel", GoImportPath: "context"}
	grpcCallOption    = protogen.GoIdent{GoName: "CallOption", GoImportPath: "google.golang.org/grpc"}
	grpcMetadataMD    = protogen.GoIdent{GoName: "MD", GoImportPath: "google.golang.org/grpc/metadata"}
	errorsNew         = protogen.GoIdent{GoName: "New", GoImportPath: "errors"}
	ioEOF             = protogen.GoIdent{GoName: "EOF", GoImportPath: "io"}
)

func servername(s *protogen.Service, m *protogen.Method, g *protogen.GeneratedFile) string {
	return fmt.Sprintf("_%s_%sSrvServerStream", s.GoName, m.GoName)
}

func clientname(s *protogen.Service, m *protogen.Method, g *protogen.GeneratedFile) string {
	return fmt.Sprintf("_%s_%sSrvClientStream", s.GoName, m.GoName)
}

func GenForFile(g *protogen.GeneratedFile, f *protogen.File) {
	for _, s := range f.Services {
		// actual implementation
		stypename := fmt.Sprintf("_%s_SrvClient", s.GoName)
		g.P("type ", stypename, " struct {")
		g.P("Server ", s.GoName, "Server")
		g.P("}")
		g.P("var _ ", s.GoName, "Client = (*", stypename, ")(nil)")

		g.P("func New", s.GoName, "SrvClient(server ", s.GoName, "Server) ", s.GoName, "Client {")
		g.P("return &", stypename, "{Server: server}")
		g.P("}")

		for _, m := range s.Methods {
			clientype := clientname(s, m, g)
			servertype := servername(s, m, g)

			switch {
			case m.Desc.IsStreamingClient() && m.Desc.IsStreamingServer():
				genClientStream(s, m, g)
				genServerStream(s, m, g)
				g.P("// do bidirectional streaming")
				g.P("func (client *", stypename, ") ", m.GoName, "(ctx ", contextContext, ", opts ...", grpcCallOption, ") (", s.GoName, "_", m.GoName, "Client, error) {")
				clientype := clientname(s, m, g)
				g.P("r := new_", clientype, "(ctx)")
				g.P("go func() {")
				g.P("err := client.Server.", m.GoName, "(&r.", servername(s, m, g), ")")
				g.P("if err != nil {")
				g.P("r.errfromsrv = err")
				g.P("r.cancel()")
				g.P("} else if r.errfromsrv == nil {")
				g.P("r.errfromsrv = ", ioEOF)
				g.P("}")
				g.P("close(r.toclient)")
				g.P("}()")
				g.P("return r, nil")
				g.P("}")
			case m.Desc.IsStreamingClient():
				genClientStream(s, m, g)
				genServerStream(s, m, g)
				g.P("// do client streaming")
				g.P("func (client *", stypename, ") ", m.GoName, "(ctx ", contextContext, ", opts ...", grpcCallOption, ") (", s.GoName, "_", m.GoName, "Client, error) {")
				g.P("r := new_", clientype, "(ctx)")
				g.P("go func() {")
				g.P("err := client.Server.", m.GoName, "(&r.", servertype, ")")
				g.P("if err != nil && r.errfromsrv == nil {")
				g.P("r.errfromsrv = err")
				g.P("}")
				g.P("if err != nil {")
				g.P("r.cancel()")
				g.P("}")
				g.P("}()")
				g.P("return r, nil")
				g.P("}")
			case m.Desc.IsStreamingServer():
				genClientStream(s, m, g)
				genServerStream(s, m, g)
				g.P("// do server streaming")
				g.P("func (client *", stypename, ") ", m.GoName, "(ctx ", contextContext, ", req *", m.Input.GoIdent, ", opts... ", grpcCallOption, ")(", s.GoName, "_", m.GoName, "Client, error) {")
				g.P("r := new_", clientype, "(ctx)")
				g.P("go func () {")
				g.P("err := client.Server.", m.GoName, "(req, &r.", servertype, ")")
				g.P("if err != nil {")
				g.P("r.errfromsrv = err")
				g.P("} else {")
				g.P("r.errfromsrv = ", ioEOF)
				g.P("}")
				g.P("close(r.toclient)")
				g.P("r.cancel()")
				g.P("} ()")
				g.P("return r, nil")
				g.P("}")
			default:
				g.P("func (client *", stypename, ") ", m.GoName, "(ctx ", contextContext, ", req *", m.Input.GoIdent, ", opts... ", grpcCallOption, ")(* ", m.Output.GoIdent, ",error) {")
				g.P("return client.Server.", m.GoName, "(ctx, req)")
				g.P("}")
			}
		}
	}
}

func genClientStream(s *protogen.Service, m *protogen.Method, g *protogen.GeneratedFile) {
	mtypename := clientname(s, m, g)
	servertype := servername(s, m, g)
	// client type
	g.P("type ", mtypename, " struct {")
	g.P(servertype)
	g.P("}")
	// interface check
	g.P("var _ ", s.GoName, "_", m.GoName, "Client = (*", mtypename, ")(nil)")

	g.P("func new_", mtypename, "(ctx ", contextContext, ") *", mtypename, "{")
	g.P("return &", mtypename, "{", servertype, ": new_", servertype, "(ctx)}")
	g.P("}")
	// Context method
	g.P("func (client *", mtypename, ") Context() ", contextContext, "{")
	g.P("return client.ctx")
	g.P("}")

	// Header
	g.P("func (client *", mtypename, ") Header() (", grpcMetadataMD, ", error){")
	g.P("return client.header, nil")
	g.P("}")

	// Trailer
	g.P("func (client *", mtypename, ") Trailer() ", grpcMetadataMD, "{")
	g.P("return client.trailer")
	g.P("}")

	// CloseSend
	g.P("func (client *", mtypename, ") CloseSend() error {")
	g.P("if client.errfromclient == nil {")
	g.P("client.errfromclient = ", ioEOF)
	g.P("}")
	g.P("close(client.fromclient)")
	g.P("return nil")
	g.P("}")

	// SendMsg
	g.P("func (client *", mtypename, ") SendMsg(m any) error {")
	g.P("return ", errorsNew, "(\"unimplemented\")")
	g.P("}")

	// RecvMsg
	g.P("func (client *", mtypename, ") RecvMsg(m any) error {")
	g.P("return ", errorsNew, "(\"unimplemented\")")
	g.P("}")

	if m.Desc.IsStreamingServer() {
		// Recv
		g.P("func (client *", mtypename, ") Recv()(*", m.Output.GoIdent, ", error) {")
		g.P("r, ok := <- client.toclient")
		g.P("if ok {")
		g.P("return r, nil")
		g.P("} else {")
		g.P("return nil, client.errfromsrv")
		g.P("}")
		g.P("}")
	} else {
		// CloseAndRecv
		g.P("func (client *", mtypename, ") CloseAndRecv()(*", m.Output.GoIdent, ", error) {")
		g.P("client.errfromclient = ", ioEOF)
		g.P("if err := client.CloseSend(); err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("r, ok := <- client.toclient")
		g.P("if ok {")
		g.P("return r, nil")
		g.P("} else {")
		g.P("return nil, client.errfromsrv")
		g.P("}")
		g.P("}")
	}

	if m.Desc.IsStreamingClient() {
		g.P("func (client *", mtypename, ") Send(m * ", m.Input.GoIdent, ") error {")
		g.P("select {") // select
		g.P("case client.fromclient <- m:")
		g.P("return nil")
		g.P("case <- client.ctx.Done():")
		g.P("client.errfromclient = client.ctx.Err()")
		g.P("client.CloseSend()")
		g.P("return client.ctx.Err()")
		g.P("}") // end select
		g.P("}")
	}
}

func genServerStream(s *protogen.Service, m *protogen.Method, g *protogen.GeneratedFile) {
	mtypename := servername(s, m, g)
	// the server type
	g.P("type ", mtypename, " struct {")
	g.P("header ", grpcMetadataMD)
	g.P("trailer ", grpcMetadataMD)
	g.P("ctx ", contextContext)
	g.P("cancel ", contextCancelFunc)
	g.P("errfromsrv error")
	g.P("errfromclient error")
	g.P("fromclient chan *", m.Input.GoIdent)
	g.P("toclient chan *", m.Output.GoIdent)
	g.P("}")

	// interface check
	g.P("var _ ", s.GoName, "_", m.GoName, "Server = (*", mtypename, ")(nil)")

	// new
	g.P("func new_", mtypename, "(ctx ", contextContext, ") ", mtypename, "{")
	g.P("newctx, cancel := ", contextWithCancel, "(ctx)")
	g.P("return ", mtypename, "{")
	g.P("ctx: newctx,")
	g.P("cancel: cancel,")
	g.P("header: ", grpcMetadataMD, "{},")
	g.P("trailer: ", grpcMetadataMD, "{},")
	g.P("fromclient: make(chan *", m.Input.GoIdent, "),")
	g.P("toclient: make(chan *", m.Output.GoIdent, "),")
	g.P("}")
	g.P("}")

	// Context
	g.P("func (server *", mtypename, ") Context() ", contextContext, "{")
	g.P("return server.ctx")
	g.P("}")

	// SetHeader
	g.P("func (server *", mtypename, ") SetHeader(m ", grpcMetadataMD, ") error {")
	g.P("for k, v := range m {")
	g.P("server.header[k] =v")
	g.P("}")
	g.P("return nil")
	g.P("}")
	// SendHeader
	g.P("func (server *", mtypename, ") SendHeader(m ", grpcMetadataMD, ") error {")
	g.P("for k, v := range m {")
	g.P("server.header[k] =v")
	g.P("}")
	g.P("return nil")
	g.P("}")

	// SetTrailer
	g.P("func (server *", mtypename, ") SetTrailer(m ", grpcMetadataMD, ") {")
	g.P("for k, v := range m {")
	g.P("server.trailer[k] =v")
	g.P("}")
	g.P("}")

	// SendMsg
	g.P("func (server *", mtypename, ") SendMsg(m any) error {")
	g.P("return ", errorsNew, "(\"unimplemented\")")
	g.P("}")

	// RecvMsg
	g.P("func (server *", mtypename, ") RecvMsg(m any) error {")
	g.P("return ", errorsNew, "(\"unimplemented\")")
	g.P("}")

	if m.Desc.IsStreamingServer() {
		// Send
		g.P("func (server *", mtypename, ") Send(m *", m.Output.GoIdent, ") error {")
		g.P("select {")
		g.P("case <- server.ctx.Done():")
		g.P("server.errfromsrv = server.ctx.Err()")
		g.P("return server.ctx.Err()")
		g.P("case server.toclient <- m:")
		g.P("return nil")
		g.P("}")
		g.P("}")
	} else {
		g.P("func (server *", mtypename, ") SendAndClose(m *", m.Output.GoIdent, ") error {")
		g.P("defer close(server.toclient)")
		g.P("select {")
		g.P("case <- server.ctx.Done():")
		g.P("server.errfromsrv = server.ctx.Err()")
		g.P("return server.ctx.Err()")
		g.P("case server.toclient <- m:")
		g.P("server.errfromsrv = ", ioEOF)
		g.P("return nil")
		g.P("}")
		g.P("}")
	}
	if m.Desc.IsStreamingClient() {
		// Recv
		g.P("func (server *", mtypename, ") Recv()( *", m.Input.GoIdent, ", error) {")
		g.P("r, ok := <- server.fromclient")
		g.P("if ok {")
		g.P("return r, nil")
		g.P("}")
		g.P("return nil, server.errfromclient")
		g.P("}")
	}
}
