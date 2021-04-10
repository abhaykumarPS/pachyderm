// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: internal/storage/chunk/chunk.proto

package chunk

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

type CompressionAlgo int32

const (
	CompressionAlgo_NONE            CompressionAlgo = 0
	CompressionAlgo_GZIP_BEST_SPEED CompressionAlgo = 1
)

// Enum value maps for CompressionAlgo.
var (
	CompressionAlgo_name = map[int32]string{
		0: "NONE",
		1: "GZIP_BEST_SPEED",
	}
	CompressionAlgo_value = map[string]int32{
		"NONE":            0,
		"GZIP_BEST_SPEED": 1,
	}
)

func (x CompressionAlgo) Enum() *CompressionAlgo {
	p := new(CompressionAlgo)
	*p = x
	return p
}

func (x CompressionAlgo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompressionAlgo) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_storage_chunk_chunk_proto_enumTypes[0].Descriptor()
}

func (CompressionAlgo) Type() protoreflect.EnumType {
	return &file_internal_storage_chunk_chunk_proto_enumTypes[0]
}

func (x CompressionAlgo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompressionAlgo.Descriptor instead.
func (CompressionAlgo) EnumDescriptor() ([]byte, []int) {
	return file_internal_storage_chunk_chunk_proto_rawDescGZIP(), []int{0}
}

type EncryptionAlgo int32

const (
	EncryptionAlgo_CHACHA20 EncryptionAlgo = 0
)

// Enum value maps for EncryptionAlgo.
var (
	EncryptionAlgo_name = map[int32]string{
		0: "CHACHA20",
	}
	EncryptionAlgo_value = map[string]int32{
		"CHACHA20": 0,
	}
)

func (x EncryptionAlgo) Enum() *EncryptionAlgo {
	p := new(EncryptionAlgo)
	*p = x
	return p
}

func (x EncryptionAlgo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncryptionAlgo) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_storage_chunk_chunk_proto_enumTypes[1].Descriptor()
}

func (EncryptionAlgo) Type() protoreflect.EnumType {
	return &file_internal_storage_chunk_chunk_proto_enumTypes[1]
}

func (x EncryptionAlgo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncryptionAlgo.Descriptor instead.
func (EncryptionAlgo) EnumDescriptor() ([]byte, []int) {
	return file_internal_storage_chunk_chunk_proto_rawDescGZIP(), []int{1}
}

// DataRef is a reference to data within a chunk.
type DataRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The chunk the referenced data is located in.
	Ref *Ref `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// The hash of the data being referenced.
	// This field is empty when it is equal to the chunk hash (the ref is the whole chunk).
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// The offset and size used for accessing the data within the chunk.
	OffsetBytes int64 `protobuf:"varint,3,opt,name=offset_bytes,json=offsetBytes,proto3" json:"offset_bytes,omitempty"`
	SizeBytes   int64 `protobuf:"varint,4,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
}

func (x *DataRef) Reset() {
	*x = DataRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_storage_chunk_chunk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRef) ProtoMessage() {}

func (x *DataRef) ProtoReflect() protoreflect.Message {
	mi := &file_internal_storage_chunk_chunk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRef.ProtoReflect.Descriptor instead.
func (*DataRef) Descriptor() ([]byte, []int) {
	return file_internal_storage_chunk_chunk_proto_rawDescGZIP(), []int{0}
}

func (x *DataRef) GetRef() *Ref {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *DataRef) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *DataRef) GetOffsetBytes() int64 {
	if x != nil {
		return x.OffsetBytes
	}
	return 0
}

func (x *DataRef) GetSizeBytes() int64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

type Ref struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              []byte          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SizeBytes       int64           `protobuf:"varint,2,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	Edge            bool            `protobuf:"varint,3,opt,name=edge,proto3" json:"edge,omitempty"`
	Dek             []byte          `protobuf:"bytes,4,opt,name=dek,proto3" json:"dek,omitempty"`
	EncryptionAlgo  EncryptionAlgo  `protobuf:"varint,5,opt,name=encryption_algo,json=encryptionAlgo,proto3,enum=chunk.EncryptionAlgo" json:"encryption_algo,omitempty"`
	CompressionAlgo CompressionAlgo `protobuf:"varint,6,opt,name=compression_algo,json=compressionAlgo,proto3,enum=chunk.CompressionAlgo" json:"compression_algo,omitempty"`
}

func (x *Ref) Reset() {
	*x = Ref{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_storage_chunk_chunk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ref) ProtoMessage() {}

func (x *Ref) ProtoReflect() protoreflect.Message {
	mi := &file_internal_storage_chunk_chunk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ref.ProtoReflect.Descriptor instead.
func (*Ref) Descriptor() ([]byte, []int) {
	return file_internal_storage_chunk_chunk_proto_rawDescGZIP(), []int{1}
}

func (x *Ref) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Ref) GetSizeBytes() int64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

func (x *Ref) GetEdge() bool {
	if x != nil {
		return x.Edge
	}
	return false
}

func (x *Ref) GetDek() []byte {
	if x != nil {
		return x.Dek
	}
	return nil
}

func (x *Ref) GetEncryptionAlgo() EncryptionAlgo {
	if x != nil {
		return x.EncryptionAlgo
	}
	return EncryptionAlgo_CHACHA20
}

func (x *Ref) GetCompressionAlgo() CompressionAlgo {
	if x != nil {
		return x.CompressionAlgo
	}
	return CompressionAlgo_NONE
}

var File_internal_storage_chunk_chunk_proto protoreflect.FileDescriptor

var file_internal_storage_chunk_chunk_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x7d, 0x0a, 0x07, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x66, 0x12, 0x1c, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x52, 0x65, 0x66, 0x52,
	0x03, 0x72, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x03, 0x52,
	0x65, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x64, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x65, 0x64, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x64, 0x65, 0x6b, 0x12, 0x3e, 0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x52, 0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x41, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x2a, 0x30, 0x0a, 0x0f, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x5a, 0x49, 0x50, 0x5f,
	0x42, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x50, 0x45, 0x45, 0x44, 0x10, 0x01, 0x2a, 0x1e, 0x0a, 0x0e,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x48, 0x41, 0x43, 0x48, 0x41, 0x32, 0x30, 0x10, 0x00, 0x42, 0x3e, 0x5a, 0x3c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x68, 0x79,
	0x64, 0x65, 0x72, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2f, 0x76,
	0x32, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_storage_chunk_chunk_proto_rawDescOnce sync.Once
	file_internal_storage_chunk_chunk_proto_rawDescData = file_internal_storage_chunk_chunk_proto_rawDesc
)

func file_internal_storage_chunk_chunk_proto_rawDescGZIP() []byte {
	file_internal_storage_chunk_chunk_proto_rawDescOnce.Do(func() {
		file_internal_storage_chunk_chunk_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_storage_chunk_chunk_proto_rawDescData)
	})
	return file_internal_storage_chunk_chunk_proto_rawDescData
}

var file_internal_storage_chunk_chunk_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_internal_storage_chunk_chunk_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_storage_chunk_chunk_proto_goTypes = []interface{}{
	(CompressionAlgo)(0), // 0: chunk.CompressionAlgo
	(EncryptionAlgo)(0),  // 1: chunk.EncryptionAlgo
	(*DataRef)(nil),      // 2: chunk.DataRef
	(*Ref)(nil),          // 3: chunk.Ref
}
var file_internal_storage_chunk_chunk_proto_depIdxs = []int32{
	3, // 0: chunk.DataRef.ref:type_name -> chunk.Ref
	1, // 1: chunk.Ref.encryption_algo:type_name -> chunk.EncryptionAlgo
	0, // 2: chunk.Ref.compression_algo:type_name -> chunk.CompressionAlgo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_storage_chunk_chunk_proto_init() }
func file_internal_storage_chunk_chunk_proto_init() {
	if File_internal_storage_chunk_chunk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_storage_chunk_chunk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRef); i {
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
		file_internal_storage_chunk_chunk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ref); i {
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
			RawDescriptor: file_internal_storage_chunk_chunk_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_storage_chunk_chunk_proto_goTypes,
		DependencyIndexes: file_internal_storage_chunk_chunk_proto_depIdxs,
		EnumInfos:         file_internal_storage_chunk_chunk_proto_enumTypes,
		MessageInfos:      file_internal_storage_chunk_chunk_proto_msgTypes,
	}.Build()
	File_internal_storage_chunk_chunk_proto = out.File
	file_internal_storage_chunk_chunk_proto_rawDesc = nil
	file_internal_storage_chunk_chunk_proto_goTypes = nil
	file_internal_storage_chunk_chunk_proto_depIdxs = nil
}
