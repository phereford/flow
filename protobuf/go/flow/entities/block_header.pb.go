// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/entities/block_header.proto

package entities

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

type BlockHeader struct {
	Id                   []byte               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId             []byte               `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Height               uint64               `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d8363da0276a74, []int{0}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *BlockHeader) GetParentId() []byte {
	if m != nil {
		return m.ParentId
	}
	return nil
}

func (m *BlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "flow.entities.BlockHeader")
}

func init() { proto.RegisterFile("flow/entities/block_header.proto", fileDescriptor_b9d8363da0276a74) }

var fileDescriptor_b9d8363da0276a74 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xcb, 0xc9, 0x2f,
	0xd7, 0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xc9, 0x4c, 0x2d, 0xd6, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x8e,
	0xcf, 0x48, 0x4d, 0x4c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x05, 0xa9,
	0xd0, 0x83, 0xa9, 0x90, 0x92, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x4b, 0x26, 0x95,
	0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0xd4, 0x2b, 0xf5, 0x30,
	0x72, 0x71, 0x3b, 0x81, 0x8c, 0xf1, 0x00, 0x9b, 0x22, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x13, 0xc4, 0x94, 0x99, 0x22, 0x24, 0xcd, 0xc5, 0x59, 0x90, 0x58, 0x94,
	0x9a, 0x57, 0x12, 0x9f, 0x99, 0x22, 0xc1, 0x04, 0x16, 0xe6, 0x80, 0x08, 0x78, 0xa6, 0x08, 0x89,
	0x71, 0xb1, 0x65, 0xa4, 0x66, 0xa6, 0x67, 0x94, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0xb0, 0x04, 0x41,
	0x79, 0x42, 0x16, 0x5c, 0x9c, 0x70, 0x7b, 0x24, 0x58, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xa4, 0xf4,
	0x20, 0x2e, 0xd1, 0x83, 0xb9, 0x44, 0x2f, 0x04, 0xa6, 0x22, 0x08, 0xa1, 0xd8, 0x89, 0x2b, 0x8a,
	0x03, 0xe6, 0xf6, 0x24, 0x36, 0xb0, 0x52, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x45,
	0x9a, 0x9b, 0xf5, 0x00, 0x00, 0x00,
}