// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: idl/math.proto

package math_service

import (
	context "context"
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

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Left  int32 `protobuf:"varint,1,opt,name=left,proto3" json:"left,omitempty"`
	Right int32 `protobuf:"varint,2,opt,name=right,proto3" json:"right,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_math_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_math_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_idl_math_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetLeft() int32 {
	if x != nil {
		return x.Left
	}
	return 0
}

func (x *AddRequest) GetRight() int32 {
	if x != nil {
		return x.Right
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_math_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_math_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_idl_math_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type SubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Left  int32 `protobuf:"varint,1,opt,name=left,proto3" json:"left,omitempty"`
	Right int32 `protobuf:"varint,2,opt,name=right,proto3" json:"right,omitempty"`
}

func (x *SubRequest) Reset() {
	*x = SubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_math_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubRequest) ProtoMessage() {}

func (x *SubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_math_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubRequest.ProtoReflect.Descriptor instead.
func (*SubRequest) Descriptor() ([]byte, []int) {
	return file_idl_math_proto_rawDescGZIP(), []int{2}
}

func (x *SubRequest) GetLeft() int32 {
	if x != nil {
		return x.Left
	}
	return 0
}

func (x *SubRequest) GetRight() int32 {
	if x != nil {
		return x.Right
	}
	return 0
}

type SubResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Diff int32 `protobuf:"varint,1,opt,name=diff,proto3" json:"diff,omitempty"`
}

func (x *SubResponse) Reset() {
	*x = SubResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_math_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubResponse) ProtoMessage() {}

func (x *SubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_math_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubResponse.ProtoReflect.Descriptor instead.
func (*SubResponse) Descriptor() ([]byte, []int) {
	return file_idl_math_proto_rawDescGZIP(), []int{3}
}

func (x *SubResponse) GetDiff() int32 {
	if x != nil {
		return x.Diff
	}
	return 0
}

var File_idl_math_proto protoreflect.FileDescriptor

var file_idl_math_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x6c, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x69, 0x64, 0x6c, 0x22, 0x36, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x22, 0x1f, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x36,
	0x0a, 0x0a, 0x53, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x65, 0x66, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x22, 0x21, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x66, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x69, 0x66, 0x66, 0x32, 0x5a, 0x0a, 0x04, 0x4d, 0x61, 0x74,
	0x68, 0x12, 0x28, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0f, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x69, 0x64, 0x6c, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x03, 0x53,
	0x75, 0x62, 0x12, 0x0f, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x53, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x53, 0x75, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x6d, 0x79, 0x5f, 0x6b, 0x69, 0x74, 0x65,
	0x78, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x61, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_math_proto_rawDescOnce sync.Once
	file_idl_math_proto_rawDescData = file_idl_math_proto_rawDesc
)

func file_idl_math_proto_rawDescGZIP() []byte {
	file_idl_math_proto_rawDescOnce.Do(func() {
		file_idl_math_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_math_proto_rawDescData)
	})
	return file_idl_math_proto_rawDescData
}

var file_idl_math_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_idl_math_proto_goTypes = []interface{}{
	(*AddRequest)(nil),  // 0: idl.AddRequest
	(*AddResponse)(nil), // 1: idl.AddResponse
	(*SubRequest)(nil),  // 2: idl.SubRequest
	(*SubResponse)(nil), // 3: idl.SubResponse
}
var file_idl_math_proto_depIdxs = []int32{
	0, // 0: idl.Math.Add:input_type -> idl.AddRequest
	2, // 1: idl.Math.Sub:input_type -> idl.SubRequest
	1, // 2: idl.Math.Add:output_type -> idl.AddResponse
	3, // 3: idl.Math.Sub:output_type -> idl.SubResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idl_math_proto_init() }
func file_idl_math_proto_init() {
	if File_idl_math_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_math_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_idl_math_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
		file_idl_math_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubRequest); i {
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
		file_idl_math_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubResponse); i {
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
			RawDescriptor: file_idl_math_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_math_proto_goTypes,
		DependencyIndexes: file_idl_math_proto_depIdxs,
		MessageInfos:      file_idl_math_proto_msgTypes,
	}.Build()
	File_idl_math_proto = out.File
	file_idl_math_proto_rawDesc = nil
	file_idl_math_proto_goTypes = nil
	file_idl_math_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.1. DO NOT EDIT.

type Math interface {
	Add(ctx context.Context, req *AddRequest) (res *AddResponse, err error)
	Sub(ctx context.Context, req *SubRequest) (res *SubResponse, err error)
}
