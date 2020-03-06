// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/location_view_service.proto

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

// Request message for [LocationViewService.GetLocationView][google.ads.googleads.v1.services.LocationViewService.GetLocationView].
type GetLocationViewRequest struct {
	// The resource name of the location view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLocationViewRequest) Reset()         { *m = GetLocationViewRequest{} }
func (m *GetLocationViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetLocationViewRequest) ProtoMessage()    {}
func (*GetLocationViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cee10bb181b86493, []int{0}
}

func (m *GetLocationViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLocationViewRequest.Unmarshal(m, b)
}
func (m *GetLocationViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLocationViewRequest.Marshal(b, m, deterministic)
}
func (m *GetLocationViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLocationViewRequest.Merge(m, src)
}
func (m *GetLocationViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetLocationViewRequest.Size(m)
}
func (m *GetLocationViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLocationViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLocationViewRequest proto.InternalMessageInfo

func (m *GetLocationViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetLocationViewRequest)(nil), "google.ads.googleads.v1.services.GetLocationViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/location_view_service.proto", fileDescriptor_cee10bb181b86493)
}

var fileDescriptor_cee10bb181b86493 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbf, 0x4a, 0xfb, 0x50,
	0x14, 0x26, 0xf9, 0xc1, 0x0f, 0x0c, 0x8a, 0x10, 0x41, 0x4b, 0x74, 0x28, 0xb5, 0x83, 0x74, 0xb8,
	0x97, 0x28, 0xa2, 0x5c, 0xed, 0x90, 0x2e, 0x75, 0x10, 0x29, 0x15, 0x32, 0x48, 0xa0, 0xc4, 0xe4,
	0x10, 0x02, 0x49, 0x6e, 0xcd, 0x49, 0xd3, 0x41, 0x5c, 0x7c, 0x05, 0xdf, 0xc0, 0xd1, 0xdd, 0x97,
	0xe8, 0xea, 0x2b, 0x38, 0xf5, 0x25, 0x94, 0xf4, 0xe6, 0xa6, 0x55, 0x5b, 0xba, 0x7d, 0x9c, 0xf3,
	0xfd, 0xb9, 0xe7, 0x4b, 0xb4, 0xcb, 0x80, 0xf3, 0x20, 0x02, 0xea, 0xfa, 0x48, 0x05, 0x2c, 0x50,
	0x6e, 0x52, 0x84, 0x34, 0x0f, 0x3d, 0x40, 0x1a, 0x71, 0xcf, 0xcd, 0x42, 0x9e, 0x0c, 0xf2, 0x10,
	0xc6, 0x83, 0x72, 0x4c, 0x86, 0x29, 0xcf, 0xb8, 0x5e, 0x17, 0x12, 0xe2, 0xfa, 0x48, 0x2a, 0x35,
	0xc9, 0x4d, 0x22, 0xd5, 0xc6, 0xe9, 0x2a, 0xff, 0x14, 0x90, 0x8f, 0xd2, 0x3f, 0x01, 0xc2, 0xd8,
	0x38, 0x90, 0xb2, 0x61, 0x48, 0xdd, 0x24, 0xe1, 0xd9, 0x8c, 0x81, 0xe5, 0x76, 0x6f, 0x61, 0xeb,
	0x45, 0x21, 0x24, 0x99, 0x58, 0x34, 0xda, 0xda, 0x6e, 0x17, 0xb2, 0xeb, 0xd2, 0xd0, 0x0e, 0x61,
	0xdc, 0x87, 0x87, 0x11, 0x60, 0xa6, 0x1f, 0x6a, 0x5b, 0x32, 0x71, 0x90, 0xb8, 0x31, 0xd4, 0x94,
	0xba, 0x72, 0xb4, 0xd1, 0xdf, 0x94, 0xc3, 0x1b, 0x37, 0x86, 0xe3, 0xa9, 0xa2, 0xed, 0x2c, 0x8a,
	0x6f, 0xc5, 0x15, 0xfa, 0xbb, 0xa2, 0x6d, 0xff, 0xf2, 0xd5, 0xcf, 0xc9, 0xba, 0xdb, 0xc9, 0xf2,
	0xa7, 0x18, 0x74, 0xa5, 0xb2, 0xea, 0x84, 0x2c, 0xea, 0x1a, 0x67, 0xcf, 0x1f, 0x9f, 0x2f, 0xaa,
	0xa9, 0xd3, 0xa2, 0xb7, 0xc7, 0x1f, 0x67, 0xb4, 0xbd, 0x11, 0x66, 0x3c, 0x86, 0x14, 0x69, 0xab,
	0x2a, 0xb2, 0x10, 0x21, 0x6d, 0x3d, 0x19, 0xfb, 0x13, 0xab, 0x36, 0x0f, 0x28, 0xd1, 0x30, 0x44,
	0xe2, 0xf1, 0xb8, 0xf3, 0xa5, 0x68, 0x4d, 0x8f, 0xc7, 0x6b, 0xcf, 0xe8, 0xd4, 0x96, 0x54, 0xd2,
	0x2b, 0xea, 0xee, 0x29, 0x77, 0x57, 0xa5, 0x3a, 0xe0, 0x91, 0x9b, 0x04, 0x84, 0xa7, 0x01, 0x0d,
	0x20, 0x99, 0x7d, 0x0c, 0x3a, 0xcf, 0x5b, 0xfd, 0x77, 0x5d, 0x48, 0xf0, 0xaa, 0xfe, 0xeb, 0x5a,
	0xd6, 0x9b, 0x5a, 0xef, 0x0a, 0x43, 0xcb, 0x47, 0x22, 0x60, 0x81, 0x6c, 0x93, 0x94, 0xc1, 0x38,
	0x91, 0x14, 0xc7, 0xf2, 0xd1, 0xa9, 0x28, 0x8e, 0x6d, 0x3a, 0x92, 0x32, 0x55, 0x9b, 0x62, 0xce,
	0x98, 0xe5, 0x23, 0x63, 0x15, 0x89, 0x31, 0xdb, 0x64, 0x4c, 0xd2, 0xee, 0xff, 0xcf, 0xde, 0x79,
	0xf2, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xee, 0x61, 0x40, 0x7a, 0x04, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LocationViewServiceClient is the client API for LocationViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LocationViewServiceClient interface {
	// Returns the requested location view in full detail.
	GetLocationView(ctx context.Context, in *GetLocationViewRequest, opts ...grpc.CallOption) (*resources.LocationView, error)
}

type locationViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationViewServiceClient(cc grpc.ClientConnInterface) LocationViewServiceClient {
	return &locationViewServiceClient{cc}
}

func (c *locationViewServiceClient) GetLocationView(ctx context.Context, in *GetLocationViewRequest, opts ...grpc.CallOption) (*resources.LocationView, error) {
	out := new(resources.LocationView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.LocationViewService/GetLocationView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationViewServiceServer is the server API for LocationViewService service.
type LocationViewServiceServer interface {
	// Returns the requested location view in full detail.
	GetLocationView(context.Context, *GetLocationViewRequest) (*resources.LocationView, error)
}

// UnimplementedLocationViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLocationViewServiceServer struct {
}

func (*UnimplementedLocationViewServiceServer) GetLocationView(ctx context.Context, req *GetLocationViewRequest) (*resources.LocationView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocationView not implemented")
}

func RegisterLocationViewServiceServer(s *grpc.Server, srv LocationViewServiceServer) {
	s.RegisterService(&_LocationViewService_serviceDesc, srv)
}

func _LocationViewService_GetLocationView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationViewServiceServer).GetLocationView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.LocationViewService/GetLocationView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationViewServiceServer).GetLocationView(ctx, req.(*GetLocationViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LocationViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.LocationViewService",
	HandlerType: (*LocationViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocationView",
			Handler:    _LocationViewService_GetLocationView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/location_view_service.proto",
}
