// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package messagepb

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

type Message struct {
	Receiver             string   `protobuf:"bytes,1,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Sender               string   `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Value                int64    `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_a29cb1ffc5d4e9fb, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *Message) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Message) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Message) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "microservice.message.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_a29cb1ffc5d4e9fb) }

var fileDescriptor_message_a29cb1ffc5d4e9fb = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc9, 0xcd, 0x4c, 0x2e, 0xca,
	0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x83, 0xca, 0x29, 0xa5, 0x73, 0xb1, 0xfb, 0x42,
	0x98, 0x42, 0x52, 0x5c, 0x1c, 0x45, 0xa9, 0xc9, 0xa9, 0x99, 0x65, 0xa9, 0x45, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x18, 0x17, 0x5b, 0x71, 0x6a, 0x5e, 0x4a, 0x6a, 0x91,
	0x04, 0x13, 0x58, 0x06, 0xca, 0x13, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60,
	0x56, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0x70, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25,
	0x58, 0xc0, 0x6a, 0xc1, 0x6c, 0x27, 0xee, 0x28, 0x4e, 0xa8, 0x9d, 0x05, 0x49, 0x49, 0x6c, 0x60,
	0x27, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x75, 0x0c, 0x00, 0xaf, 0xa3, 0x00, 0x00, 0x00,
}