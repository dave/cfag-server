// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cfag.proto

package cfagproto

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

type Person struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_73d9d5b1044fe83f, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "cfagproto.Person")
}

func init() { proto.RegisterFile("cfag.proto", fileDescriptor_73d9d5b1044fe83f) }

var fileDescriptor_73d9d5b1044fe83f = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4e, 0x4b, 0x4c,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0xb1, 0xc1, 0x4c, 0xa5, 0x28, 0x2e, 0xb6,
	0x80, 0xd4, 0xa2, 0xe2, 0xfc, 0x3c, 0x21, 0x59, 0x2e, 0xae, 0xb4, 0xcc, 0xa2, 0xe2, 0x92, 0xf8,
	0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x4e, 0xb0, 0x88, 0x5f, 0x62,
	0x6e, 0xaa, 0x90, 0x34, 0x17, 0x67, 0x4e, 0x22, 0x4c, 0x96, 0x09, 0x2c, 0xcb, 0x01, 0x12, 0x00,
	0x4b, 0x8a, 0x70, 0xb1, 0xa6, 0xe6, 0x26, 0x66, 0xe6, 0x48, 0x30, 0x83, 0x25, 0x20, 0x9c, 0x24,
	0x36, 0xb0, 0x15, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x80, 0x2b, 0xe6, 0x7b, 0x00,
	0x00, 0x00,
}
