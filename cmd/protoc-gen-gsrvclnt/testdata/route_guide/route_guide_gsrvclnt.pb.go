package routeguide

import (
	context "context"
	errors "errors"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	io "io"
)

type _RouteGuide_SrvClient struct {
	Server RouteGuideServer
}

var _ RouteGuideClient = (*_RouteGuide_SrvClient)(nil)

func NewRouteGuideSrvClient(server RouteGuideServer) RouteGuideClient {
	return &_RouteGuide_SrvClient{Server: server}
}
func (client *_RouteGuide_SrvClient) GetFeature(ctx context.Context, req *Point, opts ...grpc.CallOption) (*Feature, error) {
	return client.Server.GetFeature(ctx, req)
}

type _RouteGuide_ListFeaturesSrvClientStream struct {
	_RouteGuide_ListFeaturesSrvServerStream
}

var _ RouteGuide_ListFeaturesClient = (*_RouteGuide_ListFeaturesSrvClientStream)(nil)

func new__RouteGuide_ListFeaturesSrvClientStream(ctx context.Context) *_RouteGuide_ListFeaturesSrvClientStream {
	return &_RouteGuide_ListFeaturesSrvClientStream{_RouteGuide_ListFeaturesSrvServerStream: new__RouteGuide_ListFeaturesSrvServerStream(ctx)}
}
func (client *_RouteGuide_ListFeaturesSrvClientStream) Context() context.Context {
	return client.ctx
}
func (client *_RouteGuide_ListFeaturesSrvClientStream) Header() (metadata.MD, error) {
	return client.header, nil
}
func (client *_RouteGuide_ListFeaturesSrvClientStream) Trailer() metadata.MD {
	return client.trailer
}
func (client *_RouteGuide_ListFeaturesSrvClientStream) CloseSend() error {
	if client.errfromclient == nil {
		client.errfromclient = io.EOF
	}
	close(client.fromclient)
	return nil
}
func (client *_RouteGuide_ListFeaturesSrvClientStream) SendMsg(m any) error {
	return errors.New("unimplemented")
}
func (client *_RouteGuide_ListFeaturesSrvClientStream) RecvMsg(m any) error {
	return errors.New("unimplemented")
}
func (client *_RouteGuide_ListFeaturesSrvClientStream) Recv() (*Feature, error) {
	r, ok := <-client.toclient
	if ok {
		return r, nil
	} else {
		return nil, client.errfromsrv
	}
}

type _RouteGuide_ListFeaturesSrvServerStream struct {
	header        metadata.MD
	trailer       metadata.MD
	ctx           context.Context
	errfromsrv    error
	errfromclient error
	fromclient    chan *Rectangle
	toclient      chan *Feature
}

var _ RouteGuide_ListFeaturesServer = (*_RouteGuide_ListFeaturesSrvServerStream)(nil)

func (server *_RouteGuide_ListFeaturesSrvServerStream) Context() context.Context {
	return server.ctx
}
func new__RouteGuide_ListFeaturesSrvServerStream(ctx context.Context) _RouteGuide_ListFeaturesSrvServerStream {
	return _RouteGuide_ListFeaturesSrvServerStream{
		ctx:        ctx,
		header:     metadata.MD{},
		trailer:    metadata.MD{},
		fromclient: make(chan *Rectangle),
		toclient:   make(chan *Feature),
	}
}
func (server *_RouteGuide_ListFeaturesSrvServerStream) SetHeader(m metadata.MD) error {
	for k, v := range m {
		server.header[k] = v
	}
	return nil
}
func (server *_RouteGuide_ListFeaturesSrvServerStream) SendHeader(m metadata.MD) error {
	for k, v := range m {
		server.header[k] = v
	}
	return nil
}
func (server *_RouteGuide_ListFeaturesSrvServerStream) SetTrailer(m metadata.MD) {
	for k, v := range m {
		server.trailer[k] = v
	}
}
func (server *_RouteGuide_ListFeaturesSrvServerStream) SendMsg(m any) error {
	return errors.New("unimplemented")
}
func (server *_RouteGuide_ListFeaturesSrvServerStream) RecvMsg(m any) error {
	return errors.New("unimplemented")
}
func (server *_RouteGuide_ListFeaturesSrvServerStream) Send(m *Feature) error {
	select {
	case <-server.ctx.Done():
		server.errfromsrv = server.ctx.Err()
		return server.ctx.Err()
	case server.toclient <- m:
		return nil
	}
}

// do server streaming
func (client *_RouteGuide_SrvClient) ListFeatures(ctx context.Context, req *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error) {
	r := new__RouteGuide_ListFeaturesSrvClientStream(ctx)
	go func() {
		err := client.Server.ListFeatures(req, &r._RouteGuide_ListFeaturesSrvServerStream)
		if err != nil {
			r.errfromsrv = err
		} else {
			r.errfromsrv = io.EOF
		}
		close(r.toclient)
	}()
	return r, nil
}

type _RouteGuide_RecordRouteSrvClientStream struct {
	_RouteGuide_RecordRouteSrvServerStream
}

var _ RouteGuide_RecordRouteClient = (*_RouteGuide_RecordRouteSrvClientStream)(nil)

func new__RouteGuide_RecordRouteSrvClientStream(ctx context.Context) *_RouteGuide_RecordRouteSrvClientStream {
	return &_RouteGuide_RecordRouteSrvClientStream{_RouteGuide_RecordRouteSrvServerStream: new__RouteGuide_RecordRouteSrvServerStream(ctx)}
}
func (client *_RouteGuide_RecordRouteSrvClientStream) Context() context.Context {
	return client.ctx
}
func (client *_RouteGuide_RecordRouteSrvClientStream) Header() (metadata.MD, error) {
	return client.header, nil
}
func (client *_RouteGuide_RecordRouteSrvClientStream) Trailer() metadata.MD {
	return client.trailer
}
func (client *_RouteGuide_RecordRouteSrvClientStream) CloseSend() error {
	if client.errfromclient == nil {
		client.errfromclient = io.EOF
	}
	close(client.fromclient)
	return nil
}
func (client *_RouteGuide_RecordRouteSrvClientStream) SendMsg(m any) error {
	return errors.New("unimplemented")
}
func (client *_RouteGuide_RecordRouteSrvClientStream) RecvMsg(m any) error {
	return errors.New("unimplemented")
}
func (client *_RouteGuide_RecordRouteSrvClientStream) CloseAndRecv() (*RouteSummary, error) {
	client.errfromclient = io.EOF
	if err := client.CloseSend(); err != nil {
		return nil, err
	}
	r, ok := <-client.toclient
	if ok {
		return r, nil
	} else {
		return nil, client.errfromsrv
	}
}
func (client *_RouteGuide_RecordRouteSrvClientStream) Send(m *Point) error {
	select {
	case client.fromclient <- m:
		return nil
	case <-client.ctx.Done():
		client.errfromclient = client.ctx.Err()
		return client.CloseSend()
	}
}

type _RouteGuide_RecordRouteSrvServerStream struct {
	header        metadata.MD
	trailer       metadata.MD
	ctx           context.Context
	errfromsrv    error
	errfromclient error
	fromclient    chan *Point
	toclient      chan *RouteSummary
}

var _ RouteGuide_RecordRouteServer = (*_RouteGuide_RecordRouteSrvServerStream)(nil)

func (server *_RouteGuide_RecordRouteSrvServerStream) Context() context.Context {
	return server.ctx
}
func new__RouteGuide_RecordRouteSrvServerStream(ctx context.Context) _RouteGuide_RecordRouteSrvServerStream {
	return _RouteGuide_RecordRouteSrvServerStream{
		ctx:        ctx,
		header:     metadata.MD{},
		trailer:    metadata.MD{},
		fromclient: make(chan *Point),
		toclient:   make(chan *RouteSummary),
	}
}
func (server *_RouteGuide_RecordRouteSrvServerStream) SetHeader(m metadata.MD) error {
	for k, v := range m {
		server.header[k] = v
	}
	return nil
}
func (server *_RouteGuide_RecordRouteSrvServerStream) SendHeader(m metadata.MD) error {
	for k, v := range m {
		server.header[k] = v
	}
	return nil
}
func (server *_RouteGuide_RecordRouteSrvServerStream) SetTrailer(m metadata.MD) {
	for k, v := range m {
		server.trailer[k] = v
	}
}
func (server *_RouteGuide_RecordRouteSrvServerStream) SendMsg(m any) error {
	return errors.New("unimplemented")
}
func (server *_RouteGuide_RecordRouteSrvServerStream) RecvMsg(m any) error {
	return errors.New("unimplemented")
}
func (server *_RouteGuide_RecordRouteSrvServerStream) SendAndClose(m *RouteSummary) error {
	defer close(server.toclient)
	select {
	case <-server.ctx.Done():
		server.errfromsrv = server.ctx.Err()
		return server.ctx.Err()
	case server.toclient <- m:
		server.errfromsrv = io.EOF
		return nil
	}
}
func (server *_RouteGuide_RecordRouteSrvServerStream) Recv() (*Point, error) {
	r, ok := <-server.fromclient
	if ok {
		return r, nil
	}
	return nil, server.errfromclient
}

// do client streaming
func (client *_RouteGuide_SrvClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error) {
	r := new__RouteGuide_RecordRouteSrvClientStream(ctx)
	go func() {
		err := client.Server.RecordRoute(&r._RouteGuide_RecordRouteSrvServerStream)
		if err != nil && r.errfromsrv == nil {
			r.errfromsrv = err
		}
	}()
	return r, nil
}

type _RouteGuide_RouteChatSrvClientStream struct {
	_RouteGuide_RouteChatSrvServerStream
}

var _ RouteGuide_RouteChatClient = (*_RouteGuide_RouteChatSrvClientStream)(nil)

func new__RouteGuide_RouteChatSrvClientStream(ctx context.Context) *_RouteGuide_RouteChatSrvClientStream {
	return &_RouteGuide_RouteChatSrvClientStream{_RouteGuide_RouteChatSrvServerStream: new__RouteGuide_RouteChatSrvServerStream(ctx)}
}
func (client *_RouteGuide_RouteChatSrvClientStream) Context() context.Context {
	return client.ctx
}
func (client *_RouteGuide_RouteChatSrvClientStream) Header() (metadata.MD, error) {
	return client.header, nil
}
func (client *_RouteGuide_RouteChatSrvClientStream) Trailer() metadata.MD {
	return client.trailer
}
func (client *_RouteGuide_RouteChatSrvClientStream) CloseSend() error {
	if client.errfromclient == nil {
		client.errfromclient = io.EOF
	}
	close(client.fromclient)
	return nil
}
func (client *_RouteGuide_RouteChatSrvClientStream) SendMsg(m any) error {
	return errors.New("unimplemented")
}
func (client *_RouteGuide_RouteChatSrvClientStream) RecvMsg(m any) error {
	return errors.New("unimplemented")
}
func (client *_RouteGuide_RouteChatSrvClientStream) Recv() (*RouteNote, error) {
	r, ok := <-client.toclient
	if ok {
		return r, nil
	} else {
		return nil, client.errfromsrv
	}
}
func (client *_RouteGuide_RouteChatSrvClientStream) Send(m *RouteNote) error {
	select {
	case client.fromclient <- m:
		return nil
	case <-client.ctx.Done():
		client.errfromclient = client.ctx.Err()
		return client.CloseSend()
	}
}

type _RouteGuide_RouteChatSrvServerStream struct {
	header        metadata.MD
	trailer       metadata.MD
	ctx           context.Context
	errfromsrv    error
	errfromclient error
	fromclient    chan *RouteNote
	toclient      chan *RouteNote
}

var _ RouteGuide_RouteChatServer = (*_RouteGuide_RouteChatSrvServerStream)(nil)

func (server *_RouteGuide_RouteChatSrvServerStream) Context() context.Context {
	return server.ctx
}
func new__RouteGuide_RouteChatSrvServerStream(ctx context.Context) _RouteGuide_RouteChatSrvServerStream {
	return _RouteGuide_RouteChatSrvServerStream{
		ctx:        ctx,
		header:     metadata.MD{},
		trailer:    metadata.MD{},
		fromclient: make(chan *RouteNote),
		toclient:   make(chan *RouteNote),
	}
}
func (server *_RouteGuide_RouteChatSrvServerStream) SetHeader(m metadata.MD) error {
	for k, v := range m {
		server.header[k] = v
	}
	return nil
}
func (server *_RouteGuide_RouteChatSrvServerStream) SendHeader(m metadata.MD) error {
	for k, v := range m {
		server.header[k] = v
	}
	return nil
}
func (server *_RouteGuide_RouteChatSrvServerStream) SetTrailer(m metadata.MD) {
	for k, v := range m {
		server.trailer[k] = v
	}
}
func (server *_RouteGuide_RouteChatSrvServerStream) SendMsg(m any) error {
	return errors.New("unimplemented")
}
func (server *_RouteGuide_RouteChatSrvServerStream) RecvMsg(m any) error {
	return errors.New("unimplemented")
}
func (server *_RouteGuide_RouteChatSrvServerStream) Send(m *RouteNote) error {
	select {
	case <-server.ctx.Done():
		server.errfromsrv = server.ctx.Err()
		return server.ctx.Err()
	case server.toclient <- m:
		return nil
	}
}
func (server *_RouteGuide_RouteChatSrvServerStream) Recv() (*RouteNote, error) {
	r, ok := <-server.fromclient
	if ok {
		return r, nil
	}
	return nil, server.errfromclient
}

// do bidirectional streaming
func (client *_RouteGuide_SrvClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error) {
	r := new__RouteGuide_RouteChatSrvClientStream(ctx)
	go func() {
		err := client.Server.RouteChat(&r._RouteGuide_RouteChatSrvServerStream)
		if err != nil {
			r.errfromsrv = err
		} else if r.errfromsrv == nil {
			r.errfromsrv = io.EOF
		}
		close(r.toclient)
	}()
	return r, nil
}
