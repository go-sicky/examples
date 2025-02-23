//
// The MIT License (MIT)
//
// Copyright (c) 2021 HereweTech Co.LTD
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

//*
// @file hybrid.proto
// @author Dr.NP <np@herewe.tech>
// @since 02/23/2025

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: proto/hybrid.proto

package proto

import (
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

type HybridRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HybridRequest) Reset() {
	*x = HybridRequest{}
	mi := &file_proto_hybrid_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HybridRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HybridRequest) ProtoMessage() {}

func (x *HybridRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hybrid_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HybridRequest.ProtoReflect.Descriptor instead.
func (*HybridRequest) Descriptor() ([]byte, []int) {
	return file_proto_hybrid_proto_rawDescGZIP(), []int{0}
}

func (x *HybridRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HybridResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Answer        string                 `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HybridResponse) Reset() {
	*x = HybridResponse{}
	mi := &file_proto_hybrid_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HybridResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HybridResponse) ProtoMessage() {}

func (x *HybridResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hybrid_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HybridResponse.ProtoReflect.Descriptor instead.
func (*HybridResponse) Descriptor() ([]byte, []int) {
	return file_proto_hybrid_proto_rawDescGZIP(), []int{1}
}

func (x *HybridResponse) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

var File_proto_hybrid_proto protoreflect.FileDescriptor

var file_proto_hybrid_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x69, 0x63, 0x6b, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x22,
	0x23, 0x0a, 0x0d, 0x48, 0x79, 0x62, 0x72, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x48, 0x79, 0x62, 0x72, 0x69, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x32, 0x67,
	0x0a, 0x06, 0x48, 0x79, 0x62, 0x72, 0x69, 0x64, 0x12, 0x5d, 0x0a, 0x06, 0x48, 0x79, 0x62, 0x72,
	0x69, 0x64, 0x12, 0x28, 0x2e, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x69, 0x63, 0x6b, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x48,
	0x79, 0x62, 0x72, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x68,
	0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73,
	0x69, 0x63, 0x6b, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x48, 0x79, 0x62, 0x72, 0x69, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_hybrid_proto_rawDescOnce sync.Once
	file_proto_hybrid_proto_rawDescData []byte
)

func file_proto_hybrid_proto_rawDescGZIP() []byte {
	file_proto_hybrid_proto_rawDescOnce.Do(func() {
		file_proto_hybrid_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_hybrid_proto_rawDesc), len(file_proto_hybrid_proto_rawDesc)))
	})
	return file_proto_hybrid_proto_rawDescData
}

var file_proto_hybrid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_hybrid_proto_goTypes = []any{
	(*HybridRequest)(nil),  // 0: hybrid.examples.sicky.dev.HybridRequest
	(*HybridResponse)(nil), // 1: hybrid.examples.sicky.dev.HybridResponse
}
var file_proto_hybrid_proto_depIdxs = []int32{
	0, // 0: hybrid.examples.sicky.dev.Hybrid.Hybrid:input_type -> hybrid.examples.sicky.dev.HybridRequest
	1, // 1: hybrid.examples.sicky.dev.Hybrid.Hybrid:output_type -> hybrid.examples.sicky.dev.HybridResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_hybrid_proto_init() }
func file_proto_hybrid_proto_init() {
	if File_proto_hybrid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_hybrid_proto_rawDesc), len(file_proto_hybrid_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hybrid_proto_goTypes,
		DependencyIndexes: file_proto_hybrid_proto_depIdxs,
		MessageInfos:      file_proto_hybrid_proto_msgTypes,
	}.Build()
	File_proto_hybrid_proto = out.File
	file_proto_hybrid_proto_goTypes = nil
	file_proto_hybrid_proto_depIdxs = nil
}
