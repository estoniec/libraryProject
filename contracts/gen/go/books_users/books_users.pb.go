// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.0
// source: proto/books_users.proto

package pb_books_users

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
		mi := &file_proto_books_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_users_proto_msgTypes[0]
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
	return file_proto_books_users_proto_rawDescGZIP(), []int{0}
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Phone    string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Class    string `protobuf:"bytes,4,opt,name=class,proto3" json:"class,omitempty"`
	Status   int64  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_books_users_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *User) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type BooksUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Book     *Book `protobuf:"bytes,2,opt,name=Book,proto3" json:"Book,omitempty"`
	User     *User `protobuf:"bytes,3,opt,name=User,proto3" json:"User,omitempty"`
	IsReturn bool  `protobuf:"varint,4,opt,name=isReturn,proto3" json:"isReturn,omitempty"`
	IsGet    bool  `protobuf:"varint,5,opt,name=isGet,proto3" json:"isGet,omitempty"`
}

func (x *BooksUsers) Reset() {
	*x = BooksUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksUsers) ProtoMessage() {}

func (x *BooksUsers) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksUsers.ProtoReflect.Descriptor instead.
func (*BooksUsers) Descriptor() ([]byte, []int) {
	return file_proto_books_users_proto_rawDescGZIP(), []int{2}
}

func (x *BooksUsers) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *BooksUsers) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

func (x *BooksUsers) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *BooksUsers) GetIsReturn() bool {
	if x != nil {
		return x.IsReturn
	}
	return false
}

func (x *BooksUsers) GetIsGet() bool {
	if x != nil {
		return x.IsGet
	}
	return false
}

type RentBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookID   int64 `protobuf:"varint,1,opt,name=bookID,proto3" json:"bookID,omitempty"`
	UserID   int64 `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	ReturnAt int64 `protobuf:"varint,3,opt,name=returnAt,proto3" json:"returnAt,omitempty"`
}

func (x *RentBookRequest) Reset() {
	*x = RentBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RentBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RentBookRequest) ProtoMessage() {}

func (x *RentBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RentBookRequest.ProtoReflect.Descriptor instead.
func (*RentBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_books_users_proto_rawDescGZIP(), []int{3}
}

func (x *RentBookRequest) GetBookID() int64 {
	if x != nil {
		return x.BookID
	}
	return 0
}

func (x *RentBookRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *RentBookRequest) GetReturnAt() int64 {
	if x != nil {
		return x.ReturnAt
	}
	return 0
}

type RentBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Status int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RentBookResponse) Reset() {
	*x = RentBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RentBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RentBookResponse) ProtoMessage() {}

func (x *RentBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RentBookResponse.ProtoReflect.Descriptor instead.
func (*RentBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_books_users_proto_rawDescGZIP(), []int{4}
}

func (x *RentBookResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *RentBookResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_proto_books_users_proto protoreflect.FileDescriptor

var file_proto_books_users_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x6c, 0x0a, 0x04, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x53, 0x42, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x49, 0x53, 0x42, 0x4e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x76, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xc4, 0x01, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x39, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x39, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x47, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x69, 0x73, 0x47, 0x65, 0x74, 0x22, 0x5d, 0x0a, 0x0f, 0x52, 0x65, 0x6e, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f,
	0x6f, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x41, 0x74, 0x22, 0x40, 0x0a, 0x10, 0x52, 0x65, 0x6e, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x86, 0x01, 0x0a, 0x11, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x55, 0x73, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71,
	0x0a, 0x08, 0x52, 0x65, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x30, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x6e,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x52,
	0x65, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x25, 0x5a, 0x23, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3b, 0x70, 0x62, 0x5f, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_books_users_proto_rawDescOnce sync.Once
	file_proto_books_users_proto_rawDescData = file_proto_books_users_proto_rawDesc
)

func file_proto_books_users_proto_rawDescGZIP() []byte {
	file_proto_books_users_proto_rawDescOnce.Do(func() {
		file_proto_books_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_books_users_proto_rawDescData)
	})
	return file_proto_books_users_proto_rawDescData
}

var file_proto_books_users_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_books_users_proto_goTypes = []interface{}{
	(*Book)(nil),             // 0: books_users_service.books_users.Book
	(*User)(nil),             // 1: books_users_service.books_users.User
	(*BooksUsers)(nil),       // 2: books_users_service.books_users.BooksUsers
	(*RentBookRequest)(nil),  // 3: books_users_service.books_users.RentBookRequest
	(*RentBookResponse)(nil), // 4: books_users_service.books_users.RentBookResponse
}
var file_proto_books_users_proto_depIdxs = []int32{
	0, // 0: books_users_service.books_users.BooksUsers.Book:type_name -> books_users_service.books_users.Book
	1, // 1: books_users_service.books_users.BooksUsers.User:type_name -> books_users_service.books_users.User
	3, // 2: books_users_service.books_users.BooksUsersService.RentBook:input_type -> books_users_service.books_users.RentBookRequest
	4, // 3: books_users_service.books_users.BooksUsersService.RentBook:output_type -> books_users_service.books_users.RentBookResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_books_users_proto_init() }
func file_proto_books_users_proto_init() {
	if File_proto_books_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_books_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_books_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_books_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksUsers); i {
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
		file_proto_books_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RentBookRequest); i {
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
		file_proto_books_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RentBookResponse); i {
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
			RawDescriptor: file_proto_books_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_books_users_proto_goTypes,
		DependencyIndexes: file_proto_books_users_proto_depIdxs,
		MessageInfos:      file_proto_books_users_proto_msgTypes,
	}.Build()
	File_proto_books_users_proto = out.File
	file_proto_books_users_proto_rawDesc = nil
	file_proto_books_users_proto_goTypes = nil
	file_proto_books_users_proto_depIdxs = nil
}