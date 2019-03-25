// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: role_access_policy.proto

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

type RoleAccessPolicy struct {
	RoleId               uint64   `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	AccessPolicyId       uint64   `protobuf:"varint,2,opt,name=access_policy_id,json=accessPolicyId,proto3" json:"access_policy_id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleAccessPolicy) Reset()         { *m = RoleAccessPolicy{} }
func (m *RoleAccessPolicy) String() string { return proto.CompactTextString(m) }
func (*RoleAccessPolicy) ProtoMessage()    {}
func (*RoleAccessPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_role_access_policy_4781c94acc0ee5ef, []int{0}
}
func (m *RoleAccessPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAccessPolicy.Unmarshal(m, b)
}
func (m *RoleAccessPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAccessPolicy.Marshal(b, m, deterministic)
}
func (dst *RoleAccessPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAccessPolicy.Merge(dst, src)
}
func (m *RoleAccessPolicy) XXX_Size() int {
	return xxx_messageInfo_RoleAccessPolicy.Size(m)
}
func (m *RoleAccessPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAccessPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAccessPolicy proto.InternalMessageInfo

type RoleAccessPoliciesUpdateRequest struct {
	RoleId               uint64   `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	AccessPolicyIds      []uint64 `protobuf:"varint,2,rep,packed,name=access_policy_ids,json=accessPolicyIds" json:"access_policy_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleAccessPoliciesUpdateRequest) Reset()         { *m = RoleAccessPoliciesUpdateRequest{} }
func (m *RoleAccessPoliciesUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*RoleAccessPoliciesUpdateRequest) ProtoMessage()    {}
func (*RoleAccessPoliciesUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_role_access_policy_4781c94acc0ee5ef, []int{1}
}
func (m *RoleAccessPoliciesUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAccessPoliciesUpdateRequest.Unmarshal(m, b)
}
func (m *RoleAccessPoliciesUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAccessPoliciesUpdateRequest.Marshal(b, m, deterministic)
}
func (dst *RoleAccessPoliciesUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAccessPoliciesUpdateRequest.Merge(dst, src)
}
func (m *RoleAccessPoliciesUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_RoleAccessPoliciesUpdateRequest.Size(m)
}
func (m *RoleAccessPoliciesUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAccessPoliciesUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAccessPoliciesUpdateRequest proto.InternalMessageInfo

type RoleAccessPoliciesUpdateReply struct {
	RoleAccessPolicies   []*RoleAccessPolicy `protobuf:"bytes,1,rep,name=role_access_policies,json=roleAccessPolicies" json:"role_access_policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RoleAccessPoliciesUpdateReply) Reset()         { *m = RoleAccessPoliciesUpdateReply{} }
func (m *RoleAccessPoliciesUpdateReply) String() string { return proto.CompactTextString(m) }
func (*RoleAccessPoliciesUpdateReply) ProtoMessage()    {}
func (*RoleAccessPoliciesUpdateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_role_access_policy_4781c94acc0ee5ef, []int{2}
}
func (m *RoleAccessPoliciesUpdateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAccessPoliciesUpdateReply.Unmarshal(m, b)
}
func (m *RoleAccessPoliciesUpdateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAccessPoliciesUpdateReply.Marshal(b, m, deterministic)
}
func (dst *RoleAccessPoliciesUpdateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAccessPoliciesUpdateReply.Merge(dst, src)
}
func (m *RoleAccessPoliciesUpdateReply) XXX_Size() int {
	return xxx_messageInfo_RoleAccessPoliciesUpdateReply.Size(m)
}
func (m *RoleAccessPoliciesUpdateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAccessPoliciesUpdateReply.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAccessPoliciesUpdateReply proto.InternalMessageInfo

func (m *RoleAccessPoliciesUpdateReply) GetRoleAccessPolicies() []*RoleAccessPolicy {
	if m != nil {
		return m.RoleAccessPolicies
	}
	return nil
}

func init() {
	proto.RegisterType((*RoleAccessPolicy)(nil), "proto.RoleAccessPolicy")
	proto.RegisterType((*RoleAccessPoliciesUpdateRequest)(nil), "proto.RoleAccessPoliciesUpdateRequest")
	proto.RegisterType((*RoleAccessPoliciesUpdateReply)(nil), "proto.RoleAccessPoliciesUpdateReply")
}

func init() {
	proto.RegisterFile("role_access_policy.proto", fileDescriptor_role_access_policy_4781c94acc0ee5ef)
}

var fileDescriptor_role_access_policy_4781c94acc0ee5ef = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0x34, 0xdb, 0xba, 0xe2, 0x13, 0x74, 0x0d, 0xc2, 0x06, 0x61, 0xd9, 0xb2, 0xa7, 0x20, 0xd8,
	0x05, 0xbd, 0x79, 0xeb, 0xb1, 0x37, 0x09, 0x78, 0x2e, 0x6d, 0x12, 0x6b, 0x24, 0x92, 0xda, 0xa4,
	0x87, 0xfd, 0x03, 0x7f, 0xc1, 0x9b, 0x9f, 0x2a, 0x7d, 0xad, 0xe0, 0x76, 0xd9, 0x53, 0xfb, 0x66,
	0x26, 0x33, 0xc3, 0x00, 0x6b, 0x9d, 0xd5, 0x45, 0x29, 0xa5, 0xf6, 0xbe, 0x68, 0x9c, 0x35, 0x72,
	0x97, 0x36, 0xad, 0x0b, 0x8e, 0x9e, 0xe2, 0xe7, 0xf6, 0xbe, 0x36, 0xe1, 0xad, 0xab, 0x52, 0xe9,
	0x3e, 0xb6, 0xb5, 0xab, 0xdd, 0x16, 0xe1, 0xaa, 0x7b, 0xc5, 0x0b, 0x0f, 0xfc, 0x1b, 0x5e, 0x6d,
	0xbe, 0x09, 0x2c, 0x84, 0xb3, 0x3a, 0x43, 0xc7, 0x67, 0x34, 0xa4, 0x4b, 0x38, 0xc3, 0x18, 0xa3,
	0x18, 0x49, 0x08, 0x8f, 0xc5, 0xbc, 0x3f, 0x73, 0x45, 0x39, 0x2c, 0xf6, 0xa2, 0x7b, 0xc5, 0x0c,
	0x15, 0x97, 0xe5, 0x3f, 0x83, 0x5c, 0xd1, 0x15, 0x80, 0x6c, 0x75, 0x19, 0xb4, 0x2a, 0xca, 0xc0,
	0xa2, 0x84, 0xf0, 0x48, 0x9c, 0x8f, 0x48, 0x16, 0x7a, 0xba, 0x6b, 0xd4, 0x1f, 0x1d, 0x0f, 0xf4,
	0x88, 0x64, 0xe1, 0x29, 0xfe, 0xfa, 0x59, 0x9f, 0x6c, 0x2c, 0xac, 0x27, 0xd5, 0x8c, 0xf6, 0x2f,
	0x28, 0x12, 0xfa, 0xb3, 0xd3, 0x3e, 0x1c, 0x6f, 0x7a, 0x07, 0xd7, 0xd3, 0xa6, 0x9e, 0xcd, 0x92,
	0x88, 0xc7, 0xe2, 0x6a, 0xbf, 0xaa, 0x1f, 0xd3, 0xde, 0x61, 0x75, 0x3c, 0xad, 0xb1, 0x3b, 0x9a,
	0xc3, 0xcd, 0xc1, 0xf8, 0x46, 0x7b, 0x46, 0x92, 0x88, 0x5f, 0x3c, 0x2c, 0x87, 0x41, 0xd3, 0xe9,
	0x98, 0x82, 0xb6, 0x07, 0xae, 0xd5, 0x1c, 0xb5, 0x8f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x32,
	0x26, 0x98, 0x85, 0xce, 0x01, 0x00, 0x00,
}
