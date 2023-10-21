// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0
// source: proto/books_users.proto

package pb_books_users

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

const (
	BooksUsersService_RentBook_FullMethodName = "/books_users_service.books_users.BooksUsersService/RentBook"
)

// BooksUsersServiceClient is the client API for BooksUsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BooksUsersServiceClient interface {
	RentBook(ctx context.Context, in *RentBookRequest, opts ...grpc.CallOption) (*RentBookResponse, error)
}

type booksUsersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBooksUsersServiceClient(cc grpc.ClientConnInterface) BooksUsersServiceClient {
	return &booksUsersServiceClient{cc}
}

func (c *booksUsersServiceClient) RentBook(ctx context.Context, in *RentBookRequest, opts ...grpc.CallOption) (*RentBookResponse, error) {
	out := new(RentBookResponse)
	err := c.cc.Invoke(ctx, BooksUsersService_RentBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BooksUsersServiceServer is the server API for BooksUsersService service.
// All implementations must embed UnimplementedBooksUsersServiceServer
// for forward compatibility
type BooksUsersServiceServer interface {
	RentBook(context.Context, *RentBookRequest) (*RentBookResponse, error)
	mustEmbedUnimplementedBooksUsersServiceServer()
}

// UnimplementedBooksUsersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBooksUsersServiceServer struct {
}

func (UnimplementedBooksUsersServiceServer) RentBook(context.Context, *RentBookRequest) (*RentBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RentBook not implemented")
}
func (UnimplementedBooksUsersServiceServer) mustEmbedUnimplementedBooksUsersServiceServer() {}

// UnsafeBooksUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BooksUsersServiceServer will
// result in compilation errors.
type UnsafeBooksUsersServiceServer interface {
	mustEmbedUnimplementedBooksUsersServiceServer()
}

func RegisterBooksUsersServiceServer(s grpc.ServiceRegistrar, srv BooksUsersServiceServer) {
	s.RegisterService(&BooksUsersService_ServiceDesc, srv)
}

func _BooksUsersService_RentBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RentBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksUsersServiceServer).RentBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BooksUsersService_RentBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksUsersServiceServer).RentBook(ctx, req.(*RentBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BooksUsersService_ServiceDesc is the grpc.ServiceDesc for BooksUsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BooksUsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "books_users_service.books_users.BooksUsersService",
	HandlerType: (*BooksUsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RentBook",
			Handler:    _BooksUsersService_RentBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/books_users.proto",
}
