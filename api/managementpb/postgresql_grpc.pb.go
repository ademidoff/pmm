// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: managementpb/postgresql.proto

package managementpb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PostgreSQL_AddPostgreSQL_FullMethodName = "/management.PostgreSQL/AddPostgreSQL"
)

// PostgreSQLClient is the client API for PostgreSQL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// PostgreSQL service provides public Management API methods for PostgreSQL Service.
type PostgreSQLClient interface {
	// AddPostgreSQL adds PostgreSQL Service and starts postgres exporter.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "postgres_exporter" with provided "pmm_agent_id" and other parameters.
	AddPostgreSQL(ctx context.Context, in *AddPostgreSQLRequest, opts ...grpc.CallOption) (*AddPostgreSQLResponse, error)
}

type postgreSQLClient struct {
	cc grpc.ClientConnInterface
}

func NewPostgreSQLClient(cc grpc.ClientConnInterface) PostgreSQLClient {
	return &postgreSQLClient{cc}
}

func (c *postgreSQLClient) AddPostgreSQL(ctx context.Context, in *AddPostgreSQLRequest, opts ...grpc.CallOption) (*AddPostgreSQLResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddPostgreSQLResponse)
	err := c.cc.Invoke(ctx, PostgreSQL_AddPostgreSQL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostgreSQLServer is the server API for PostgreSQL service.
// All implementations must embed UnimplementedPostgreSQLServer
// for forward compatibility.
//
// PostgreSQL service provides public Management API methods for PostgreSQL Service.
type PostgreSQLServer interface {
	// AddPostgreSQL adds PostgreSQL Service and starts postgres exporter.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "postgres_exporter" with provided "pmm_agent_id" and other parameters.
	AddPostgreSQL(context.Context, *AddPostgreSQLRequest) (*AddPostgreSQLResponse, error)
	mustEmbedUnimplementedPostgreSQLServer()
}

// UnimplementedPostgreSQLServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPostgreSQLServer struct{}

func (UnimplementedPostgreSQLServer) AddPostgreSQL(context.Context, *AddPostgreSQLRequest) (*AddPostgreSQLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPostgreSQL not implemented")
}
func (UnimplementedPostgreSQLServer) mustEmbedUnimplementedPostgreSQLServer() {}
func (UnimplementedPostgreSQLServer) testEmbeddedByValue()                    {}

// UnsafePostgreSQLServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostgreSQLServer will
// result in compilation errors.
type UnsafePostgreSQLServer interface {
	mustEmbedUnimplementedPostgreSQLServer()
}

func RegisterPostgreSQLServer(s grpc.ServiceRegistrar, srv PostgreSQLServer) {
	// If the following call pancis, it indicates UnimplementedPostgreSQLServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PostgreSQL_ServiceDesc, srv)
}

func _PostgreSQL_AddPostgreSQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPostgreSQLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgreSQLServer).AddPostgreSQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostgreSQL_AddPostgreSQL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgreSQLServer).AddPostgreSQL(ctx, req.(*AddPostgreSQLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostgreSQL_ServiceDesc is the grpc.ServiceDesc for PostgreSQL service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostgreSQL_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.PostgreSQL",
	HandlerType: (*PostgreSQLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPostgreSQL",
			Handler:    _PostgreSQL_AddPostgreSQL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/postgresql.proto",
}
