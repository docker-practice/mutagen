// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session/configuration.proto

package session

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	behavior "github.com/havoc-io/mutagen/pkg/filesystem/behavior"
	sync "github.com/havoc-io/mutagen/pkg/sync"
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

// Configuration encodes session configuration parameters. It is used for create
// commands to specify configuration options, for loading global configuration
// options, and for storing a merged configuration inside sessions.
type Configuration struct {
	// SynchronizationMode specifies the synchronization mode that should be
	// used in synchronization.
	SynchronizationMode sync.SynchronizationMode `protobuf:"varint,11,opt,name=synchronizationMode,proto3,enum=sync.SynchronizationMode" json:"synchronizationMode,omitempty"`
	// MaximumEntryCount specifies the maximum number of filesystem entries that
	// endpoints will tolerate managing. A zero value indicates no limit.
	MaximumEntryCount uint64 `protobuf:"varint,12,opt,name=maximumEntryCount,proto3" json:"maximumEntryCount,omitempty"`
	// MaximumStagingFileSize is the maximum (individual) file size that
	// endpoints will stage. A zero value indicates no limit.
	MaximumStagingFileSize uint64 `protobuf:"varint,13,opt,name=maximumStagingFileSize,proto3" json:"maximumStagingFileSize,omitempty"`
	// ProbeMode specifies the filesystem probing mode.
	ProbeMode behavior.ProbeMode `protobuf:"varint,14,opt,name=probeMode,proto3,enum=behavior.ProbeMode" json:"probeMode,omitempty"`
	// StageMode specifies the file staging mode.
	StageMode StageMode `protobuf:"varint,16,opt,name=stageMode,proto3,enum=session.StageMode" json:"stageMode,omitempty"`
	// SymlinkMode specifies the symlink mode that should be used in
	// synchronization.
	SymlinkMode sync.SymlinkMode `protobuf:"varint,1,opt,name=symlinkMode,proto3,enum=sync.SymlinkMode" json:"symlinkMode,omitempty"`
	// WatchMode specifies the filesystem watching mode.
	WatchMode WatchMode `protobuf:"varint,21,opt,name=watchMode,proto3,enum=session.WatchMode" json:"watchMode,omitempty"`
	// WatchPollingInterval specifies the interval (in seconds) for poll-based
	// file monitoring. A value of 0 specifies that the default interval should
	// be used.
	WatchPollingInterval uint32 `protobuf:"varint,22,opt,name=watchPollingInterval,proto3" json:"watchPollingInterval,omitempty"`
	// DefaultIgnores specifies the ignore patterns brought in from the global
	// configuration.
	// DEPRECATED: This field is no longer used when loading from global
	// configuration. Instead, ignores provided by global configuration are
	// simply merged into the ignore list of the main configuration. However,
	// older sessions still use this field.
	DefaultIgnores []string `protobuf:"bytes,31,rep,name=defaultIgnores,proto3" json:"defaultIgnores,omitempty"`
	// Ignores specifies the ignore patterns brought in from the create request.
	Ignores []string `protobuf:"bytes,32,rep,name=ignores,proto3" json:"ignores,omitempty"`
	// IgnoreVCSMode specifies the VCS ignore mode that should be used in
	// synchronization.
	IgnoreVCSMode sync.IgnoreVCSMode `protobuf:"varint,33,opt,name=ignoreVCSMode,proto3,enum=sync.IgnoreVCSMode" json:"ignoreVCSMode,omitempty"`
	// DefaultFileMode specifies the default permission mode to use for new
	// files in "portable" permission propagation mode.
	DefaultFileMode uint32 `protobuf:"varint,63,opt,name=defaultFileMode,proto3" json:"defaultFileMode,omitempty"`
	// DefaultDirectoryMode specifies the default permission mode to use for new
	// files in "portable" permission propagation mode.
	DefaultDirectoryMode uint32 `protobuf:"varint,64,opt,name=defaultDirectoryMode,proto3" json:"defaultDirectoryMode,omitempty"`
	// DefaultOwner specifies the default owner identifier to use when setting
	// ownership of new files and directories in "portable" permission
	// propagation mode.
	DefaultOwner string `protobuf:"bytes,65,opt,name=defaultOwner,proto3" json:"defaultOwner,omitempty"`
	// DefaultGroup specifies the default group identifier to use when setting
	// ownership of new files and directories in "portable" permission
	// propagation mode.
	DefaultGroup         string   `protobuf:"bytes,66,opt,name=defaultGroup,proto3" json:"defaultGroup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_34b104152623cac3, []int{0}
}

func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetSynchronizationMode() sync.SynchronizationMode {
	if m != nil {
		return m.SynchronizationMode
	}
	return sync.SynchronizationMode_SynchronizationModeDefault
}

func (m *Configuration) GetMaximumEntryCount() uint64 {
	if m != nil {
		return m.MaximumEntryCount
	}
	return 0
}

func (m *Configuration) GetMaximumStagingFileSize() uint64 {
	if m != nil {
		return m.MaximumStagingFileSize
	}
	return 0
}

func (m *Configuration) GetProbeMode() behavior.ProbeMode {
	if m != nil {
		return m.ProbeMode
	}
	return behavior.ProbeMode_ProbeModeDefault
}

func (m *Configuration) GetStageMode() StageMode {
	if m != nil {
		return m.StageMode
	}
	return StageMode_StageModeDefault
}

func (m *Configuration) GetSymlinkMode() sync.SymlinkMode {
	if m != nil {
		return m.SymlinkMode
	}
	return sync.SymlinkMode_SymlinkModeDefault
}

func (m *Configuration) GetWatchMode() WatchMode {
	if m != nil {
		return m.WatchMode
	}
	return WatchMode_WatchModeDefault
}

func (m *Configuration) GetWatchPollingInterval() uint32 {
	if m != nil {
		return m.WatchPollingInterval
	}
	return 0
}

func (m *Configuration) GetDefaultIgnores() []string {
	if m != nil {
		return m.DefaultIgnores
	}
	return nil
}

func (m *Configuration) GetIgnores() []string {
	if m != nil {
		return m.Ignores
	}
	return nil
}

func (m *Configuration) GetIgnoreVCSMode() sync.IgnoreVCSMode {
	if m != nil {
		return m.IgnoreVCSMode
	}
	return sync.IgnoreVCSMode_IgnoreVCSModeDefault
}

func (m *Configuration) GetDefaultFileMode() uint32 {
	if m != nil {
		return m.DefaultFileMode
	}
	return 0
}

func (m *Configuration) GetDefaultDirectoryMode() uint32 {
	if m != nil {
		return m.DefaultDirectoryMode
	}
	return 0
}

func (m *Configuration) GetDefaultOwner() string {
	if m != nil {
		return m.DefaultOwner
	}
	return ""
}

func (m *Configuration) GetDefaultGroup() string {
	if m != nil {
		return m.DefaultGroup
	}
	return ""
}

func init() {
	proto.RegisterType((*Configuration)(nil), "session.Configuration")
}

func init() { proto.RegisterFile("session/configuration.proto", fileDescriptor_34b104152623cac3) }

var fileDescriptor_34b104152623cac3 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xdd, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0x15, 0xf1, 0x31, 0xd5, 0x5b, 0x3b, 0xea, 0xc2, 0x14, 0xc6, 0xc3, 0xc2, 0x1e, 0x20,
	0x48, 0x90, 0xc0, 0x26, 0x21, 0xf1, 0x04, 0xac, 0x7c, 0xa8, 0x42, 0x88, 0xc9, 0x91, 0x40, 0xe2,
	0x2d, 0x4d, 0xdd, 0xc4, 0x5a, 0xe2, 0x5b, 0xd9, 0x4e, 0x47, 0xf6, 0xff, 0xf2, 0x7f, 0xa0, 0xdc,
	0xda, 0xf4, 0x73, 0x6f, 0xf5, 0x39, 0xbf, 0x53, 0x5f, 0x9f, 0xdb, 0x92, 0x27, 0x9a, 0x6b, 0x2d,
	0x40, 0xc6, 0x19, 0xc8, 0xa9, 0xc8, 0x6b, 0x95, 0x1a, 0x01, 0x32, 0x9a, 0x29, 0x30, 0x40, 0xf7,
	0xac, 0x79, 0x7c, 0x32, 0x15, 0x25, 0xd7, 0x8d, 0x36, 0xbc, 0x8a, 0xc7, 0xbc, 0x48, 0xe7, 0x02,
	0x54, 0x3c, 0x53, 0x30, 0xe6, 0x0b, 0xf2, 0x78, 0xe0, 0xbe, 0x46, 0x9b, 0x34, 0xdf, 0x12, 0xaf,
	0x53, 0x93, 0x15, 0x56, 0xec, 0xeb, 0x46, 0x66, 0xb1, 0xc8, 0x25, 0x28, 0xc7, 0x1d, 0xa2, 0x54,
	0xc1, 0xc4, 0x09, 0x14, 0x05, 0xdd, 0x54, 0xa5, 0x90, 0x57, 0x0b, 0xed, 0xf4, 0xef, 0x3d, 0xd2,
	0x1d, 0xae, 0xce, 0x48, 0xbf, 0x91, 0x41, 0xcb, 0x15, 0x0a, 0xa4, 0xb8, 0x41, 0xe9, 0x3b, 0x4c,
	0xb8, 0xbf, 0x1f, 0x78, 0x61, 0xef, 0xec, 0x71, 0xd4, 0x7a, 0x51, 0xb2, 0x0d, 0xb0, 0x5d, 0x29,
	0xfa, 0x92, 0xf4, 0xab, 0xf4, 0x8f, 0xa8, 0xea, 0xea, 0xb3, 0x34, 0xaa, 0x19, 0x42, 0x2d, 0x8d,
	0x7f, 0x10, 0x78, 0xe1, 0x5d, 0xb6, 0x6d, 0xd0, 0xb7, 0xe4, 0xc8, 0x8a, 0x89, 0x49, 0x73, 0x21,
	0xf3, 0x2f, 0xa2, 0xe4, 0x89, 0xb8, 0xe1, 0x7e, 0x17, 0x23, 0xb7, 0xb8, 0xf4, 0x0d, 0xe9, 0x60,
	0x6b, 0x38, 0x68, 0x0f, 0x07, 0x1d, 0x44, 0xae, 0xd0, 0xe8, 0xd2, 0x59, 0x6c, 0x49, 0xd1, 0xd7,
	0xa4, 0x83, 0x9d, 0x62, 0xe4, 0x01, 0x46, 0x68, 0x64, 0x8b, 0x8d, 0x12, 0xe7, 0xb0, 0x25, 0x44,
	0xcf, 0xc9, 0xbe, 0xad, 0x0e, 0x33, 0x1e, 0x66, 0xfa, 0xae, 0x8f, 0xff, 0x06, 0x5b, 0xa5, 0xda,
	0x6b, 0x70, 0x4b, 0x18, 0x79, 0xb4, 0x71, 0xcd, 0x2f, 0xe7, 0xb0, 0x25, 0x44, 0xcf, 0xc8, 0x43,
	0x3c, 0x5c, 0x42, 0x59, 0x0a, 0x99, 0x8f, 0xa4, 0xe1, 0x6a, 0x9e, 0x96, 0xfe, 0x51, 0xe0, 0x85,
	0x5d, 0xb6, 0xd3, 0xa3, 0xcf, 0x48, 0x6f, 0xc2, 0xa7, 0x69, 0x5d, 0x9a, 0x11, 0xfe, 0x00, 0xb4,
	0x7f, 0x12, 0xdc, 0x09, 0x3b, 0x6c, 0x43, 0xa5, 0x3e, 0xd9, 0x13, 0x16, 0x08, 0x10, 0x70, 0x47,
	0xfa, 0x8e, 0x74, 0x17, 0x1f, 0x7f, 0x0e, 0x13, 0x9c, 0xf5, 0xa9, 0x6d, 0x11, 0x9f, 0x37, 0x5a,
	0xb5, 0xd8, 0x3a, 0x49, 0x43, 0x72, 0x68, 0xaf, 0x69, 0xf7, 0x81, 0xe1, 0xf7, 0x38, 0xeb, 0xa6,
	0xdc, 0x3e, 0xcd, 0x4a, 0x9f, 0x84, 0xe2, 0x99, 0x01, 0xd5, 0x20, 0xfe, 0x61, 0xf1, 0xb4, 0x5d,
	0x1e, 0x3d, 0x25, 0x07, 0x56, 0xff, 0x71, 0x2d, 0xb9, 0xf2, 0x3f, 0x06, 0x5e, 0xd8, 0x61, 0x6b,
	0xda, 0x0a, 0xf3, 0x55, 0x41, 0x3d, 0xf3, 0x2f, 0xd6, 0x18, 0xd4, 0x2e, 0x5e, 0xfc, 0x7e, 0x9e,
	0x0b, 0x53, 0xd4, 0xe3, 0x28, 0x83, 0x2a, 0x2e, 0xd2, 0x39, 0x64, 0xaf, 0x04, 0xc4, 0x55, 0xdd,
	0xee, 0x57, 0xc6, 0xb3, 0xab, 0x3c, 0xb6, 0x6b, 0x19, 0xdf, 0xc7, 0x7f, 0xc6, 0xf9, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x60, 0x9e, 0xc4, 0x8f, 0xc4, 0x03, 0x00, 0x00,
}
