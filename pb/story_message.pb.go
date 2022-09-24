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

type StoryCreateRequest struct {
	// UUID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Title of the story
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Text contents
	TextContent string `protobuf:"bytes,3,opt,name=textContent,proto3" json:"textContent,omitempty"`
	// Story can have multiple images and videos
	//
	// Types that are valid to be assigned to Data:
	//	*StoryCreateRequest_Info
	//	*StoryCreateRequest_ChunkData
	Data isStoryCreateRequest_Data `protobuf_oneof:"data"`
	// Story can have formatted texts, like code
	FormattedText []string `protobuf:"bytes,5,rep,name=formattedText,proto3" json:"formattedText,omitempty"`
	// Name of the author
	WrittenBy string `protobuf:"bytes,6,opt,name=WrittenBy,proto3" json:"WrittenBy,omitempty"`
	// Is the story draft and not published
	IsDraft bool `protobuf:"varint,7,opt,name=isDraft,proto3" json:"isDraft,omitempty"`
	// Create time
	CreateTime *timestamp.Timestamp `protobuf:"bytes,8,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// Update time
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,9,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StoryCreateRequest) Reset()         { *m = StoryCreateRequest{} }
func (m *StoryCreateRequest) String() string { return proto.CompactTextString(m) }
func (*StoryCreateRequest) ProtoMessage()    {}
func (*StoryCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc724024df710507, []int{0}
}

func (m *StoryCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoryCreateRequest.Unmarshal(m, b)
}
func (m *StoryCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoryCreateRequest.Marshal(b, m, deterministic)
}
func (m *StoryCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoryCreateRequest.Merge(m, src)
}
func (m *StoryCreateRequest) XXX_Size() int {
	return xxx_messageInfo_StoryCreateRequest.Size(m)
}
func (m *StoryCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoryCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoryCreateRequest proto.InternalMessageInfo

func (m *StoryCreateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StoryCreateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *StoryCreateRequest) GetTextContent() string {
	if m != nil {
		return m.TextContent
	}
	return ""
}

type isStoryCreateRequest_Data interface {
	isStoryCreateRequest_Data()
}

type StoryCreateRequest_Info struct {
	Info *ImageInfo `protobuf:"bytes,4,opt,name=info,proto3,oneof"`
}

type StoryCreateRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,10,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*StoryCreateRequest_Info) isStoryCreateRequest_Data() {}

func (*StoryCreateRequest_ChunkData) isStoryCreateRequest_Data() {}

func (m *StoryCreateRequest) GetData() isStoryCreateRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *StoryCreateRequest) GetInfo() *ImageInfo {
	if x, ok := m.GetData().(*StoryCreateRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (m *StoryCreateRequest) GetChunkData() []byte {
	if x, ok := m.GetData().(*StoryCreateRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

func (m *StoryCreateRequest) GetFormattedText() []string {
	if m != nil {
		return m.FormattedText
	}
	return nil
}

func (m *StoryCreateRequest) GetWrittenBy() string {
	if m != nil {
		return m.WrittenBy
	}
	return ""
}

func (m *StoryCreateRequest) GetIsDraft() bool {
	if m != nil {
		return m.IsDraft
	}
	return false
}

func (m *StoryCreateRequest) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *StoryCreateRequest) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StoryCreateRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StoryCreateRequest_Info)(nil),
		(*StoryCreateRequest_ChunkData)(nil),
	}
}

type ImageInfo struct {
	LaptopId             string   `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	ImageType            string   `protobuf:"bytes,2,opt,name=image_type,json=imageType,proto3" json:"image_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageInfo) Reset()         { *m = ImageInfo{} }
func (m *ImageInfo) String() string { return proto.CompactTextString(m) }
func (*ImageInfo) ProtoMessage()    {}
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc724024df710507, []int{1}
}

func (m *ImageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageInfo.Unmarshal(m, b)
}
func (m *ImageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageInfo.Marshal(b, m, deterministic)
}
func (m *ImageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageInfo.Merge(m, src)
}
func (m *ImageInfo) XXX_Size() int {
	return xxx_messageInfo_ImageInfo.Size(m)
}
func (m *ImageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ImageInfo proto.InternalMessageInfo

func (m *ImageInfo) GetLaptopId() string {
	if m != nil {
		return m.LaptopId
	}
	return ""
}

func (m *ImageInfo) GetImageType() string {
	if m != nil {
		return m.ImageType
	}
	return ""
}

type StoryCreateResponse struct {
	// UUID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoryCreateResponse) Reset()         { *m = StoryCreateResponse{} }
func (m *StoryCreateResponse) String() string { return proto.CompactTextString(m) }
func (*StoryCreateResponse) ProtoMessage()    {}
func (*StoryCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc724024df710507, []int{2}
}

func (m *StoryCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoryCreateResponse.Unmarshal(m, b)
}
func (m *StoryCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoryCreateResponse.Marshal(b, m, deterministic)
}
func (m *StoryCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoryCreateResponse.Merge(m, src)
}
func (m *StoryCreateResponse) XXX_Size() int {
	return xxx_messageInfo_StoryCreateResponse.Size(m)
}
func (m *StoryCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoryCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoryCreateResponse proto.InternalMessageInfo

func (m *StoryCreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*StoryCreateRequest)(nil), "StoryCreateRequest")
	proto.RegisterType((*ImageInfo)(nil), "ImageInfo")
	proto.RegisterType((*StoryCreateResponse)(nil), "StoryCreateResponse")
}

func init() { proto.RegisterFile("story_message.proto", fileDescriptor_dc724024df710507) }

var fileDescriptor_dc724024df710507 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x8b, 0xd4, 0x40,
	0x10, 0xc5, 0x37, 0xd9, 0x6c, 0x76, 0xba, 0x56, 0x3d, 0xf4, 0x7a, 0x68, 0x56, 0x65, 0xc3, 0xa0,
	0x90, 0x53, 0x06, 0xf4, 0xa6, 0xb7, 0x99, 0x01, 0x67, 0xae, 0x31, 0x20, 0x78, 0x09, 0x3d, 0x93,
	0x4a, 0x6c, 0x9c, 0xa4, 0xdb, 0x74, 0x05, 0x26, 0x1f, 0x5e, 0x90, 0x74, 0xc8, 0xfc, 0xf1, 0xe2,
	0xb1, 0x7e, 0xf5, 0x1e, 0x54, 0xbd, 0x07, 0x8f, 0x96, 0x74, 0xdb, 0xe7, 0x35, 0x5a, 0x2b, 0x2b,
	0x4c, 0x4c, 0xab, 0x49, 0x3f, 0x3d, 0x57, 0x5a, 0x57, 0x07, 0x5c, 0xb8, 0x69, 0xd7, 0x95, 0x0b,
	0x52, 0x35, 0x5a, 0x92, 0xb5, 0x19, 0x05, 0xf3, 0x3f, 0x3e, 0xf0, 0x6f, 0x83, 0x71, 0xd5, 0xa2,
	0x24, 0x4c, 0xf1, 0x77, 0x87, 0x96, 0xf8, 0x2b, 0xf0, 0x55, 0x21, 0xbc, 0xc8, 0x8b, 0x59, 0xea,
	0xab, 0x82, 0xbf, 0x86, 0x3b, 0x52, 0x74, 0x40, 0xe1, 0x3b, 0x34, 0x0e, 0x3c, 0x82, 0x07, 0xc2,
	0x23, 0xad, 0x74, 0x43, 0xd8, 0x90, 0xb8, 0x75, 0xbb, 0x4b, 0xc4, 0x23, 0x08, 0x54, 0x53, 0x6a,
	0x11, 0x44, 0x5e, 0xfc, 0xf0, 0x11, 0x92, 0x6d, 0x2d, 0x2b, 0xdc, 0x36, 0xa5, 0xde, 0xdc, 0xa4,
	0x6e, 0xc3, 0x9f, 0x01, 0xf6, 0x3f, 0xbb, 0xe6, 0x57, 0x5e, 0x48, 0x92, 0x02, 0x22, 0x2f, 0x7e,
	0xb1, 0xb9, 0x49, 0x99, 0x63, 0x6b, 0x49, 0x92, 0xbf, 0x87, 0x97, 0xa5, 0x6e, 0x6b, 0x49, 0x84,
	0x45, 0x86, 0x47, 0x12, 0x77, 0xd1, 0x6d, 0xcc, 0xd2, 0x6b, 0xc8, 0xdf, 0x02, 0xfb, 0xde, 0x2a,
	0x22, 0x6c, 0x96, 0xbd, 0x08, 0xdd, 0x21, 0x67, 0xc0, 0x05, 0xdc, 0x2b, 0xbb, 0x6e, 0x65, 0x49,
	0xe2, 0x3e, 0xf2, 0xe2, 0x59, 0x3a, 0x8d, 0xfc, 0x33, 0xc0, 0xde, 0x7d, 0x9e, 0xa9, 0x1a, 0xc5,
	0xcc, 0x9d, 0xf9, 0x94, 0x8c, 0xa9, 0x25, 0x53, 0x6a, 0x49, 0x36, 0xa5, 0x96, 0x5e, 0xa8, 0x07,
	0x6f, 0x67, 0x8a, 0xc9, 0xcb, 0xfe, 0xef, 0x3d, 0xab, 0x97, 0x21, 0x04, 0xc3, 0xc3, 0xf3, 0xaf,
	0xc0, 0x4e, 0x99, 0xf0, 0x37, 0xc0, 0x0e, 0xd2, 0x90, 0x36, 0xf9, 0x29, 0xfc, 0xd9, 0x08, 0xb6,
	0x05, 0x7f, 0x07, 0xa0, 0x06, 0x65, 0x4e, 0xbd, 0x99, 0x7a, 0x60, 0x8e, 0x64, 0xbd, 0xc1, 0xf9,
	0x07, 0x78, 0xbc, 0xea, 0xd1, 0x1a, 0xdd, 0x58, 0xfc, 0xb7, 0xc8, 0x65, 0xf8, 0x23, 0x48, 0xbe,
	0x98, 0xdd, 0x2e, 0x74, 0xf7, 0x7d, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x83, 0xbf, 0x5d, 0x13,
	0x36, 0x02, 0x00, 0x00,
}
