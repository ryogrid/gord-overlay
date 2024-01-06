// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: server/public.proto

package serverconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	server "github.com/ryogrid/gord-overlay/server"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ExternalServiceName is the fully-qualified name of the ExternalService service.
	ExternalServiceName = "server.ExternalService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ExternalServiceFindHostForKeyProcedure is the fully-qualified name of the ExternalService's
	// FindHostForKey RPC.
	ExternalServiceFindHostForKeyProcedure = "/server.ExternalService/FindHostForKey"
	// ExternalServicePutValueProcedure is the fully-qualified name of the ExternalService's PutValue
	// RPC.
	ExternalServicePutValueProcedure = "/server.ExternalService/PutValue"
	// ExternalServiceGetValueProcedure is the fully-qualified name of the ExternalService's GetValue
	// RPC.
	ExternalServiceGetValueProcedure = "/server.ExternalService/GetValue"
	// ExternalServiceDeleteValueProcedure is the fully-qualified name of the ExternalService's
	// DeleteValue RPC.
	ExternalServiceDeleteValueProcedure = "/server.ExternalService/DeleteValue"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	externalServiceServiceDescriptor              = server.File_server_public_proto.Services().ByName("ExternalService")
	externalServiceFindHostForKeyMethodDescriptor = externalServiceServiceDescriptor.Methods().ByName("FindHostForKey")
	externalServicePutValueMethodDescriptor       = externalServiceServiceDescriptor.Methods().ByName("PutValue")
	externalServiceGetValueMethodDescriptor       = externalServiceServiceDescriptor.Methods().ByName("GetValue")
	externalServiceDeleteValueMethodDescriptor    = externalServiceServiceDescriptor.Methods().ByName("DeleteValue")
)

// ExternalServiceClient is a client for the server.ExternalService service.
type ExternalServiceClient interface {
	FindHostForKey(context.Context, *connect.Request[server.FindHostRequest]) (*connect.Response[server.Node], error)
	PutValue(context.Context, *connect.Request[server.PutValueRequest]) (*connect.Response[server.PutValueResponse], error)
	GetValue(context.Context, *connect.Request[server.GetValueRequest]) (*connect.Response[server.GetValueResponse], error)
	DeleteValue(context.Context, *connect.Request[server.DeleteValueRequest]) (*connect.Response[server.DeleteValueResponse], error)
}

// NewExternalServiceClient constructs a client for the server.ExternalService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewExternalServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ExternalServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &externalServiceClient{
		findHostForKey: connect.NewClient[server.FindHostRequest, server.Node](
			httpClient,
			baseURL+ExternalServiceFindHostForKeyProcedure,
			connect.WithSchema(externalServiceFindHostForKeyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		putValue: connect.NewClient[server.PutValueRequest, server.PutValueResponse](
			httpClient,
			baseURL+ExternalServicePutValueProcedure,
			connect.WithSchema(externalServicePutValueMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getValue: connect.NewClient[server.GetValueRequest, server.GetValueResponse](
			httpClient,
			baseURL+ExternalServiceGetValueProcedure,
			connect.WithSchema(externalServiceGetValueMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteValue: connect.NewClient[server.DeleteValueRequest, server.DeleteValueResponse](
			httpClient,
			baseURL+ExternalServiceDeleteValueProcedure,
			connect.WithSchema(externalServiceDeleteValueMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// externalServiceClient implements ExternalServiceClient.
type externalServiceClient struct {
	findHostForKey *connect.Client[server.FindHostRequest, server.Node]
	putValue       *connect.Client[server.PutValueRequest, server.PutValueResponse]
	getValue       *connect.Client[server.GetValueRequest, server.GetValueResponse]
	deleteValue    *connect.Client[server.DeleteValueRequest, server.DeleteValueResponse]
}

// FindHostForKey calls server.ExternalService.FindHostForKey.
func (c *externalServiceClient) FindHostForKey(ctx context.Context, req *connect.Request[server.FindHostRequest]) (*connect.Response[server.Node], error) {
	return c.findHostForKey.CallUnary(ctx, req)
}

// PutValue calls server.ExternalService.PutValue.
func (c *externalServiceClient) PutValue(ctx context.Context, req *connect.Request[server.PutValueRequest]) (*connect.Response[server.PutValueResponse], error) {
	return c.putValue.CallUnary(ctx, req)
}

// GetValue calls server.ExternalService.GetValue.
func (c *externalServiceClient) GetValue(ctx context.Context, req *connect.Request[server.GetValueRequest]) (*connect.Response[server.GetValueResponse], error) {
	return c.getValue.CallUnary(ctx, req)
}

// DeleteValue calls server.ExternalService.DeleteValue.
func (c *externalServiceClient) DeleteValue(ctx context.Context, req *connect.Request[server.DeleteValueRequest]) (*connect.Response[server.DeleteValueResponse], error) {
	return c.deleteValue.CallUnary(ctx, req)
}

// ExternalServiceHandler is an implementation of the server.ExternalService service.
type ExternalServiceHandler interface {
	FindHostForKey(context.Context, *connect.Request[server.FindHostRequest]) (*connect.Response[server.Node], error)
	PutValue(context.Context, *connect.Request[server.PutValueRequest]) (*connect.Response[server.PutValueResponse], error)
	GetValue(context.Context, *connect.Request[server.GetValueRequest]) (*connect.Response[server.GetValueResponse], error)
	DeleteValue(context.Context, *connect.Request[server.DeleteValueRequest]) (*connect.Response[server.DeleteValueResponse], error)
}

// NewExternalServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewExternalServiceHandler(svc ExternalServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	externalServiceFindHostForKeyHandler := connect.NewUnaryHandler(
		ExternalServiceFindHostForKeyProcedure,
		svc.FindHostForKey,
		connect.WithSchema(externalServiceFindHostForKeyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	externalServicePutValueHandler := connect.NewUnaryHandler(
		ExternalServicePutValueProcedure,
		svc.PutValue,
		connect.WithSchema(externalServicePutValueMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	externalServiceGetValueHandler := connect.NewUnaryHandler(
		ExternalServiceGetValueProcedure,
		svc.GetValue,
		connect.WithSchema(externalServiceGetValueMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	externalServiceDeleteValueHandler := connect.NewUnaryHandler(
		ExternalServiceDeleteValueProcedure,
		svc.DeleteValue,
		connect.WithSchema(externalServiceDeleteValueMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/server.ExternalService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ExternalServiceFindHostForKeyProcedure:
			externalServiceFindHostForKeyHandler.ServeHTTP(w, r)
		case ExternalServicePutValueProcedure:
			externalServicePutValueHandler.ServeHTTP(w, r)
		case ExternalServiceGetValueProcedure:
			externalServiceGetValueHandler.ServeHTTP(w, r)
		case ExternalServiceDeleteValueProcedure:
			externalServiceDeleteValueHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedExternalServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedExternalServiceHandler struct{}

func (UnimplementedExternalServiceHandler) FindHostForKey(context.Context, *connect.Request[server.FindHostRequest]) (*connect.Response[server.Node], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("server.ExternalService.FindHostForKey is not implemented"))
}

func (UnimplementedExternalServiceHandler) PutValue(context.Context, *connect.Request[server.PutValueRequest]) (*connect.Response[server.PutValueResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("server.ExternalService.PutValue is not implemented"))
}

func (UnimplementedExternalServiceHandler) GetValue(context.Context, *connect.Request[server.GetValueRequest]) (*connect.Response[server.GetValueResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("server.ExternalService.GetValue is not implemented"))
}

func (UnimplementedExternalServiceHandler) DeleteValue(context.Context, *connect.Request[server.DeleteValueRequest]) (*connect.Response[server.DeleteValueResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("server.ExternalService.DeleteValue is not implemented"))
}
