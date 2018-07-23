// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unconfirmed_db.proto

package vitepb

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

type UnconfirmedMeta struct {
	AccountId            []byte       `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	TotalNumber          []byte       `protobuf:"bytes,2,opt,name=totalNumber,proto3" json:"totalNumber,omitempty"`
	TokenInfoList        []*TokenInfo `protobuf:"bytes,3,rep,name=tokenInfoList,proto3" json:"tokenInfoList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UnconfirmedMeta) Reset()         { *m = UnconfirmedMeta{} }
func (m *UnconfirmedMeta) String() string { return proto.CompactTextString(m) }
func (*UnconfirmedMeta) ProtoMessage()    {}
func (*UnconfirmedMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_unconfirmed_db_6b7874ca432b3a31, []int{0}
}
func (m *UnconfirmedMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnconfirmedMeta.Unmarshal(m, b)
}
func (m *UnconfirmedMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnconfirmedMeta.Marshal(b, m, deterministic)
}
func (dst *UnconfirmedMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnconfirmedMeta.Merge(dst, src)
}
func (m *UnconfirmedMeta) XXX_Size() int {
	return xxx_messageInfo_UnconfirmedMeta.Size(m)
}
func (m *UnconfirmedMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_UnconfirmedMeta.DiscardUnknown(m)
}

var xxx_messageInfo_UnconfirmedMeta proto.InternalMessageInfo

func (m *UnconfirmedMeta) GetAccountId() []byte {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *UnconfirmedMeta) GetTotalNumber() []byte {
	if m != nil {
		return m.TotalNumber
	}
	return nil
}

func (m *UnconfirmedMeta) GetTokenInfoList() []*TokenInfo {
	if m != nil {
		return m.TokenInfoList
	}
	return nil
}

type TokenInfo struct {
	TokenId              []byte   `protobuf:"bytes,1,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	TotalAmount          []byte   `protobuf:"bytes,2,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenInfo) Reset()         { *m = TokenInfo{} }
func (m *TokenInfo) String() string { return proto.CompactTextString(m) }
func (*TokenInfo) ProtoMessage()    {}
func (*TokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_unconfirmed_db_6b7874ca432b3a31, []int{1}
}
func (m *TokenInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenInfo.Unmarshal(m, b)
}
func (m *TokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenInfo.Marshal(b, m, deterministic)
}
func (dst *TokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenInfo.Merge(dst, src)
}
func (m *TokenInfo) XXX_Size() int {
	return xxx_messageInfo_TokenInfo.Size(m)
}
func (m *TokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TokenInfo proto.InternalMessageInfo

func (m *TokenInfo) GetTokenId() []byte {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *TokenInfo) GetTotalAmount() []byte {
	if m != nil {
		return m.TotalAmount
	}
	return nil
}

type HashList struct {
	HashList             [][]byte `protobuf:"bytes,1,rep,name=hashList,proto3" json:"hashList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashList) Reset()         { *m = HashList{} }
func (m *HashList) String() string { return proto.CompactTextString(m) }
func (*HashList) ProtoMessage()    {}
func (*HashList) Descriptor() ([]byte, []int) {
	return fileDescriptor_unconfirmed_db_6b7874ca432b3a31, []int{2}
}
func (m *HashList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashList.Unmarshal(m, b)
}
func (m *HashList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashList.Marshal(b, m, deterministic)
}
func (dst *HashList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashList.Merge(dst, src)
}
func (m *HashList) XXX_Size() int {
	return xxx_messageInfo_HashList.Size(m)
}
func (m *HashList) XXX_DiscardUnknown() {
	xxx_messageInfo_HashList.DiscardUnknown(m)
}

var xxx_messageInfo_HashList proto.InternalMessageInfo

func (m *HashList) GetHashList() [][]byte {
	if m != nil {
		return m.HashList
	}
	return nil
}

func init() {
	proto.RegisterType((*UnconfirmedMeta)(nil), "vitepb.UnconfirmedMeta")
	proto.RegisterType((*TokenInfo)(nil), "vitepb.TokenInfo")
	proto.RegisterType((*HashList)(nil), "vitepb.HashList")
}

func init() {
	proto.RegisterFile("unconfirmed_db.proto", fileDescriptor_unconfirmed_db_6b7874ca432b3a31)
}

var fileDescriptor_unconfirmed_db_6b7874ca432b3a31 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xcd, 0x4b, 0xce,
	0xcf, 0x4b, 0xcb, 0x2c, 0xca, 0x4d, 0x4d, 0x89, 0x4f, 0x49, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x2b, 0xcb, 0x2c, 0x49, 0x2d, 0x48, 0x52, 0xea, 0x62, 0xe4, 0xe2, 0x0f, 0x45, 0x28,
	0xf0, 0x4d, 0x2d, 0x49, 0x14, 0x92, 0xe1, 0xe2, 0x4c, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xf1,
	0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x42, 0x08, 0x08, 0x29, 0x70, 0x71, 0x97, 0xe4,
	0x97, 0x24, 0xe6, 0xf8, 0x95, 0xe6, 0x26, 0xa5, 0x16, 0x49, 0x30, 0x81, 0xe5, 0x91, 0x85, 0x84,
	0xcc, 0xb9, 0x78, 0x4b, 0xf2, 0xb3, 0x53, 0xf3, 0x3c, 0xf3, 0xd2, 0xf2, 0x7d, 0x32, 0x8b, 0x4b,
	0x24, 0x98, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x04, 0xf5, 0x20, 0x76, 0xea, 0x85, 0xc0, 0x24, 0x83,
	0x50, 0xd5, 0x29, 0xb9, 0x73, 0x71, 0xc2, 0xe5, 0x84, 0x24, 0xb8, 0xd8, 0x21, 0xb2, 0x30, 0x37,
	0xc0, 0xb8, 0x70, 0x17, 0x38, 0xe6, 0x82, 0x9c, 0x84, 0xe2, 0x02, 0x88, 0x90, 0x92, 0x1a, 0x17,
	0x87, 0x47, 0x62, 0x71, 0x06, 0xc8, 0x50, 0x21, 0x29, 0x2e, 0x8e, 0x0c, 0x28, 0x5b, 0x82, 0x51,
	0x81, 0x59, 0x83, 0x27, 0x08, 0xce, 0x4f, 0x62, 0x03, 0x07, 0x86, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xd7, 0x93, 0x05, 0x50, 0x24, 0x01, 0x00, 0x00,
}