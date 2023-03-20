// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/permission/common.proto

package types

import (
	fmt "fmt"
	common "github.com/bnb-chain/greenfield/types/common"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ActionType defines the operations you can execute in greenfield storage network
type ActionType int32

const (
	ACTION_UPDATE_BUCKET_INFO  ActionType = 0
	ACTION_DELETE_BUCKET       ActionType = 1
	ACTION_CREATE_OBJECT       ActionType = 2
	ACTION_DELETE_OBJECT       ActionType = 3
	ACTION_COPY_OBJECT         ActionType = 4
	ACTION_GET_OBJECT          ActionType = 5
	ACTION_EXECUTE_OBJECT      ActionType = 6
	ACTION_LIST_OBJECT         ActionType = 7
	ACTION_UPDATE_GROUP_MEMBER ActionType = 8
	ACTION_DELETE_GROUP        ActionType = 9
	ACTION_GROUP_MEMBER        ActionType = 10
	ACTION_TYPE_ALL            ActionType = 99
)

var ActionType_name = map[int32]string{
	0:  "ACTION_UPDATE_BUCKET_INFO",
	1:  "ACTION_DELETE_BUCKET",
	2:  "ACTION_CREATE_OBJECT",
	3:  "ACTION_DELETE_OBJECT",
	4:  "ACTION_COPY_OBJECT",
	5:  "ACTION_GET_OBJECT",
	6:  "ACTION_EXECUTE_OBJECT",
	7:  "ACTION_LIST_OBJECT",
	8:  "ACTION_UPDATE_GROUP_MEMBER",
	9:  "ACTION_DELETE_GROUP",
	10: "ACTION_GROUP_MEMBER",
	99: "ACTION_TYPE_ALL",
}

var ActionType_value = map[string]int32{
	"ACTION_UPDATE_BUCKET_INFO":  0,
	"ACTION_DELETE_BUCKET":       1,
	"ACTION_CREATE_OBJECT":       2,
	"ACTION_DELETE_OBJECT":       3,
	"ACTION_COPY_OBJECT":         4,
	"ACTION_GET_OBJECT":          5,
	"ACTION_EXECUTE_OBJECT":      6,
	"ACTION_LIST_OBJECT":         7,
	"ACTION_UPDATE_GROUP_MEMBER": 8,
	"ACTION_DELETE_GROUP":        9,
	"ACTION_GROUP_MEMBER":        10,
	"ACTION_TYPE_ALL":            99,
}

func (x ActionType) String() string {
	return proto.EnumName(ActionType_name, int32(x))
}

func (ActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33a4d646aee30990, []int{0}
}

// Effect define the effect of the operation permission, include Allow or deny
type Effect int32

const (
	EFFECT_ALLOW Effect = 0
	EFFECT_DENY  Effect = 1
	// Use internally, means skipped. there is no explicit ALLOW or DENY, and further permission checks are required.
	EFFECT_PASS Effect = 2
)

var Effect_name = map[int32]string{
	0: "EFFECT_ALLOW",
	1: "EFFECT_DENY",
	2: "EFFECT_PASS",
}

var Effect_value = map[string]int32{
	"EFFECT_ALLOW": 0,
	"EFFECT_DENY":  1,
	"EFFECT_PASS":  2,
}

func (x Effect) String() string {
	return proto.EnumName(Effect_name, int32(x))
}

func (Effect) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33a4d646aee30990, []int{1}
}

type PrincipalType int32

const (
	// Reserved for extended use
	TYPE_GNFD_ACCOUNT PrincipalType = 0
	TYPE_GNFD_GROUP   PrincipalType = 1
)

var PrincipalType_name = map[int32]string{
	0: "TYPE_GNFD_ACCOUNT",
	1: "TYPE_GNFD_GROUP",
}

var PrincipalType_value = map[string]int32{
	"TYPE_GNFD_ACCOUNT": 0,
	"TYPE_GNFD_GROUP":   1,
}

func (x PrincipalType) String() string {
	return proto.EnumName(PrincipalType_name, int32(x))
}

func (PrincipalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33a4d646aee30990, []int{2}
}

type Statement struct {
	// effect define the impact of permissions, which can be Allow/Deny
	Effect Effect `protobuf:"varint,1,opt,name=effect,proto3,enum=bnbchain.greenfield.permission.Effect" json:"effect,omitempty"`
	// action_type define the operation type you can act. greenfield defines a set of permission
	// that you can specify in a permissionInfo. see ActionType enum for detail.
	Actions []ActionType `protobuf:"varint,2,rep,packed,name=actions,proto3,enum=bnbchain.greenfield.permission.ActionType" json:"actions,omitempty"`
	// resources define the resource list you can operate.
	// CAN ONLY USED IN bucket level. Support fuzzy match and limit to 5
	Resources []string `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources,omitempty"`
	// expiration_time defines how long the permission is valid
	ExpirationTime *time.Time `protobuf:"bytes,4,opt,name=expiration_time,json=expirationTime,proto3,stdtime" json:"expiration_time,omitempty"`
	// limit_size defines the total data size that is allowed to operate
	LimitSize *common.UInt64Value `protobuf:"bytes,5,opt,name=limit_size,json=limitSize,proto3" json:"limit_size,omitempty"`
}

func (m *Statement) Reset()         { *m = Statement{} }
func (m *Statement) String() string { return proto.CompactTextString(m) }
func (*Statement) ProtoMessage()    {}
func (*Statement) Descriptor() ([]byte, []int) {
	return fileDescriptor_33a4d646aee30990, []int{0}
}
func (m *Statement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Statement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Statement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Statement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statement.Merge(m, src)
}
func (m *Statement) XXX_Size() int {
	return m.Size()
}
func (m *Statement) XXX_DiscardUnknown() {
	xxx_messageInfo_Statement.DiscardUnknown(m)
}

var xxx_messageInfo_Statement proto.InternalMessageInfo

func (m *Statement) GetEffect() Effect {
	if m != nil {
		return m.Effect
	}
	return EFFECT_ALLOW
}

func (m *Statement) GetActions() []ActionType {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *Statement) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Statement) GetExpirationTime() *time.Time {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

func (m *Statement) GetLimitSize() *common.UInt64Value {
	if m != nil {
		return m.LimitSize
	}
	return nil
}

// Principal define the roles that can grant permissions. Currently, it can be account or group.
type Principal struct {
	Type  PrincipalType `protobuf:"varint,1,opt,name=type,proto3,enum=bnbchain.greenfield.permission.PrincipalType" json:"type,omitempty"`
	Value string        `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_33a4d646aee30990, []int{1}
}
func (m *Principal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(m, src)
}
func (m *Principal) XXX_Size() int {
	return m.Size()
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

func (m *Principal) GetType() PrincipalType {
	if m != nil {
		return m.Type
	}
	return TYPE_GNFD_ACCOUNT
}

func (m *Principal) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("bnbchain.greenfield.permission.ActionType", ActionType_name, ActionType_value)
	proto.RegisterEnum("bnbchain.greenfield.permission.Effect", Effect_name, Effect_value)
	proto.RegisterEnum("bnbchain.greenfield.permission.PrincipalType", PrincipalType_name, PrincipalType_value)
	proto.RegisterType((*Statement)(nil), "bnbchain.greenfield.permission.Statement")
	proto.RegisterType((*Principal)(nil), "bnbchain.greenfield.permission.Principal")
}

func init() {
	proto.RegisterFile("greenfield/permission/common.proto", fileDescriptor_33a4d646aee30990)
}

var fileDescriptor_33a4d646aee30990 = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdd, 0x6e, 0x93, 0x50,
	0x1c, 0x87, 0xb6, 0xeb, 0xc6, 0x99, 0x6e, 0x78, 0xb6, 0x29, 0x6b, 0x94, 0x35, 0xbb, 0xd0, 0x66,
	0xc9, 0x20, 0xa9, 0xc6, 0x4b, 0x13, 0xa0, 0xa7, 0x4b, 0xb5, 0x2b, 0x0d, 0xa5, 0xea, 0xbc, 0x21,
	0xc0, 0x4e, 0xbb, 0x93, 0x14, 0x0e, 0x01, 0xaa, 0xdb, 0x9e, 0xc0, 0xcb, 0xbd, 0x83, 0x2f, 0xb3,
	0x0b, 0x2f, 0x76, 0xe9, 0x95, 0x9a, 0xed, 0x45, 0x0c, 0x1f, 0x2d, 0x6d, 0x34, 0xee, 0xae, 0xe7,
	0xf7, 0xc5, 0x9f, 0x1f, 0xff, 0x53, 0xb0, 0x3f, 0x0e, 0x31, 0xf6, 0x47, 0x04, 0x4f, 0x4e, 0xe5,
	0x00, 0x87, 0x1e, 0x89, 0x22, 0x42, 0x7d, 0xd9, 0xa5, 0x9e, 0x47, 0x7d, 0x29, 0x08, 0x69, 0x4c,
	0xa1, 0xe8, 0xf8, 0x8e, 0x7b, 0x66, 0x13, 0x5f, 0x2a, 0xc4, 0x52, 0x21, 0xae, 0xed, 0xba, 0x34,
	0xf2, 0x68, 0x64, 0xa5, 0x6a, 0x39, 0x3b, 0x64, 0xd6, 0xda, 0xf6, 0x98, 0x8e, 0x69, 0x86, 0x27,
	0xbf, 0x72, 0x74, 0x6f, 0x4c, 0xe9, 0x78, 0x82, 0xe5, 0xf4, 0xe4, 0x4c, 0x47, 0x72, 0x4c, 0x3c,
	0x1c, 0xc5, 0xb6, 0x17, 0xcc, 0x05, 0xc5, 0x54, 0xd9, 0x28, 0xf2, 0x97, 0xd0, 0x0e, 0x02, 0x1c,
	0x66, 0x82, 0xfd, 0xef, 0x25, 0xc0, 0x0d, 0x62, 0x3b, 0xc6, 0x1e, 0xf6, 0x63, 0xf8, 0x06, 0x54,
	0xf1, 0x68, 0x84, 0xdd, 0x58, 0x60, 0xeb, 0x6c, 0x63, 0xa3, 0xf9, 0x5c, 0xfa, 0xff, 0xc4, 0x12,
	0x4a, 0xd5, 0x46, 0xee, 0x82, 0x2d, 0xb0, 0x6a, 0xbb, 0x31, 0xa1, 0x7e, 0x24, 0x94, 0xea, 0xe5,
	0xc6, 0x46, 0xf3, 0xe0, 0xbe, 0x00, 0x25, 0x95, 0x9b, 0x17, 0x01, 0x36, 0x66, 0x56, 0xf8, 0x14,
	0x70, 0x21, 0x8e, 0xe8, 0x34, 0x74, 0x71, 0x24, 0x94, 0xeb, 0xe5, 0x06, 0x67, 0x14, 0x00, 0x3c,
	0x06, 0x9b, 0xf8, 0x3c, 0x20, 0xa1, 0x9d, 0x88, 0xad, 0xe4, 0x85, 0x85, 0x4a, 0x9d, 0x6d, 0xac,
	0x37, 0x6b, 0x52, 0xd6, 0x86, 0x34, 0x6b, 0x43, 0x32, 0x67, 0x6d, 0xa8, 0x6b, 0xd7, 0x3f, 0xf7,
	0xd8, 0xab, 0x5f, 0x7b, 0xac, 0xb1, 0x51, 0x98, 0x13, 0x1a, 0x76, 0x01, 0x98, 0x10, 0x8f, 0xc4,
	0x56, 0x44, 0x2e, 0xb1, 0xb0, 0x92, 0x26, 0xbd, 0xf8, 0xe7, 0xd4, 0xf9, 0xa7, 0x1c, 0x76, 0xfc,
	0xf8, 0xf5, 0xab, 0xf7, 0xf6, 0x64, 0x8a, 0xd5, 0x4a, 0x12, 0x6b, 0x70, 0x69, 0xc0, 0x80, 0x5c,
	0xe2, 0xfd, 0x53, 0xc0, 0xf5, 0x43, 0xe2, 0xbb, 0x24, 0xb0, 0x27, 0x50, 0x01, 0x95, 0xf8, 0x22,
	0xc0, 0x79, 0x97, 0x87, 0xf7, 0x55, 0x31, 0x37, 0xa6, 0x6d, 0xa4, 0x56, 0xb8, 0x0d, 0x56, 0x3e,
	0x27, 0x4f, 0x12, 0x4a, 0x75, 0xb6, 0xc1, 0x19, 0xd9, 0xe1, 0xe0, 0xba, 0x04, 0x40, 0x51, 0x1c,
	0x7c, 0x06, 0x76, 0x15, 0xcd, 0xec, 0xe8, 0x3d, 0x6b, 0xd8, 0x6f, 0x29, 0x26, 0xb2, 0xd4, 0xa1,
	0xf6, 0x0e, 0x99, 0x56, 0xa7, 0xd7, 0xd6, 0x79, 0x06, 0x0a, 0x60, 0x3b, 0xa7, 0x5b, 0xa8, 0x8b,
	0xe6, 0x34, 0xcf, 0x2e, 0x30, 0x9a, 0x81, 0x12, 0xa3, 0xae, 0xbe, 0x45, 0x9a, 0xc9, 0x97, 0xfe,
	0xf6, 0xe4, 0x4c, 0x19, 0x3e, 0x06, 0x70, 0xe6, 0xd1, 0xfb, 0x27, 0x33, 0xbc, 0x02, 0x77, 0xc0,
	0xa3, 0x1c, 0x3f, 0x42, 0xe6, 0x0c, 0x5e, 0x81, 0xbb, 0x60, 0x27, 0x87, 0xd1, 0x47, 0xa4, 0x0d,
	0x8b, 0xa4, 0xea, 0x42, 0x52, 0xb7, 0x33, 0x98, 0x5b, 0x56, 0xa1, 0x08, 0x6a, 0xcb, 0xaf, 0x73,
	0x64, 0xe8, 0xc3, 0xbe, 0x75, 0x8c, 0x8e, 0x55, 0x64, 0xf0, 0x6b, 0xf0, 0x09, 0xd8, 0x5a, 0x9e,
	0x2d, 0xe5, 0x79, 0x6e, 0x81, 0x58, 0x72, 0x00, 0xb8, 0x05, 0x36, 0x73, 0xc2, 0x3c, 0xe9, 0x23,
	0x4b, 0xe9, 0x76, 0x79, 0xb7, 0x56, 0xf9, 0xfa, 0x4d, 0x64, 0x0e, 0x54, 0x50, 0xcd, 0x76, 0x18,
	0xf2, 0xe0, 0x01, 0x6a, 0xb7, 0x91, 0x66, 0x26, 0xbc, 0xfe, 0x81, 0x67, 0xe0, 0x26, 0x58, 0xcf,
	0x91, 0x16, 0xea, 0x9d, 0xf0, 0xec, 0x02, 0xd0, 0x57, 0x06, 0x03, 0xbe, 0x94, 0x67, 0x28, 0xe0,
	0xe1, 0xd2, 0xb7, 0x4b, 0xba, 0x48, 0x1f, 0x74, 0xd4, 0x6b, 0xb7, 0x2c, 0x45, 0xd3, 0xf4, 0x61,
	0xcf, 0xe4, 0x99, 0x64, 0x8c, 0x02, 0xce, 0x86, 0x66, 0xb3, 0x08, 0xb5, 0x7b, 0x7d, 0x2b, 0xb2,
	0x37, 0xb7, 0x22, 0xfb, 0xfb, 0x56, 0x64, 0xaf, 0xee, 0x44, 0xe6, 0xe6, 0x4e, 0x64, 0x7e, 0xdc,
	0x89, 0xcc, 0xa7, 0xe6, 0x98, 0xc4, 0x67, 0x53, 0x27, 0xd9, 0x3e, 0xd9, 0xf1, 0x9d, 0xc3, 0x74,
	0x83, 0xe4, 0x85, 0x6b, 0x7d, 0xbe, 0xf8, 0x77, 0x93, 0x2c, 0x4d, 0xe4, 0x54, 0xd3, 0x1b, 0xf0,
	0xf2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0x28, 0x46, 0x4d, 0x94, 0x04, 0x00, 0x00,
}

func (m *Statement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Statement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Statement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LimitSize != nil {
		{
			size, err := m.LimitSize.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.ExpirationTime != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.ExpirationTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.ExpirationTime):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintCommon(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Resources) > 0 {
		for iNdEx := len(m.Resources) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Resources[iNdEx])
			copy(dAtA[i:], m.Resources[iNdEx])
			i = encodeVarintCommon(dAtA, i, uint64(len(m.Resources[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Actions) > 0 {
		dAtA4 := make([]byte, len(m.Actions)*10)
		var j3 int
		for _, num := range m.Actions {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintCommon(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x12
	}
	if m.Effect != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Effect))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Principal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Principal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Principal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Statement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Effect != 0 {
		n += 1 + sovCommon(uint64(m.Effect))
	}
	if len(m.Actions) > 0 {
		l = 0
		for _, e := range m.Actions {
			l += sovCommon(uint64(e))
		}
		n += 1 + sovCommon(uint64(l)) + l
	}
	if len(m.Resources) > 0 {
		for _, s := range m.Resources {
			l = len(s)
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	if m.ExpirationTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.ExpirationTime)
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.LimitSize != nil {
		l = m.LimitSize.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *Principal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCommon(uint64(m.Type))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Statement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Statement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Statement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Effect", wireType)
			}
			m.Effect = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Effect |= Effect(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v ActionType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ActionType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Actions = append(m.Actions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthCommon
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthCommon
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Actions) == 0 {
					m.Actions = make([]ActionType, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ActionType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCommon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ActionType(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Actions = append(m.Actions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resources = append(m.Resources, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpirationTime == nil {
				m.ExpirationTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.ExpirationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitSize", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LimitSize == nil {
				m.LimitSize = &common.UInt64Value{}
			}
			if err := m.LimitSize.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Principal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Principal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Principal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= PrincipalType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
