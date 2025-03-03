//
//Copyright 2021 The Kubernetes Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/grpc/bpfrecorder/api.proto

package api_bpfrecorder

import (
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

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_bpfrecorder_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_bpfrecorder_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_bpfrecorder_api_proto_rawDescGZIP(), []int{0}
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_bpfrecorder_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_bpfrecorder_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_bpfrecorder_api_proto_rawDescGZIP(), []int{1}
}

type ProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProfileRequest) Reset() {
	*x = ProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_bpfrecorder_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileRequest) ProtoMessage() {}

func (x *ProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_bpfrecorder_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileRequest.ProtoReflect.Descriptor instead.
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_bpfrecorder_api_proto_rawDescGZIP(), []int{2}
}

func (x *ProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SyscallsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Syscalls []string `protobuf:"bytes,1,rep,name=syscalls,proto3" json:"syscalls,omitempty"`
	GoArch   string   `protobuf:"bytes,2,opt,name=go_arch,json=goArch,proto3" json:"go_arch,omitempty"`
}

func (x *SyscallsResponse) Reset() {
	*x = SyscallsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_bpfrecorder_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyscallsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyscallsResponse) ProtoMessage() {}

func (x *SyscallsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_bpfrecorder_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyscallsResponse.ProtoReflect.Descriptor instead.
func (*SyscallsResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_bpfrecorder_api_proto_rawDescGZIP(), []int{3}
}

func (x *SyscallsResponse) GetSyscalls() []string {
	if x != nil {
		return x.Syscalls
	}
	return nil
}

func (x *SyscallsResponse) GetGoArch() string {
	if x != nil {
		return x.GoArch
	}
	return ""
}

var File_api_grpc_bpfrecorder_api_proto protoreflect.FileDescriptor

var file_api_grpc_bpfrecorder_api_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x70, 0x66, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x61, 0x70, 0x69, 0x5f, 0x62, 0x70, 0x66, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x63,
	0x61, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x6f, 0x5f, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x6f, 0x41, 0x72, 0x63,
	0x68, 0x32, 0xfc, 0x01, 0x0a, 0x0b, 0x42, 0x70, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x48, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x5f, 0x62, 0x70, 0x66, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x5f,
	0x62, 0x70, 0x66, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x04, 0x53,
	0x74, 0x6f, 0x70, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x62, 0x70, 0x66, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x62, 0x70, 0x66, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x12, 0x53, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73,
	0x46, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x5f, 0x62, 0x70, 0x66, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70,
	0x69, 0x5f, 0x62, 0x70, 0x66, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x79,
	0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x12, 0x5a, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x62, 0x70, 0x66, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_bpfrecorder_api_proto_rawDescOnce sync.Once
	file_api_grpc_bpfrecorder_api_proto_rawDescData = file_api_grpc_bpfrecorder_api_proto_rawDesc
)

func file_api_grpc_bpfrecorder_api_proto_rawDescGZIP() []byte {
	file_api_grpc_bpfrecorder_api_proto_rawDescOnce.Do(func() {
		file_api_grpc_bpfrecorder_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_bpfrecorder_api_proto_rawDescData)
	})
	return file_api_grpc_bpfrecorder_api_proto_rawDescData
}

var file_api_grpc_bpfrecorder_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_grpc_bpfrecorder_api_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),     // 0: api_bpfrecorder.EmptyRequest
	(*EmptyResponse)(nil),    // 1: api_bpfrecorder.EmptyResponse
	(*ProfileRequest)(nil),   // 2: api_bpfrecorder.ProfileRequest
	(*SyscallsResponse)(nil), // 3: api_bpfrecorder.SyscallsResponse
}
var file_api_grpc_bpfrecorder_api_proto_depIdxs = []int32{
	0, // 0: api_bpfrecorder.BpfRecorder.Start:input_type -> api_bpfrecorder.EmptyRequest
	0, // 1: api_bpfrecorder.BpfRecorder.Stop:input_type -> api_bpfrecorder.EmptyRequest
	2, // 2: api_bpfrecorder.BpfRecorder.SyscallsForProfile:input_type -> api_bpfrecorder.ProfileRequest
	1, // 3: api_bpfrecorder.BpfRecorder.Start:output_type -> api_bpfrecorder.EmptyResponse
	1, // 4: api_bpfrecorder.BpfRecorder.Stop:output_type -> api_bpfrecorder.EmptyResponse
	3, // 5: api_bpfrecorder.BpfRecorder.SyscallsForProfile:output_type -> api_bpfrecorder.SyscallsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_grpc_bpfrecorder_api_proto_init() }
func file_api_grpc_bpfrecorder_api_proto_init() {
	if File_api_grpc_bpfrecorder_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_bpfrecorder_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_api_grpc_bpfrecorder_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_api_grpc_bpfrecorder_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileRequest); i {
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
		file_api_grpc_bpfrecorder_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyscallsResponse); i {
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
			RawDescriptor: file_api_grpc_bpfrecorder_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_bpfrecorder_api_proto_goTypes,
		DependencyIndexes: file_api_grpc_bpfrecorder_api_proto_depIdxs,
		MessageInfos:      file_api_grpc_bpfrecorder_api_proto_msgTypes,
	}.Build()
	File_api_grpc_bpfrecorder_api_proto = out.File
	file_api_grpc_bpfrecorder_api_proto_rawDesc = nil
	file_api_grpc_bpfrecorder_api_proto_goTypes = nil
	file_api_grpc_bpfrecorder_api_proto_depIdxs = nil
}
