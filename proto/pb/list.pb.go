// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v5.29.2
// source: list.proto

package __

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

type CreateListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId string `protobuf:"bytes,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateListRequest) Reset() {
	*x = CreateListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateListRequest) ProtoMessage() {}

func (x *CreateListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateListRequest.ProtoReflect.Descriptor instead.
func (*CreateListRequest) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{0}
}

func (x *CreateListRequest) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *CreateListRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListId string `protobuf:"bytes,1,opt,name=list_id,json=listId,proto3" json:"list_id,omitempty"`
}

func (x *CreateListResponse) Reset() {
	*x = CreateListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateListResponse) ProtoMessage() {}

func (x *CreateListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateListResponse.ProtoReflect.Descriptor instead.
func (*CreateListResponse) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{1}
}

func (x *CreateListResponse) GetListId() string {
	if x != nil {
		return x.ListId
	}
	return ""
}

type GetListsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId string `protobuf:"bytes,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
}

func (x *GetListsRequest) Reset() {
	*x = GetListsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListsRequest) ProtoMessage() {}

func (x *GetListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListsRequest.ProtoReflect.Descriptor instead.
func (*GetListsRequest) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{2}
}

func (x *GetListsRequest) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

type GetListsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lists []*List `protobuf:"bytes,1,rep,name=lists,proto3" json:"lists,omitempty"`
}

func (x *GetListsResponse) Reset() {
	*x = GetListsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListsResponse) ProtoMessage() {}

func (x *GetListsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListsResponse.ProtoReflect.Descriptor instead.
func (*GetListsResponse) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{3}
}

func (x *GetListsResponse) GetLists() []*List {
	if x != nil {
		return x.Lists
	}
	return nil
}

type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	BoardId string `protobuf:"bytes,3,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{4}
}

func (x *List) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *List) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *List) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

var File_list_proto protoreflect.FileDescriptor

var file_list_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x44, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x2d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x6c, 0x69, 0x73,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x22, 0x47, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x49, 0x64, 0x32, 0x89, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x17, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x73, 0x12, 0x15, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_list_proto_rawDescOnce sync.Once
	file_list_proto_rawDescData = file_list_proto_rawDesc
)

func file_list_proto_rawDescGZIP() []byte {
	file_list_proto_rawDescOnce.Do(func() {
		file_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_list_proto_rawDescData)
	})
	return file_list_proto_rawDescData
}

var file_list_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_list_proto_goTypes = []interface{}{
	(*CreateListRequest)(nil),  // 0: list.CreateListRequest
	(*CreateListResponse)(nil), // 1: list.CreateListResponse
	(*GetListsRequest)(nil),    // 2: list.GetListsRequest
	(*GetListsResponse)(nil),   // 3: list.GetListsResponse
	(*List)(nil),               // 4: list.List
}
var file_list_proto_depIdxs = []int32{
	4, // 0: list.GetListsResponse.lists:type_name -> list.List
	0, // 1: list.ListService.CreateList:input_type -> list.CreateListRequest
	2, // 2: list.ListService.GetLists:input_type -> list.GetListsRequest
	1, // 3: list.ListService.CreateList:output_type -> list.CreateListResponse
	3, // 4: list.ListService.GetLists:output_type -> list.GetListsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_list_proto_init() }
func file_list_proto_init() {
	if File_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateListRequest); i {
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
		file_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateListResponse); i {
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
		file_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListsRequest); i {
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
		file_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListsResponse); i {
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
		file_list_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*List); i {
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
			RawDescriptor: file_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_list_proto_goTypes,
		DependencyIndexes: file_list_proto_depIdxs,
		MessageInfos:      file_list_proto_msgTypes,
	}.Build()
	File_list_proto = out.File
	file_list_proto_rawDesc = nil
	file_list_proto_goTypes = nil
	file_list_proto_depIdxs = nil
}
