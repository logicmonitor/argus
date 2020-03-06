// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/travel_placeholder_field.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Possible values for Travel placeholder fields.
type TravelPlaceholderFieldEnum_TravelPlaceholderField int32

const (
	// Not specified.
	TravelPlaceholderFieldEnum_UNSPECIFIED TravelPlaceholderFieldEnum_TravelPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	TravelPlaceholderFieldEnum_UNKNOWN TravelPlaceholderFieldEnum_TravelPlaceholderField = 1
	// Data Type: STRING. Required. Destination id. Example: PAR, LON.
	// For feed items that only have destination id, destination id must be a
	// unique key. For feed items that have both destination id and origin id,
	// then the combination must be a unique key.
	TravelPlaceholderFieldEnum_DESTINATION_ID TravelPlaceholderFieldEnum_TravelPlaceholderField = 2
	// Data Type: STRING. Origin id. Example: PAR, LON.
	// Combination of DESTINATION_ID and ORIGIN_ID must be
	// unique per offer.
	TravelPlaceholderFieldEnum_ORIGIN_ID TravelPlaceholderFieldEnum_TravelPlaceholderField = 3
	// Data Type: STRING. Required. Main headline with name to be shown in
	// dynamic ad.
	TravelPlaceholderFieldEnum_TITLE TravelPlaceholderFieldEnum_TravelPlaceholderField = 4
	// Data Type: STRING. The destination name. Shorter names are recommended.
	TravelPlaceholderFieldEnum_DESTINATION_NAME TravelPlaceholderFieldEnum_TravelPlaceholderField = 5
	// Data Type: STRING. Origin name. Shorter names are recommended.
	TravelPlaceholderFieldEnum_ORIGIN_NAME TravelPlaceholderFieldEnum_TravelPlaceholderField = 6
	// Data Type: STRING. Price to be shown in the ad. Highly recommended for
	// dynamic ads.
	// Example: "100.00 USD"
	TravelPlaceholderFieldEnum_PRICE TravelPlaceholderFieldEnum_TravelPlaceholderField = 7
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	TravelPlaceholderFieldEnum_FORMATTED_PRICE TravelPlaceholderFieldEnum_TravelPlaceholderField = 8
	// Data Type: STRING. Sale price to be shown in the ad.
	// Example: "80.00 USD"
	TravelPlaceholderFieldEnum_SALE_PRICE TravelPlaceholderFieldEnum_TravelPlaceholderField = 9
	// Data Type: STRING. Formatted sale price to be shown in the ad.
	// Example: "On sale for $80.00", "$60 - $80"
	TravelPlaceholderFieldEnum_FORMATTED_SALE_PRICE TravelPlaceholderFieldEnum_TravelPlaceholderField = 10
	// Data Type: URL. Image to be displayed in the ad.
	TravelPlaceholderFieldEnum_IMAGE_URL TravelPlaceholderFieldEnum_TravelPlaceholderField = 11
	// Data Type: STRING. Category of travel offer used to group like items
	// together for recommendation engine.
	TravelPlaceholderFieldEnum_CATEGORY TravelPlaceholderFieldEnum_TravelPlaceholderField = 12
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	TravelPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS TravelPlaceholderFieldEnum_TravelPlaceholderField = 13
	// Data Type: STRING. Address of travel offer, including postal code.
	TravelPlaceholderFieldEnum_DESTINATION_ADDRESS TravelPlaceholderFieldEnum_TravelPlaceholderField = 14
	// Data Type: URL_LIST. Required. Final URLs to be used in ad, when using
	// Upgraded URLs; the more specific the better (e.g. the individual URL of a
	// specific travel offer and its location).
	TravelPlaceholderFieldEnum_FINAL_URL TravelPlaceholderFieldEnum_TravelPlaceholderField = 15
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	TravelPlaceholderFieldEnum_FINAL_MOBILE_URLS TravelPlaceholderFieldEnum_TravelPlaceholderField = 16
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	TravelPlaceholderFieldEnum_TRACKING_URL TravelPlaceholderFieldEnum_TravelPlaceholderField = 17
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	TravelPlaceholderFieldEnum_ANDROID_APP_LINK TravelPlaceholderFieldEnum_TravelPlaceholderField = 18
	// Data Type: STRING_LIST. List of recommended destination IDs to show
	// together with this item.
	TravelPlaceholderFieldEnum_SIMILAR_DESTINATION_IDS TravelPlaceholderFieldEnum_TravelPlaceholderField = 19
	// Data Type: STRING. iOS app link.
	TravelPlaceholderFieldEnum_IOS_APP_LINK TravelPlaceholderFieldEnum_TravelPlaceholderField = 20
	// Data Type: INT64. iOS app store ID.
	TravelPlaceholderFieldEnum_IOS_APP_STORE_ID TravelPlaceholderFieldEnum_TravelPlaceholderField = 21
)

var TravelPlaceholderFieldEnum_TravelPlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "DESTINATION_ID",
	3:  "ORIGIN_ID",
	4:  "TITLE",
	5:  "DESTINATION_NAME",
	6:  "ORIGIN_NAME",
	7:  "PRICE",
	8:  "FORMATTED_PRICE",
	9:  "SALE_PRICE",
	10: "FORMATTED_SALE_PRICE",
	11: "IMAGE_URL",
	12: "CATEGORY",
	13: "CONTEXTUAL_KEYWORDS",
	14: "DESTINATION_ADDRESS",
	15: "FINAL_URL",
	16: "FINAL_MOBILE_URLS",
	17: "TRACKING_URL",
	18: "ANDROID_APP_LINK",
	19: "SIMILAR_DESTINATION_IDS",
	20: "IOS_APP_LINK",
	21: "IOS_APP_STORE_ID",
}

var TravelPlaceholderFieldEnum_TravelPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":             0,
	"UNKNOWN":                 1,
	"DESTINATION_ID":          2,
	"ORIGIN_ID":               3,
	"TITLE":                   4,
	"DESTINATION_NAME":        5,
	"ORIGIN_NAME":             6,
	"PRICE":                   7,
	"FORMATTED_PRICE":         8,
	"SALE_PRICE":              9,
	"FORMATTED_SALE_PRICE":    10,
	"IMAGE_URL":               11,
	"CATEGORY":                12,
	"CONTEXTUAL_KEYWORDS":     13,
	"DESTINATION_ADDRESS":     14,
	"FINAL_URL":               15,
	"FINAL_MOBILE_URLS":       16,
	"TRACKING_URL":            17,
	"ANDROID_APP_LINK":        18,
	"SIMILAR_DESTINATION_IDS": 19,
	"IOS_APP_LINK":            20,
	"IOS_APP_STORE_ID":        21,
}

func (x TravelPlaceholderFieldEnum_TravelPlaceholderField) String() string {
	return proto.EnumName(TravelPlaceholderFieldEnum_TravelPlaceholderField_name, int32(x))
}

func (TravelPlaceholderFieldEnum_TravelPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c5525e64a37cd678, []int{0, 0}
}

// Values for Travel placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type TravelPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TravelPlaceholderFieldEnum) Reset()         { *m = TravelPlaceholderFieldEnum{} }
func (m *TravelPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*TravelPlaceholderFieldEnum) ProtoMessage()    {}
func (*TravelPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5525e64a37cd678, []int{0}
}

func (m *TravelPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TravelPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *TravelPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TravelPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *TravelPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TravelPlaceholderFieldEnum.Merge(m, src)
}
func (m *TravelPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_TravelPlaceholderFieldEnum.Size(m)
}
func (m *TravelPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_TravelPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_TravelPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.TravelPlaceholderFieldEnum_TravelPlaceholderField", TravelPlaceholderFieldEnum_TravelPlaceholderField_name, TravelPlaceholderFieldEnum_TravelPlaceholderField_value)
	proto.RegisterType((*TravelPlaceholderFieldEnum)(nil), "google.ads.googleads.v3.enums.TravelPlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/travel_placeholder_field.proto", fileDescriptor_c5525e64a37cd678)
}

var fileDescriptor_c5525e64a37cd678 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0x6e, 0xda, 0x30,
	0x14, 0x1e, 0xb0, 0xfe, 0x60, 0x28, 0xb8, 0x86, 0xae, 0x53, 0xbb, 0x5e, 0xb4, 0x0f, 0x90, 0x5c,
	0xe4, 0x2e, 0xdb, 0x8d, 0x21, 0x26, 0xb2, 0x08, 0x4e, 0x14, 0x1b, 0xba, 0x4e, 0x48, 0x51, 0xd6,
	0x64, 0x19, 0x52, 0x48, 0x10, 0xa1, 0x3c, 0xd0, 0x2e, 0xf7, 0x02, 0x7b, 0x87, 0x3d, 0xc0, 0x1e,
	0x62, 0x57, 0x7b, 0x84, 0xc9, 0x0e, 0x7f, 0x93, 0xba, 0xde, 0x44, 0xe7, 0x9c, 0xef, 0xc7, 0x27,
	0xf6, 0x07, 0x3e, 0x24, 0x79, 0x9e, 0xa4, 0xb1, 0x1e, 0x46, 0x85, 0x5e, 0x96, 0xb2, 0x5a, 0x1b,
	0x7a, 0x9c, 0x3d, 0xcd, 0x0b, 0x7d, 0xb5, 0x0c, 0xd7, 0x71, 0x1a, 0x2c, 0xd2, 0xf0, 0x31, 0xfe,
	0x9a, 0xa7, 0x51, 0xbc, 0x0c, 0xbe, 0xcc, 0xe2, 0x34, 0xd2, 0x16, 0xcb, 0x7c, 0x95, 0xa3, 0x9b,
	0x52, 0xa2, 0x85, 0x51, 0xa1, 0xed, 0xd4, 0xda, 0xda, 0xd0, 0x94, 0xfa, 0xea, 0xdd, 0xd6, 0x7c,
	0x31, 0xd3, 0xc3, 0x2c, 0xcb, 0x57, 0xe1, 0x6a, 0x96, 0x67, 0x45, 0x29, 0xbe, 0xfb, 0x55, 0x03,
	0x57, 0x42, 0xf9, 0x7b, 0x7b, 0xfb, 0x81, 0x74, 0x27, 0xd9, 0xd3, 0xfc, 0xee, 0x47, 0x0d, 0xbc,
	0x79, 0x1e, 0x46, 0x6d, 0xd0, 0x18, 0x33, 0xee, 0x91, 0x3e, 0x1d, 0x50, 0x62, 0xc1, 0x57, 0xa8,
	0x01, 0x4e, 0xc6, 0x6c, 0xc8, 0xdc, 0x7b, 0x06, 0x2b, 0x08, 0x81, 0x96, 0x45, 0xb8, 0xa0, 0x0c,
	0x0b, 0xea, 0xb2, 0x80, 0x5a, 0xb0, 0x8a, 0xce, 0x40, 0xdd, 0xf5, 0xa9, 0x4d, 0x55, 0x5b, 0x43,
	0x75, 0x70, 0x24, 0xa8, 0x70, 0x08, 0x7c, 0x8d, 0xba, 0x00, 0x1e, 0xb2, 0x19, 0x1e, 0x11, 0x78,
	0x24, 0x4f, 0xd8, 0xf0, 0xd5, 0xe0, 0x58, 0x2a, 0x3c, 0x9f, 0xf6, 0x09, 0x3c, 0x41, 0x1d, 0xd0,
	0x1e, 0xb8, 0xfe, 0x08, 0x0b, 0x41, 0xac, 0xa0, 0x1c, 0x9e, 0xa2, 0x16, 0x00, 0x1c, 0x3b, 0x64,
	0xd3, 0xd7, 0xd1, 0x5b, 0xd0, 0xdd, 0x93, 0x0e, 0x10, 0x20, 0x57, 0xa1, 0x23, 0x6c, 0x93, 0x60,
	0xec, 0x3b, 0xb0, 0x81, 0x9a, 0xe0, 0xb4, 0x8f, 0x05, 0xb1, 0x5d, 0xff, 0x01, 0x36, 0xd1, 0x25,
	0xe8, 0xf4, 0x5d, 0x26, 0xc8, 0x47, 0x31, 0xc6, 0x4e, 0x30, 0x24, 0x0f, 0xf7, 0xae, 0x6f, 0x71,
	0x78, 0x26, 0x81, 0xc3, 0x35, 0xb1, 0x65, 0xf9, 0x84, 0x73, 0xd8, 0x92, 0x76, 0x03, 0xca, 0xb0,
	0xa3, 0xec, 0xda, 0xe8, 0x02, 0x9c, 0x97, 0xed, 0xc8, 0xed, 0x51, 0x47, 0x1d, 0xc2, 0x21, 0x44,
	0x10, 0x34, 0x85, 0x8f, 0xfb, 0x43, 0xca, 0x6c, 0x45, 0x3c, 0x97, 0xff, 0x8d, 0x99, 0xe5, 0xbb,
	0xd4, 0x0a, 0xb0, 0xe7, 0x05, 0x0e, 0x65, 0x43, 0x88, 0xd0, 0x35, 0xb8, 0xe4, 0x74, 0x44, 0x1d,
	0xec, 0x07, 0xff, 0xde, 0x21, 0x87, 0x1d, 0x69, 0x42, 0x5d, 0xbe, 0xa7, 0x77, 0xa5, 0xc9, 0x76,
	0xc2, 0x85, 0xeb, 0x13, 0x79, 0xbb, 0x17, 0xbd, 0x3f, 0x15, 0x70, 0xfb, 0x98, 0xcf, 0xb5, 0x17,
	0xc3, 0xd1, 0xbb, 0x7e, 0xfe, 0x71, 0x3d, 0x99, 0x0d, 0xaf, 0xf2, 0xa9, 0xb7, 0x51, 0x27, 0x79,
	0x1a, 0x66, 0x89, 0x96, 0x2f, 0x13, 0x3d, 0x89, 0x33, 0x95, 0x9c, 0x6d, 0x50, 0x17, 0xb3, 0xe2,
	0x3f, 0xb9, 0x7d, 0xaf, 0xbe, 0xdf, 0xaa, 0x35, 0x1b, 0xe3, 0xef, 0xd5, 0x1b, 0xbb, 0xb4, 0xc2,
	0x51, 0xa1, 0x95, 0xa5, 0xac, 0x26, 0x86, 0x26, 0x73, 0x56, 0xfc, 0xdc, 0xe2, 0x53, 0x1c, 0x15,
	0xd3, 0x1d, 0x3e, 0x9d, 0x18, 0x53, 0x85, 0xff, 0xae, 0xde, 0x96, 0x43, 0xd3, 0xc4, 0x51, 0x61,
	0x9a, 0x3b, 0x86, 0x69, 0x4e, 0x0c, 0xd3, 0x54, 0x9c, 0xcf, 0xc7, 0x6a, 0x31, 0xe3, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1f, 0x98, 0x06, 0xeb, 0x4f, 0x03, 0x00, 0x00,
}
