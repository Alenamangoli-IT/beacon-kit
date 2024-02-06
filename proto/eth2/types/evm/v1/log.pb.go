// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
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
// source: types/evm/v1/log.proto

package v1

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

// Log represents an protobuf compatible Ethereum Log that defines a contract
// log event. These events are generated by the LOG opcode and stored/indexed by
// the node.
//
// NOTE: address, topics and data are consensus fields. The rest of the fields
// are derived, i.e. filled in by the nodes, but not secured by consensus.
type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address of the contract that generated the event
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// topics is a list of topics provided by the contract.
	Topics [][]byte `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
	// data which is supplied by the contract, usually ABI-encoded
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// block_number of the block in which the transaction was included
	BlockNumber uint64 `protobuf:"varint,4,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	// tx_hash is the transaction hash
	TxHash []byte `protobuf:"bytes,5,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	// tx_index of the transaction in the block
	TxIndex uint64 `protobuf:"varint,6,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	// block_hash of the block in which the transaction was included
	BlockHash []byte `protobuf:"bytes,7,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	// index of the log in the block
	Index uint64 `protobuf:"varint,8,opt,name=index,proto3" json:"index,omitempty"`
	// removed is true if this log was reverted due to a chain
	// reorganisation. You must pay attention to this field if you receive logs
	// through a filter query.
	Removed bool `protobuf:"varint,9,opt,name=removed,proto3" json:"removed,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_evm_v1_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_types_evm_v1_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_types_evm_v1_log_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Log) GetTopics() [][]byte {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *Log) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Log) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *Log) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *Log) GetTxIndex() uint64 {
	if x != nil {
		return x.TxIndex
	}
	return 0
}

func (x *Log) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *Log) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Log) GetRemoved() bool {
	if x != nil {
		return x.Removed
	}
	return false
}

var File_types_evm_v1_log_proto protoreflect.FileDescriptor

var file_types_evm_v1_log_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x76, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x65, 0x76, 0x6d, 0x2e, 0x76, 0x31, 0x22, 0xf1, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x74, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x73, 0x64, 0x65, 0x76, 0x62,
	0x65, 0x61, 0x72, 0x2f, 0x62, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x76, 0x6d,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_evm_v1_log_proto_rawDescOnce sync.Once
	file_types_evm_v1_log_proto_rawDescData = file_types_evm_v1_log_proto_rawDesc
)

func file_types_evm_v1_log_proto_rawDescGZIP() []byte {
	file_types_evm_v1_log_proto_rawDescOnce.Do(func() {
		file_types_evm_v1_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_evm_v1_log_proto_rawDescData)
	})
	return file_types_evm_v1_log_proto_rawDescData
}

var file_types_evm_v1_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_evm_v1_log_proto_goTypes = []interface{}{
	(*Log)(nil), // 0: types.evm.v1.Log
}
var file_types_evm_v1_log_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_evm_v1_log_proto_init() }
func file_types_evm_v1_log_proto_init() {
	if File_types_evm_v1_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_evm_v1_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_types_evm_v1_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_evm_v1_log_proto_goTypes,
		DependencyIndexes: file_types_evm_v1_log_proto_depIdxs,
		MessageInfos:      file_types_evm_v1_log_proto_msgTypes,
	}.Build()
	File_types_evm_v1_log_proto = out.File
	file_types_evm_v1_log_proto_rawDesc = nil
	file_types_evm_v1_log_proto_goTypes = nil
	file_types_evm_v1_log_proto_depIdxs = nil
}
