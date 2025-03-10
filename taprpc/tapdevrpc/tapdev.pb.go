// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: tapdevrpc/tapdev.proto

package tapdevrpc

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

type ImportProofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProofFile    []byte `protobuf:"bytes,1,opt,name=proof_file,json=proofFile,proto3" json:"proof_file,omitempty"`
	GenesisPoint string `protobuf:"bytes,2,opt,name=genesis_point,json=genesisPoint,proto3" json:"genesis_point,omitempty"`
}

func (x *ImportProofRequest) Reset() {
	*x = ImportProofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tapdevrpc_tapdev_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportProofRequest) ProtoMessage() {}

func (x *ImportProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tapdevrpc_tapdev_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportProofRequest.ProtoReflect.Descriptor instead.
func (*ImportProofRequest) Descriptor() ([]byte, []int) {
	return file_tapdevrpc_tapdev_proto_rawDescGZIP(), []int{0}
}

func (x *ImportProofRequest) GetProofFile() []byte {
	if x != nil {
		return x.ProofFile
	}
	return nil
}

func (x *ImportProofRequest) GetGenesisPoint() string {
	if x != nil {
		return x.GenesisPoint
	}
	return ""
}

type ImportProofResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ImportProofResponse) Reset() {
	*x = ImportProofResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tapdevrpc_tapdev_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportProofResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportProofResponse) ProtoMessage() {}

func (x *ImportProofResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tapdevrpc_tapdev_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportProofResponse.ProtoReflect.Descriptor instead.
func (*ImportProofResponse) Descriptor() ([]byte, []int) {
	return file_tapdevrpc_tapdev_proto_rawDescGZIP(), []int{1}
}

var File_tapdevrpc_tapdev_proto protoreflect.FileDescriptor

var file_tapdevrpc_tapdev_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x61, 0x70, 0x64, 0x65, 0x76, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x61, 0x70, 0x64,
	0x65, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x61, 0x70, 0x64, 0x65, 0x76,
	0x72, 0x70, 0x63, 0x22, 0x58, 0x0a, 0x12, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65,
	0x73, 0x69, 0x73, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x15, 0x0a,
	0x13, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x56, 0x0a, 0x06, 0x54, 0x61, 0x70, 0x44, 0x65, 0x76, 0x12, 0x4c,
	0x0a, 0x0b, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1d, 0x2e,
	0x74, 0x61, 0x70, 0x64, 0x65, 0x76, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74,
	0x61, 0x70, 0x64, 0x65, 0x76, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x6e, 0x69, 0x6e, 0x67, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x6f, 0x74,
	0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x61, 0x70, 0x72, 0x70, 0x63, 0x2f, 0x74,
	0x61, 0x70, 0x64, 0x65, 0x76, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tapdevrpc_tapdev_proto_rawDescOnce sync.Once
	file_tapdevrpc_tapdev_proto_rawDescData = file_tapdevrpc_tapdev_proto_rawDesc
)

func file_tapdevrpc_tapdev_proto_rawDescGZIP() []byte {
	file_tapdevrpc_tapdev_proto_rawDescOnce.Do(func() {
		file_tapdevrpc_tapdev_proto_rawDescData = protoimpl.X.CompressGZIP(file_tapdevrpc_tapdev_proto_rawDescData)
	})
	return file_tapdevrpc_tapdev_proto_rawDescData
}

var file_tapdevrpc_tapdev_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tapdevrpc_tapdev_proto_goTypes = []interface{}{
	(*ImportProofRequest)(nil),  // 0: tapdevrpc.ImportProofRequest
	(*ImportProofResponse)(nil), // 1: tapdevrpc.ImportProofResponse
}
var file_tapdevrpc_tapdev_proto_depIdxs = []int32{
	0, // 0: tapdevrpc.TapDev.ImportProof:input_type -> tapdevrpc.ImportProofRequest
	1, // 1: tapdevrpc.TapDev.ImportProof:output_type -> tapdevrpc.ImportProofResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tapdevrpc_tapdev_proto_init() }
func file_tapdevrpc_tapdev_proto_init() {
	if File_tapdevrpc_tapdev_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tapdevrpc_tapdev_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportProofRequest); i {
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
		file_tapdevrpc_tapdev_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportProofResponse); i {
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
			RawDescriptor: file_tapdevrpc_tapdev_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tapdevrpc_tapdev_proto_goTypes,
		DependencyIndexes: file_tapdevrpc_tapdev_proto_depIdxs,
		MessageInfos:      file_tapdevrpc_tapdev_proto_msgTypes,
	}.Build()
	File_tapdevrpc_tapdev_proto = out.File
	file_tapdevrpc_tapdev_proto_rawDesc = nil
	file_tapdevrpc_tapdev_proto_goTypes = nil
	file_tapdevrpc_tapdev_proto_depIdxs = nil
}
