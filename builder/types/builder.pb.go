// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: builder/v1alpha1/builder.proto

package types

import (
	v1 "github.com/itsdevbear/bolaris/types/consensus/v1"
	github_com_itsdevbear_bolaris_types_consensus_primitives "github.com/itsdevbear/bolaris/types/consensus/primitives"
	_ "github.com/prysmaticlabs/prysm/v4/proto/eth/ext"
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

// RequestBestBlockRequest is the request to the RequestBestBlock RPC.
type RequestBestBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// slot is the slot of the block to request.
	Slot github_com_itsdevbear_bolaris_types_consensus_primitives.Slot `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/itsdevbear/bolaris/types/consensus/primitives.Slot"`
	// blinded is true if the block should be blinded.
	Blinded bool `protobuf:"varint,2,opt,name=blinded,proto3" json:"blinded,omitempty"`
}

func (x *RequestBestBlockRequest) Reset() {
	*x = RequestBestBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builder_v1alpha1_builder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBestBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBestBlockRequest) ProtoMessage() {}

func (x *RequestBestBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_builder_v1alpha1_builder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestBestBlockRequest.ProtoReflect.Descriptor instead.
func (*RequestBestBlockRequest) Descriptor() ([]byte, []int) {
	return file_builder_v1alpha1_builder_proto_rawDescGZIP(), []int{0}
}

func (x *RequestBestBlockRequest) GetSlot() github_com_itsdevbear_bolaris_types_consensus_primitives.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_itsdevbear_bolaris_types_consensus_primitives.Slot(0)
}

func (x *RequestBestBlockRequest) GetBlinded() bool {
	if x != nil {
		return x.Blinded
	}
	return false
}

// RequestBestBlockResponse is the response to the RequestBestBlock RPC.
type RequestBestBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// override is true if the block should override the current best block.
	// this field should only be taken into account if the block builder is
	// a TRUSTED party.
	Override bool `protobuf:"varint,1,opt,name=override,proto3" json:"override,omitempty"`
	// block_container is the block requested.
	BlockContainer *v1.BeaconKitBlockContainer `protobuf:"bytes,2,opt,name=block_container,json=blockContainer,proto3" json:"block_container,omitempty"`
}

func (x *RequestBestBlockResponse) Reset() {
	*x = RequestBestBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builder_v1alpha1_builder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBestBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBestBlockResponse) ProtoMessage() {}

func (x *RequestBestBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_builder_v1alpha1_builder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestBestBlockResponse.ProtoReflect.Descriptor instead.
func (*RequestBestBlockResponse) Descriptor() ([]byte, []int) {
	return file_builder_v1alpha1_builder_proto_rawDescGZIP(), []int{1}
}

func (x *RequestBestBlockResponse) GetOverride() bool {
	if x != nil {
		return x.Override
	}
	return false
}

func (x *RequestBestBlockResponse) GetBlockContainer() *v1.BeaconKitBlockContainer {
	if x != nil {
		return x.BlockContainer
	}
	return nil
}

var File_builder_v1alpha1_builder_proto protoreflect.FileDescriptor

var file_builder_v1alpha1_builder_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x1e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x65, 0x74, 0x68,
	0x2f, 0x65, 0x78, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x55, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x41, 0x82, 0xb5, 0x18, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x74, 0x73, 0x64, 0x65, 0x76, 0x62, 0x65, 0x61, 0x72, 0x2f, 0x62, 0x6f, 0x6c, 0x61,
	0x72, 0x69, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x53,
	0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x69,
	0x6e, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x6c, 0x69, 0x6e,
	0x64, 0x65, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x54, 0x0a, 0x0f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f,
	0x6e, 0x4b, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x32, 0x7b, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x29, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x65,
	0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74,
	0x73, 0x64, 0x65, 0x76, 0x62, 0x65, 0x61, 0x72, 0x2f, 0x62, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_builder_v1alpha1_builder_proto_rawDescOnce sync.Once
	file_builder_v1alpha1_builder_proto_rawDescData = file_builder_v1alpha1_builder_proto_rawDesc
)

func file_builder_v1alpha1_builder_proto_rawDescGZIP() []byte {
	file_builder_v1alpha1_builder_proto_rawDescOnce.Do(func() {
		file_builder_v1alpha1_builder_proto_rawDescData = protoimpl.X.CompressGZIP(file_builder_v1alpha1_builder_proto_rawDescData)
	})
	return file_builder_v1alpha1_builder_proto_rawDescData
}

var file_builder_v1alpha1_builder_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_builder_v1alpha1_builder_proto_goTypes = []interface{}{
	(*RequestBestBlockRequest)(nil),    // 0: builder.v1alpha1.RequestBestBlockRequest
	(*RequestBestBlockResponse)(nil),   // 1: builder.v1alpha1.RequestBestBlockResponse
	(*v1.BeaconKitBlockContainer)(nil), // 2: types.consensus.v1.BeaconKitBlockContainer
}
var file_builder_v1alpha1_builder_proto_depIdxs = []int32{
	2, // 0: builder.v1alpha1.RequestBestBlockResponse.block_container:type_name -> types.consensus.v1.BeaconKitBlockContainer
	0, // 1: builder.v1alpha1.BuilderService.RequestBestBlock:input_type -> builder.v1alpha1.RequestBestBlockRequest
	1, // 2: builder.v1alpha1.BuilderService.RequestBestBlock:output_type -> builder.v1alpha1.RequestBestBlockResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_builder_v1alpha1_builder_proto_init() }
func file_builder_v1alpha1_builder_proto_init() {
	if File_builder_v1alpha1_builder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_builder_v1alpha1_builder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestBestBlockRequest); i {
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
		file_builder_v1alpha1_builder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestBestBlockResponse); i {
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
			RawDescriptor: file_builder_v1alpha1_builder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_builder_v1alpha1_builder_proto_goTypes,
		DependencyIndexes: file_builder_v1alpha1_builder_proto_depIdxs,
		MessageInfos:      file_builder_v1alpha1_builder_proto_msgTypes,
	}.Build()
	File_builder_v1alpha1_builder_proto = out.File
	file_builder_v1alpha1_builder_proto_rawDesc = nil
	file_builder_v1alpha1_builder_proto_goTypes = nil
	file_builder_v1alpha1_builder_proto_depIdxs = nil
}
