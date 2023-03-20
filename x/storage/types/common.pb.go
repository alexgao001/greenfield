// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/storage/common.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SourceType int32

const (
	SOURCE_TYPE_ORIGIN          SourceType = 0
	SOURCE_TYPE_BSC_CROSS_CHAIN SourceType = 1
	SOURCE_TYPE_MIRROR_PENDING  SourceType = 2
)

var SourceType_name = map[int32]string{
	0: "SOURCE_TYPE_ORIGIN",
	1: "SOURCE_TYPE_BSC_CROSS_CHAIN",
	2: "SOURCE_TYPE_MIRROR_PENDING",
}

var SourceType_value = map[string]int32{
	"SOURCE_TYPE_ORIGIN":          0,
	"SOURCE_TYPE_BSC_CROSS_CHAIN": 1,
	"SOURCE_TYPE_MIRROR_PENDING":  2,
}

func (x SourceType) String() string {
	return proto.EnumName(SourceType_name, int32(x))
}

func (SourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{0}
}

type RedundancyType int32

const (
	REDUNDANCY_EC_TYPE      RedundancyType = 0
	REDUNDANCY_REPLICA_TYPE RedundancyType = 1
)

var RedundancyType_name = map[int32]string{
	0: "REDUNDANCY_EC_TYPE",
	1: "REDUNDANCY_REPLICA_TYPE",
}

var RedundancyType_value = map[string]int32{
	"REDUNDANCY_EC_TYPE":      0,
	"REDUNDANCY_REPLICA_TYPE": 1,
}

func (x RedundancyType) String() string {
	return proto.EnumName(RedundancyType_name, int32(x))
}

func (RedundancyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{1}
}

type ObjectStatus int32

const (
	OBJECT_STATUS_CREATED ObjectStatus = 0
	OBJECT_STATUS_SEALED  ObjectStatus = 1
)

var ObjectStatus_name = map[int32]string{
	0: "OBJECT_STATUS_CREATED",
	1: "OBJECT_STATUS_SEALED",
}

var ObjectStatus_value = map[string]int32{
	"OBJECT_STATUS_CREATED": 0,
	"OBJECT_STATUS_SEALED":  1,
}

func (x ObjectStatus) String() string {
	return proto.EnumName(ObjectStatus_name, int32(x))
}

func (ObjectStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{2}
}

type Approval struct {
	ExpiredHeight uint64 `protobuf:"varint,1,opt,name=expired_height,json=expiredHeight,proto3" json:"expired_height,omitempty"`
	Sig           []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *Approval) Reset()         { *m = Approval{} }
func (m *Approval) String() string { return proto.CompactTextString(m) }
func (*Approval) ProtoMessage()    {}
func (*Approval) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{0}
}
func (m *Approval) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Approval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Approval.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Approval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Approval.Merge(m, src)
}
func (m *Approval) XXX_Size() int {
	return m.Size()
}
func (m *Approval) XXX_DiscardUnknown() {
	xxx_messageInfo_Approval.DiscardUnknown(m)
}

var xxx_messageInfo_Approval proto.InternalMessageInfo

func (m *Approval) GetExpiredHeight() uint64 {
	if m != nil {
		return m.ExpiredHeight
	}
	return 0
}

func (m *Approval) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

// SecondarySpSignDoc used to generate seal signature of secondary SP
type SecondarySpSignDoc struct {
	SpAddress string `protobuf:"bytes,1,opt,name=sp_address,json=spAddress,proto3" json:"sp_address,omitempty"`
	ObjectId  Uint   `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3,customtype=Uint" json:"object_id"`
	Checksum  []byte `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (m *SecondarySpSignDoc) Reset()         { *m = SecondarySpSignDoc{} }
func (m *SecondarySpSignDoc) String() string { return proto.CompactTextString(m) }
func (*SecondarySpSignDoc) ProtoMessage()    {}
func (*SecondarySpSignDoc) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{1}
}
func (m *SecondarySpSignDoc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SecondarySpSignDoc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SecondarySpSignDoc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SecondarySpSignDoc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecondarySpSignDoc.Merge(m, src)
}
func (m *SecondarySpSignDoc) XXX_Size() int {
	return m.Size()
}
func (m *SecondarySpSignDoc) XXX_DiscardUnknown() {
	xxx_messageInfo_SecondarySpSignDoc.DiscardUnknown(m)
}

var xxx_messageInfo_SecondarySpSignDoc proto.InternalMessageInfo

func (m *SecondarySpSignDoc) GetSpAddress() string {
	if m != nil {
		return m.SpAddress
	}
	return ""
}

func (m *SecondarySpSignDoc) GetChecksum() []byte {
	if m != nil {
		return m.Checksum
	}
	return nil
}

func init() {
	proto.RegisterEnum("bnbchain.greenfield.storage.SourceType", SourceType_name, SourceType_value)
	proto.RegisterEnum("bnbchain.greenfield.storage.RedundancyType", RedundancyType_name, RedundancyType_value)
	proto.RegisterEnum("bnbchain.greenfield.storage.ObjectStatus", ObjectStatus_name, ObjectStatus_value)
	proto.RegisterType((*Approval)(nil), "bnbchain.greenfield.storage.Approval")
	proto.RegisterType((*SecondarySpSignDoc)(nil), "bnbchain.greenfield.storage.SecondarySpSignDoc")
}

func init() { proto.RegisterFile("greenfield/storage/common.proto", fileDescriptor_4eff6c0fa4aaf4c9) }

var fileDescriptor_4eff6c0fa4aaf4c9 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xc1, 0x6a, 0xdb, 0x4c,
	0x14, 0x85, 0xa5, 0x24, 0xfc, 0xc4, 0xf3, 0xa7, 0x41, 0x0c, 0x6e, 0xeb, 0xd8, 0x20, 0x87, 0x40,
	0x21, 0x04, 0x6c, 0x2d, 0xba, 0x68, 0xb7, 0xb2, 0x34, 0x24, 0x6a, 0x53, 0xc9, 0xcc, 0xc8, 0x8b,
	0x74, 0x33, 0x48, 0xa3, 0xa9, 0xac, 0x26, 0x9e, 0x11, 0x1a, 0xb9, 0xc4, 0x6f, 0xd0, 0x65, 0xdf,
	0x21, 0xaf, 0x90, 0x87, 0xc8, 0x32, 0x64, 0x55, 0xba, 0x08, 0xc5, 0x7e, 0x91, 0x62, 0x49, 0x49,
	0xdd, 0xdd, 0xdc, 0xef, 0x9e, 0x7b, 0xee, 0x1d, 0x38, 0xa0, 0x9f, 0x16, 0x9c, 0x8b, 0x2f, 0x19,
	0xbf, 0x4a, 0x2c, 0x55, 0xca, 0x22, 0x4a, 0xb9, 0xc5, 0xe4, 0x6c, 0x26, 0xc5, 0x30, 0x2f, 0x64,
	0x29, 0x61, 0x2f, 0x16, 0x31, 0x9b, 0x46, 0x99, 0x18, 0xfe, 0x55, 0x0e, 0x1b, 0x65, 0xf7, 0x80,
	0x49, 0x35, 0x93, 0x8a, 0x56, 0x52, 0xab, 0x2e, 0xea, 0xb9, 0x6e, 0x3b, 0x95, 0xa9, 0xac, 0xf9,
	0xfa, 0x55, 0xd3, 0x23, 0x07, 0xec, 0xda, 0x79, 0x5e, 0xc8, 0x6f, 0xd1, 0x15, 0x7c, 0x03, 0xf6,
	0xf9, 0x75, 0x9e, 0x15, 0x3c, 0xa1, 0x53, 0x9e, 0xa5, 0xd3, 0xb2, 0xa3, 0x1f, 0xea, 0xc7, 0x3b,
	0xf8, 0x45, 0x43, 0xcf, 0x2a, 0x08, 0x0d, 0xb0, 0xad, 0xb2, 0xb4, 0xb3, 0x75, 0xa8, 0x1f, 0xef,
	0xe1, 0xf5, 0xf3, 0xe8, 0x46, 0x07, 0x90, 0x70, 0x26, 0x45, 0x12, 0x15, 0x0b, 0x92, 0x93, 0x2c,
	0x15, 0xae, 0x64, 0xf0, 0x1d, 0x00, 0x2a, 0xa7, 0x51, 0x92, 0x14, 0x5c, 0xa9, 0xca, 0xab, 0x35,
	0xea, 0x3c, 0xdc, 0x0e, 0xda, 0xcd, 0x5d, 0x76, 0xdd, 0x21, 0x65, 0x91, 0x89, 0x14, 0xb7, 0x54,
	0xde, 0x00, 0xf8, 0x1e, 0xb4, 0x64, 0xfc, 0x95, 0xb3, 0x92, 0x66, 0x49, 0xb5, 0xa7, 0x35, 0xea,
	0xdd, 0x3d, 0xf6, 0xb5, 0x5f, 0x8f, 0xfd, 0x9d, 0x49, 0x26, 0xca, 0x87, 0xdb, 0xc1, 0xff, 0x8d,
	0xc7, 0xba, 0xc4, 0xbb, 0xb5, 0xda, 0x4b, 0x60, 0x17, 0xec, 0xb2, 0x29, 0x67, 0x97, 0x6a, 0x3e,
	0xeb, 0x6c, 0x57, 0x07, 0x3e, 0xd7, 0x27, 0x97, 0x00, 0x10, 0x39, 0x2f, 0x18, 0x0f, 0x17, 0x39,
	0x87, 0xaf, 0x00, 0x24, 0xc1, 0x04, 0x3b, 0x88, 0x86, 0x17, 0x63, 0x44, 0x03, 0xec, 0x9d, 0x7a,
	0xbe, 0xa1, 0xc1, 0x3e, 0xe8, 0x6d, 0xf2, 0x11, 0x71, 0xa8, 0x83, 0x03, 0x42, 0xa8, 0x73, 0x66,
	0x7b, 0xbe, 0xa1, 0x43, 0x13, 0x74, 0x37, 0x05, 0x9f, 0x3c, 0x8c, 0x03, 0x4c, 0xc7, 0xc8, 0x77,
	0x3d, 0xff, 0xd4, 0xd8, 0xea, 0xee, 0x7c, 0xbf, 0x31, 0xb5, 0x93, 0x8f, 0x60, 0x1f, 0xf3, 0x64,
	0x2e, 0x92, 0x48, 0xb0, 0xc5, 0xd3, 0x42, 0x8c, 0xdc, 0x89, 0xef, 0xda, 0xbe, 0x73, 0x41, 0x91,
	0x53, 0x8d, 0x1b, 0x1a, 0xec, 0x81, 0xd7, 0x1b, 0x1c, 0xa3, 0xf1, 0xb9, 0xe7, 0xd8, 0x75, 0x53,
	0x6f, 0xcc, 0x3c, 0xb0, 0x17, 0x54, 0x3f, 0x24, 0x65, 0x54, 0xce, 0x15, 0x3c, 0x00, 0x2f, 0x83,
	0xd1, 0x07, 0xe4, 0x84, 0x94, 0x84, 0x76, 0x38, 0x21, 0xd4, 0xc1, 0xc8, 0x0e, 0x91, 0x6b, 0x68,
	0xb0, 0x03, 0xda, 0xff, 0xb6, 0x08, 0xb2, 0xcf, 0x91, 0xfb, 0x64, 0x35, 0xf2, 0xee, 0x96, 0xa6,
	0x7e, 0xbf, 0x34, 0xf5, 0xdf, 0x4b, 0x53, 0xff, 0xb1, 0x32, 0xb5, 0xfb, 0x95, 0xa9, 0xfd, 0x5c,
	0x99, 0xda, 0x67, 0x2b, 0xcd, 0xca, 0xe9, 0x3c, 0x1e, 0x32, 0x39, 0xb3, 0x62, 0x11, 0x0f, 0xaa,
	0x8c, 0x59, 0x1b, 0x69, 0xbc, 0x7e, 0xce, 0x63, 0xb9, 0xc8, 0xb9, 0x8a, 0xff, 0xab, 0x12, 0xf4,
	0xf6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x79, 0x7a, 0xb6, 0xb2, 0x02, 0x00, 0x00,
}

func (m *Approval) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Approval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Approval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sig) > 0 {
		i -= len(m.Sig)
		copy(dAtA[i:], m.Sig)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Sig)))
		i--
		dAtA[i] = 0x12
	}
	if m.ExpiredHeight != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.ExpiredHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SecondarySpSignDoc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SecondarySpSignDoc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SecondarySpSignDoc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.ObjectId.Size()
		i -= size
		if _, err := m.ObjectId.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCommon(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.SpAddress) > 0 {
		i -= len(m.SpAddress)
		copy(dAtA[i:], m.SpAddress)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.SpAddress)))
		i--
		dAtA[i] = 0xa
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
func (m *Approval) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExpiredHeight != 0 {
		n += 1 + sovCommon(uint64(m.ExpiredHeight))
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *SecondarySpSignDoc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SpAddress)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = m.ObjectId.Size()
	n += 1 + l + sovCommon(uint64(l))
	l = len(m.Checksum)
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
func (m *Approval) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Approval: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Approval: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiredHeight", wireType)
			}
			m.ExpiredHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiredHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
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
func (m *SecondarySpSignDoc) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SecondarySpSignDoc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SecondarySpSignDoc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpAddress", wireType)
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
			m.SpAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectId", wireType)
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
			if err := m.ObjectId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = append(m.Checksum[:0], dAtA[iNdEx:postIndex]...)
			if m.Checksum == nil {
				m.Checksum = []byte{}
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
