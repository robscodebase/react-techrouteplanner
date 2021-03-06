// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/technicianRoutePlanner.proto

/*
Package technicianRoutePlanner is a generated protocol buffer package.

It is generated from these files:
	proto/technicianRoutePlanner.proto

It has these top-level messages:
	RoutePlannerRequest
	RoutePlannerReply
*/
package technicianRoutePlanner

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RoutePlannerRequest struct {
	RouteName string `protobuf:"bytes,1,opt,name=routeName" json:"routeName,omitempty"`
}

func (m *RoutePlannerRequest) Reset()                    { *m = RoutePlannerRequest{} }
func (m *RoutePlannerRequest) String() string            { return proto.CompactTextString(m) }
func (*RoutePlannerRequest) ProtoMessage()               {}
func (*RoutePlannerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RoutePlannerRequest) GetRouteName() string {
	if m != nil {
		return m.RouteName
	}
	return ""
}

type RoutePlannerReply struct {
	Coordinates []float64 `protobuf:"fixed64,1,rep,packed,name=coordinates" json:"coordinates,omitempty"`
	SizeOfArray float64   `protobuf:"fixed64,2,opt,name=sizeOfArray" json:"sizeOfArray,omitempty"`
}

func (m *RoutePlannerReply) Reset()                    { *m = RoutePlannerReply{} }
func (m *RoutePlannerReply) String() string            { return proto.CompactTextString(m) }
func (*RoutePlannerReply) ProtoMessage()               {}
func (*RoutePlannerReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RoutePlannerReply) GetCoordinates() []float64 {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *RoutePlannerReply) GetSizeOfArray() float64 {
	if m != nil {
		return m.SizeOfArray
	}
	return 0
}

func init() {
	proto.RegisterType((*RoutePlannerRequest)(nil), "technicianRoutePlanner.RoutePlannerRequest")
	proto.RegisterType((*RoutePlannerReply)(nil), "technicianRoutePlanner.RoutePlannerReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TechnicianRoutePlanner service

type TechnicianRoutePlannerClient interface {
	RoutePlanner(ctx context.Context, in *RoutePlannerRequest, opts ...grpc.CallOption) (*RoutePlannerReply, error)
}

type technicianRoutePlannerClient struct {
	cc *grpc.ClientConn
}

func NewTechnicianRoutePlannerClient(cc *grpc.ClientConn) TechnicianRoutePlannerClient {
	return &technicianRoutePlannerClient{cc}
}

func (c *technicianRoutePlannerClient) RoutePlanner(ctx context.Context, in *RoutePlannerRequest, opts ...grpc.CallOption) (*RoutePlannerReply, error) {
	out := new(RoutePlannerReply)
	err := grpc.Invoke(ctx, "/technicianRoutePlanner.technicianRoutePlanner/RoutePlanner", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TechnicianRoutePlanner service

type TechnicianRoutePlannerServer interface {
	RoutePlanner(context.Context, *RoutePlannerRequest) (*RoutePlannerReply, error)
}

func RegisterTechnicianRoutePlannerServer(s *grpc.Server, srv TechnicianRoutePlannerServer) {
	s.RegisterService(&_TechnicianRoutePlanner_serviceDesc, srv)
}

func _TechnicianRoutePlanner_RoutePlanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoutePlannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TechnicianRoutePlannerServer).RoutePlanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/technicianRoutePlanner.technicianRoutePlanner/RoutePlanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TechnicianRoutePlannerServer).RoutePlanner(ctx, req.(*RoutePlannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TechnicianRoutePlanner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "technicianRoutePlanner.technicianRoutePlanner",
	HandlerType: (*TechnicianRoutePlannerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RoutePlanner",
			Handler:    _TechnicianRoutePlanner_RoutePlanner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/technicianRoutePlanner.proto",
}

func init() { proto.RegisterFile("proto/technicianRoutePlanner.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x49, 0x4d, 0xce, 0xc8, 0xcb, 0x4c, 0xce, 0x4c, 0xcc, 0x0b, 0xca, 0x2f, 0x2d,
	0x49, 0x0d, 0xc8, 0x49, 0xcc, 0xcb, 0x4b, 0x2d, 0xd2, 0x03, 0x4b, 0x0a, 0x89, 0x61, 0x97, 0x55,
	0x32, 0xe6, 0x12, 0x46, 0xe6, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xc9, 0x70, 0x71,
	0x16, 0x81, 0x84, 0xfd, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x10, 0x02,
	0x4a, 0xe1, 0x5c, 0x82, 0xa8, 0x9a, 0x0a, 0x72, 0x2a, 0x85, 0x14, 0xb8, 0xb8, 0x93, 0xf3, 0xf3,
	0x8b, 0x52, 0x32, 0xf3, 0x12, 0x4b, 0x52, 0x8b, 0x25, 0x18, 0x15, 0x98, 0x35, 0x18, 0x83, 0x90,
	0x85, 0x40, 0x2a, 0x8a, 0x33, 0xab, 0x52, 0xfd, 0xd3, 0x1c, 0x8b, 0x8a, 0x12, 0x2b, 0x25, 0x98,
	0x14, 0x18, 0x41, 0x2a, 0x90, 0x84, 0x8c, 0x9a, 0x18, 0xb9, 0x70, 0x38, 0x54, 0x28, 0x83, 0x8b,
	0x07, 0x85, 0xaf, 0xad, 0x87, 0xc3, 0xbf, 0x58, 0xbc, 0x23, 0xa5, 0x49, 0x9c, 0xe2, 0x82, 0x9c,
	0x4a, 0x25, 0x86, 0x24, 0x36, 0x70, 0x88, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x13, 0xcf,
	0xc6, 0x41, 0x57, 0x01, 0x00, 0x00,
}
