// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: da-client/disperser.proto

package da_client

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Disperser_DisperseBlob_FullMethodName  = "/disperser.Disperser/DisperseBlob"
	Disperser_GetBlobStatus_FullMethodName = "/disperser.Disperser/GetBlobStatus"
	Disperser_RetrieveBlob_FullMethodName  = "/disperser.Disperser/RetrieveBlob"
	Disperser_GetStatus_FullMethodName     = "/disperser.Disperser/GetStatus"
)

// DisperserClient is the client API for Disperser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Disperser defines the public APIs for dispersing blobs.
type DisperserClient interface {
	// This API accepts blob to disperse from clients.
	// This executes the dispersal async, i.e. it returns once the request
	// is accepted. The client could use GetBlobStatus() API to poll the the
	// processing status of the blob.
	DisperseBlob(ctx context.Context, in *DisperseBlobRequest, opts ...grpc.CallOption) (*DisperseBlobReply, error)
	// This API is meant to be polled for the blob status.
	GetBlobStatus(ctx context.Context, in *BlobStatusRequest, opts ...grpc.CallOption) (*BlobStatusReply, error)
	// This retrieves the requested blob from the Disperser's backend.
	// This is a more efficient way to retrieve blobs than directly retrieving
	// from the DA Nodes (see detail about this approach in
	// api/proto/retriever/retriever.proto).
	// The blob should have been initially dispersed via this Disperser service
	// for this API to work.
	RetrieveBlob(ctx context.Context, in *RetrieveBlobRequest, opts ...grpc.CallOption) (*RetrieveBlobReply, error)
	// This checks the disperser status
	GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetStatusReply, error)
}

type disperserClient struct {
	cc grpc.ClientConnInterface
}

func NewDisperserClient(cc grpc.ClientConnInterface) DisperserClient {
	return &disperserClient{cc}
}

func (c *disperserClient) DisperseBlob(ctx context.Context, in *DisperseBlobRequest, opts ...grpc.CallOption) (*DisperseBlobReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DisperseBlobReply)
	err := c.cc.Invoke(ctx, Disperser_DisperseBlob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disperserClient) GetBlobStatus(ctx context.Context, in *BlobStatusRequest, opts ...grpc.CallOption) (*BlobStatusReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlobStatusReply)
	err := c.cc.Invoke(ctx, Disperser_GetBlobStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disperserClient) RetrieveBlob(ctx context.Context, in *RetrieveBlobRequest, opts ...grpc.CallOption) (*RetrieveBlobReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RetrieveBlobReply)
	err := c.cc.Invoke(ctx, Disperser_RetrieveBlob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disperserClient) GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetStatusReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatusReply)
	err := c.cc.Invoke(ctx, Disperser_GetStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DisperserServer is the server API for Disperser service.
// All implementations must embed UnimplementedDisperserServer
// for forward compatibility
//
// Disperser defines the public APIs for dispersing blobs.
type DisperserServer interface {
	// This API accepts blob to disperse from clients.
	// This executes the dispersal async, i.e. it returns once the request
	// is accepted. The client could use GetBlobStatus() API to poll the the
	// processing status of the blob.
	DisperseBlob(context.Context, *DisperseBlobRequest) (*DisperseBlobReply, error)
	// This API is meant to be polled for the blob status.
	GetBlobStatus(context.Context, *BlobStatusRequest) (*BlobStatusReply, error)
	// This retrieves the requested blob from the Disperser's backend.
	// This is a more efficient way to retrieve blobs than directly retrieving
	// from the DA Nodes (see detail about this approach in
	// api/proto/retriever/retriever.proto).
	// The blob should have been initially dispersed via this Disperser service
	// for this API to work.
	RetrieveBlob(context.Context, *RetrieveBlobRequest) (*RetrieveBlobReply, error)
	// This checks the disperser status
	GetStatus(context.Context, *Empty) (*GetStatusReply, error)
	mustEmbedUnimplementedDisperserServer()
}

// UnimplementedDisperserServer must be embedded to have forward compatible implementations.
type UnimplementedDisperserServer struct {
}

func (UnimplementedDisperserServer) DisperseBlob(context.Context, *DisperseBlobRequest) (*DisperseBlobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisperseBlob not implemented")
}
func (UnimplementedDisperserServer) GetBlobStatus(context.Context, *BlobStatusRequest) (*BlobStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlobStatus not implemented")
}
func (UnimplementedDisperserServer) RetrieveBlob(context.Context, *RetrieveBlobRequest) (*RetrieveBlobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveBlob not implemented")
}
func (UnimplementedDisperserServer) GetStatus(context.Context, *Empty) (*GetStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedDisperserServer) mustEmbedUnimplementedDisperserServer() {}

// UnsafeDisperserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DisperserServer will
// result in compilation errors.
type UnsafeDisperserServer interface {
	mustEmbedUnimplementedDisperserServer()
}

func RegisterDisperserServer(s grpc.ServiceRegistrar, srv DisperserServer) {
	s.RegisterService(&Disperser_ServiceDesc, srv)
}

func _Disperser_DisperseBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisperseBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisperserServer).DisperseBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Disperser_DisperseBlob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisperserServer).DisperseBlob(ctx, req.(*DisperseBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disperser_GetBlobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisperserServer).GetBlobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Disperser_GetBlobStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisperserServer).GetBlobStatus(ctx, req.(*BlobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disperser_RetrieveBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisperserServer).RetrieveBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Disperser_RetrieveBlob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisperserServer).RetrieveBlob(ctx, req.(*RetrieveBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disperser_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisperserServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Disperser_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisperserServer).GetStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Disperser_ServiceDesc is the grpc.ServiceDesc for Disperser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Disperser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "disperser.Disperser",
	HandlerType: (*DisperserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DisperseBlob",
			Handler:    _Disperser_DisperseBlob_Handler,
		},
		{
			MethodName: "GetBlobStatus",
			Handler:    _Disperser_GetBlobStatus_Handler,
		},
		{
			MethodName: "RetrieveBlob",
			Handler:    _Disperser_RetrieveBlob_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Disperser_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "da-client/disperser.proto",
}