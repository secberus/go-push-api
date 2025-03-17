//
// Copyright 2018-2025 Secberus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: types/v1/table.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TableSyncType int32

const (
	TableSyncType_TABLE_SYNC_TYPE_UNSPECIFIED TableSyncType = 0
	TableSyncType_TABLE_SYNC_TYPE_TRUNCATE    TableSyncType = 1
	TableSyncType_TABLE_SYNC_TYPE_APPEND      TableSyncType = 2
)

// Enum value maps for TableSyncType.
var (
	TableSyncType_name = map[int32]string{
		0: "TABLE_SYNC_TYPE_UNSPECIFIED",
		1: "TABLE_SYNC_TYPE_TRUNCATE",
		2: "TABLE_SYNC_TYPE_APPEND",
	}
	TableSyncType_value = map[string]int32{
		"TABLE_SYNC_TYPE_UNSPECIFIED": 0,
		"TABLE_SYNC_TYPE_TRUNCATE":    1,
		"TABLE_SYNC_TYPE_APPEND":      2,
	}
)

func (x TableSyncType) Enum() *TableSyncType {
	p := new(TableSyncType)
	*p = x
	return p
}

func (x TableSyncType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TableSyncType) Descriptor() protoreflect.EnumDescriptor {
	return file_types_v1_table_proto_enumTypes[0].Descriptor()
}

func (TableSyncType) Type() protoreflect.EnumType {
	return &file_types_v1_table_proto_enumTypes[0]
}

func (x TableSyncType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TableSyncType.Descriptor instead.
func (TableSyncType) EnumDescriptor() ([]byte, []int) {
	return file_types_v1_table_proto_rawDescGZIP(), []int{0}
}

// CreateTableInput is the input parameters for creating a new table in the Secberus data warehouse
type Table struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SyncType      TableSyncType          `protobuf:"varint,2,opt,name=sync_type,json=syncType,proto3,enum=types.v1.TableSyncType" json:"sync_type,omitempty"`
	Columns       []*Column              `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Table) Reset() {
	*x = Table{}
	mi := &file_types_v1_table_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_table_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_types_v1_table_proto_rawDescGZIP(), []int{0}
}

func (x *Table) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Table) GetSyncType() TableSyncType {
	if x != nil {
		return x.SyncType
	}
	return TableSyncType_TABLE_SYNC_TYPE_UNSPECIFIED
}

func (x *Table) GetColumns() []*Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

// TableStats is a read-only message sent from the server containing table metrics.
type TableStats struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TotalSize     int64                  `protobuf:"varint,1,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	RecordCount   int64                  `protobuf:"varint,2,opt,name=record_count,json=recordCount,proto3" json:"record_count,omitempty"`
	LastSyncTime  int64                  `protobuf:"varint,4,opt,name=last_sync_time,json=lastSyncTime,proto3" json:"last_sync_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TableStats) Reset() {
	*x = TableStats{}
	mi := &file_types_v1_table_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TableStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableStats) ProtoMessage() {}

func (x *TableStats) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_table_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableStats.ProtoReflect.Descriptor instead.
func (*TableStats) Descriptor() ([]byte, []int) {
	return file_types_v1_table_proto_rawDescGZIP(), []int{1}
}

func (x *TableStats) GetTotalSize() int64 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *TableStats) GetRecordCount() int64 {
	if x != nil {
		return x.RecordCount
	}
	return 0
}

func (x *TableStats) GetLastSyncTime() int64 {
	if x != nil {
		return x.LastSyncTime
	}
	return 0
}

var File_types_v1_table_proto protoreflect.FileDescriptor

var file_types_v1_table_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x15, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1d,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0x48,
	0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x3f, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a,
	0x09, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22,
	0x74, 0x0a, 0x0a, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x79, 0x6e,
	0x63, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x6a, 0x0a, 0x0d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x79,
	0x6e, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f,
	0x53, 0x59, 0x4e, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x41, 0x42, 0x4c, 0x45,
	0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x55, 0x4e, 0x43,
	0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x53,
	0x59, 0x4e, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x45, 0x4e, 0x44, 0x10,
	0x02, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x65, 0x63, 0x62, 0x65, 0x72, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x75, 0x73, 0x68,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_types_v1_table_proto_rawDescOnce sync.Once
	file_types_v1_table_proto_rawDescData []byte
)

func file_types_v1_table_proto_rawDescGZIP() []byte {
	file_types_v1_table_proto_rawDescOnce.Do(func() {
		file_types_v1_table_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_types_v1_table_proto_rawDesc), len(file_types_v1_table_proto_rawDesc)))
	})
	return file_types_v1_table_proto_rawDescData
}

var file_types_v1_table_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_v1_table_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_types_v1_table_proto_goTypes = []any{
	(TableSyncType)(0), // 0: types.v1.TableSyncType
	(*Table)(nil),      // 1: types.v1.Table
	(*TableStats)(nil), // 2: types.v1.TableStats
	(*Column)(nil),     // 3: types.v1.Column
}
var file_types_v1_table_proto_depIdxs = []int32{
	0, // 0: types.v1.Table.sync_type:type_name -> types.v1.TableSyncType
	3, // 1: types.v1.Table.columns:type_name -> types.v1.Column
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_types_v1_table_proto_init() }
func file_types_v1_table_proto_init() {
	if File_types_v1_table_proto != nil {
		return
	}
	file_types_v1_column_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_types_v1_table_proto_rawDesc), len(file_types_v1_table_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_table_proto_goTypes,
		DependencyIndexes: file_types_v1_table_proto_depIdxs,
		EnumInfos:         file_types_v1_table_proto_enumTypes,
		MessageInfos:      file_types_v1_table_proto_msgTypes,
	}.Build()
	File_types_v1_table_proto = out.File
	file_types_v1_table_proto_goTypes = nil
	file_types_v1_table_proto_depIdxs = nil
}
