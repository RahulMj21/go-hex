// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: numberMsg.proto

package pb

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

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_numberMsg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_numberMsg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_numberMsg_proto_rawDescGZIP(), []int{0}
}

func (x *Answer) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type OperationParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *OperationParameters) Reset() {
	*x = OperationParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_numberMsg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationParameters) ProtoMessage() {}

func (x *OperationParameters) ProtoReflect() protoreflect.Message {
	mi := &file_numberMsg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationParameters.ProtoReflect.Descriptor instead.
func (*OperationParameters) Descriptor() ([]byte, []int) {
	return file_numberMsg_proto_rawDescGZIP(), []int{1}
}

func (x *OperationParameters) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *OperationParameters) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

var File_numberMsg_proto protoreflect.FileDescriptor

var file_numberMsg_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x1e, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x31, 0x0a, 0x13, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x0c, 0x0a, 0x01,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_numberMsg_proto_rawDescOnce sync.Once
	file_numberMsg_proto_rawDescData = file_numberMsg_proto_rawDesc
)

func file_numberMsg_proto_rawDescGZIP() []byte {
	file_numberMsg_proto_rawDescOnce.Do(func() {
		file_numberMsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_numberMsg_proto_rawDescData)
	})
	return file_numberMsg_proto_rawDescData
}

var file_numberMsg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_numberMsg_proto_goTypes = []interface{}{
	(*Answer)(nil),              // 0: pb.Answer
	(*OperationParameters)(nil), // 1: pb.OperationParameters
}
var file_numberMsg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_numberMsg_proto_init() }
func file_numberMsg_proto_init() {
	if File_numberMsg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_numberMsg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
		file_numberMsg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationParameters); i {
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
			RawDescriptor: file_numberMsg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_numberMsg_proto_goTypes,
		DependencyIndexes: file_numberMsg_proto_depIdxs,
		MessageInfos:      file_numberMsg_proto_msgTypes,
	}.Build()
	File_numberMsg_proto = out.File
	file_numberMsg_proto_rawDesc = nil
	file_numberMsg_proto_goTypes = nil
	file_numberMsg_proto_depIdxs = nil
}
