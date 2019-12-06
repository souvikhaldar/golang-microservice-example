// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sum.proto

package sum

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SumRequest struct {
	First                int32    `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	Second               int32    `protobuf:"varint,2,opt,name=second,proto3" json:"second,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62743f9cdc99b9fd, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetFirst() int32 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *SumRequest) GetSecond() int32 {
	if m != nil {
		return m.Second
	}
	return 0
}

type SumResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62743f9cdc99b9fd, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "SumRequest")
	proto.RegisterType((*SumResponse)(nil), "SumResponse")
}

func init() { proto.RegisterFile("sum.proto", fileDescriptor_62743f9cdc99b9fd) }

var fileDescriptor_62743f9cdc99b9fd = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2e, 0xcd, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe2, 0xe2, 0x0a, 0x2e, 0xcd, 0x0d, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x4d, 0xcb, 0x2c, 0x2a, 0x2e, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0d, 0x82, 0x70, 0x84, 0xc4, 0xb8, 0xd8, 0x8a, 0x53, 0x93, 0xf3, 0xf3, 0x52, 0x24,
	0x98, 0xc0, 0xc2, 0x50, 0x9e, 0x92, 0x2a, 0x17, 0x37, 0x58, 0x6f, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0x2a, 0x48, 0x59, 0x51, 0x6a, 0x71, 0x69, 0x0e, 0x4c, 0x37, 0x94, 0x67, 0xa4, 0xcd, 0xc5, 0xea,
	0x98, 0x92, 0x92, 0x5a, 0x24, 0xa4, 0xc4, 0xc5, 0x1c, 0x5c, 0x9a, 0x2b, 0xc4, 0xad, 0x87, 0xb0,
	0x51, 0x8a, 0x47, 0x0f, 0xc9, 0x08, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xb3, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xb8, 0xcf, 0x12, 0x3c, 0xa3, 0x00, 0x00, 0x00,
}