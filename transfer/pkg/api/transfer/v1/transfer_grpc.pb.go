// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: transfer.proto

package transfer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TransferClient is the client API for Transfer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferClient interface {
	CreateMoneyTransfer(ctx context.Context, in *CreateMoneyTransferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type transferClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferClient(cc grpc.ClientConnInterface) TransferClient {
	return &transferClient{cc}
}

func (c *transferClient) CreateMoneyTransfer(ctx context.Context, in *CreateMoneyTransferRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/github.com.volnistii11.accounting_service.transfer.Transfer/CreateMoneyTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferServer is the server API for Transfer service.
// All implementations must embed UnimplementedTransferServer
// for forward compatibility
type TransferServer interface {
	CreateMoneyTransfer(context.Context, *CreateMoneyTransferRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTransferServer()
}

// UnimplementedTransferServer must be embedded to have forward compatible implementations.
type UnimplementedTransferServer struct {
}

func (UnimplementedTransferServer) CreateMoneyTransfer(context.Context, *CreateMoneyTransferRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMoneyTransfer not implemented")
}
func (UnimplementedTransferServer) mustEmbedUnimplementedTransferServer() {}

// UnsafeTransferServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferServer will
// result in compilation errors.
type UnsafeTransferServer interface {
	mustEmbedUnimplementedTransferServer()
}

func RegisterTransferServer(s grpc.ServiceRegistrar, srv TransferServer) {
	s.RegisterService(&Transfer_ServiceDesc, srv)
}

func _Transfer_CreateMoneyTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMoneyTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServer).CreateMoneyTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.volnistii11.accounting_service.transfer.Transfer/CreateMoneyTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServer).CreateMoneyTransfer(ctx, req.(*CreateMoneyTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Transfer_ServiceDesc is the grpc.ServiceDesc for Transfer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transfer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.volnistii11.accounting_service.transfer.Transfer",
	HandlerType: (*TransferServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMoneyTransfer",
			Handler:    _Transfer_CreateMoneyTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfer.proto",
}
