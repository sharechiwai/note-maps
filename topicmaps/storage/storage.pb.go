// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package storage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TopicMapInfo struct {
	TopicMap             uint64   `protobuf:"varint,1,opt,name=TopicMap,proto3" json:"TopicMap,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicMapInfo) Reset()         { *m = TopicMapInfo{} }
func (m *TopicMapInfo) String() string { return proto.CompactTextString(m) }
func (*TopicMapInfo) ProtoMessage()    {}
func (*TopicMapInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_e8588f6e87e9117e, []int{0}
}
func (m *TopicMapInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicMapInfo.Unmarshal(m, b)
}
func (m *TopicMapInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicMapInfo.Marshal(b, m, deterministic)
}
func (dst *TopicMapInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicMapInfo.Merge(dst, src)
}
func (m *TopicMapInfo) XXX_Size() int {
	return xxx_messageInfo_TopicMapInfo.Size(m)
}
func (m *TopicMapInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicMapInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TopicMapInfo proto.InternalMessageInfo

func (m *TopicMapInfo) GetTopicMap() uint64 {
	if m != nil {
		return m.TopicMap
	}
	return 0
}

type TopicRefList struct {
	ItemIdentifiers      []string `protobuf:"bytes,1,rep,name=item_identifiers,json=itemIdentifiers,proto3" json:"item_identifiers,omitempty"`
	SubjectIndicators    []string `protobuf:"bytes,2,rep,name=subject_indicators,json=subjectIndicators,proto3" json:"subject_indicators,omitempty"`
	SubjectLocators      []string `protobuf:"bytes,3,rep,name=subject_locators,json=subjectLocators,proto3" json:"subject_locators,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicRefList) Reset()         { *m = TopicRefList{} }
func (m *TopicRefList) String() string { return proto.CompactTextString(m) }
func (*TopicRefList) ProtoMessage()    {}
func (*TopicRefList) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_e8588f6e87e9117e, []int{1}
}
func (m *TopicRefList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicRefList.Unmarshal(m, b)
}
func (m *TopicRefList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicRefList.Marshal(b, m, deterministic)
}
func (dst *TopicRefList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicRefList.Merge(dst, src)
}
func (m *TopicRefList) XXX_Size() int {
	return xxx_messageInfo_TopicRefList.Size(m)
}
func (m *TopicRefList) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicRefList.DiscardUnknown(m)
}

var xxx_messageInfo_TopicRefList proto.InternalMessageInfo

func (m *TopicRefList) GetItemIdentifiers() []string {
	if m != nil {
		return m.ItemIdentifiers
	}
	return nil
}

func (m *TopicRefList) GetSubjectIndicators() []string {
	if m != nil {
		return m.SubjectIndicators
	}
	return nil
}

func (m *TopicRefList) GetSubjectLocators() []string {
	if m != nil {
		return m.SubjectLocators
	}
	return nil
}

type Name struct {
	Topic                uint64   `protobuf:"varint,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Name) Reset()         { *m = Name{} }
func (m *Name) String() string { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()    {}
func (*Name) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_e8588f6e87e9117e, []int{2}
}
func (m *Name) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Name.Unmarshal(m, b)
}
func (m *Name) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Name.Marshal(b, m, deterministic)
}
func (dst *Name) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Name.Merge(dst, src)
}
func (m *Name) XXX_Size() int {
	return xxx_messageInfo_Name.Size(m)
}
func (m *Name) XXX_DiscardUnknown() {
	xxx_messageInfo_Name.DiscardUnknown(m)
}

var xxx_messageInfo_Name proto.InternalMessageInfo

func (m *Name) GetTopic() uint64 {
	if m != nil {
		return m.Topic
	}
	return 0
}

func (m *Name) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Occurrence struct {
	Topic                uint64   `protobuf:"varint,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Occurrence) Reset()         { *m = Occurrence{} }
func (m *Occurrence) String() string { return proto.CompactTextString(m) }
func (*Occurrence) ProtoMessage()    {}
func (*Occurrence) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_e8588f6e87e9117e, []int{3}
}
func (m *Occurrence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Occurrence.Unmarshal(m, b)
}
func (m *Occurrence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Occurrence.Marshal(b, m, deterministic)
}
func (dst *Occurrence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Occurrence.Merge(dst, src)
}
func (m *Occurrence) XXX_Size() int {
	return xxx_messageInfo_Occurrence.Size(m)
}
func (m *Occurrence) XXX_DiscardUnknown() {
	xxx_messageInfo_Occurrence.DiscardUnknown(m)
}

var xxx_messageInfo_Occurrence proto.InternalMessageInfo

func (m *Occurrence) GetTopic() uint64 {
	if m != nil {
		return m.Topic
	}
	return 0
}

func (m *Occurrence) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*TopicMapInfo)(nil), "TopicMapInfo")
	proto.RegisterType((*TopicRefList)(nil), "TopicRefList")
	proto.RegisterType((*Name)(nil), "Name")
	proto.RegisterType((*Occurrence)(nil), "Occurrence")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_storage_e8588f6e87e9117e) }

var fileDescriptor_storage_e8588f6e87e9117e = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x40, 0x49, 0x5b, 0xc5, 0x0e, 0x8a, 0x1a, 0x3c, 0x04, 0x4f, 0x4b, 0x4e, 0xab, 0xa0, 0x07,
	0xbd, 0xf8, 0x0b, 0x0b, 0x55, 0x21, 0x78, 0x2f, 0x69, 0x3a, 0x2b, 0x23, 0x6d, 0xb2, 0x24, 0xb3,
	0xfe, 0x88, 0x3f, 0x2c, 0xd9, 0xcd, 0xae, 0x67, 0x8f, 0xef, 0xcd, 0xcb, 0x30, 0x04, 0x2e, 0x12,
	0x87, 0x68, 0x3f, 0xf1, 0xb1, 0x8b, 0x81, 0x83, 0xbe, 0x87, 0xf3, 0x8f, 0xd0, 0x91, 0x7b, 0xb5,
	0x5d, 0xe3, 0xdb, 0x20, 0x6f, 0xe1, 0x6c, 0x62, 0x25, 0x2a, 0x51, 0xaf, 0xcc, 0xcc, 0xfa, 0x47,
	0x94, 0xd8, 0x60, 0xbb, 0xa1, 0xc4, 0xf2, 0x0e, 0xae, 0x88, 0xf1, 0xb8, 0xa5, 0x3d, 0x7a, 0xa6,
	0x96, 0x30, 0x26, 0x25, 0xaa, 0x65, 0xbd, 0x36, 0x97, 0xd9, 0x37, 0x7f, 0x5a, 0x3e, 0x80, 0x4c,
	0xfd, 0xee, 0x0b, 0x1d, 0x6f, 0xc9, 0xef, 0xc9, 0x59, 0x0e, 0x31, 0xa9, 0xc5, 0x10, 0x5f, 0x97,
	0x49, 0x33, 0x0f, 0xf2, 0xe6, 0x29, 0x3f, 0x84, 0x12, 0x2f, 0xc7, 0xcd, 0xc5, 0x6f, 0x8a, 0xd6,
	0x4f, 0xb0, 0x7a, 0xb3, 0x47, 0x94, 0x37, 0x70, 0xc2, 0xf9, 0xb8, 0x72, 0xf6, 0x08, 0xd9, 0x7e,
	0xdb, 0x43, 0x8f, 0x6a, 0x51, 0x89, 0x7a, 0x6d, 0x46, 0xd0, 0x2f, 0x00, 0xef, 0xce, 0xf5, 0x31,
	0xa2, 0x77, 0xff, 0x7a, 0xb9, 0x3b, 0x1d, 0xbe, 0xed, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x6b,
	0xa8, 0x30, 0x9b, 0x47, 0x01, 0x00, 0x00,
}
