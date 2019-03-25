// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: access_policy.proto

package proto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AccessPolicy struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            int64    `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	State                string   `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicy) Reset()         { *m = AccessPolicy{} }
func (m *AccessPolicy) String() string { return proto.CompactTextString(m) }
func (*AccessPolicy) ProtoMessage()    {}
func (*AccessPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_policy_11fdf3d11cd97400, []int{0}
}
func (m *AccessPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicy.Unmarshal(m, b)
}
func (m *AccessPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicy.Marshal(b, m, deterministic)
}
func (dst *AccessPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicy.Merge(dst, src)
}
func (m *AccessPolicy) XXX_Size() int {
	return xxx_messageInfo_AccessPolicy.Size(m)
}
func (m *AccessPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicy proto.InternalMessageInfo

type AccessPolicyCreateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicyCreateRequest) Reset()         { *m = AccessPolicyCreateRequest{} }
func (m *AccessPolicyCreateRequest) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyCreateRequest) ProtoMessage()    {}
func (*AccessPolicyCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_policy_11fdf3d11cd97400, []int{1}
}
func (m *AccessPolicyCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyCreateRequest.Unmarshal(m, b)
}
func (m *AccessPolicyCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyCreateRequest.Marshal(b, m, deterministic)
}
func (dst *AccessPolicyCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyCreateRequest.Merge(dst, src)
}
func (m *AccessPolicyCreateRequest) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyCreateRequest.Size(m)
}
func (m *AccessPolicyCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyCreateRequest proto.InternalMessageInfo

type AccessPolicyCreateReply struct {
	AccessPolicy         *AccessPolicy `protobuf:"bytes,1,opt,name=access_policy,json=accessPolicy" json:"access_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AccessPolicyCreateReply) Reset()         { *m = AccessPolicyCreateReply{} }
func (m *AccessPolicyCreateReply) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyCreateReply) ProtoMessage()    {}
func (*AccessPolicyCreateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_policy_11fdf3d11cd97400, []int{2}
}
func (m *AccessPolicyCreateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyCreateReply.Unmarshal(m, b)
}
func (m *AccessPolicyCreateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyCreateReply.Marshal(b, m, deterministic)
}
func (dst *AccessPolicyCreateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyCreateReply.Merge(dst, src)
}
func (m *AccessPolicyCreateReply) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyCreateReply.Size(m)
}
func (m *AccessPolicyCreateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyCreateReply.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyCreateReply proto.InternalMessageInfo

func (m *AccessPolicyCreateReply) GetAccessPolicy() *AccessPolicy {
	if m != nil {
		return m.AccessPolicy
	}
	return nil
}

type AccessPolicyUpdateRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicyUpdateRequest) Reset()         { *m = AccessPolicyUpdateRequest{} }
func (m *AccessPolicyUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyUpdateRequest) ProtoMessage()    {}
func (*AccessPolicyUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_policy_11fdf3d11cd97400, []int{3}
}
func (m *AccessPolicyUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyUpdateRequest.Unmarshal(m, b)
}
func (m *AccessPolicyUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyUpdateRequest.Marshal(b, m, deterministic)
}
func (dst *AccessPolicyUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyUpdateRequest.Merge(dst, src)
}
func (m *AccessPolicyUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyUpdateRequest.Size(m)
}
func (m *AccessPolicyUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyUpdateRequest proto.InternalMessageInfo

type AccessPolicyUpdateReply struct {
	AccessPolicy         *AccessPolicy `protobuf:"bytes,1,opt,name=access_policy,json=accessPolicy" json:"access_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AccessPolicyUpdateReply) Reset()         { *m = AccessPolicyUpdateReply{} }
func (m *AccessPolicyUpdateReply) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyUpdateReply) ProtoMessage()    {}
func (*AccessPolicyUpdateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_policy_11fdf3d11cd97400, []int{4}
}
func (m *AccessPolicyUpdateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyUpdateReply.Unmarshal(m, b)
}
func (m *AccessPolicyUpdateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyUpdateReply.Marshal(b, m, deterministic)
}
func (dst *AccessPolicyUpdateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyUpdateReply.Merge(dst, src)
}
func (m *AccessPolicyUpdateReply) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyUpdateReply.Size(m)
}
func (m *AccessPolicyUpdateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyUpdateReply.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyUpdateReply proto.InternalMessageInfo

func (m *AccessPolicyUpdateReply) GetAccessPolicy() *AccessPolicy {
	if m != nil {
		return m.AccessPolicy
	}
	return nil
}

type AccessPolicyDeleteRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicyDeleteRequest) Reset()         { *m = AccessPolicyDeleteRequest{} }
func (m *AccessPolicyDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyDeleteRequest) ProtoMessage()    {}
func (*AccessPolicyDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_policy_11fdf3d11cd97400, []int{5}
}
func (m *AccessPolicyDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyDeleteRequest.Unmarshal(m, b)
}
func (m *AccessPolicyDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyDeleteRequest.Marshal(b, m, deterministic)
}
func (dst *AccessPolicyDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyDeleteRequest.Merge(dst, src)
}
func (m *AccessPolicyDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyDeleteRequest.Size(m)
}
func (m *AccessPolicyDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyDeleteRequest proto.InternalMessageInfo

type AccessPolicyDeleteReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicyDeleteReply) Reset()         { *m = AccessPolicyDeleteReply{} }
func (m *AccessPolicyDeleteReply) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyDeleteReply) ProtoMessage()    {}
func (*AccessPolicyDeleteReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_policy_11fdf3d11cd97400, []int{6}
}
func (m *AccessPolicyDeleteReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyDeleteReply.Unmarshal(m, b)
}
func (m *AccessPolicyDeleteReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyDeleteReply.Marshal(b, m, deterministic)
}
func (dst *AccessPolicyDeleteReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyDeleteReply.Merge(dst, src)
}
func (m *AccessPolicyDeleteReply) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyDeleteReply.Size(m)
}
func (m *AccessPolicyDeleteReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyDeleteReply.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyDeleteReply proto.InternalMessageInfo

func (m *AccessPolicyDeleteReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type AccessPolicyListRequest struct {
	Ids                  []uint64 `protobuf:"varint,1,rep,packed,name=ids" json:"ids,omitempty"`
	RoleId               uint64   `protobuf:"varint,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicyListRequest) Reset()         { *m = AccessPolicyListRequest{} }
func (m *AccessPolicyListRequest) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyListRequest) ProtoMessage()    {}
func (*AccessPolicyListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_policy_11fdf3d11cd97400, []int{7}
}
func (m *AccessPolicyListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyListRequest.Unmarshal(m, b)
}
func (m *AccessPolicyListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyListRequest.Marshal(b, m, deterministic)
}
func (dst *AccessPolicyListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyListRequest.Merge(dst, src)
}
func (m *AccessPolicyListRequest) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyListRequest.Size(m)
}
func (m *AccessPolicyListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyListRequest proto.InternalMessageInfo

type AccessPolicyListReply struct {
	AccessPolicies       []*AccessPolicy `protobuf:"bytes,1,rep,name=access_policies,json=accessPolicies" json:"access_policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AccessPolicyListReply) Reset()         { *m = AccessPolicyListReply{} }
func (m *AccessPolicyListReply) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyListReply) ProtoMessage()    {}
func (*AccessPolicyListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_policy_11fdf3d11cd97400, []int{8}
}
func (m *AccessPolicyListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyListReply.Unmarshal(m, b)
}
func (m *AccessPolicyListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyListReply.Marshal(b, m, deterministic)
}
func (dst *AccessPolicyListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyListReply.Merge(dst, src)
}
func (m *AccessPolicyListReply) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyListReply.Size(m)
}
func (m *AccessPolicyListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyListReply.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyListReply proto.InternalMessageInfo

func (m *AccessPolicyListReply) GetAccessPolicies() []*AccessPolicy {
	if m != nil {
		return m.AccessPolicies
	}
	return nil
}

type AccessPolicyBlockRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicyBlockRequest) Reset()         { *m = AccessPolicyBlockRequest{} }
func (m *AccessPolicyBlockRequest) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyBlockRequest) ProtoMessage()    {}
func (*AccessPolicyBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_policy_11fdf3d11cd97400, []int{9}
}
func (m *AccessPolicyBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyBlockRequest.Unmarshal(m, b)
}
func (m *AccessPolicyBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyBlockRequest.Marshal(b, m, deterministic)
}
func (dst *AccessPolicyBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyBlockRequest.Merge(dst, src)
}
func (m *AccessPolicyBlockRequest) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyBlockRequest.Size(m)
}
func (m *AccessPolicyBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyBlockRequest proto.InternalMessageInfo

type AccessPolicyBlockReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicyBlockReply) Reset()         { *m = AccessPolicyBlockReply{} }
func (m *AccessPolicyBlockReply) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyBlockReply) ProtoMessage()    {}
func (*AccessPolicyBlockReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_policy_11fdf3d11cd97400, []int{10}
}
func (m *AccessPolicyBlockReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyBlockReply.Unmarshal(m, b)
}
func (m *AccessPolicyBlockReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyBlockReply.Marshal(b, m, deterministic)
}
func (dst *AccessPolicyBlockReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyBlockReply.Merge(dst, src)
}
func (m *AccessPolicyBlockReply) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyBlockReply.Size(m)
}
func (m *AccessPolicyBlockReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyBlockReply.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyBlockReply proto.InternalMessageInfo

func (m *AccessPolicyBlockReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type AccessPolicyActivateRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicyActivateRequest) Reset()         { *m = AccessPolicyActivateRequest{} }
func (m *AccessPolicyActivateRequest) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyActivateRequest) ProtoMessage()    {}
func (*AccessPolicyActivateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_policy_11fdf3d11cd97400, []int{11}
}
func (m *AccessPolicyActivateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyActivateRequest.Unmarshal(m, b)
}
func (m *AccessPolicyActivateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyActivateRequest.Marshal(b, m, deterministic)
}
func (dst *AccessPolicyActivateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyActivateRequest.Merge(dst, src)
}
func (m *AccessPolicyActivateRequest) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyActivateRequest.Size(m)
}
func (m *AccessPolicyActivateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyActivateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyActivateRequest proto.InternalMessageInfo

type AccessPolicyActivateReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicyActivateReply) Reset()         { *m = AccessPolicyActivateReply{} }
func (m *AccessPolicyActivateReply) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyActivateReply) ProtoMessage()    {}
func (*AccessPolicyActivateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_policy_11fdf3d11cd97400, []int{12}
}
func (m *AccessPolicyActivateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyActivateReply.Unmarshal(m, b)
}
func (m *AccessPolicyActivateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyActivateReply.Marshal(b, m, deterministic)
}
func (dst *AccessPolicyActivateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyActivateReply.Merge(dst, src)
}
func (m *AccessPolicyActivateReply) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyActivateReply.Size(m)
}
func (m *AccessPolicyActivateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyActivateReply.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyActivateReply proto.InternalMessageInfo

func (m *AccessPolicyActivateReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*AccessPolicy)(nil), "proto.AccessPolicy")
	proto.RegisterType((*AccessPolicyCreateRequest)(nil), "proto.AccessPolicyCreateRequest")
	proto.RegisterType((*AccessPolicyCreateReply)(nil), "proto.AccessPolicyCreateReply")
	proto.RegisterType((*AccessPolicyUpdateRequest)(nil), "proto.AccessPolicyUpdateRequest")
	proto.RegisterType((*AccessPolicyUpdateReply)(nil), "proto.AccessPolicyUpdateReply")
	proto.RegisterType((*AccessPolicyDeleteRequest)(nil), "proto.AccessPolicyDeleteRequest")
	proto.RegisterType((*AccessPolicyDeleteReply)(nil), "proto.AccessPolicyDeleteReply")
	proto.RegisterType((*AccessPolicyListRequest)(nil), "proto.AccessPolicyListRequest")
	proto.RegisterType((*AccessPolicyListReply)(nil), "proto.AccessPolicyListReply")
	proto.RegisterType((*AccessPolicyBlockRequest)(nil), "proto.AccessPolicyBlockRequest")
	proto.RegisterType((*AccessPolicyBlockReply)(nil), "proto.AccessPolicyBlockReply")
	proto.RegisterType((*AccessPolicyActivateRequest)(nil), "proto.AccessPolicyActivateRequest")
	proto.RegisterType((*AccessPolicyActivateReply)(nil), "proto.AccessPolicyActivateReply")
}

func init() { proto.RegisterFile("access_policy.proto", fileDescriptor_access_policy_11fdf3d11cd97400) }

var fileDescriptor_access_policy_11fdf3d11cd97400 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xcd, 0x4e, 0x83, 0x40,
	0x14, 0x85, 0xa5, 0xd0, 0x6a, 0xaf, 0xb5, 0x9a, 0xa9, 0xb6, 0xa8, 0x31, 0x12, 0x56, 0x6c, 0xec,
	0x8f, 0x8d, 0x89, 0x31, 0x6e, 0xf0, 0x67, 0xa1, 0x71, 0x61, 0xc6, 0x74, 0xdd, 0x50, 0x18, 0xeb,
	0x44, 0xea, 0x20, 0x33, 0x98, 0xf4, 0x0d, 0x5c, 0xf8, 0x10, 0x3e, 0xaa, 0x61, 0x00, 0x0b, 0xb1,
	0xad, 0x26, 0xae, 0x98, 0x33, 0x37, 0xe7, 0xdc, 0x6f, 0x0e, 0xd0, 0x70, 0x5c, 0x97, 0x70, 0x3e,
	0x0c, 0x98, 0x4f, 0xdd, 0x69, 0x3b, 0x08, 0x99, 0x60, 0xa8, 0x2c, 0x3f, 0x7b, 0x47, 0x63, 0x2a,
	0x9e, 0xa2, 0x51, 0xdb, 0x65, 0x93, 0xce, 0x98, 0x8d, 0x59, 0x47, 0x5e, 0x8f, 0xa2, 0x47, 0xa9,
	0xa4, 0x90, 0xa7, 0xc4, 0x65, 0x7e, 0x28, 0x50, 0xb3, 0x65, 0xda, 0xbd, 0x0c, 0x43, 0x75, 0x28,
	0x51, 0x4f, 0x57, 0x0c, 0xc5, 0xd2, 0x70, 0x89, 0x7a, 0x08, 0x81, 0xf6, 0xe2, 0x4c, 0x88, 0x5e,
	0x32, 0x14, 0xab, 0x8a, 0xe5, 0x19, 0x1d, 0x00, 0xb8, 0x21, 0x71, 0x04, 0xf1, 0x86, 0x8e, 0xd0,
	0x55, 0x43, 0xb1, 0x54, 0x5c, 0x4d, 0x6f, 0x6c, 0x11, 0x8f, 0xa3, 0xc0, 0xcb, 0xc6, 0x5a, 0x32,
	0x4e, 0x6f, 0x6c, 0x81, 0xb6, 0xa1, 0xcc, 0x85, 0x23, 0x88, 0x5e, 0x96, 0x91, 0x89, 0x38, 0xd3,
	0xde, 0x3f, 0x0f, 0x57, 0xcc, 0x13, 0xd8, 0xcd, 0xd3, 0x5c, 0xca, 0x4c, 0x4c, 0x5e, 0x23, 0xc2,
	0xc5, 0x37, 0x8a, 0x32, 0x43, 0x49, 0x6d, 0x0f, 0xd0, 0x9a, 0x67, 0x0b, 0xfc, 0x29, 0x3a, 0x85,
	0x8d, 0x42, 0x5b, 0xd2, 0xbd, 0x7e, 0xdc, 0x48, 0xde, 0xdf, 0xce, 0xdb, 0x70, 0xcd, 0xc9, 0x29,
	0xf3, 0xba, 0xc8, 0x32, 0x90, 0x0f, 0xc8, 0x58, 0xfe, 0x50, 0xd3, 0x7c, 0xb6, 0x2c, 0xe6, 0x7f,
	0x6c, 0xbd, 0x22, 0xdb, 0x15, 0xf1, 0xc9, 0x42, 0xb6, 0x94, 0xa3, 0x57, 0xe4, 0xc8, 0x2c, 0x31,
	0x47, 0x13, 0x2a, 0x21, 0xe1, 0x91, 0x2f, 0xa4, 0x69, 0x0d, 0xa7, 0xca, 0xbc, 0x2d, 0x5a, 0xee,
	0x28, 0x17, 0xd9, 0x8e, 0x2d, 0x50, 0xa9, 0xc7, 0x75, 0xc5, 0x50, 0x2d, 0x0d, 0xc7, 0x47, 0xd4,
	0x82, 0xd5, 0x90, 0xf9, 0x64, 0x48, 0x3d, 0x59, 0x82, 0x86, 0x2b, 0xb1, 0xbc, 0xc9, 0xd6, 0x0f,
	0x60, 0xe7, 0x67, 0x56, 0xbc, 0xfc, 0x1c, 0x36, 0xf3, 0x25, 0x50, 0x92, 0xa4, 0x2e, 0xa8, 0xa1,
	0x9e, 0xab, 0x81, 0x12, 0x6e, 0x76, 0x41, 0xcf, 0xcf, 0x2f, 0x7c, 0xe6, 0x3e, 0x2f, 0xef, 0xa1,
	0x0b, 0xcd, 0x39, 0x8e, 0x65, 0x35, 0xf4, 0x61, 0x3f, 0xef, 0xb0, 0x5d, 0x41, 0xdf, 0x9c, 0xdf,
	0xea, 0xee, 0x17, 0xff, 0xd0, 0xcc, 0xb4, 0x64, 0xd3, 0xa8, 0x22, 0x5f, 0xdc, 0xff, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0xcb, 0x97, 0x43, 0x7a, 0xe1, 0x03, 0x00, 0x00,
}