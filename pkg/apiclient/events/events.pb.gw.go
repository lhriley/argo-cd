// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: server/application/events.proto

/*
Package events is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package events

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage
var _ = metadata.Join

var (
	filter_Eventing_StartEventSource_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_Eventing_StartEventSource_0(ctx context.Context, marshaler runtime.Marshaler, client EventingClient, req *http.Request, pathParams map[string]string) (Eventing_StartEventSourceClient, runtime.ServerMetadata, error) {
	var protoReq EventSource
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_Eventing_StartEventSource_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.StartEventSource(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterEventingHandlerServer registers the http handlers for service Eventing to "mux".
// UnaryRPC     :call EventingServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterEventingHandlerFromEndpoint instead.
func RegisterEventingHandlerServer(ctx context.Context, mux *runtime.ServeMux, server EventingServer) error {

	mux.Handle("GET", pattern_Eventing_StartEventSource_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterEventingHandlerFromEndpoint is same as RegisterEventingHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterEventingHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterEventingHandler(ctx, mux, conn)
}

// RegisterEventingHandler registers the http handlers for service Eventing to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterEventingHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterEventingHandlerClient(ctx, mux, NewEventingClient(conn))
}

// RegisterEventingHandlerClient registers the http handlers for service Eventing
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "EventingClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "EventingClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "EventingClient" to call the correct interceptors.
func RegisterEventingHandlerClient(ctx context.Context, mux *runtime.ServeMux, client EventingClient) error {

	mux.Handle("GET", pattern_Eventing_StartEventSource_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Eventing_StartEventSource_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Eventing_StartEventSource_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Eventing_StartEventSource_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v1", "stream", "events"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_Eventing_StartEventSource_0 = runtime.ForwardResponseStream
)
