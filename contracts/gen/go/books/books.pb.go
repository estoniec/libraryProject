// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.0
// source: proto/books.proto

package pb_books

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

type FindByISBNRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ISBN string `protobuf:"bytes,1,opt,name=ISBN,proto3" json:"ISBN,omitempty"`
}

func (x *FindByISBNRequest) Reset() {
	*x = FindByISBNRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByISBNRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByISBNRequest) ProtoMessage() {}

func (x *FindByISBNRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByISBNRequest.ProtoReflect.Descriptor instead.
func (*FindByISBNRequest) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{0}
}

func (x *FindByISBNRequest) GetISBN() string {
	if x != nil {
		return x.ISBN
	}
	return ""
}

type FindByISBNResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Status int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Isbn   string `protobuf:"bytes,3,opt,name=isbn,proto3" json:"isbn,omitempty"`
	Count  int64  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *FindByISBNResponse) Reset() {
	*x = FindByISBNResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByISBNResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByISBNResponse) ProtoMessage() {}

func (x *FindByISBNResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByISBNResponse.ProtoReflect.Descriptor instead.
func (*FindByISBNResponse) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{1}
}

func (x *FindByISBNResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindByISBNResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindByISBNResponse) GetIsbn() string {
	if x != nil {
		return x.Isbn
	}
	return ""
}

func (x *FindByISBNResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_proto_books_proto protoreflect.FileDescriptor

var file_proto_books_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x49, 0x53, 0x42, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x49, 0x53, 0x42, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x53, 0x42,
	0x4e, 0x22, 0x6c, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x53, 0x42, 0x4e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32,
	0x6f, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5f, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x53, 0x42, 0x4e, 0x12, 0x26, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x53, 0x42, 0x4e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x49, 0x53, 0x42, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x3b, 0x70, 0x62, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_books_proto_rawDescOnce sync.Once
	file_proto_books_proto_rawDescData = file_proto_books_proto_rawDesc
)

func file_proto_books_proto_rawDescGZIP() []byte {
	file_proto_books_proto_rawDescOnce.Do(func() {
		file_proto_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_books_proto_rawDescData)
	})
	return file_proto_books_proto_rawDescData
}

var file_proto_books_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_books_proto_goTypes = []interface{}{
	(*FindByISBNRequest)(nil),  // 0: books_service.books.FindByISBNRequest
	(*FindByISBNResponse)(nil), // 1: books_service.books.FindByISBNResponse
}
var file_proto_books_proto_depIdxs = []int32{
	0, // 0: books_service.books.BooksService.FindByISBN:input_type -> books_service.books.FindByISBNRequest
	1, // 1: books_service.books.BooksService.FindByISBN:output_type -> books_service.books.FindByISBNResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_books_proto_init() }
func file_proto_books_proto_init() {
	if File_proto_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByISBNRequest); i {
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
		file_proto_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByISBNResponse); i {
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
			RawDescriptor: file_proto_books_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_books_proto_goTypes,
		DependencyIndexes: file_proto_books_proto_depIdxs,
		MessageInfos:      file_proto_books_proto_msgTypes,
	}.Build()
	File_proto_books_proto = out.File
	file_proto_books_proto_rawDesc = nil
	file_proto_books_proto_goTypes = nil
	file_proto_books_proto_depIdxs = nil
}
