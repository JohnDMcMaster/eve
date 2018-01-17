// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devcommon.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UUIDandVersion struct {
	Uuid    string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *UUIDandVersion) Reset()                    { *m = UUIDandVersion{} }
func (m *UUIDandVersion) String() string            { return proto.CompactTextString(m) }
func (*UUIDandVersion) ProtoMessage()               {}
func (*UUIDandVersion) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *UUIDandVersion) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *UUIDandVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*UUIDandVersion)(nil), "UUIDandVersion")
}

func init() { proto.RegisterFile("devcommon.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x49, 0x2d, 0x4b,
	0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe3, 0xe2, 0x0b,
	0x0d, 0xf5, 0x74, 0x49, 0xcc, 0x4b, 0x09, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x13, 0x12, 0xe2,
	0x62, 0x29, 0x2d, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x24,
	0xb8, 0xd8, 0xcb, 0x20, 0xd2, 0x12, 0x4c, 0x60, 0x61, 0x18, 0xd7, 0xc9, 0x81, 0x4b, 0x3e, 0x39,
	0x3f, 0x57, 0xaf, 0x2a, 0x35, 0x25, 0x35, 0x25, 0x51, 0x2f, 0x39, 0x27, 0xbf, 0x34, 0x45, 0xaf,
	0xb4, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x15, 0x62, 0x45, 0x94, 0x6c, 0x7a, 0x66, 0x49, 0x46,
	0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x44, 0x9d, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x55, 0x72,
	0x7e, 0x5e, 0x5a, 0x66, 0x7a, 0x12, 0x1b, 0x58, 0x95, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa8,
	0x4a, 0x2f, 0x1c, 0x9b, 0x00, 0x00, 0x00,
}
