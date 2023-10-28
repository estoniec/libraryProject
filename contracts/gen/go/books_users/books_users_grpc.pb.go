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
	BooksUsersService_RentBook_FullMethodName      = "/books_users_service.books_users.BooksUsersService/RentBook"
	BooksUsersService_ConfirmRent_FullMethodName   = "/books_users_service.books_users.BooksUsersService/ConfirmRent"
	BooksUsersService_ConfirmReturn_FullMethodName = "/books_users_service.books_users.BooksUsersService/ConfirmReturn"
	BooksUsersService_FindBook_FullMethodName      = "/books_users_service.books_users.BooksUsersService/FindBook"
	BooksUsersService_GetDebt_FullMethodName       = "/books_users_service.books_users.BooksUsersService/GetDebt"
)

// BooksUsersServiceClient is the client API for BooksUsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BooksUsersServiceClient interface {
	RentBook(ctx context.Context, in *RentBookRequest, opts ...grpc.CallOption) (*RentBookResponse, error)
	ConfirmRent(ctx context.Context, in *ConfirmRentRequest, opts ...grpc.CallOption) (*ConfirmRentResponse, error)
	ConfirmReturn(ctx context.Context, in *ConfirmReturnRequest, opts ...grpc.CallOption) (*ConfirmReturnResponse, error)
	FindBook(ctx context.Context, in *FindBookRequest, opts ...grpc.CallOption) (*FindBookResponse, error)
	GetDebt(ctx context.Context, in *GetDebtRequest, opts ...grpc.CallOption) (*GetDebtResponse, error)
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

func (c *booksUsersServiceClient) ConfirmRent(ctx context.Context, in *ConfirmRentRequest, opts ...grpc.CallOption) (*ConfirmRentResponse, error) {
	out := new(ConfirmRentResponse)
	err := c.cc.Invoke(ctx, BooksUsersService_ConfirmRent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksUsersServiceClient) ConfirmReturn(ctx context.Context, in *ConfirmReturnRequest, opts ...grpc.CallOption) (*ConfirmReturnResponse, error) {
	out := new(ConfirmReturnResponse)
	err := c.cc.Invoke(ctx, BooksUsersService_ConfirmReturn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksUsersServiceClient) FindBook(ctx context.Context, in *FindBookRequest, opts ...grpc.CallOption) (*FindBookResponse, error) {
	out := new(FindBookResponse)
	err := c.cc.Invoke(ctx, BooksUsersService_FindBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksUsersServiceClient) GetDebt(ctx context.Context, in *GetDebtRequest, opts ...grpc.CallOption) (*GetDebtResponse, error) {
	out := new(GetDebtResponse)
	err := c.cc.Invoke(ctx, BooksUsersService_GetDebt_FullMethodName, in, out, opts...)
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
	ConfirmRent(context.Context, *ConfirmRentRequest) (*ConfirmRentResponse, error)
	ConfirmReturn(context.Context, *ConfirmReturnRequest) (*ConfirmReturnResponse, error)
	FindBook(context.Context, *FindBookRequest) (*FindBookResponse, error)
	GetDebt(context.Context, *GetDebtRequest) (*GetDebtResponse, error)
	mustEmbedUnimplementedBooksUsersServiceServer()
}

// UnimplementedBooksUsersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBooksUsersServiceServer struct {
}

func (UnimplementedBooksUsersServiceServer) RentBook(context.Context, *RentBookRequest) (*RentBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RentBook not implemented")
}
func (UnimplementedBooksUsersServiceServer) ConfirmRent(context.Context, *ConfirmRentRequest) (*ConfirmRentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmRent not implemented")
}
func (UnimplementedBooksUsersServiceServer) ConfirmReturn(context.Context, *ConfirmReturnRequest) (*ConfirmReturnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmReturn not implemented")
}
func (UnimplementedBooksUsersServiceServer) FindBook(context.Context, *FindBookRequest) (*FindBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBook not implemented")
}
func (UnimplementedBooksUsersServiceServer) GetDebt(context.Context, *GetDebtRequest) (*GetDebtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDebt not implemented")
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

func _BooksUsersService_ConfirmRent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmRentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksUsersServiceServer).ConfirmRent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BooksUsersService_ConfirmRent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksUsersServiceServer).ConfirmRent(ctx, req.(*ConfirmRentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksUsersService_ConfirmReturn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmReturnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksUsersServiceServer).ConfirmReturn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BooksUsersService_ConfirmReturn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksUsersServiceServer).ConfirmReturn(ctx, req.(*ConfirmReturnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksUsersService_FindBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksUsersServiceServer).FindBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BooksUsersService_FindBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksUsersServiceServer).FindBook(ctx, req.(*FindBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksUsersService_GetDebt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDebtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksUsersServiceServer).GetDebt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BooksUsersService_GetDebt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksUsersServiceServer).GetDebt(ctx, req.(*GetDebtRequest))
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
		{
			MethodName: "ConfirmRent",
			Handler:    _BooksUsersService_ConfirmRent_Handler,
		},
		{
			MethodName: "ConfirmReturn",
			Handler:    _BooksUsersService_ConfirmReturn_Handler,
		},
		{
			MethodName: "FindBook",
			Handler:    _BooksUsersService_FindBook_Handler,
		},
		{
			MethodName: "GetDebt",
			Handler:    _BooksUsersService_GetDebt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/books_users.proto",
}
