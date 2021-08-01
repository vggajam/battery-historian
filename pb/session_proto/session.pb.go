// Code generated by protoc-gen-go.
// source: github.com/vggajam/battery-historian/pb/session_proto/session.proto
// DO NOT EDIT!

/*
Package session is a generated protocol buffer package.

It is generated from these files:
	github.com/vggajam/battery-historian/pb/session_proto/session.proto

It has these top-level messages:
	Checkin
*/
package session

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import usagestats1 "github.com/vggajam/battery-historian/pb/usagestats_proto"
import usagestats "github.com/vggajam/battery-historian/pb/usagestats_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Aggregated checkin stats that we will use to compute deltas.
type Checkin struct {
	AndroidId *int64 `protobuf:"varint,1,opt,name=android_id" json:"android_id,omitempty"`
	// Build fingerprint
	BuildFingerprint *string `protobuf:"bytes,2,opt,name=build_fingerprint" json:"build_fingerprint,omitempty"`
	// Device
	Device *string `protobuf:"bytes,3,opt,name=device" json:"device,omitempty"`
	// Timestamp of measurement
	BucketSnapshotMsec *int64 `protobuf:"varint,4,opt,name=bucket_snapshot_msec" json:"bucket_snapshot_msec,omitempty"`
	// Duration of the collection
	BucketDurationMsec *int64 `protobuf:"varint,5,opt,name=bucket_duration_msec" json:"bucket_duration_msec,omitempty"`
	// Checkin
	Checkin *string `protobuf:"bytes,6,opt,name=checkin" json:"checkin,omitempty"`
	// Installed packages
	Packages []*usagestats.PackageInfo `protobuf:"bytes,7,rep,name=packages" json:"packages,omitempty"`
	// Checkin/OTA groups for the device
	//    e.g. auto.droidfood, auto.googlefood.lmp
	Groups []string `protobuf:"bytes,10,rep,name=groups" json:"groups,omitempty"`
	// System info uploaded directly from the device.
	SystemInfo       *usagestats1.SystemInfo `protobuf:"bytes,8,opt,name=system_info" json:"system_info,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *Checkin) Reset()                    { *m = Checkin{} }
func (m *Checkin) String() string            { return proto.CompactTextString(m) }
func (*Checkin) ProtoMessage()               {}
func (*Checkin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Checkin) GetAndroidId() int64 {
	if m != nil && m.AndroidId != nil {
		return *m.AndroidId
	}
	return 0
}

func (m *Checkin) GetBuildFingerprint() string {
	if m != nil && m.BuildFingerprint != nil {
		return *m.BuildFingerprint
	}
	return ""
}

func (m *Checkin) GetDevice() string {
	if m != nil && m.Device != nil {
		return *m.Device
	}
	return ""
}

func (m *Checkin) GetBucketSnapshotMsec() int64 {
	if m != nil && m.BucketSnapshotMsec != nil {
		return *m.BucketSnapshotMsec
	}
	return 0
}

func (m *Checkin) GetBucketDurationMsec() int64 {
	if m != nil && m.BucketDurationMsec != nil {
		return *m.BucketDurationMsec
	}
	return 0
}

func (m *Checkin) GetCheckin() string {
	if m != nil && m.Checkin != nil {
		return *m.Checkin
	}
	return ""
}

func (m *Checkin) GetPackages() []*usagestats.PackageInfo {
	if m != nil {
		return m.Packages
	}
	return nil
}

func (m *Checkin) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *Checkin) GetSystemInfo() *usagestats1.SystemInfo {
	if m != nil {
		return m.SystemInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Checkin)(nil), "session.Checkin")
}

var fileDescriptor0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0x45, 0xa9, 0xd5, 0x4e, 0x9b, 0x82, 0xe2, 0x20, 0x1a, 0x8b, 0x8b, 0xe2, 0xaa, 0x22, 0xce,
	0x40, 0x3f, 0x41, 0x57, 0xee, 0xc4, 0x7e, 0x40, 0xc8, 0x64, 0xd2, 0xcc, 0xa3, 0x6d, 0x12, 0xf2,
	0x5e, 0x84, 0x7e, 0xbc, 0xe0, 0x34, 0x33, 0xa5, 0x2e, 0x05, 0x97, 0x97, 0x73, 0x73, 0x78, 0x37,
	0xec, 0xd5, 0x00, 0x35, 0xb1, 0x2a, 0x94, 0xdb, 0x95, 0xc6, 0x39, 0xb3, 0xd5, 0x65, 0x25, 0x89,
	0x74, 0xd8, 0xbf, 0x34, 0x80, 0xe4, 0x02, 0x48, 0x5b, 0xfa, 0xaa, 0x44, 0x8d, 0x08, 0xce, 0x0a,
	0x1f, 0x1c, 0xb9, 0x63, 0x2a, 0x52, 0xca, 0xb3, 0x3e, 0xce, 0x56, 0x7f, 0x94, 0x45, 0x94, 0x46,
	0x23, 0x49, 0xc2, 0xde, 0x27, 0x6d, 0x1d, 0x1c, 0xd4, 0xa2, 0x6f, 0x8b, 0x54, 0xe8, 0xec, 0xb3,
	0xcf, 0xff, 0x4a, 0xbd, 0x54, 0x9b, 0x16, 0x09, 0xb0, 0x6b, 0xd7, 0x39, 0x1f, 0xbf, 0x07, 0x2c,
	0x7b, 0x6b, 0xb4, 0xda, 0x80, 0xcd, 0x73, 0xc6, 0x8e, 0x4d, 0xa8, 0xf9, 0x60, 0x3e, 0x58, 0x0c,
	0xf3, 0x7b, 0x76, 0x5d, 0x45, 0xd8, 0xd6, 0x62, 0x0d, 0xd6, 0xe8, 0xe0, 0x03, 0x58, 0xe2, 0x67,
	0x2d, 0x9a, 0xe4, 0x97, 0x6c, 0x54, 0xeb, 0x2f, 0x50, 0x9a, 0x0f, 0x53, 0x7e, 0x60, 0x37, 0x55,
	0x54, 0x1b, 0x4d, 0x02, 0xad, 0xf4, 0xd8, 0x38, 0x12, 0x3b, 0xd4, 0x8a, 0x9f, 0x27, 0xd1, 0x89,
	0xd6, 0x31, 0x48, 0x3a, 0xfc, 0x60, 0xa2, 0x17, 0x89, 0x5e, 0xb1, 0x4c, 0x75, 0x57, 0xf0, 0x51,
	0x92, 0x3d, 0xb1, 0x71, 0x7f, 0x2d, 0xf2, 0x6c, 0x3e, 0x5c, 0x4c, 0x97, 0x77, 0xc5, 0x69, 0x57,
	0xf1, 0xd1, 0xb1, 0xf7, 0x76, 0xc8, 0xe1, 0x0e, 0x13, 0x5c, 0xf4, 0xc8, 0x59, 0x5b, 0x9c, 0xe4,
	0xcf, 0x6c, 0x8a, 0x7b, 0x24, 0xbd, 0x4b, 0x3b, 0xf9, 0xb8, 0xf5, 0x4d, 0x97, 0xb7, 0xbf, 0x5f,
	0xaf, 0x12, 0x3e, 0x3c, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x76, 0x98, 0xfc, 0x1e, 0xf6, 0x01,
	0x00, 0x00,
}
