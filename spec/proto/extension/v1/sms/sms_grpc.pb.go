// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: sms.proto

package sms

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

// SmsServiceClient is the client API for SmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsServiceClient interface {
	// Send the SMS message.
	SendSmsWithTemplate(ctx context.Context, in *SendSmsWithTemplateRequest, opts ...grpc.CallOption) (*SendSmsWithTemplateResponse, error)
}

type smsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsServiceClient(cc grpc.ClientConnInterface) SmsServiceClient {
	return &smsServiceClient{cc}
}

func (c *smsServiceClient) SendSmsWithTemplate(ctx context.Context, in *SendSmsWithTemplateRequest, opts ...grpc.CallOption) (*SendSmsWithTemplateResponse, error) {
	out := new(SendSmsWithTemplateResponse)
	err := c.cc.Invoke(ctx, "/spec.proto.extension.v1.sms.SmsService/SendSmsWithTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServiceServer is the server API for SmsService service.
// All implementations should embed UnimplementedSmsServiceServer
// for forward compatibility
type SmsServiceServer interface {
	// Send the SMS message.
	SendSmsWithTemplate(context.Context, *SendSmsWithTemplateRequest) (*SendSmsWithTemplateResponse, error)
}

// UnimplementedSmsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSmsServiceServer struct {
}

func (UnimplementedSmsServiceServer) SendSmsWithTemplate(context.Context, *SendSmsWithTemplateRequest) (*SendSmsWithTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmsWithTemplate not implemented")
}

// UnsafeSmsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsServiceServer will
// result in compilation errors.
type UnsafeSmsServiceServer interface {
	mustEmbedUnimplementedSmsServiceServer()
}

func RegisterSmsServiceServer(s grpc.ServiceRegistrar, srv SmsServiceServer) {
	s.RegisterService(&SmsService_ServiceDesc, srv)
}

func _SmsService_SendSmsWithTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsWithTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServiceServer).SendSmsWithTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.proto.extension.v1.sms.SmsService/SendSmsWithTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServiceServer).SendSmsWithTemplate(ctx, req.(*SendSmsWithTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmsService_ServiceDesc is the grpc.ServiceDesc for SmsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spec.proto.extension.v1.sms.SmsService",
	HandlerType: (*SmsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSmsWithTemplate",
			Handler:    _SmsService_SendSmsWithTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms.proto",
}
