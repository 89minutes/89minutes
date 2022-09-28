// Code generated by protoc-gen-go. DO NOT EDIT.
// source: story_message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Story struct {
	// UUID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Title of the story
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Text contents
	TextContent string `protobuf:"bytes,3,opt,name=textContent,proto3" json:"textContent,omitempty"`
	// Story can have multiple images and videos
	FileContent [][]byte `protobuf:"bytes,4,rep,name=fileContent,proto3" json:"fileContent,omitempty"`
	// Story can have formatted texts, like code
	FormattedText []string `protobuf:"bytes,5,rep,name=formattedText,proto3" json:"formattedText,omitempty"`
	// Name of the author
	WrittenBy string `protobuf:"bytes,6,opt,name=WrittenBy,proto3" json:"WrittenBy,omitempty"`
	// Is the story draft and not published
	IsDraft bool `protobuf:"varint,7,opt,name=isDraft,proto3" json:"isDraft,omitempty"`
	// Tags for categorizing stories
	Tag []*Tag `protobuf:"bytes,8,rep,name=Tag,proto3" json:"Tag,omitempty"`
	// Create time
	CreateTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// Update time
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,10,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Story) Reset()         { *m = Story{} }
func (m *Story) String() string { return proto.CompactTextString(m) }
func (*Story) ProtoMessage()    {}
func (*Story) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc724024df710507, []int{0}
}

func (m *Story) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Story.Unmarshal(m, b)
}
func (m *Story) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Story.Marshal(b, m, deterministic)
}
func (m *Story) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Story.Merge(m, src)
}
func (m *Story) XXX_Size() int {
	return xxx_messageInfo_Story.Size(m)
}
func (m *Story) XXX_DiscardUnknown() {
	xxx_messageInfo_Story.DiscardUnknown(m)
}

var xxx_messageInfo_Story proto.InternalMessageInfo

func (m *Story) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Story) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Story) GetTextContent() string {
	if m != nil {
		return m.TextContent
	}
	return ""
}

func (m *Story) GetFileContent() [][]byte {
	if m != nil {
		return m.FileContent
	}
	return nil
}

func (m *Story) GetFormattedText() []string {
	if m != nil {
		return m.FormattedText
	}
	return nil
}

func (m *Story) GetWrittenBy() string {
	if m != nil {
		return m.WrittenBy
	}
	return ""
}

func (m *Story) GetIsDraft() bool {
	if m != nil {
		return m.IsDraft
	}
	return false
}

func (m *Story) GetTag() []*Tag {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *Story) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Story) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type StoryGetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoryGetRequest) Reset()         { *m = StoryGetRequest{} }
func (m *StoryGetRequest) String() string { return proto.CompactTextString(m) }
func (*StoryGetRequest) ProtoMessage()    {}
func (*StoryGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc724024df710507, []int{1}
}

func (m *StoryGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoryGetRequest.Unmarshal(m, b)
}
func (m *StoryGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoryGetRequest.Marshal(b, m, deterministic)
}
func (m *StoryGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoryGetRequest.Merge(m, src)
}
func (m *StoryGetRequest) XXX_Size() int {
	return xxx_messageInfo_StoryGetRequest.Size(m)
}
func (m *StoryGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoryGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoryGetRequest proto.InternalMessageInfo

type StoryGetResponse struct {
	// Story will be streamed
	Top                  *Story   `protobuf:"bytes,1,opt,name=top,proto3" json:"top,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoryGetResponse) Reset()         { *m = StoryGetResponse{} }
func (m *StoryGetResponse) String() string { return proto.CompactTextString(m) }
func (*StoryGetResponse) ProtoMessage()    {}
func (*StoryGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc724024df710507, []int{2}
}

func (m *StoryGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoryGetResponse.Unmarshal(m, b)
}
func (m *StoryGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoryGetResponse.Marshal(b, m, deterministic)
}
func (m *StoryGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoryGetResponse.Merge(m, src)
}
func (m *StoryGetResponse) XXX_Size() int {
	return xxx_messageInfo_StoryGetResponse.Size(m)
}
func (m *StoryGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoryGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoryGetResponse proto.InternalMessageInfo

func (m *StoryGetResponse) GetTop() *Story {
	if m != nil {
		return m.Top
	}
	return nil
}

type Tag struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TagName              string   `protobuf:"bytes,2,opt,name=tagName,proto3" json:"tagName,omitempty"`
	TagVal               string   `protobuf:"bytes,3,opt,name=tagVal,proto3" json:"tagVal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc724024df710507, []int{3}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tag) GetTagName() string {
	if m != nil {
		return m.TagName
	}
	return ""
}

func (m *Tag) GetTagVal() string {
	if m != nil {
		return m.TagVal
	}
	return ""
}

type UpdateStoryRequest struct {
	// UUID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Title of the story
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Text contents
	TextContent string `protobuf:"bytes,3,opt,name=textContent,proto3" json:"textContent,omitempty"`
	// Update time
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,9,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpdateStoryRequest) Reset()         { *m = UpdateStoryRequest{} }
func (m *UpdateStoryRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateStoryRequest) ProtoMessage()    {}
func (*UpdateStoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc724024df710507, []int{4}
}

func (m *UpdateStoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStoryRequest.Unmarshal(m, b)
}
func (m *UpdateStoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStoryRequest.Marshal(b, m, deterministic)
}
func (m *UpdateStoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStoryRequest.Merge(m, src)
}
func (m *UpdateStoryRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateStoryRequest.Size(m)
}
func (m *UpdateStoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStoryRequest proto.InternalMessageInfo

func (m *UpdateStoryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateStoryRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateStoryRequest) GetTextContent() string {
	if m != nil {
		return m.TextContent
	}
	return ""
}

func (m *UpdateStoryRequest) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type UpdateStoryResponse struct {
	// UUID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateStoryResponse) Reset()         { *m = UpdateStoryResponse{} }
func (m *UpdateStoryResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateStoryResponse) ProtoMessage()    {}
func (*UpdateStoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc724024df710507, []int{5}
}

func (m *UpdateStoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStoryResponse.Unmarshal(m, b)
}
func (m *UpdateStoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStoryResponse.Marshal(b, m, deterministic)
}
func (m *UpdateStoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStoryResponse.Merge(m, src)
}
func (m *UpdateStoryResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateStoryResponse.Size(m)
}
func (m *UpdateStoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStoryResponse proto.InternalMessageInfo

func (m *UpdateStoryResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StoryDelete struct {
	// UUID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoryDelete) Reset()         { *m = StoryDelete{} }
func (m *StoryDelete) String() string { return proto.CompactTextString(m) }
func (*StoryDelete) ProtoMessage()    {}
func (*StoryDelete) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc724024df710507, []int{6}
}

func (m *StoryDelete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoryDelete.Unmarshal(m, b)
}
func (m *StoryDelete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoryDelete.Marshal(b, m, deterministic)
}
func (m *StoryDelete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoryDelete.Merge(m, src)
}
func (m *StoryDelete) XXX_Size() int {
	return xxx_messageInfo_StoryDelete.Size(m)
}
func (m *StoryDelete) XXX_DiscardUnknown() {
	xxx_messageInfo_StoryDelete.DiscardUnknown(m)
}

var xxx_messageInfo_StoryDelete proto.InternalMessageInfo

func (m *StoryDelete) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetLatestStoryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLatestStoryRequest) Reset()         { *m = GetLatestStoryRequest{} }
func (m *GetLatestStoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetLatestStoryRequest) ProtoMessage()    {}
func (*GetLatestStoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc724024df710507, []int{7}
}

func (m *GetLatestStoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLatestStoryRequest.Unmarshal(m, b)
}
func (m *GetLatestStoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLatestStoryRequest.Marshal(b, m, deterministic)
}
func (m *GetLatestStoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestStoryRequest.Merge(m, src)
}
func (m *GetLatestStoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetLatestStoryRequest.Size(m)
}
func (m *GetLatestStoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestStoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestStoryRequest proto.InternalMessageInfo

type GetLatestStoryResponse struct {
	Story                *Story   `protobuf:"bytes,1,opt,name=story,proto3" json:"story,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLatestStoryResponse) Reset()         { *m = GetLatestStoryResponse{} }
func (m *GetLatestStoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetLatestStoryResponse) ProtoMessage()    {}
func (*GetLatestStoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc724024df710507, []int{8}
}

func (m *GetLatestStoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLatestStoryResponse.Unmarshal(m, b)
}
func (m *GetLatestStoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLatestStoryResponse.Marshal(b, m, deterministic)
}
func (m *GetLatestStoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestStoryResponse.Merge(m, src)
}
func (m *GetLatestStoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetLatestStoryResponse.Size(m)
}
func (m *GetLatestStoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestStoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestStoryResponse proto.InternalMessageInfo

func (m *GetLatestStoryResponse) GetStory() *Story {
	if m != nil {
		return m.Story
	}
	return nil
}

func init() {
	proto.RegisterType((*Story)(nil), "Story")
	proto.RegisterType((*StoryGetRequest)(nil), "StoryGetRequest")
	proto.RegisterType((*StoryGetResponse)(nil), "StoryGetResponse")
	proto.RegisterType((*Tag)(nil), "Tag")
	proto.RegisterType((*UpdateStoryRequest)(nil), "UpdateStoryRequest")
	proto.RegisterType((*UpdateStoryResponse)(nil), "UpdateStoryResponse")
	proto.RegisterType((*StoryDelete)(nil), "StoryDelete")
	proto.RegisterType((*GetLatestStoryRequest)(nil), "GetLatestStoryRequest")
	proto.RegisterType((*GetLatestStoryResponse)(nil), "GetLatestStoryResponse")
}

func init() { proto.RegisterFile("story_message.proto", fileDescriptor_dc724024df710507) }

var fileDescriptor_dc724024df710507 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x8a, 0xd4, 0x40,
	0x10, 0xc6, 0xc9, 0x64, 0x26, 0xb3, 0xa9, 0xf8, 0xb7, 0x57, 0xc7, 0x66, 0x59, 0x31, 0x04, 0x85,
	0x1c, 0x24, 0x0b, 0x23, 0x78, 0xd0, 0xdb, 0xba, 0x30, 0x17, 0xf1, 0xd0, 0x46, 0x05, 0x2f, 0xd2,
	0x63, 0x2a, 0xa1, 0x21, 0x49, 0xc7, 0x74, 0x0d, 0xec, 0xbe, 0x89, 0x6f, 0xe7, 0xab, 0x48, 0x3a,
	0x09, 0x93, 0x99, 0x3d, 0x88, 0xb0, 0xc7, 0xfa, 0xea, 0xab, 0xf4, 0x57, 0xbf, 0x22, 0x70, 0x6a,
	0x48, 0xb7, 0x37, 0x3f, 0x2a, 0x34, 0x46, 0x16, 0x98, 0x34, 0xad, 0x26, 0x7d, 0xf6, 0xa2, 0xd0,
	0xba, 0x28, 0xf1, 0xc2, 0x56, 0xdb, 0x5d, 0x7e, 0x41, 0xaa, 0x42, 0x43, 0xb2, 0x6a, 0x7a, 0x43,
	0xf4, 0x67, 0x06, 0x8b, 0xcf, 0xdd, 0x20, 0x7b, 0x00, 0x33, 0x95, 0x71, 0x27, 0x74, 0x62, 0x5f,
	0xcc, 0x54, 0xc6, 0x9e, 0xc0, 0x82, 0x14, 0x95, 0xc8, 0x67, 0x56, 0xea, 0x0b, 0x16, 0x42, 0x40,
	0x78, 0x4d, 0x1f, 0x74, 0x4d, 0x58, 0x13, 0x77, 0x6d, 0x6f, 0x2a, 0x75, 0x8e, 0x5c, 0x95, 0x38,
	0x3a, 0xe6, 0xa1, 0x1b, 0xdf, 0x13, 0x53, 0x89, 0xbd, 0x84, 0xfb, 0xb9, 0x6e, 0x2b, 0x49, 0x84,
	0x59, 0x8a, 0xd7, 0xc4, 0x17, 0xa1, 0x1b, 0xfb, 0xe2, 0x50, 0x64, 0xe7, 0xe0, 0x7f, 0x6b, 0x15,
	0x11, 0xd6, 0x97, 0x37, 0xdc, 0xb3, 0xef, 0xec, 0x05, 0xc6, 0x61, 0xa9, 0xcc, 0x55, 0x2b, 0x73,
	0xe2, 0xcb, 0xd0, 0x89, 0x4f, 0xc4, 0x58, 0xb2, 0x15, 0xb8, 0xa9, 0x2c, 0xf8, 0x49, 0xe8, 0xc6,
	0xc1, 0x7a, 0x9e, 0xa4, 0xb2, 0x10, 0x9d, 0xc0, 0xde, 0x01, 0xfc, 0x6c, 0x51, 0x12, 0xa6, 0xaa,
	0x42, 0xee, 0x87, 0x4e, 0x1c, 0xac, 0xcf, 0x92, 0x9e, 0x4f, 0x32, 0xf2, 0x49, 0xd2, 0x91, 0x8f,
	0x98, 0xb8, 0xbb, 0xd9, 0x5d, 0x93, 0x8d, 0xb3, 0xf0, 0xef, 0xd9, 0xbd, 0x3b, 0x7a, 0x0c, 0x0f,
	0x2d, 0xe0, 0x0d, 0x92, 0xc0, 0x5f, 0x3b, 0x34, 0x14, 0xbd, 0x86, 0x47, 0x7b, 0xc9, 0x34, 0xba,
	0x36, 0xc8, 0x38, 0xb8, 0xa4, 0x1b, 0xcb, 0x3f, 0x58, 0x7b, 0x89, 0xed, 0x8b, 0x4e, 0x8a, 0x36,
	0x76, 0xa1, 0x5b, 0xf7, 0xe1, 0xb0, 0x24, 0x59, 0x7c, 0x92, 0xd5, 0x78, 0xa1, 0xb1, 0x64, 0x2b,
	0xf0, 0x48, 0x16, 0x5f, 0x65, 0x39, 0x9c, 0x67, 0xa8, 0xa2, 0xdf, 0x0e, 0xb0, 0x2f, 0x36, 0x58,
	0xff, 0xf5, 0x3e, 0xcd, 0x9d, 0x1d, 0xfe, 0x10, 0x92, 0xff, 0x5f, 0x90, 0x5e, 0xc1, 0xe9, 0x41,
	0xb2, 0x01, 0xca, 0x51, 0xb4, 0xe8, 0x39, 0x04, 0xd6, 0x70, 0x85, 0x25, 0xd2, 0xed, 0xf6, 0x33,
	0x78, 0xba, 0x41, 0xfa, 0x28, 0x09, 0x0d, 0x4d, 0x57, 0x8c, 0xde, 0xc2, 0xea, 0xb8, 0x31, 0xbc,
	0x70, 0x0e, 0x0b, 0xfb, 0xdf, 0x1c, 0x81, 0xef, 0xc5, 0x4b, 0xef, 0xfb, 0x3c, 0x79, 0xdf, 0x6c,
	0xb7, 0x9e, 0x8d, 0xff, 0xe6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0x19, 0xf1, 0x89, 0x64,
	0x03, 0x00, 0x00,
}
