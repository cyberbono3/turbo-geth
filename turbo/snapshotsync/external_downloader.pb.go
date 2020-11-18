// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: external_downloader.proto

package snapshotsync

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SnapshotType int32

const (
	SnapshotType_headers  SnapshotType = 0
	SnapshotType_bodies   SnapshotType = 1
	SnapshotType_state    SnapshotType = 2
	SnapshotType_receipts SnapshotType = 3
)

// Enum value maps for SnapshotType.
var (
	SnapshotType_name = map[int32]string{
		0: "headers",
		1: "bodies",
		2: "state",
		3: "receipts",
	}
	SnapshotType_value = map[string]int32{
		"headers":  0,
		"bodies":   1,
		"state":    2,
		"receipts": 3,
	}
)

func (x SnapshotType) Enum() *SnapshotType {
	p := new(SnapshotType)
	*p = x
	return p
}

func (x SnapshotType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SnapshotType) Descriptor() protoreflect.EnumDescriptor {
	return file_external_downloader_proto_enumTypes[0].Descriptor()
}

func (SnapshotType) Type() protoreflect.EnumType {
	return &file_external_downloader_proto_enumTypes[0]
}

func (x SnapshotType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SnapshotType.Descriptor instead.
func (SnapshotType) EnumDescriptor() ([]byte, []int) {
	return file_external_downloader_proto_rawDescGZIP(), []int{0}
}

type DownloadSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId uint64         `protobuf:"varint,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	Type      []SnapshotType `protobuf:"varint,2,rep,packed,name=type,proto3,enum=snapshotsync.SnapshotType" json:"type,omitempty"`
}

func (x *DownloadSnapshotRequest) Reset() {
	*x = DownloadSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_downloader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadSnapshotRequest) ProtoMessage() {}

func (x *DownloadSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_downloader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadSnapshotRequest.ProtoReflect.Descriptor instead.
func (*DownloadSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_external_downloader_proto_rawDescGZIP(), []int{0}
}

func (x *DownloadSnapshotRequest) GetNetworkId() uint64 {
	if x != nil {
		return x.NetworkId
	}
	return 0
}

func (x *DownloadSnapshotRequest) GetType() []SnapshotType {
	if x != nil {
		return x.Type
	}
	return nil
}

type SnapshotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId uint64 `protobuf:"varint,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
}

func (x *SnapshotsRequest) Reset() {
	*x = SnapshotsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_downloader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotsRequest) ProtoMessage() {}

func (x *SnapshotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_downloader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotsRequest.ProtoReflect.Descriptor instead.
func (*SnapshotsRequest) Descriptor() ([]byte, []int) {
	return file_external_downloader_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotsRequest) GetNetworkId() uint64 {
	if x != nil {
		return x.NetworkId
	}
	return 0
}

type SnapshotsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          SnapshotType `protobuf:"varint,1,opt,name=type,proto3,enum=snapshotsync.SnapshotType" json:"type,omitempty"`
	GotInfoByte   bool         `protobuf:"varint,2,opt,name=gotInfoByte,proto3" json:"gotInfoByte,omitempty"`
	Readiness     int32        `protobuf:"varint,3,opt,name=readiness,proto3" json:"readiness,omitempty"`
	SnapshotBlock uint64       `protobuf:"varint,4,opt,name=snapshotBlock,proto3" json:"snapshotBlock,omitempty"`
	Dbpath        string       `protobuf:"bytes,5,opt,name=dbpath,proto3" json:"dbpath,omitempty"`
}

func (x *SnapshotsInfo) Reset() {
	*x = SnapshotsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_downloader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotsInfo) ProtoMessage() {}

func (x *SnapshotsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_external_downloader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotsInfo.ProtoReflect.Descriptor instead.
func (*SnapshotsInfo) Descriptor() ([]byte, []int) {
	return file_external_downloader_proto_rawDescGZIP(), []int{2}
}

func (x *SnapshotsInfo) GetType() SnapshotType {
	if x != nil {
		return x.Type
	}
	return SnapshotType_headers
}

func (x *SnapshotsInfo) GetGotInfoByte() bool {
	if x != nil {
		return x.GotInfoByte
	}
	return false
}

func (x *SnapshotsInfo) GetReadiness() int32 {
	if x != nil {
		return x.Readiness
	}
	return 0
}

func (x *SnapshotsInfo) GetSnapshotBlock() uint64 {
	if x != nil {
		return x.SnapshotBlock
	}
	return 0
}

func (x *SnapshotsInfo) GetDbpath() string {
	if x != nil {
		return x.Dbpath
	}
	return ""
}

type SnapshotsInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*SnapshotsInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *SnapshotsInfoReply) Reset() {
	*x = SnapshotsInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_downloader_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotsInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotsInfoReply) ProtoMessage() {}

func (x *SnapshotsInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_external_downloader_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotsInfoReply.ProtoReflect.Descriptor instead.
func (*SnapshotsInfoReply) Descriptor() ([]byte, []int) {
	return file_external_downloader_proto_rawDescGZIP(), []int{3}
}

func (x *SnapshotsInfoReply) GetInfo() []*SnapshotsInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_external_downloader_proto protoreflect.FileDescriptor

var file_external_downloader_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x79, 0x6e, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x17, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x31, 0x0a, 0x10, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x64, 0x22, 0xbd, 0x01, 0x0a, 0x0d, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x79,
	0x6e, 0x63, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x79, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x67, 0x6f, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x79, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x73, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x62, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x70,
	0x61, 0x74, 0x68, 0x22, 0x45, 0x0a, 0x12, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x2a, 0x40, 0x0a, 0x0c, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x62, 0x6f, 0x64, 0x69, 0x65,
	0x73, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x10, 0x03, 0x32, 0xaa, 0x01, 0x0a,
	0x0a, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x08, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x25, 0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x09, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x3b,
	0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x79, 0x6e, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_downloader_proto_rawDescOnce sync.Once
	file_external_downloader_proto_rawDescData = file_external_downloader_proto_rawDesc
)

func file_external_downloader_proto_rawDescGZIP() []byte {
	file_external_downloader_proto_rawDescOnce.Do(func() {
		file_external_downloader_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_downloader_proto_rawDescData)
	})
	return file_external_downloader_proto_rawDescData
}

var file_external_downloader_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_external_downloader_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_external_downloader_proto_goTypes = []interface{}{
	(SnapshotType)(0),               // 0: snapshotsync.SnapshotType
	(*DownloadSnapshotRequest)(nil), // 1: snapshotsync.DownloadSnapshotRequest
	(*SnapshotsRequest)(nil),        // 2: snapshotsync.SnapshotsRequest
	(*SnapshotsInfo)(nil),           // 3: snapshotsync.SnapshotsInfo
	(*SnapshotsInfoReply)(nil),      // 4: snapshotsync.SnapshotsInfoReply
	(*empty.Empty)(nil),             // 5: google.protobuf.Empty
}
var file_external_downloader_proto_depIdxs = []int32{
	0, // 0: snapshotsync.DownloadSnapshotRequest.type:type_name -> snapshotsync.SnapshotType
	0, // 1: snapshotsync.SnapshotsInfo.type:type_name -> snapshotsync.SnapshotType
	3, // 2: snapshotsync.SnapshotsInfoReply.info:type_name -> snapshotsync.SnapshotsInfo
	1, // 3: snapshotsync.Downloader.Download:input_type -> snapshotsync.DownloadSnapshotRequest
	2, // 4: snapshotsync.Downloader.Snapshots:input_type -> snapshotsync.SnapshotsRequest
	5, // 5: snapshotsync.Downloader.Download:output_type -> google.protobuf.Empty
	4, // 6: snapshotsync.Downloader.Snapshots:output_type -> snapshotsync.SnapshotsInfoReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_external_downloader_proto_init() }
func file_external_downloader_proto_init() {
	if File_external_downloader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_downloader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadSnapshotRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_downloader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_downloader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotsInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_downloader_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotsInfoReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_downloader_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_downloader_proto_goTypes,
		DependencyIndexes: file_external_downloader_proto_depIdxs,
		EnumInfos:         file_external_downloader_proto_enumTypes,
		MessageInfos:      file_external_downloader_proto_msgTypes,
	}.Build()
	File_external_downloader_proto = out.File
	file_external_downloader_proto_rawDesc = nil
	file_external_downloader_proto_goTypes = nil
	file_external_downloader_proto_depIdxs = nil
}
