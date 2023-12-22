// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: proto/book/book.proto

package bookv1

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type NewBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title         string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author        []string             `protobuf:"bytes,2,rep,name=author,proto3" json:"author,omitempty"`
	Language      string               `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
	PublishedDate *timestamp.Timestamp `protobuf:"bytes,4,opt,name=published_date,json=publishedDate,proto3" json:"published_date,omitempty"`
	CoverPage     []byte               `protobuf:"bytes,5,opt,name=cover_page,json=coverPage,proto3" json:"cover_page,omitempty"`
}

func (x *NewBookRequest) Reset() {
	*x = NewBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBookRequest) ProtoMessage() {}

func (x *NewBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBookRequest.ProtoReflect.Descriptor instead.
func (*NewBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_book_proto_rawDescGZIP(), []int{0}
}

func (x *NewBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewBookRequest) GetAuthor() []string {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *NewBookRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *NewBookRequest) GetPublishedDate() *timestamp.Timestamp {
	if x != nil {
		return x.PublishedDate
	}
	return nil
}

func (x *NewBookRequest) GetCoverPage() []byte {
	if x != nil {
		return x.CoverPage
	}
	return nil
}

type NewBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *NewBookResponse) Reset() {
	*x = NewBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBookResponse) ProtoMessage() {}

func (x *NewBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBookResponse.ProtoReflect.Descriptor instead.
func (*NewBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_book_book_proto_rawDescGZIP(), []int{1}
}

func (x *NewBookResponse) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author        []string             `protobuf:"bytes,3,rep,name=author,proto3" json:"author,omitempty"`
	Language      string               `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	PublishedDate *timestamp.Timestamp `protobuf:"bytes,5,opt,name=published_date,json=publishedDate,proto3" json:"published_date,omitempty"`
	CoverPage     []byte               `protobuf:"bytes,6,opt,name=cover_page,json=coverPage,proto3" json:"cover_page,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_book_proto_msgTypes[2]
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
	return file_proto_book_book_proto_rawDescGZIP(), []int{2}
}

func (x *Book) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() []string {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Book) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Book) GetPublishedDate() *timestamp.Timestamp {
	if x != nil {
		return x.PublishedDate
	}
	return nil
}

func (x *Book) GetCoverPage() []byte {
	if x != nil {
		return x.CoverPage
	}
	return nil
}

var File_proto_book_book_proto protoreflect.FileDescriptor

var file_proto_book_book_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc,
	0x01, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a,
	0x0f, 0x4e, 0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b,
	0x22, 0xc2, 0x01, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x50, 0x61, 0x67, 0x65, 0x32, 0x45, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4e, 0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4e, 0x65, 0x77,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a,
	0x70, 0x64, 0x70, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x76, 0x31, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_book_book_proto_rawDescOnce sync.Once
	file_proto_book_book_proto_rawDescData = file_proto_book_book_proto_rawDesc
)

func file_proto_book_book_proto_rawDescGZIP() []byte {
	file_proto_book_book_proto_rawDescOnce.Do(func() {
		file_proto_book_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_book_book_proto_rawDescData)
	})
	return file_proto_book_book_proto_rawDescData
}

var file_proto_book_book_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_book_book_proto_goTypes = []interface{}{
	(*NewBookRequest)(nil),      // 0: book.NewBookRequest
	(*NewBookResponse)(nil),     // 1: book.NewBookResponse
	(*Book)(nil),                // 2: book.Book
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_proto_book_book_proto_depIdxs = []int32{
	3, // 0: book.NewBookRequest.published_date:type_name -> google.protobuf.Timestamp
	2, // 1: book.NewBookResponse.book:type_name -> book.Book
	3, // 2: book.Book.published_date:type_name -> google.protobuf.Timestamp
	0, // 3: book.BookService.NewBook:input_type -> book.NewBookRequest
	1, // 4: book.BookService.NewBook:output_type -> book.NewBookResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_book_book_proto_init() }
func file_proto_book_book_proto_init() {
	if File_proto_book_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_book_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewBookRequest); i {
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
		file_proto_book_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewBookResponse); i {
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
		file_proto_book_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_book_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_book_book_proto_goTypes,
		DependencyIndexes: file_proto_book_book_proto_depIdxs,
		MessageInfos:      file_proto_book_book_proto_msgTypes,
	}.Build()
	File_proto_book_book_proto = out.File
	file_proto_book_book_proto_rawDesc = nil
	file_proto_book_book_proto_goTypes = nil
	file_proto_book_book_proto_depIdxs = nil
}
