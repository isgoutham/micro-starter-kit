// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/emailer/proto/emailer/emailer.proto

package go_micro_srv_emailer

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

type Message struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Subject              string   `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_81737eb8d95d8ad1, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Message) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.emailer.Message")
}

func init() {
	proto.RegisterFile("srv/emailer/proto/emailer/emailer.proto", fileDescriptor_81737eb8d95d8ad1)
}

var fileDescriptor_81737eb8d95d8ad1 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x2e, 0x2a, 0xd3,
	0x4f, 0xcd, 0x4d, 0xcc, 0xcc, 0x49, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0xf3, 0xa0,
	0xb4, 0x1e, 0x58, 0x54, 0x48, 0x24, 0x3d, 0x5f, 0x2f, 0x37, 0x33, 0xb9, 0x28, 0x5f, 0xaf, 0xb8,
	0xa8, 0x4c, 0x0f, 0x2a, 0xa7, 0x14, 0xcd, 0xc5, 0xee, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x2a,
	0xc4, 0xc7, 0xc5, 0x54, 0x92, 0x2f, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x54, 0x92, 0x2f,
	0x24, 0xc4, 0xc5, 0x92, 0x56, 0x94, 0x9f, 0x2b, 0xc1, 0x04, 0x16, 0x01, 0xb3, 0x85, 0x24, 0xb8,
	0xd8, 0x8b, 0x4b, 0x93, 0xb2, 0x52, 0x93, 0x4b, 0x24, 0x98, 0xc1, 0xc2, 0x30, 0x2e, 0x48, 0x75,
	0x52, 0x7e, 0x4a, 0xa5, 0x04, 0x0b, 0x44, 0x35, 0x88, 0x9d, 0xc4, 0x06, 0xb6, 0xd9, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x4a, 0xe8, 0x64, 0x2b, 0xa4, 0x00, 0x00, 0x00,
}