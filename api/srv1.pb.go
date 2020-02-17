// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv1.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type GithubRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GithubRequest) Reset()         { *m = GithubRequest{} }
func (m *GithubRequest) String() string { return proto.CompactTextString(m) }
func (*GithubRequest) ProtoMessage()    {}
func (*GithubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6330b32150a8abd, []int{0}
}

func (m *GithubRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GithubRequest.Unmarshal(m, b)
}
func (m *GithubRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GithubRequest.Marshal(b, m, deterministic)
}
func (m *GithubRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GithubRequest.Merge(m, src)
}
func (m *GithubRequest) XXX_Size() int {
	return xxx_messageInfo_GithubRequest.Size(m)
}
func (m *GithubRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GithubRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GithubRequest proto.InternalMessageInfo

func (m *GithubRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GithubResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GithubResponse) Reset()         { *m = GithubResponse{} }
func (m *GithubResponse) String() string { return proto.CompactTextString(m) }
func (*GithubResponse) ProtoMessage()    {}
func (*GithubResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6330b32150a8abd, []int{1}
}

func (m *GithubResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GithubResponse.Unmarshal(m, b)
}
func (m *GithubResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GithubResponse.Marshal(b, m, deterministic)
}
func (m *GithubResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GithubResponse.Merge(m, src)
}
func (m *GithubResponse) XXX_Size() int {
	return xxx_messageInfo_GithubResponse.Size(m)
}
func (m *GithubResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GithubResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GithubResponse proto.InternalMessageInfo

func (m *GithubResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GithubRequest)(nil), "api.GithubRequest")
	proto.RegisterType((*GithubResponse)(nil), "api.GithubResponse")
}

func init() { proto.RegisterFile("srv1.proto", fileDescriptor_f6330b32150a8abd) }

var fileDescriptor_f6330b32150a8abd = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2e, 0x2a, 0x33,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92, 0x49, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9,
	0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x51, 0xd2, 0xe4, 0xe2, 0x75, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0x0a,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe0, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
	0x4f, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x95, 0xb4, 0xb8, 0xf8, 0x60, 0x4a,
	0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x71, 0xab, 0x35, 0x8a, 0xe3, 0xe2, 0x84, 0xaa, 0x0d, 0x70,
	0x16, 0x0a, 0xe4, 0x12, 0x74, 0x4f, 0x2d, 0x41, 0xd3, 0x2b, 0xa4, 0x97, 0x58, 0x90, 0xa9, 0x87,
	0x62, 0xb7, 0x94, 0x30, 0x8a, 0x18, 0x44, 0xa1, 0x92, 0x70, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0x78,
	0x95, 0x38, 0xf4, 0xa1, 0x86, 0x5b, 0x31, 0x6a, 0x25, 0xb1, 0x81, 0x5d, 0x6f, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xdb, 0x49, 0xcf, 0x4f, 0xee, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GithubRPCClient is the client API for GithubRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GithubRPCClient interface {
	GetGithubResponse(ctx context.Context, in *GithubRequest, opts ...grpc.CallOption) (*GithubResponse, error)
}

type githubRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewGithubRPCClient(cc grpc.ClientConnInterface) GithubRPCClient {
	return &githubRPCClient{cc}
}

func (c *githubRPCClient) GetGithubResponse(ctx context.Context, in *GithubRequest, opts ...grpc.CallOption) (*GithubResponse, error) {
	out := new(GithubResponse)
	err := c.cc.Invoke(ctx, "/api.GithubRPC/GetGithubResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GithubRPCServer is the server API for GithubRPC service.
type GithubRPCServer interface {
	GetGithubResponse(context.Context, *GithubRequest) (*GithubResponse, error)
}

// UnimplementedGithubRPCServer can be embedded to have forward compatible implementations.
type UnimplementedGithubRPCServer struct {
}

func (*UnimplementedGithubRPCServer) GetGithubResponse(ctx context.Context, req *GithubRequest) (*GithubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGithubResponse not implemented")
}

func RegisterGithubRPCServer(s *grpc.Server, srv GithubRPCServer) {
	s.RegisterService(&_GithubRPC_serviceDesc, srv)
}

func _GithubRPC_GetGithubResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GithubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubRPCServer).GetGithubResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GithubRPC/GetGithubResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubRPCServer).GetGithubResponse(ctx, req.(*GithubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GithubRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.GithubRPC",
	HandlerType: (*GithubRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGithubResponse",
			Handler:    _GithubRPC_GetGithubResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "srv1.proto",
}