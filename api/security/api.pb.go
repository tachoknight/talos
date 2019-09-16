// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// The request message containing the process name.
type CertificateRequest struct {
	Csr                  []byte   `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateRequest) Reset()         { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()    {}
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *CertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequest.Unmarshal(m, b)
}

func (m *CertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequest.Marshal(b, m, deterministic)
}

func (m *CertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequest.Merge(m, src)
}

func (m *CertificateRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateRequest.Size(m)
}

func (m *CertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequest proto.InternalMessageInfo

func (m *CertificateRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

// The response message containing the requested logs.
type CertificateResponse struct {
	Ca                   []byte   `protobuf:"bytes,1,opt,name=ca,proto3" json:"ca,omitempty"`
	Crt                  []byte   `protobuf:"bytes,2,opt,name=crt,proto3" json:"crt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateResponse.Unmarshal(m, b)
}

func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
}

func (m *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(m, src)
}

func (m *CertificateResponse) XXX_Size() int {
	return xxx_messageInfo_CertificateResponse.Size(m)
}

func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetCa() []byte {
	if m != nil {
		return m.Ca
	}
	return nil
}

func (m *CertificateResponse) GetCrt() []byte {
	if m != nil {
		return m.Crt
	}
	return nil
}

// The request message for reading a file on disk.
type ReadFileRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadFileRequest) Reset()         { *m = ReadFileRequest{} }
func (m *ReadFileRequest) String() string { return proto.CompactTextString(m) }
func (*ReadFileRequest) ProtoMessage()    {}
func (*ReadFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *ReadFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadFileRequest.Unmarshal(m, b)
}

func (m *ReadFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadFileRequest.Marshal(b, m, deterministic)
}

func (m *ReadFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadFileRequest.Merge(m, src)
}

func (m *ReadFileRequest) XXX_Size() int {
	return xxx_messageInfo_ReadFileRequest.Size(m)
}

func (m *ReadFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadFileRequest proto.InternalMessageInfo

func (m *ReadFileRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// The response message for reading a file on disk.
type ReadFileResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadFileResponse) Reset()         { *m = ReadFileResponse{} }
func (m *ReadFileResponse) String() string { return proto.CompactTextString(m) }
func (*ReadFileResponse) ProtoMessage()    {}
func (*ReadFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *ReadFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadFileResponse.Unmarshal(m, b)
}

func (m *ReadFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadFileResponse.Marshal(b, m, deterministic)
}

func (m *ReadFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadFileResponse.Merge(m, src)
}

func (m *ReadFileResponse) XXX_Size() int {
	return xxx_messageInfo_ReadFileResponse.Size(m)
}

func (m *ReadFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadFileResponse proto.InternalMessageInfo

func (m *ReadFileResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// The request message containing the process name.
type WriteFileRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Perm                 int32    `protobuf:"varint,3,opt,name=perm,proto3" json:"perm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteFileRequest) Reset()         { *m = WriteFileRequest{} }
func (m *WriteFileRequest) String() string { return proto.CompactTextString(m) }
func (*WriteFileRequest) ProtoMessage()    {}
func (*WriteFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *WriteFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteFileRequest.Unmarshal(m, b)
}

func (m *WriteFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteFileRequest.Marshal(b, m, deterministic)
}

func (m *WriteFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteFileRequest.Merge(m, src)
}

func (m *WriteFileRequest) XXX_Size() int {
	return xxx_messageInfo_WriteFileRequest.Size(m)
}

func (m *WriteFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteFileRequest proto.InternalMessageInfo

func (m *WriteFileRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *WriteFileRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *WriteFileRequest) GetPerm() int32 {
	if m != nil {
		return m.Perm
	}
	return 0
}

// The response message containing the requested logs.
type WriteFileResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteFileResponse) Reset()         { *m = WriteFileResponse{} }
func (m *WriteFileResponse) String() string { return proto.CompactTextString(m) }
func (*WriteFileResponse) ProtoMessage()    {}
func (*WriteFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *WriteFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteFileResponse.Unmarshal(m, b)
}

func (m *WriteFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteFileResponse.Marshal(b, m, deterministic)
}

func (m *WriteFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteFileResponse.Merge(m, src)
}

func (m *WriteFileResponse) XXX_Size() int {
	return xxx_messageInfo_WriteFileResponse.Size(m)
}

func (m *WriteFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteFileResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CertificateRequest)(nil), "proto.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "proto.CertificateResponse")
	proto.RegisterType((*ReadFileRequest)(nil), "proto.ReadFileRequest")
	proto.RegisterType((*ReadFileResponse)(nil), "proto.ReadFileResponse")
	proto.RegisterType((*WriteFileRequest)(nil), "proto.WriteFileRequest")
	proto.RegisterType((*WriteFileResponse)(nil), "proto.WriteFileResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0xfa, 0x87, 0x66, 0x14, 0xad, 0x53, 0xd0, 0x98, 0x83, 0x94, 0x05, 0x4b, 0x4f,
	0x3d, 0xe8, 0xc1, 0x83, 0x20, 0x58, 0xc5, 0xa3, 0x94, 0x45, 0x10, 0xbc, 0xad, 0x9b, 0x15, 0x17,
	0xac, 0x59, 0x77, 0xa7, 0x9f, 0xd2, 0x2f, 0x25, 0xdd, 0xdd, 0xa4, 0xb5, 0x11, 0x3c, 0xcd, 0x23,
	0xf3, 0x7e, 0x2f, 0x99, 0x47, 0x20, 0x13, 0x46, 0xcf, 0x8c, 0xad, 0xa8, 0xc2, 0x9e, 0x1f, 0x6c,
	0x02, 0x78, 0xa7, 0x2c, 0xe9, 0x37, 0x2d, 0x05, 0x29, 0xae, 0xbe, 0x56, 0xca, 0x11, 0x0e, 0xa1,
	0x23, 0x9d, 0xcd, 0x93, 0x71, 0x32, 0xdd, 0xe7, 0x6b, 0xc9, 0xae, 0x60, 0xf4, 0xcb, 0xe7, 0x4c,
	0xf5, 0xe9, 0x14, 0x1e, 0x40, 0x2a, 0x45, 0xf4, 0xa5, 0x52, 0x78, 0xd0, 0x52, 0x9e, 0x46, 0xd0,
	0x12, 0x3b, 0x87, 0x43, 0xae, 0x44, 0xf9, 0xa0, 0x3f, 0x9a, 0x74, 0x84, 0xae, 0x11, 0xf4, 0xee,
	0xb1, 0x8c, 0x7b, 0xcd, 0x26, 0x30, 0xdc, 0xd8, 0x62, 0x38, 0x42, 0xb7, 0x14, 0x54, 0xc7, 0x7b,
	0xcd, 0x1e, 0x61, 0xf8, 0x6c, 0x35, 0xa9, 0x7f, 0xf2, 0x1a, 0x36, 0xdd, 0xb0, 0xde, 0xa7, 0xec,
	0x32, 0xef, 0x8c, 0x93, 0x69, 0x8f, 0x7b, 0xcd, 0x46, 0x70, 0xb4, 0x95, 0x17, 0x5e, 0x7c, 0xf1,
	0x9d, 0x40, 0xff, 0xc9, 0xae, 0x1c, 0x95, 0x78, 0x0f, 0x7b, 0x5b, 0x77, 0xe3, 0x69, 0x68, 0x6f,
	0xd6, 0xee, 0xac, 0x28, 0xfe, 0x5a, 0xc5, 0x4b, 0xae, 0x61, 0x50, 0x5f, 0x87, 0xc7, 0xd1, 0xb7,
	0xd3, 0x4a, 0x71, 0xd2, 0x7a, 0x1e, 0xe1, 0x1b, 0xc8, 0x9a, 0x4f, 0xc4, 0xda, 0xb5, 0x5b, 0x42,
	0x91, 0xb7, 0x17, 0x81, 0x9f, 0x9f, 0x41, 0x26, 0xab, 0x65, 0x58, 0xcf, 0x07, 0xb7, 0x46, 0x2f,
	0xd6, 0x6a, 0x91, 0xbc, 0x84, 0x5f, 0xe0, 0xb5, 0xef, 0xc7, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x59, 0x5a, 0x3f, 0x46, 0x1d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TrustdClient is the client API for Trustd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrustdClient interface {
	Certificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	ReadFile(ctx context.Context, in *ReadFileRequest, opts ...grpc.CallOption) (*ReadFileResponse, error)
	WriteFile(ctx context.Context, in *WriteFileRequest, opts ...grpc.CallOption) (*WriteFileResponse, error)
}

type trustdClient struct {
	cc *grpc.ClientConn
}

func NewTrustdClient(cc *grpc.ClientConn) TrustdClient {
	return &trustdClient{cc}
}

func (c *trustdClient) Certificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/proto.Trustd/Certificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustdClient) ReadFile(ctx context.Context, in *ReadFileRequest, opts ...grpc.CallOption) (*ReadFileResponse, error) {
	out := new(ReadFileResponse)
	err := c.cc.Invoke(ctx, "/proto.Trustd/ReadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustdClient) WriteFile(ctx context.Context, in *WriteFileRequest, opts ...grpc.CallOption) (*WriteFileResponse, error) {
	out := new(WriteFileResponse)
	err := c.cc.Invoke(ctx, "/proto.Trustd/WriteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrustdServer is the server API for Trustd service.
type TrustdServer interface {
	Certificate(context.Context, *CertificateRequest) (*CertificateResponse, error)
	ReadFile(context.Context, *ReadFileRequest) (*ReadFileResponse, error)
	WriteFile(context.Context, *WriteFileRequest) (*WriteFileResponse, error)
}

func RegisterTrustdServer(s *grpc.Server, srv TrustdServer) {
	s.RegisterService(&_Trustd_serviceDesc, srv)
}

func _Trustd_Certificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustdServer).Certificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Trustd/Certificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustdServer).Certificate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trustd_ReadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustdServer).ReadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Trustd/ReadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustdServer).ReadFile(ctx, req.(*ReadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trustd_WriteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustdServer).WriteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Trustd/WriteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustdServer).WriteFile(ctx, req.(*WriteFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Trustd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Trustd",
	HandlerType: (*TrustdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Certificate",
			Handler:    _Trustd_Certificate_Handler,
		},
		{
			MethodName: "ReadFile",
			Handler:    _Trustd_ReadFile_Handler,
		},
		{
			MethodName: "WriteFile",
			Handler:    _Trustd_WriteFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
