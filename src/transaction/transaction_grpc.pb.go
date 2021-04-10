// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package transaction

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

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIClient interface {
	// Transaction rpcs
	BatchTransaction(ctx context.Context, in *BatchTransactionRequest, opts ...grpc.CallOption) (*TransactionInfo, error)
	StartTransaction(ctx context.Context, in *StartTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
	InspectTransaction(ctx context.Context, in *InspectTransactionRequest, opts ...grpc.CallOption) (*TransactionInfo, error)
	DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListTransaction(ctx context.Context, in *ListTransactionRequest, opts ...grpc.CallOption) (*TransactionInfos, error)
	FinishTransaction(ctx context.Context, in *FinishTransactionRequest, opts ...grpc.CallOption) (*TransactionInfo, error)
	DeleteAll(ctx context.Context, in *DeleteAllRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) BatchTransaction(ctx context.Context, in *BatchTransactionRequest, opts ...grpc.CallOption) (*TransactionInfo, error) {
	out := new(TransactionInfo)
	err := c.cc.Invoke(ctx, "/transaction.API/BatchTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) StartTransaction(ctx context.Context, in *StartTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/transaction.API/StartTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) InspectTransaction(ctx context.Context, in *InspectTransactionRequest, opts ...grpc.CallOption) (*TransactionInfo, error) {
	out := new(TransactionInfo)
	err := c.cc.Invoke(ctx, "/transaction.API/InspectTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/transaction.API/DeleteTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListTransaction(ctx context.Context, in *ListTransactionRequest, opts ...grpc.CallOption) (*TransactionInfos, error) {
	out := new(TransactionInfos)
	err := c.cc.Invoke(ctx, "/transaction.API/ListTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) FinishTransaction(ctx context.Context, in *FinishTransactionRequest, opts ...grpc.CallOption) (*TransactionInfo, error) {
	out := new(TransactionInfo)
	err := c.cc.Invoke(ctx, "/transaction.API/FinishTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteAll(ctx context.Context, in *DeleteAllRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/transaction.API/DeleteAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
// All implementations should embed UnimplementedAPIServer
// for forward compatibility
type APIServer interface {
	// Transaction rpcs
	BatchTransaction(context.Context, *BatchTransactionRequest) (*TransactionInfo, error)
	StartTransaction(context.Context, *StartTransactionRequest) (*Transaction, error)
	InspectTransaction(context.Context, *InspectTransactionRequest) (*TransactionInfo, error)
	DeleteTransaction(context.Context, *DeleteTransactionRequest) (*emptypb.Empty, error)
	ListTransaction(context.Context, *ListTransactionRequest) (*TransactionInfos, error)
	FinishTransaction(context.Context, *FinishTransactionRequest) (*TransactionInfo, error)
	DeleteAll(context.Context, *DeleteAllRequest) (*emptypb.Empty, error)
}

// UnimplementedAPIServer should be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (UnimplementedAPIServer) BatchTransaction(context.Context, *BatchTransactionRequest) (*TransactionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchTransaction not implemented")
}
func (UnimplementedAPIServer) StartTransaction(context.Context, *StartTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartTransaction not implemented")
}
func (UnimplementedAPIServer) InspectTransaction(context.Context, *InspectTransactionRequest) (*TransactionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InspectTransaction not implemented")
}
func (UnimplementedAPIServer) DeleteTransaction(context.Context, *DeleteTransactionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransaction not implemented")
}
func (UnimplementedAPIServer) ListTransaction(context.Context, *ListTransactionRequest) (*TransactionInfos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransaction not implemented")
}
func (UnimplementedAPIServer) FinishTransaction(context.Context, *FinishTransactionRequest) (*TransactionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishTransaction not implemented")
}
func (UnimplementedAPIServer) DeleteAll(context.Context, *DeleteAllRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAll not implemented")
}

// UnsafeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServer will
// result in compilation errors.
type UnsafeAPIServer interface {
	mustEmbedUnimplementedAPIServer()
}

func RegisterAPIServer(s grpc.ServiceRegistrar, srv APIServer) {
	s.RegisterService(&API_ServiceDesc, srv)
}

func _API_BatchTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).BatchTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.API/BatchTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).BatchTransaction(ctx, req.(*BatchTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_StartTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).StartTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.API/StartTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).StartTransaction(ctx, req.(*StartTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_InspectTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InspectTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).InspectTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.API/InspectTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).InspectTransaction(ctx, req.(*InspectTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_DeleteTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DeleteTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.API/DeleteTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DeleteTransaction(ctx, req.(*DeleteTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.API/ListTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListTransaction(ctx, req.(*ListTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_FinishTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).FinishTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.API/FinishTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).FinishTransaction(ctx, req.(*FinishTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_DeleteAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DeleteAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.API/DeleteAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DeleteAll(ctx, req.(*DeleteAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// API_ServiceDesc is the grpc.ServiceDesc for API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BatchTransaction",
			Handler:    _API_BatchTransaction_Handler,
		},
		{
			MethodName: "StartTransaction",
			Handler:    _API_StartTransaction_Handler,
		},
		{
			MethodName: "InspectTransaction",
			Handler:    _API_InspectTransaction_Handler,
		},
		{
			MethodName: "DeleteTransaction",
			Handler:    _API_DeleteTransaction_Handler,
		},
		{
			MethodName: "ListTransaction",
			Handler:    _API_ListTransaction_Handler,
		},
		{
			MethodName: "FinishTransaction",
			Handler:    _API_FinishTransaction_Handler,
		},
		{
			MethodName: "DeleteAll",
			Handler:    _API_DeleteAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction/transaction.proto",
}
