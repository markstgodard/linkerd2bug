// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello/service.proto

package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SayHelloRequest struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloRequest) Reset()         { *m = SayHelloRequest{} }
func (m *SayHelloRequest) String() string { return proto.CompactTextString(m) }
func (*SayHelloRequest) ProtoMessage()    {}
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_a2f82f4c147fb45b, []int{0}
}
func (m *SayHelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloRequest.Unmarshal(m, b)
}
func (m *SayHelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloRequest.Marshal(b, m, deterministic)
}
func (dst *SayHelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloRequest.Merge(dst, src)
}
func (m *SayHelloRequest) XXX_Size() int {
	return xxx_messageInfo_SayHelloRequest.Size(m)
}
func (m *SayHelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloRequest proto.InternalMessageInfo

func (m *SayHelloRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type SayHelloResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloResponse) Reset()         { *m = SayHelloResponse{} }
func (m *SayHelloResponse) String() string { return proto.CompactTextString(m) }
func (*SayHelloResponse) ProtoMessage()    {}
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_a2f82f4c147fb45b, []int{1}
}
func (m *SayHelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloResponse.Unmarshal(m, b)
}
func (m *SayHelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloResponse.Marshal(b, m, deterministic)
}
func (dst *SayHelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloResponse.Merge(dst, src)
}
func (m *SayHelloResponse) XXX_Size() int {
	return xxx_messageInfo_SayHelloResponse.Size(m)
}
func (m *SayHelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloResponse proto.InternalMessageInfo

func (m *SayHelloResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*SayHelloRequest)(nil), "baz.hello.SayHelloRequest")
	proto.RegisterType((*SayHelloResponse)(nil), "baz.hello.SayHelloResponse")
}

func init() { proto.RegisterFile("hello/service.proto", fileDescriptor_service_a2f82f4c147fb45b) }

var fileDescriptor_service_a2f82f4c147fb45b = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4c, 0x4a, 0xac, 0xd2, 0x03, 0x4b, 0x28, 0x29, 0x73, 0xf1, 0x07, 0x27, 0x56, 0x7a, 0x80,
	0xd8, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x92, 0x0a, 0x97, 0x00, 0x42, 0x51, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0x2a, 0xa6, 0x2a, 0x23, 0x3f, 0x2e, 0x2e, 0xa7, 0xc4, 0xaa, 0x60, 0x88, 0x4d,
	0x42, 0x0e, 0x5c, 0xac, 0x60, 0x0d, 0x42, 0x52, 0x7a, 0x70, 0xdb, 0xf4, 0xd0, 0xac, 0x92, 0x92,
	0xc6, 0x2a, 0x07, 0xb1, 0xc1, 0x89, 0x3d, 0x8a, 0x15, 0x2c, 0x93, 0xc4, 0x06, 0x76, 0xb5, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x8b, 0xe7, 0xe8, 0xcc, 0x00, 0x00, 0x00,
}
