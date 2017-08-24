// Code generated by protoc-gen-go.
// source: github.com/google/gapid/gapis/api/gvr/resolvables.proto
// DO NOT EDIT!

/*
Package gvr is a generated protocol buffer package.

It is generated from these files:
	github.com/google/gapid/gapis/api/gvr/resolvables.proto

It has these top-level messages:
	FrameBindingsResolvable
*/
package gvr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import path "github.com/google/gapid/gapis/service/path"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// GAPIS internal structure.
type FrameBindingsResolvable struct {
	Capture *path.Capture `protobuf:"bytes,1,opt,name=capture" json:"capture,omitempty"`
}

func (m *FrameBindingsResolvable) Reset()                    { *m = FrameBindingsResolvable{} }
func (m *FrameBindingsResolvable) String() string            { return proto.CompactTextString(m) }
func (*FrameBindingsResolvable) ProtoMessage()               {}
func (*FrameBindingsResolvable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FrameBindingsResolvable) GetCapture() *path.Capture {
	if m != nil {
		return m.Capture
	}
	return nil
}

func init() {
	proto.RegisterType((*FrameBindingsResolvable)(nil), "gvr.FrameBindingsResolvable")
}

func init() {
	proto.RegisterFile("github.com/google/gapid/gapis/api/gvr/resolvables.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x4f,
	0x2c, 0xc8, 0x4c, 0x01, 0x93, 0xc5, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0xe9, 0x65, 0x45, 0xfa, 0x45,
	0xa9, 0xc5, 0xf9, 0x39, 0x65, 0x89, 0x49, 0x39, 0xa9, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0xcc, 0xe9, 0x65, 0x45, 0x52, 0xb2, 0x10, 0x55, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9,
	0xfa, 0x05, 0x89, 0x25, 0x19, 0x60, 0x02, 0xa2, 0x46, 0xc9, 0x89, 0x4b, 0xdc, 0xad, 0x28, 0x31,
	0x37, 0xd5, 0x29, 0x33, 0x2f, 0x25, 0x33, 0x2f, 0xbd, 0x38, 0x08, 0x6e, 0x8a, 0x90, 0x3a, 0x17,
	0x7b, 0x72, 0x62, 0x41, 0x49, 0x69, 0x51, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x11, 0xaf,
	0x1e, 0x58, 0xa3, 0x33, 0x44, 0x30, 0x08, 0x26, 0x9b, 0xc4, 0x06, 0x36, 0xca, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xd7, 0xb8, 0xc9, 0xb1, 0xa9, 0x00, 0x00, 0x00,
}