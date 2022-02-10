// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: content.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ContentClient is the client API for Content service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentClient interface {
	// Insert new post, id should be left empty.
	// Returns id of newly created post.
	InsertPost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*InsertPostResponse, error)
	// Update existing post, id must be filled.
	// Returns update success status.
	UpdatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*UpdatePostResponse, error)
	// Get post by id.
	// Returns Post object.
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error)
}

type contentClient struct {
	cc grpc.ClientConnInterface
}

func NewContentClient(cc grpc.ClientConnInterface) ContentClient {
	return &contentClient{cc}
}

func (c *contentClient) InsertPost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*InsertPostResponse, error) {
	out := new(InsertPostResponse)
	err := c.cc.Invoke(ctx, "/content.Content/InsertPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) UpdatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*UpdatePostResponse, error) {
	out := new(UpdatePostResponse)
	err := c.cc.Invoke(ctx, "/content.Content/UpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/content.Content/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServer is the server API for Content service.
// All implementations must embed UnimplementedContentServer
// for forward compatibility
type ContentServer interface {
	// Insert new post, id should be left empty.
	// Returns id of newly created post.
	InsertPost(context.Context, *Post) (*InsertPostResponse, error)
	// Update existing post, id must be filled.
	// Returns update success status.
	UpdatePost(context.Context, *Post) (*UpdatePostResponse, error)
	// Get post by id.
	// Returns Post object.
	GetPost(context.Context, *GetPostRequest) (*Post, error)
	mustEmbedUnimplementedContentServer()
}

// UnimplementedContentServer must be embedded to have forward compatible implementations.
type UnimplementedContentServer struct {
}

func (UnimplementedContentServer) InsertPost(context.Context, *Post) (*InsertPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertPost not implemented")
}
func (UnimplementedContentServer) UpdatePost(context.Context, *Post) (*UpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedContentServer) GetPost(context.Context, *GetPostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedContentServer) mustEmbedUnimplementedContentServer() {}

// UnsafeContentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServer will
// result in compilation errors.
type UnsafeContentServer interface {
	mustEmbedUnimplementedContentServer()
}

func RegisterContentServer(s grpc.ServiceRegistrar, srv ContentServer) {
	s.RegisterService(&Content_ServiceDesc, srv)
}

func _Content_InsertPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).InsertPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.Content/InsertPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).InsertPost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.Content/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).UpdatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.Content/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Content_ServiceDesc is the grpc.ServiceDesc for Content service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Content_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content.Content",
	HandlerType: (*ContentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertPost",
			Handler:    _Content_InsertPost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _Content_UpdatePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _Content_GetPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content.proto",
}