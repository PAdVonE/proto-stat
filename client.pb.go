// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

package stat_repository

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Client struct {
	Uuid                 string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Referral             int32                `protobuf:"varint,2,opt,name=referral,proto3" json:"referral,omitempty"`
	Stream               int32                `protobuf:"varint,3,opt,name=stream,proto3" json:"stream,omitempty"`
	Service              int32                `protobuf:"varint,4,opt,name=service,proto3" json:"service,omitempty"`
	UtmSource            string               `protobuf:"bytes,5,opt,name=utm_source,json=utmSource,proto3" json:"utm_source,omitempty"`
	UtmMedium            string               `protobuf:"bytes,6,opt,name=utm_medium,json=utmMedium,proto3" json:"utm_medium,omitempty"`
	UtmCampaign          string               `protobuf:"bytes,7,opt,name=utm_campaign,json=utmCampaign,proto3" json:"utm_campaign,omitempty"`
	UtmContent           string               `protobuf:"bytes,8,opt,name=utm_content,json=utmContent,proto3" json:"utm_content,omitempty"`
	UtmTerm              string               `protobuf:"bytes,9,opt,name=utm_term,json=utmTerm,proto3" json:"utm_term,omitempty"`
	Sub1                 string               `protobuf:"bytes,10,opt,name=sub1,proto3" json:"sub1,omitempty"`
	Sub2                 string               `protobuf:"bytes,11,opt,name=sub2,proto3" json:"sub2,omitempty"`
	Sub3                 string               `protobuf:"bytes,12,opt,name=sub3,proto3" json:"sub3,omitempty"`
	Sub4                 string               `protobuf:"bytes,14,opt,name=sub4,proto3" json:"sub4,omitempty"`
	Sub5                 string               `protobuf:"bytes,15,opt,name=sub5,proto3" json:"sub5,omitempty"`
	Gender               string               `protobuf:"bytes,16,opt,name=gender,proto3" json:"gender,omitempty"`
	Referrer             string               `protobuf:"bytes,17,opt,name=referrer,proto3" json:"referrer,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,18,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LocationIp           string               `protobuf:"bytes,19,opt,name=location_ip,json=locationIp,proto3" json:"location_ip,omitempty"`
	UserAgent            string               `protobuf:"bytes,20,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{0}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Client) GetReferral() int32 {
	if m != nil {
		return m.Referral
	}
	return 0
}

func (m *Client) GetStream() int32 {
	if m != nil {
		return m.Stream
	}
	return 0
}

func (m *Client) GetService() int32 {
	if m != nil {
		return m.Service
	}
	return 0
}

func (m *Client) GetUtmSource() string {
	if m != nil {
		return m.UtmSource
	}
	return ""
}

func (m *Client) GetUtmMedium() string {
	if m != nil {
		return m.UtmMedium
	}
	return ""
}

func (m *Client) GetUtmCampaign() string {
	if m != nil {
		return m.UtmCampaign
	}
	return ""
}

func (m *Client) GetUtmContent() string {
	if m != nil {
		return m.UtmContent
	}
	return ""
}

func (m *Client) GetUtmTerm() string {
	if m != nil {
		return m.UtmTerm
	}
	return ""
}

func (m *Client) GetSub1() string {
	if m != nil {
		return m.Sub1
	}
	return ""
}

func (m *Client) GetSub2() string {
	if m != nil {
		return m.Sub2
	}
	return ""
}

func (m *Client) GetSub3() string {
	if m != nil {
		return m.Sub3
	}
	return ""
}

func (m *Client) GetSub4() string {
	if m != nil {
		return m.Sub4
	}
	return ""
}

func (m *Client) GetSub5() string {
	if m != nil {
		return m.Sub5
	}
	return ""
}

func (m *Client) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *Client) GetReferrer() string {
	if m != nil {
		return m.Referrer
	}
	return ""
}

func (m *Client) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Client) GetLocationIp() string {
	if m != nil {
		return m.LocationIp
	}
	return ""
}

func (m *Client) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

type ClientRequest struct {
	Uuid                 string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Referral             int32                `protobuf:"varint,2,opt,name=referral,proto3" json:"referral,omitempty"`
	Stream               int32                `protobuf:"varint,3,opt,name=stream,proto3" json:"stream,omitempty"`
	Service              int32                `protobuf:"varint,4,opt,name=service,proto3" json:"service,omitempty"`
	UtmSource            string               `protobuf:"bytes,5,opt,name=utm_source,json=utmSource,proto3" json:"utm_source,omitempty"`
	UtmMedium            string               `protobuf:"bytes,6,opt,name=utm_medium,json=utmMedium,proto3" json:"utm_medium,omitempty"`
	UtmCampaign          string               `protobuf:"bytes,7,opt,name=utm_campaign,json=utmCampaign,proto3" json:"utm_campaign,omitempty"`
	UtmContent           string               `protobuf:"bytes,8,opt,name=utm_content,json=utmContent,proto3" json:"utm_content,omitempty"`
	UtmTerm              string               `protobuf:"bytes,9,opt,name=utm_term,json=utmTerm,proto3" json:"utm_term,omitempty"`
	Sub1                 string               `protobuf:"bytes,10,opt,name=sub1,proto3" json:"sub1,omitempty"`
	Sub2                 string               `protobuf:"bytes,11,opt,name=sub2,proto3" json:"sub2,omitempty"`
	Sub3                 string               `protobuf:"bytes,12,opt,name=sub3,proto3" json:"sub3,omitempty"`
	Sub4                 string               `protobuf:"bytes,14,opt,name=sub4,proto3" json:"sub4,omitempty"`
	Sub5                 string               `protobuf:"bytes,15,opt,name=sub5,proto3" json:"sub5,omitempty"`
	Gender               string               `protobuf:"bytes,16,opt,name=gender,proto3" json:"gender,omitempty"`
	Referrer             string               `protobuf:"bytes,17,opt,name=referrer,proto3" json:"referrer,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,18,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LocationIp           string               `protobuf:"bytes,19,opt,name=location_ip,json=locationIp,proto3" json:"location_ip,omitempty"`
	UserAgent            string               `protobuf:"bytes,20,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClientRequest) Reset()         { *m = ClientRequest{} }
func (m *ClientRequest) String() string { return proto.CompactTextString(m) }
func (*ClientRequest) ProtoMessage()    {}
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{1}
}

func (m *ClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientRequest.Unmarshal(m, b)
}
func (m *ClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientRequest.Marshal(b, m, deterministic)
}
func (m *ClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientRequest.Merge(m, src)
}
func (m *ClientRequest) XXX_Size() int {
	return xxx_messageInfo_ClientRequest.Size(m)
}
func (m *ClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientRequest proto.InternalMessageInfo

func (m *ClientRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ClientRequest) GetReferral() int32 {
	if m != nil {
		return m.Referral
	}
	return 0
}

func (m *ClientRequest) GetStream() int32 {
	if m != nil {
		return m.Stream
	}
	return 0
}

func (m *ClientRequest) GetService() int32 {
	if m != nil {
		return m.Service
	}
	return 0
}

func (m *ClientRequest) GetUtmSource() string {
	if m != nil {
		return m.UtmSource
	}
	return ""
}

func (m *ClientRequest) GetUtmMedium() string {
	if m != nil {
		return m.UtmMedium
	}
	return ""
}

func (m *ClientRequest) GetUtmCampaign() string {
	if m != nil {
		return m.UtmCampaign
	}
	return ""
}

func (m *ClientRequest) GetUtmContent() string {
	if m != nil {
		return m.UtmContent
	}
	return ""
}

func (m *ClientRequest) GetUtmTerm() string {
	if m != nil {
		return m.UtmTerm
	}
	return ""
}

func (m *ClientRequest) GetSub1() string {
	if m != nil {
		return m.Sub1
	}
	return ""
}

func (m *ClientRequest) GetSub2() string {
	if m != nil {
		return m.Sub2
	}
	return ""
}

func (m *ClientRequest) GetSub3() string {
	if m != nil {
		return m.Sub3
	}
	return ""
}

func (m *ClientRequest) GetSub4() string {
	if m != nil {
		return m.Sub4
	}
	return ""
}

func (m *ClientRequest) GetSub5() string {
	if m != nil {
		return m.Sub5
	}
	return ""
}

func (m *ClientRequest) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *ClientRequest) GetReferrer() string {
	if m != nil {
		return m.Referrer
	}
	return ""
}

func (m *ClientRequest) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ClientRequest) GetLocationIp() string {
	if m != nil {
		return m.LocationIp
	}
	return ""
}

func (m *ClientRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func init() {
	proto.RegisterType((*Client)(nil), "stat.repository.Client")
	proto.RegisterType((*ClientRequest)(nil), "stat.repository.ClientRequest")
}

func init() { proto.RegisterFile("client.proto", fileDescriptor_014de31d7ac8c57c) }

var fileDescriptor_014de31d7ac8c57c = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x92, 0x4f, 0x6f, 0x13, 0x31,
	0x10, 0xc5, 0x15, 0x48, 0xf3, 0xc7, 0x09, 0x14, 0x0c, 0x42, 0x43, 0x24, 0xd4, 0xd0, 0x53, 0x4e,
	0x5b, 0x91, 0xb4, 0x07, 0x8e, 0x55, 0x4e, 0x1c, 0xb8, 0x84, 0xde, 0x23, 0x67, 0x33, 0x5d, 0x59,
	0x5a, 0xaf, 0x97, 0xf1, 0x18, 0x89, 0xaf, 0xc2, 0x8d, 0x6f, 0x8a, 0x3c, 0xeb, 0x5d, 0xa4, 0x7e,
	0x86, 0xde, 0xfc, 0x7e, 0xef, 0xad, 0x57, 0x7e, 0x33, 0x6a, 0x59, 0xd6, 0x16, 0x1b, 0x2e, 0x5a,
	0xf2, 0xec, 0xf5, 0x65, 0x60, 0xc3, 0x05, 0x61, 0xeb, 0x83, 0x65, 0x4f, 0xbf, 0x57, 0x57, 0x95,
	0xf7, 0x55, 0x8d, 0x37, 0x62, 0x9f, 0xe2, 0xe3, 0x0d, 0x5b, 0x87, 0x81, 0x8d, 0x6b, 0xbb, 0x2f,
	0xae, 0xff, 0x8c, 0xd5, 0x64, 0x2f, 0x57, 0x68, 0xad, 0xc6, 0x31, 0xda, 0x33, 0x8c, 0xd6, 0xa3,
	0xcd, 0xfc, 0x20, 0x67, 0xbd, 0x52, 0x33, 0xc2, 0x47, 0x24, 0x32, 0x35, 0xbc, 0x58, 0x8f, 0x36,
	0x17, 0x87, 0x41, 0xeb, 0x0f, 0x6a, 0x12, 0x98, 0xd0, 0x38, 0x78, 0x29, 0x4e, 0x56, 0x1a, 0xd4,
	0x34, 0x20, 0xfd, 0xb2, 0x25, 0xc2, 0x58, 0x8c, 0x5e, 0xea, 0x4f, 0x4a, 0x45, 0x76, 0xc7, 0xe0,
	0x23, 0x95, 0x08, 0x17, 0xf2, 0x9f, 0x79, 0x64, 0xf7, 0x43, 0x40, 0x6f, 0x3b, 0x3c, 0xdb, 0xe8,
	0x60, 0x32, 0xd8, 0xdf, 0x05, 0xe8, 0xcf, 0x6a, 0x99, 0xec, 0xd2, 0xb8, 0xd6, 0xd8, 0xaa, 0x81,
	0xa9, 0x04, 0x16, 0x91, 0xdd, 0x3e, 0x23, 0x7d, 0xa5, 0x16, 0x12, 0xf1, 0x0d, 0x63, 0xc3, 0x30,
	0x93, 0x44, 0xba, 0x74, 0xdf, 0x11, 0xfd, 0x51, 0xcd, 0x52, 0x80, 0x91, 0x1c, 0xcc, 0xc5, 0x9d,
	0x46, 0x76, 0x0f, 0x48, 0x2e, 0x3d, 0x3f, 0xc4, 0xd3, 0x17, 0x50, 0xdd, 0xf3, 0xd3, 0x39, 0xb3,
	0x2d, 0x2c, 0x06, 0xb6, 0xcd, 0x6c, 0x07, 0xcb, 0x81, 0xed, 0x32, 0xbb, 0x85, 0xd7, 0x03, 0xbb,
	0xcd, 0xec, 0x0e, 0x2e, 0x07, 0x76, 0x97, 0x2a, 0xab, 0xb0, 0x39, 0x23, 0xc1, 0x1b, 0xa1, 0x59,
	0xfd, 0xaf, 0x19, 0x09, 0xde, 0x8a, 0x33, 0x68, 0xfd, 0x55, 0xa9, 0x92, 0xd0, 0x30, 0x9e, 0x8f,
	0x86, 0x41, 0xaf, 0x47, 0x9b, 0xc5, 0x76, 0x55, 0x74, 0x73, 0x2d, 0xfa, 0xb9, 0x16, 0x0f, 0xfd,
	0x5c, 0x0f, 0xf3, 0x9c, 0xbe, 0xe7, 0x54, 0x47, 0xed, 0x4b, 0xc3, 0xd6, 0x37, 0x47, 0xdb, 0xc2,
	0xbb, 0xae, 0x8e, 0x1e, 0x7d, 0x6b, 0xa5, 0xf1, 0x80, 0x74, 0x34, 0x55, 0xaa, 0xeb, 0x7d, 0x6e,
	0x3c, 0x20, 0xdd, 0x27, 0x70, 0xfd, 0x77, 0xac, 0x5e, 0x75, 0xcb, 0x71, 0xc0, 0x9f, 0x11, 0xc3,
	0xf3, 0x8e, 0x3c, 0xef, 0xc8, 0x93, 0x1d, 0x39, 0x4d, 0xe4, 0xfa, 0xdd, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc3, 0x34, 0xa4, 0x81, 0x89, 0x04, 0x00, 0x00,
}
