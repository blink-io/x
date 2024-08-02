// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.1
// source: i18n.proto

package grpc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LanguageEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Language string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Valid    bool   `protobuf:"varint,3,opt,name=valid,proto3" json:"valid,omitempty"`
	Payload  []byte `protobuf:"bytes,20,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *LanguageEntry) Reset() {
	*x = LanguageEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_i18n_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageEntry) ProtoMessage() {}

func (x *LanguageEntry) ProtoReflect() protoreflect.Message {
	mi := &file_i18n_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageEntry.ProtoReflect.Descriptor instead.
func (*LanguageEntry) Descriptor() ([]byte, []int) {
	return file_i18n_proto_rawDescGZIP(), []int{0}
}

func (x *LanguageEntry) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *LanguageEntry) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *LanguageEntry) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *LanguageEntry) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type ListLanguagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Languages []string `protobuf:"bytes,1,rep,name=languages,proto3" json:"languages,omitempty"`
}

func (x *ListLanguagesRequest) Reset() {
	*x = ListLanguagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_i18n_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLanguagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLanguagesRequest) ProtoMessage() {}

func (x *ListLanguagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_i18n_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLanguagesRequest.ProtoReflect.Descriptor instead.
func (*ListLanguagesRequest) Descriptor() ([]byte, []int) {
	return file_i18n_proto_rawDescGZIP(), []int{1}
}

func (x *ListLanguagesRequest) GetLanguages() []string {
	if x != nil {
		return x.Languages
	}
	return nil
}

type ListLanguagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries   map[string]*LanguageEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Timestamp int64                     `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ListLanguagesResponse) Reset() {
	*x = ListLanguagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_i18n_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLanguagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLanguagesResponse) ProtoMessage() {}

func (x *ListLanguagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_i18n_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLanguagesResponse.ProtoReflect.Descriptor instead.
func (*ListLanguagesResponse) Descriptor() ([]byte, []int) {
	return file_i18n_proto_rawDescGZIP(), []int{2}
}

func (x *ListLanguagesResponse) GetEntries() map[string]*LanguageEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *ListLanguagesResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_i18n_proto protoreflect.FileDescriptor

var file_i18n_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x31, 0x38, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x0d,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x34, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x4a, 0x0a, 0x0c, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x46, 0x0a, 0x04, 0x49, 0x31, 0x38, 0x4e, 0x12, 0x3e,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x24,
	0x48, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x66, 0x75, 0x6e, 0x2f, 0x78, 0x2f, 0x69, 0x31, 0x38, 0x6e, 0x3b,
	0x69, 0x31, 0x38, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_i18n_proto_rawDescOnce sync.Once
	file_i18n_proto_rawDescData = file_i18n_proto_rawDesc
)

func file_i18n_proto_rawDescGZIP() []byte {
	file_i18n_proto_rawDescOnce.Do(func() {
		file_i18n_proto_rawDescData = protoimpl.X.CompressGZIP(file_i18n_proto_rawDescData)
	})
	return file_i18n_proto_rawDescData
}

var file_i18n_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_i18n_proto_goTypes = []interface{}{
	(*LanguageEntry)(nil),         // 0: LanguageEntry
	(*ListLanguagesRequest)(nil),  // 1: ListLanguagesRequest
	(*ListLanguagesResponse)(nil), // 2: ListLanguagesResponse
	nil,                           // 3: ListLanguagesResponse.EntriesEntry
}
var file_i18n_proto_depIdxs = []int32{
	3, // 0: ListLanguagesResponse.entries:type_name -> ListLanguagesResponse.EntriesEntry
	0, // 1: ListLanguagesResponse.EntriesEntry.value:type_name -> LanguageEntry
	1, // 2: I18N.ListLanguages:input_type -> ListLanguagesRequest
	2, // 3: I18N.ListLanguages:output_type -> ListLanguagesResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_i18n_proto_init() }
func file_i18n_proto_init() {
	if File_i18n_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_i18n_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageEntry); i {
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
		file_i18n_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLanguagesRequest); i {
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
		file_i18n_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLanguagesResponse); i {
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
			RawDescriptor: file_i18n_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_i18n_proto_goTypes,
		DependencyIndexes: file_i18n_proto_depIdxs,
		MessageInfos:      file_i18n_proto_msgTypes,
	}.Build()
	File_i18n_proto = out.File
	file_i18n_proto_rawDesc = nil
	file_i18n_proto_goTypes = nil
	file_i18n_proto_depIdxs = nil
}
