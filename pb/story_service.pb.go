// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: story_service.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateStoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Story *Story `protobuf:"bytes,1,opt,name=story,proto3" json:"story,omitempty"`
}

func (x *CreateStoryRequest) Reset() {
	*x = CreateStoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_story_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStoryRequest) ProtoMessage() {}

func (x *CreateStoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_story_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStoryRequest.ProtoReflect.Descriptor instead.
func (*CreateStoryRequest) Descriptor() ([]byte, []int) {
	return file_story_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStoryRequest) GetStory() *Story {
	if x != nil {
		return x.Story
	}
	return nil
}

type CreateStoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateStoryResponse) Reset() {
	*x = CreateStoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_story_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStoryResponse) ProtoMessage() {}

func (x *CreateStoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_story_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStoryResponse.ProtoReflect.Descriptor instead.
func (*CreateStoryResponse) Descriptor() ([]byte, []int) {
	return file_story_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStoryResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_story_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_story_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_story_service_proto_rawDescGZIP(), []int{2}
}

func (x *Ping) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_story_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_story_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_story_service_proto_rawDescGZIP(), []int{3}
}

func (x *Pong) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

var File_story_service_proto protoreflect.FileDescriptor

var file_story_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x25, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22,
	0x1a, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x32, 0x96, 0x02, 0x0a, 0x0c,
	0x53, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x08,
	0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x05, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a,
	0x05, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12,
	0x2f, 0x76, 0x31, 0x2f, 0x38, 0x39, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x69,
	0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x12, 0x54, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x38, 0x39, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65,
	0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x40, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x35, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_story_service_proto_rawDescOnce sync.Once
	file_story_service_proto_rawDescData = file_story_service_proto_rawDesc
)

func file_story_service_proto_rawDescGZIP() []byte {
	file_story_service_proto_rawDescOnce.Do(func() {
		file_story_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_story_service_proto_rawDescData)
	})
	return file_story_service_proto_rawDescData
}

var file_story_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_story_service_proto_goTypes = []interface{}{
	(*CreateStoryRequest)(nil),     // 0: CreateStoryRequest
	(*CreateStoryResponse)(nil),    // 1: CreateStoryResponse
	(*Ping)(nil),                   // 2: Ping
	(*Pong)(nil),                   // 3: Pong
	(*Story)(nil),                  // 4: Story
	(*GetLatestStoryRequest)(nil),  // 5: GetLatestStoryRequest
	(*UpdateStoryRequest)(nil),     // 6: UpdateStoryRequest
	(*GetLatestStoryResponse)(nil), // 7: GetLatestStoryResponse
	(*UpdateStoryResponse)(nil),    // 8: UpdateStoryResponse
}
var file_story_service_proto_depIdxs = []int32{
	4, // 0: CreateStoryRequest.story:type_name -> Story
	2, // 1: StoryService.PingPong:input_type -> Ping
	0, // 2: StoryService.Create:input_type -> CreateStoryRequest
	5, // 3: StoryService.GetLatest:input_type -> GetLatestStoryRequest
	6, // 4: StoryService.Update:input_type -> UpdateStoryRequest
	3, // 5: StoryService.PingPong:output_type -> Pong
	1, // 6: StoryService.Create:output_type -> CreateStoryResponse
	7, // 7: StoryService.GetLatest:output_type -> GetLatestStoryResponse
	8, // 8: StoryService.Update:output_type -> UpdateStoryResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_story_service_proto_init() }
func file_story_service_proto_init() {
	if File_story_service_proto != nil {
		return
	}
	file_story_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_story_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStoryRequest); i {
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
		file_story_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStoryResponse); i {
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
		file_story_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_story_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
			RawDescriptor: file_story_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_story_service_proto_goTypes,
		DependencyIndexes: file_story_service_proto_depIdxs,
		MessageInfos:      file_story_service_proto_msgTypes,
	}.Build()
	File_story_service_proto = out.File
	file_story_service_proto_rawDesc = nil
	file_story_service_proto_goTypes = nil
	file_story_service_proto_depIdxs = nil
}
