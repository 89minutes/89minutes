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

func init() {
	proto.RegisterType((*Story)(nil), "Story")
	proto.RegisterType((*StoryGetRequest)(nil), "StoryGetRequest")
	proto.RegisterType((*StoryGetResponse)(nil), "StoryGetResponse")
	proto.RegisterType((*Tag)(nil), "Tag")
	proto.RegisterType((*UpdateStoryRequest)(nil), "UpdateStoryRequest")
	proto.RegisterType((*UpdateStoryResponse)(nil), "UpdateStoryResponse")
	proto.RegisterType((*StoryDelete)(nil), "StoryDelete")
}

func init() { proto.RegisterFile("story_message.proto", fileDescriptor_dc724024df710507) }

var fileDescriptor_dc724024df710507 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x4f, 0xab, 0xd3, 0x40,
	0x14, 0xc5, 0x49, 0xd2, 0xa6, 0x2f, 0x37, 0xfe, 0x9d, 0x27, 0x8f, 0xa1, 0x28, 0x86, 0xa0, 0x90,
	0x85, 0xa4, 0x50, 0x77, 0xba, 0xab, 0x85, 0xee, 0x5c, 0x8c, 0x51, 0xc1, 0x8d, 0x4c, 0xcd, 0x4d,
	0x18, 0x48, 0x32, 0x31, 0x73, 0x0b, 0xed, 0x37, 0xf1, 0xdb, 0xf9, 0x55, 0x24, 0x93, 0x86, 0xb6,
	0x76, 0x21, 0xc2, 0x5b, 0x9e, 0x73, 0xcf, 0xcd, 0xdc, 0xfc, 0x0e, 0xdc, 0x1a, 0xd2, 0xdd, 0xe1,
	0x7b, 0x8d, 0xc6, 0xc8, 0x12, 0xd3, 0xb6, 0xd3, 0xa4, 0xe7, 0x2f, 0x4b, 0xad, 0xcb, 0x0a, 0x17,
	0x56, 0x6d, 0x77, 0xc5, 0x82, 0x54, 0x8d, 0x86, 0x64, 0xdd, 0x0e, 0x81, 0xf8, 0xb7, 0x0b, 0xd3,
	0x4f, 0xfd, 0x22, 0x7b, 0x04, 0xae, 0xca, 0xb9, 0x13, 0x39, 0x49, 0x20, 0x5c, 0x95, 0xb3, 0x67,
	0x30, 0x25, 0x45, 0x15, 0x72, 0xd7, 0x5a, 0x83, 0x60, 0x11, 0x84, 0x84, 0x7b, 0xfa, 0xa0, 0x1b,
	0xc2, 0x86, 0xb8, 0x67, 0x67, 0xe7, 0x56, 0x9f, 0x28, 0x54, 0x85, 0x63, 0x62, 0x12, 0x79, 0xc9,
	0x03, 0x71, 0x6e, 0xb1, 0x57, 0xf0, 0xb0, 0xd0, 0x5d, 0x2d, 0x89, 0x30, 0xcf, 0x70, 0x4f, 0x7c,
	0x1a, 0x79, 0x49, 0x20, 0x2e, 0x4d, 0xf6, 0x1c, 0x82, 0xaf, 0x9d, 0x22, 0xc2, 0x66, 0x75, 0xe0,
	0xbe, 0x7d, 0xe7, 0x64, 0x30, 0x0e, 0x33, 0x65, 0xd6, 0x9d, 0x2c, 0x88, 0xcf, 0x22, 0x27, 0xb9,
	0x11, 0xa3, 0x64, 0x77, 0xe0, 0x65, 0xb2, 0xe4, 0x37, 0x91, 0x97, 0x84, 0xcb, 0x49, 0x9a, 0xc9,
	0x52, 0xf4, 0x06, 0x7b, 0x07, 0xf0, 0xa3, 0x43, 0x49, 0x98, 0xa9, 0x1a, 0x79, 0x10, 0x39, 0x49,
	0xb8, 0x9c, 0xa7, 0x03, 0x9f, 0x74, 0xe4, 0x93, 0x66, 0x23, 0x1f, 0x71, 0x96, 0xee, 0x77, 0x77,
	0x6d, 0x3e, 0xee, 0xc2, 0xbf, 0x77, 0x4f, 0xe9, 0xf8, 0x29, 0x3c, 0xb6, 0x80, 0x37, 0x48, 0x02,
	0x7f, 0xee, 0xd0, 0x50, 0xfc, 0x06, 0x9e, 0x9c, 0x2c, 0xd3, 0xea, 0xc6, 0x20, 0xe3, 0xe0, 0x91,
	0x6e, 0x2d, 0xff, 0x70, 0xe9, 0xa7, 0x76, 0x2e, 0x7a, 0x2b, 0xde, 0xd8, 0x1f, 0xba, 0xea, 0x87,
	0xc3, 0x8c, 0x64, 0xf9, 0x51, 0xd6, 0x63, 0x43, 0xa3, 0x64, 0x77, 0xe0, 0x93, 0x2c, 0xbf, 0xc8,
	0xea, 0x58, 0xcf, 0x51, 0xc5, 0xbf, 0x1c, 0x60, 0x9f, 0xed, 0x61, 0xc3, 0xd7, 0x87, 0x6b, 0xee,
	0xad, 0xf8, 0x4b, 0x48, 0xc1, 0x7f, 0x41, 0x7a, 0x0d, 0xb7, 0x17, 0x97, 0x1d, 0xa1, 0xfc, 0x75,
	0x5a, 0xfc, 0x02, 0x42, 0x1b, 0x58, 0x63, 0x85, 0x74, 0x35, 0x5e, 0xf9, 0xdf, 0x26, 0xe9, 0xfb,
	0x76, 0xbb, 0xf5, 0xed, 0x6b, 0x6f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x37, 0xf4, 0xba, 0x83,
	0x13, 0x03, 0x00, 0x00,
}
