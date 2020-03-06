// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/operating_system_version_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for
// [OperatingSystemVersionConstantService.GetOperatingSystemVersionConstant][google.ads.googleads.v1.services.OperatingSystemVersionConstantService.GetOperatingSystemVersionConstant].
type GetOperatingSystemVersionConstantRequest struct {
	// Resource name of the OS version to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperatingSystemVersionConstantRequest) Reset() {
	*m = GetOperatingSystemVersionConstantRequest{}
}
func (m *GetOperatingSystemVersionConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperatingSystemVersionConstantRequest) ProtoMessage()    {}
func (*GetOperatingSystemVersionConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a0893223528a69, []int{0}
}

func (m *GetOperatingSystemVersionConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Unmarshal(m, b)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Merge(m, src)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Size(m)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperatingSystemVersionConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperatingSystemVersionConstantRequest proto.InternalMessageInfo

func (m *GetOperatingSystemVersionConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetOperatingSystemVersionConstantRequest)(nil), "google.ads.googleads.v1.services.GetOperatingSystemVersionConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/operating_system_version_constant_service.proto", fileDescriptor_97a0893223528a69)
}

var fileDescriptor_97a0893223528a69 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x6a, 0xdb, 0x40,
	0x1c, 0xc7, 0x91, 0x0a, 0x85, 0x8a, 0x76, 0xd1, 0x52, 0xa3, 0x76, 0x70, 0x5d, 0x17, 0x8c, 0x87,
	0x3b, 0xd4, 0x52, 0x0a, 0x57, 0x3c, 0xc8, 0x19, 0x9c, 0x64, 0x88, 0x8d, 0x0d, 0x1a, 0x82, 0x40,
	0x5c, 0xa4, 0x43, 0x08, 0xac, 0x3b, 0x45, 0xbf, 0xb3, 0x20, 0x84, 0x2c, 0x79, 0x83, 0x90, 0x37,
	0xc8, 0x98, 0x27, 0xc8, 0x33, 0x78, 0xcd, 0x2b, 0x64, 0xca, 0x96, 0x37, 0x08, 0xf2, 0xe9, 0xe4,
	0x64, 0xf0, 0x9f, 0xed, 0xcb, 0xdd, 0xf7, 0xbe, 0x9f, 0xdf, 0x9f, 0xb3, 0x26, 0x89, 0x10, 0xc9,
	0x9c, 0x61, 0x1a, 0x03, 0x56, 0xb2, 0x52, 0xa5, 0x8b, 0x81, 0x15, 0x65, 0x1a, 0x31, 0xc0, 0x22,
	0x67, 0x05, 0x95, 0x29, 0x4f, 0x42, 0xb8, 0x00, 0xc9, 0xb2, 0xb0, 0x64, 0x05, 0xa4, 0x82, 0x87,
	0x91, 0xe0, 0x20, 0x29, 0x97, 0x61, 0x6d, 0x45, 0x79, 0x21, 0xa4, 0xb0, 0xdb, 0x2a, 0x06, 0xd1,
	0x18, 0x50, 0x93, 0x88, 0x4a, 0x17, 0xe9, 0x44, 0xe7, 0x68, 0x13, 0xb3, 0x60, 0x20, 0x16, 0xc5,
	0x5e, 0x50, 0x05, 0x73, 0xbe, 0xeb, 0xa8, 0x3c, 0xc5, 0x94, 0x73, 0x21, 0xa9, 0x4c, 0x05, 0x87,
	0xfa, 0xf6, 0xeb, 0x9b, 0xdb, 0x68, 0x9e, 0x32, 0xfd, 0xac, 0x33, 0xb6, 0x7a, 0x23, 0x26, 0xc7,
	0x1a, 0x32, 0x5b, 0x31, 0x7c, 0x85, 0x38, 0xa8, 0x09, 0x53, 0x76, 0xbe, 0x60, 0x20, 0xed, 0x9f,
	0xd6, 0x17, 0x5d, 0x57, 0xc8, 0x69, 0xc6, 0x5a, 0x46, 0xdb, 0xe8, 0x7d, 0x9a, 0x7e, 0xd6, 0x87,
	0x27, 0x34, 0x63, 0xbf, 0x1f, 0x4c, 0xeb, 0xd7, 0xf6, 0xb8, 0x99, 0xea, 0xde, 0x7e, 0x31, 0xac,
	0x1f, 0x3b, 0xd9, 0xf6, 0x31, 0xda, 0x35, 0x45, 0xb4, 0x6f, 0x03, 0x8e, 0xb7, 0x31, 0xab, 0x99,
	0x37, 0xda, 0x9e, 0xd4, 0x19, 0x5c, 0x3f, 0x3e, 0xdd, 0x9a, 0xff, 0xec, 0xbf, 0xd5, 0x96, 0x2e,
	0xdf, 0x8d, 0x63, 0x20, 0xb6, 0x3e, 0x05, 0xdc, 0xbf, 0x72, 0xbe, 0x2d, 0xbd, 0xd6, 0x1a, 0x5c,
	0xab, 0x3c, 0x05, 0x14, 0x89, 0x6c, 0x78, 0x63, 0x5a, 0xdd, 0x48, 0x64, 0x3b, 0x1b, 0x1e, 0xf6,
	0xf7, 0x1a, 0xf0, 0xa4, 0x5a, 0xf0, 0xc4, 0x38, 0x3d, 0xac, 0xf3, 0x12, 0x31, 0xa7, 0x3c, 0x41,
	0xa2, 0x48, 0x70, 0xc2, 0xf8, 0x6a, 0xfd, 0x78, 0x5d, 0xc1, 0xe6, 0x7f, 0xff, 0x5f, 0x8b, 0x3b,
	0xf3, 0xc3, 0xc8, 0xf3, 0xee, 0xcd, 0xf6, 0x48, 0x05, 0x7a, 0x31, 0x20, 0x25, 0x2b, 0xe5, 0xbb,
	0xa8, 0x06, 0xc3, 0x52, 0x5b, 0x02, 0x2f, 0x86, 0xa0, 0xb1, 0x04, 0xbe, 0x1b, 0x68, 0xcb, 0xb3,
	0xd9, 0x55, 0xe7, 0x84, 0x78, 0x31, 0x10, 0xd2, 0x98, 0x08, 0xf1, 0x5d, 0x42, 0xb4, 0xed, 0xec,
	0xe3, 0xaa, 0xce, 0x3f, 0xaf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x52, 0x64, 0xed, 0xdf, 0x9e, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OperatingSystemVersionConstantServiceClient is the client API for OperatingSystemVersionConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperatingSystemVersionConstantServiceClient interface {
	// Returns the requested OS version constant in full detail.
	GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error)
}

type operatingSystemVersionConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatingSystemVersionConstantServiceClient(cc grpc.ClientConnInterface) OperatingSystemVersionConstantServiceClient {
	return &operatingSystemVersionConstantServiceClient{cc}
}

func (c *operatingSystemVersionConstantServiceClient) GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error) {
	out := new(resources.OperatingSystemVersionConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatingSystemVersionConstantServiceServer is the server API for OperatingSystemVersionConstantService service.
type OperatingSystemVersionConstantServiceServer interface {
	// Returns the requested OS version constant in full detail.
	GetOperatingSystemVersionConstant(context.Context, *GetOperatingSystemVersionConstantRequest) (*resources.OperatingSystemVersionConstant, error)
}

// UnimplementedOperatingSystemVersionConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOperatingSystemVersionConstantServiceServer struct {
}

func (*UnimplementedOperatingSystemVersionConstantServiceServer) GetOperatingSystemVersionConstant(ctx context.Context, req *GetOperatingSystemVersionConstantRequest) (*resources.OperatingSystemVersionConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatingSystemVersionConstant not implemented")
}

func RegisterOperatingSystemVersionConstantServiceServer(s *grpc.Server, srv OperatingSystemVersionConstantServiceServer) {
	s.RegisterService(&_OperatingSystemVersionConstantService_serviceDesc, srv)
}

func _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatingSystemVersionConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, req.(*GetOperatingSystemVersionConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OperatingSystemVersionConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.OperatingSystemVersionConstantService",
	HandlerType: (*OperatingSystemVersionConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOperatingSystemVersionConstant",
			Handler:    _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/operating_system_version_constant_service.proto",
}
