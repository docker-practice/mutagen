// Code generated by protoc-gen-go. DO NOT EDIT.
// source: forwarding/session.proto

package forwarding

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	url "github.com/mutagen-io/mutagen/pkg/url"
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

type Session struct {
	// Identifier is the (unique) session identifier. It is static. It cannot be
	// empty.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Version is the session version. It is static.
	Version Version `protobuf:"varint,2,opt,name=version,proto3,enum=forwarding.Version" json:"version,omitempty"`
	// CreationTime is the creation time of the session. It is static. It cannot
	// be nil.
	CreationTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	// CreatingVersionMajor is the major version component of the version of
	// Mutagen which created the session. It is static.
	CreatingVersionMajor uint32 `protobuf:"varint,4,opt,name=creatingVersionMajor,proto3" json:"creatingVersionMajor,omitempty"`
	// CreatingVersionMinor is the minor version component of the version of
	// Mutagen which created the session. It is static.
	CreatingVersionMinor uint32 `protobuf:"varint,5,opt,name=creatingVersionMinor,proto3" json:"creatingVersionMinor,omitempty"`
	// CreatingVersionPatch is the patch version component of the version of
	// Mutagen which created the session. It is static.
	CreatingVersionPatch uint32 `protobuf:"varint,6,opt,name=creatingVersionPatch,proto3" json:"creatingVersionPatch,omitempty"`
	// Source is the source endpoint URL. It is static. It cannot be nil.
	Source *url.URL `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	// Destination is the destination endpoint URL. It is static. It cannot be
	// nil.
	Destination *url.URL `protobuf:"bytes,8,opt,name=destination,proto3" json:"destination,omitempty"`
	// Configuration is the flattened session configuration. It is static. It
	// cannot be nil.
	Configuration *Configuration `protobuf:"bytes,9,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// ConfigurationSource are the source-specific session configuration
	// overrides. It is static.
	ConfigurationSource *Configuration `protobuf:"bytes,10,opt,name=configurationSource,proto3" json:"configurationSource,omitempty"`
	// ConfigurationDestination are the destination-specific session
	// configuration overrides. It is static.
	ConfigurationDestination *Configuration `protobuf:"bytes,11,opt,name=configurationDestination,proto3" json:"configurationDestination,omitempty"`
	// Labels are the session labels. They are static.
	Labels map[string]string `protobuf:"bytes,12,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Paused indicates whether or not the session is marked as paused.
	Paused               bool     `protobuf:"varint,13,opt,name=paused,proto3" json:"paused,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_017cb915c0809617, []int{0}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Session) GetVersion() Version {
	if m != nil {
		return m.Version
	}
	return Version_Invalid
}

func (m *Session) GetCreationTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Session) GetCreatingVersionMajor() uint32 {
	if m != nil {
		return m.CreatingVersionMajor
	}
	return 0
}

func (m *Session) GetCreatingVersionMinor() uint32 {
	if m != nil {
		return m.CreatingVersionMinor
	}
	return 0
}

func (m *Session) GetCreatingVersionPatch() uint32 {
	if m != nil {
		return m.CreatingVersionPatch
	}
	return 0
}

func (m *Session) GetSource() *url.URL {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Session) GetDestination() *url.URL {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *Session) GetConfiguration() *Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *Session) GetConfigurationSource() *Configuration {
	if m != nil {
		return m.ConfigurationSource
	}
	return nil
}

func (m *Session) GetConfigurationDestination() *Configuration {
	if m != nil {
		return m.ConfigurationDestination
	}
	return nil
}

func (m *Session) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Session) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

func init() {
	proto.RegisterType((*Session)(nil), "forwarding.Session")
	proto.RegisterMapType((map[string]string)(nil), "forwarding.Session.LabelsEntry")
}

func init() { proto.RegisterFile("forwarding/session.proto", fileDescriptor_017cb915c0809617) }

var fileDescriptor_017cb915c0809617 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x6f, 0xd3, 0x30,
	0x14, 0x57, 0x56, 0x9a, 0xb6, 0x2f, 0x2b, 0x42, 0xde, 0x84, 0x4c, 0x0e, 0x5b, 0xc4, 0x29, 0x42,
	0xcc, 0x91, 0xca, 0x81, 0x8f, 0x03, 0x48, 0x7c, 0x9c, 0x18, 0x12, 0xf2, 0x36, 0x0e, 0xdc, 0xdc,
	0xd4, 0xcd, 0xcc, 0x12, 0xbb, 0x72, 0xec, 0xa1, 0xfd, 0xdd, 0xfc, 0x03, 0xa8, 0x76, 0x4a, 0x1d,
	0x94, 0x69, 0xb7, 0xbc, 0xf7, 0xfb, 0x78, 0xef, 0xe7, 0x3c, 0xc0, 0x6b, 0xa5, 0x7f, 0x33, 0xbd,
	0x12, 0xb2, 0x2a, 0x5a, 0xde, 0xb6, 0x42, 0x49, 0xb2, 0xd1, 0xca, 0x28, 0x04, 0x7b, 0x24, 0x3d,
	0xad, 0x94, 0xaa, 0x6a, 0x5e, 0x38, 0x64, 0x69, 0xd7, 0x85, 0x11, 0x0d, 0x6f, 0x0d, 0x6b, 0x36,
	0x9e, 0x9c, 0x9e, 0x04, 0x36, 0xa5, 0x92, 0x6b, 0x51, 0x59, 0xcd, 0xcc, 0x3f, 0xb3, 0x34, 0x1c,
	0x73, 0xcb, 0xf5, 0x7e, 0x4c, 0x3a, 0xb7, 0xba, 0x2e, 0xac, 0xae, 0x7d, 0xf9, 0xfc, 0xcf, 0x18,
	0x26, 0x17, 0x7e, 0x0f, 0x74, 0x02, 0x20, 0x56, 0x5c, 0x1a, 0xb1, 0x16, 0x5c, 0xe3, 0x28, 0x8b,
	0xf2, 0x19, 0x0d, 0x3a, 0xe8, 0x0c, 0x26, 0x9d, 0x17, 0x3e, 0xc8, 0xa2, 0xfc, 0xf1, 0xe2, 0x88,
	0xec, 0xc7, 0x90, 0x1f, 0x1e, 0xa2, 0x3b, 0x0e, 0x7a, 0x0f, 0x87, 0xa5, 0xe6, 0x6e, 0xab, 0x4b,
	0xd1, 0x70, 0x3c, 0xca, 0xa2, 0x3c, 0x59, 0xa4, 0xc4, 0x67, 0x23, 0xbb, 0x6c, 0xe4, 0x72, 0x97,
	0x8d, 0xf6, 0xf8, 0x68, 0x01, 0xc7, 0xbe, 0x96, 0x55, 0xe7, 0xfd, 0x8d, 0xfd, 0x52, 0x1a, 0x3f,
	0xca, 0xa2, 0x7c, 0x4e, 0x07, 0xb1, 0x21, 0x8d, 0x90, 0x4a, 0xe3, 0xf1, 0xb0, 0x66, 0x8b, 0x0d,
	0x68, 0xbe, 0x33, 0x53, 0x5e, 0xe3, 0x78, 0x50, 0xe3, 0x30, 0x94, 0x41, 0xdc, 0x2a, 0xab, 0x4b,
	0x8e, 0x27, 0x2e, 0xd5, 0x94, 0x6c, 0x9f, 0xf4, 0x8a, 0x9e, 0xd3, 0xae, 0x8f, 0x5e, 0x40, 0xb2,
	0xe2, 0xad, 0x11, 0xd2, 0x05, 0xc2, 0xd3, 0xff, 0x68, 0x21, 0x88, 0x3e, 0xc0, 0xbc, 0xf7, 0x13,
	0xf1, 0xcc, 0xb1, 0x9f, 0x85, 0xcf, 0xfb, 0x29, 0x24, 0xd0, 0x3e, 0x1f, 0x7d, 0x85, 0xa3, 0x5e,
	0xe3, 0xc2, 0xef, 0x06, 0x0f, 0xd9, 0x0c, 0xa9, 0xd0, 0x15, 0xe0, 0x5e, 0xfb, 0x73, 0x10, 0x23,
	0x79, 0xc8, 0xf1, 0x5e, 0x29, 0x7a, 0x0d, 0x71, 0xcd, 0x96, 0xbc, 0x6e, 0xf1, 0x61, 0x36, 0xca,
	0x93, 0xc5, 0x69, 0x68, 0xd2, 0x9d, 0x20, 0x39, 0x77, 0x8c, 0x2f, 0xd2, 0xe8, 0x3b, 0xda, 0xd1,
	0xd1, 0x53, 0x88, 0x37, 0xcc, 0xb6, 0x7c, 0x85, 0xe7, 0x59, 0x94, 0x4f, 0x69, 0x57, 0xa5, 0x6f,
	0x21, 0x09, 0xe8, 0xe8, 0x09, 0x8c, 0x6e, 0xf8, 0x5d, 0x77, 0xb6, 0xdb, 0x4f, 0x74, 0x0c, 0xe3,
	0x5b, 0x56, 0x5b, 0xee, 0xae, 0x75, 0x46, 0x7d, 0xf1, 0xee, 0xe0, 0x4d, 0xf4, 0x91, 0xfc, 0x7c,
	0x59, 0x09, 0x73, 0x6d, 0x97, 0xa4, 0x54, 0x4d, 0xd1, 0x58, 0xc3, 0x2a, 0x2e, 0xcf, 0x84, 0xda,
	0x7d, 0x16, 0x9b, 0x9b, 0xaa, 0xd8, 0xaf, 0xb7, 0x8c, 0xdd, 0xb1, 0xbe, 0xfa, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x42, 0xed, 0xe6, 0xc9, 0xbe, 0x03, 0x00, 0x00,
}
