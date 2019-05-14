// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filesystem/behavior/probe.proto

package behavior

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

// ProbeMode specifies the mode for filesystem probing.
type ProbeMode int32

const (
	// ProbeMode_ProbeModeDefault represents an unspecified probe mode. It
	// should be converted to one of the following values based on the desired
	// default behavior.
	ProbeMode_ProbeModeDefault ProbeMode = 0
	// ProbeMode_ProbeModeProbe specifies that filesystem behavior should be
	// determined using temporary files or, if possible, a "fast-path" mechanism
	// (such as filesystem format detection) that provides quick but certain
	// determination of filesystem behavior.
	ProbeMode_ProbeModeProbe ProbeMode = 1
	// ProbeMode_ProbeModeAssume specifies that filesystem behavior should be
	// assumed based on the underlying platform. This is not as accurate as
	// ProbeMode_ProbeModeProbe.
	ProbeMode_ProbeModeAssume ProbeMode = 2
)

var ProbeMode_name = map[int32]string{
	0: "ProbeModeDefault",
	1: "ProbeModeProbe",
	2: "ProbeModeAssume",
}

var ProbeMode_value = map[string]int32{
	"ProbeModeDefault": 0,
	"ProbeModeProbe":   1,
	"ProbeModeAssume":  2,
}

func (x ProbeMode) String() string {
	return proto.EnumName(ProbeMode_name, int32(x))
}

func (ProbeMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e720967b497c2121, []int{0}
}

func init() {
	proto.RegisterEnum("behavior.ProbeMode", ProbeMode_name, ProbeMode_value)
}

func init() { proto.RegisterFile("filesystem/behavior/probe.proto", fileDescriptor_e720967b497c2121) }

var fileDescriptor_e720967b497c2121 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcb, 0xcc, 0x49,
	0x2d, 0xae, 0x2c, 0x2e, 0x49, 0xcd, 0xd5, 0x4f, 0x4a, 0xcd, 0x48, 0x2c, 0xcb, 0xcc, 0x2f, 0xd2,
	0x2f, 0x28, 0xca, 0x4f, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x89, 0x6a,
	0x79, 0x71, 0x71, 0x06, 0x80, 0x24, 0x7c, 0xf3, 0x53, 0x52, 0x85, 0x44, 0xb8, 0x04, 0xe0, 0x1c,
	0x97, 0xd4, 0xb4, 0xc4, 0xd2, 0x9c, 0x12, 0x01, 0x06, 0x21, 0x21, 0x2e, 0x3e, 0xb8, 0x28, 0x98,
	0x21, 0xc0, 0x28, 0x24, 0xcc, 0xc5, 0x0f, 0x17, 0x73, 0x2c, 0x2e, 0x2e, 0xcd, 0x4d, 0x15, 0x60,
	0x72, 0x32, 0x8d, 0x32, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf,
	0x48, 0x2c, 0xcb, 0x4f, 0xd6, 0xcd, 0xcc, 0xd7, 0xcf, 0x2d, 0x2d, 0x49, 0x4c, 0x4f, 0xcd, 0xd3,
	0x2f, 0xc8, 0x4e, 0xd7, 0xc7, 0xe2, 0xb0, 0x24, 0x36, 0xb0, 0x9b, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xb5, 0x4f, 0x8d, 0x9a, 0xb6, 0x00, 0x00, 0x00,
}
