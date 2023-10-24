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

type FindBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ISBN   string `protobuf:"bytes,1,opt,name=ISBN,proto3" json:"ISBN,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *FindBook) Reset() {
	*x = FindBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindBook) ProtoMessage() {}

func (x *FindBook) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FindBook.ProtoReflect.Descriptor instead.
func (*FindBook) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{0}
}

func (x *FindBook) GetISBN() string {
	if x != nil {
		return x.ISBN
	}
	return ""
}

func (x *FindBook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindBook) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ISBN   string `protobuf:"bytes,1,opt,name=ISBN,proto3" json:"ISBN,omitempty"`
	Count  int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Author string `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	ID     int64  `protobuf:"varint,5,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetISBN() string {
	if x != nil {
		return x.ISBN
	}
	return ""
}

func (x *Book) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type FindByRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64     `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Find   *FindBook `protobuf:"bytes,2,opt,name=find,proto3" json:"find,omitempty"`
}

func (x *FindByRequest) Reset() {
	*x = FindByRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByRequest) ProtoMessage() {}

func (x *FindByRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByRequest.ProtoReflect.Descriptor instead.
func (*FindByRequest) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{2}
}

func (x *FindByRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FindByRequest) GetFind() *FindBook {
	if x != nil {
		return x.Find
	}
	return nil
}

type FindByResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  string  `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Status int64   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Book   []*Book `protobuf:"bytes,3,rep,name=book,proto3" json:"book,omitempty"`
}

func (x *FindByResponse) Reset() {
	*x = FindByResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByResponse) ProtoMessage() {}

func (x *FindByResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByResponse.ProtoReflect.Descriptor instead.
func (*FindByResponse) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{3}
}

func (x *FindByResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindByResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindByResponse) GetBook() []*Book {
	if x != nil {
		return x.Book
	}
	return nil
}

type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBookRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

type CreateBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Status int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Books  *Book  `protobuf:"bytes,3,opt,name=books,proto3" json:"books,omitempty"`
}

func (x *CreateBookResponse) Reset() {
	*x = CreateBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookResponse) ProtoMessage() {}

func (x *CreateBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookResponse.ProtoReflect.Descriptor instead.
func (*CreateBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{5}
}

func (x *CreateBookResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateBookResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateBookResponse) GetBooks() *Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type DeleteBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ISBN string `protobuf:"bytes,1,opt,name=ISBN,proto3" json:"ISBN,omitempty"`
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBookRequest) GetISBN() string {
	if x != nil {
		return x.ISBN
	}
	return ""
}

type DeleteBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Status int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteBookResponse) Reset() {
	*x = DeleteBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookResponse) ProtoMessage() {}

func (x *DeleteBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookResponse.ProtoReflect.Descriptor instead.
func (*DeleteBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBookResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *DeleteBookResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type EditCountBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ISBN  string `protobuf:"bytes,1,opt,name=ISBN,proto3" json:"ISBN,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *EditCountBookRequest) Reset() {
	*x = EditCountBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditCountBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditCountBookRequest) ProtoMessage() {}

func (x *EditCountBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditCountBookRequest.ProtoReflect.Descriptor instead.
func (*EditCountBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{8}
}

func (x *EditCountBookRequest) GetISBN() string {
	if x != nil {
		return x.ISBN
	}
	return ""
}

func (x *EditCountBookRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type EditCountBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Status int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *EditCountBookResponse) Reset() {
	*x = EditCountBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditCountBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditCountBookResponse) ProtoMessage() {}

func (x *EditCountBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditCountBookResponse.ProtoReflect.Descriptor instead.
func (*EditCountBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{9}
}

func (x *EditCountBookResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *EditCountBookResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_proto_books_proto protoreflect.FileDescriptor

var file_proto_books_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x4a, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x53, 0x42, 0x4e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x49, 0x53, 0x42, 0x4e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x22, 0x6c, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x49, 0x53, 0x42, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x53, 0x42, 0x4e,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x44, 0x22, 0x5a, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x66,
	0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x66, 0x69, 0x6e, 0x64, 0x22, 0x6d,
	0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d,
	0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x42, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f,
	0x6b, 0x22, 0x73, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x49,
	0x53, 0x42, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x53, 0x42, 0x4e, 0x22,
	0x42, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x40, 0x0a, 0x14, 0x45, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x49,
	0x53, 0x42, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x53, 0x42, 0x4e, 0x12,
	0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x15, 0x45, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x8f, 0x03, 0x0a,
	0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a,
	0x06, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x12, 0x22, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x26, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x26, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x0d, 0x45, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x29, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x45, 0x64, 0x69, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19,
	0x5a, 0x17, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x3b, 0x70, 0x62, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_proto_books_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_books_proto_goTypes = []interface{}{
	(*FindBook)(nil),              // 0: books_service.books.FindBook
	(*Book)(nil),                  // 1: books_service.books.Book
	(*FindByRequest)(nil),         // 2: books_service.books.FindByRequest
	(*FindByResponse)(nil),        // 3: books_service.books.FindByResponse
	(*CreateBookRequest)(nil),     // 4: books_service.books.CreateBookRequest
	(*CreateBookResponse)(nil),    // 5: books_service.books.CreateBookResponse
	(*DeleteBookRequest)(nil),     // 6: books_service.books.DeleteBookRequest
	(*DeleteBookResponse)(nil),    // 7: books_service.books.DeleteBookResponse
	(*EditCountBookRequest)(nil),  // 8: books_service.books.EditCountBookRequest
	(*EditCountBookResponse)(nil), // 9: books_service.books.EditCountBookResponse
}
var file_proto_books_proto_depIdxs = []int32{
	0, // 0: books_service.books.FindByRequest.find:type_name -> books_service.books.FindBook
	1, // 1: books_service.books.FindByResponse.book:type_name -> books_service.books.Book
	1, // 2: books_service.books.CreateBookRequest.book:type_name -> books_service.books.Book
	1, // 3: books_service.books.CreateBookResponse.books:type_name -> books_service.books.Book
	2, // 4: books_service.books.BooksService.FindBy:input_type -> books_service.books.FindByRequest
	4, // 5: books_service.books.BooksService.CreateBook:input_type -> books_service.books.CreateBookRequest
	6, // 6: books_service.books.BooksService.DeleteBook:input_type -> books_service.books.DeleteBookRequest
	8, // 7: books_service.books.BooksService.EditCountBook:input_type -> books_service.books.EditCountBookRequest
	3, // 8: books_service.books.BooksService.FindBy:output_type -> books_service.books.FindByResponse
	5, // 9: books_service.books.BooksService.CreateBook:output_type -> books_service.books.CreateBookResponse
	7, // 10: books_service.books.BooksService.DeleteBook:output_type -> books_service.books.DeleteBookResponse
	9, // 11: books_service.books.BooksService.EditCountBook:output_type -> books_service.books.EditCountBookResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_books_proto_init() }
func file_proto_books_proto_init() {
	if File_proto_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindBook); i {
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
			switch v := v.(*Book); i {
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
		file_proto_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByRequest); i {
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
		file_proto_books_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByResponse); i {
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
		file_proto_books_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookRequest); i {
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
		file_proto_books_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookResponse); i {
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
		file_proto_books_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookRequest); i {
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
		file_proto_books_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookResponse); i {
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
		file_proto_books_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditCountBookRequest); i {
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
		file_proto_books_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditCountBookResponse); i {
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
			NumMessages:   10,
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
