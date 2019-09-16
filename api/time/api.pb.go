// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The response message containing the ntp server
type TimeRequest struct {
	Server               string   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeRequest) Reset()         { *m = TimeRequest{} }
func (m *TimeRequest) String() string { return proto.CompactTextString(m) }
func (*TimeRequest) ProtoMessage()    {}
func (*TimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *TimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRequest.Unmarshal(m, b)
}

func (m *TimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRequest.Marshal(b, m, deterministic)
}

func (m *TimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRequest.Merge(m, src)
}

func (m *TimeRequest) XXX_Size() int {
	return xxx_messageInfo_TimeRequest.Size(m)
}

func (m *TimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRequest proto.InternalMessageInfo

func (m *TimeRequest) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

// The response message containing the ntp server, time, and offset
type TimeReply struct {
	Server               string               `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Localtime            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=localtime,proto3" json:"localtime,omitempty"`
	Remotetime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=remotetime,proto3" json:"remotetime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeReply) Reset()         { *m = TimeReply{} }
func (m *TimeReply) String() string { return proto.CompactTextString(m) }
func (*TimeReply) ProtoMessage()    {}
func (*TimeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *TimeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeReply.Unmarshal(m, b)
}

func (m *TimeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeReply.Marshal(b, m, deterministic)
}

func (m *TimeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeReply.Merge(m, src)
}

func (m *TimeReply) XXX_Size() int {
	return xxx_messageInfo_TimeReply.Size(m)
}

func (m *TimeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeReply.DiscardUnknown(m)
}

var xxx_messageInfo_TimeReply proto.InternalMessageInfo

func (m *TimeReply) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *TimeReply) GetLocaltime() *timestamp.Timestamp {
	if m != nil {
		return m.Localtime
	}
	return nil
}

func (m *TimeReply) GetRemotetime() *timestamp.Timestamp {
	if m != nil {
		return m.Remotetime
	}
	return nil
}

func init() {
	proto.RegisterType((*TimeRequest)(nil), "proto.TimeRequest")
	proto.RegisterType((*TimeReply)(nil), "proto.TimeReply")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0xd2, 0xe9, 0xf9, 0xf9, 0xe9, 0x39,
	0xa9, 0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x6a, 0x6e, 0x41, 0x49, 0x25, 0x44, 0x8d, 0x94,
	0x3c, 0xba, 0x64, 0x49, 0x66, 0x6e, 0x6a, 0x71, 0x49, 0x62, 0x6e, 0x01, 0x44, 0x81, 0x92, 0x2a,
	0x17, 0x77, 0x48, 0x66, 0x6e, 0x6a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x18, 0x17,
	0x5b, 0x71, 0x6a, 0x51, 0x59, 0x6a, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x94, 0xa7,
	0x34, 0x93, 0x91, 0x8b, 0x13, 0xa2, 0xae, 0x20, 0xa7, 0x12, 0x97, 0x2a, 0x21, 0x0b, 0x2e, 0xce,
	0x9c, 0xfc, 0xe4, 0xc4, 0x1c, 0x90, 0x25, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x52, 0x7a,
	0x10, 0x17, 0xe8, 0xc1, 0x5c, 0xa0, 0x17, 0x02, 0x73, 0x41, 0x10, 0x42, 0xb1, 0x90, 0x15, 0x17,
	0x57, 0x51, 0x6a, 0x6e, 0x7e, 0x49, 0x2a, 0x58, 0x2b, 0x33, 0x41, 0xad, 0x48, 0xaa, 0x8d, 0xb2,
	0xb9, 0x58, 0xfc, 0x4a, 0x0a, 0x52, 0x84, 0x0c, 0xb8, 0x58, 0x40, 0x0a, 0x84, 0xc4, 0x30, 0xf4,
	0xb9, 0x82, 0x42, 0x44, 0x4a, 0x00, 0x22, 0xa0, 0x87, 0xf0, 0x87, 0x21, 0xc4, 0x53, 0xce, 0x19,
	0xa9, 0xc9, 0xd9, 0x42, 0x42, 0x28, 0xd2, 0xe0, 0xe0, 0xc0, 0xd4, 0xe2, 0x24, 0xc7, 0xc5, 0x99,
	0x9c, 0x9f, 0x0b, 0x11, 0x76, 0xe2, 0x70, 0x2c, 0xc8, 0x0c, 0x00, 0xb1, 0x02, 0x18, 0xa3, 0x20,
	0xb1, 0x91, 0xc4, 0x06, 0xa6, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x98, 0xd1, 0xa5, 0x14,
	0xa8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NtpdClient is the client API for Ntpd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NtpdClient interface {
	Time(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TimeReply, error)
	TimeCheck(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*TimeReply, error)
}

type ntpdClient struct {
	cc *grpc.ClientConn
}

func NewNtpdClient(cc *grpc.ClientConn) NtpdClient {
	return &ntpdClient{cc}
}

func (c *ntpdClient) Time(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TimeReply, error) {
	out := new(TimeReply)
	err := c.cc.Invoke(ctx, "/proto.Ntpd/Time", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ntpdClient) TimeCheck(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*TimeReply, error) {
	out := new(TimeReply)
	err := c.cc.Invoke(ctx, "/proto.Ntpd/TimeCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NtpdServer is the server API for Ntpd service.
type NtpdServer interface {
	Time(context.Context, *empty.Empty) (*TimeReply, error)
	TimeCheck(context.Context, *TimeRequest) (*TimeReply, error)
}

func RegisterNtpdServer(s *grpc.Server, srv NtpdServer) {
	s.RegisterService(&_Ntpd_serviceDesc, srv)
}

func _Ntpd_Time_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NtpdServer).Time(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Ntpd/Time",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NtpdServer).Time(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ntpd_TimeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NtpdServer).TimeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Ntpd/TimeCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NtpdServer).TimeCheck(ctx, req.(*TimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ntpd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Ntpd",
	HandlerType: (*NtpdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Time",
			Handler:    _Ntpd_Time_Handler,
		},
		{
			MethodName: "TimeCheck",
			Handler:    _Ntpd_TimeCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
