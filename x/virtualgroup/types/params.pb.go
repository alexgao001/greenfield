// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/virtualgroup/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	// deposit_denom defines the staking coin denomination.
	DepositDenom string `protobuf:"bytes,1,opt,name=deposit_denom,json=depositDenom,proto3" json:"deposit_denom,omitempty"`
	// store price, in bnb wei per charge byte
	GvgStakingPrice github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=gvg_staking_price,json=gvgStakingPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"gvg_staking_price"`
	// the max number of lvg which allowed in a bucket
	MaxLocalVirtualGroupNumPerBucket uint32 `protobuf:"varint,3,opt,name=max_local_virtual_group_num_per_bucket,json=maxLocalVirtualGroupNumPerBucket,proto3" json:"max_local_virtual_group_num_per_bucket,omitempty"`
	// the max number of gvg which can exist in a family
	MaxGlobalVirtualGroupNumPerFamily uint32 `protobuf:"varint,4,opt,name=max_global_virtual_group_num_per_family,json=maxGlobalVirtualGroupNumPerFamily,proto3" json:"max_global_virtual_group_num_per_family,omitempty"`
	// if the store size reach the exceed, the family is not allowed to sever more buckets
	MaxStoreSizePerFamily uint64 `protobuf:"varint,5,opt,name=max_store_size_per_family,json=maxStoreSizePerFamily,proto3" json:"max_store_size_per_family,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ecf89dd5128885, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDepositDenom() string {
	if m != nil {
		return m.DepositDenom
	}
	return ""
}

func (m *Params) GetMaxLocalVirtualGroupNumPerBucket() uint32 {
	if m != nil {
		return m.MaxLocalVirtualGroupNumPerBucket
	}
	return 0
}

func (m *Params) GetMaxGlobalVirtualGroupNumPerFamily() uint32 {
	if m != nil {
		return m.MaxGlobalVirtualGroupNumPerFamily
	}
	return 0
}

func (m *Params) GetMaxStoreSizePerFamily() uint64 {
	if m != nil {
		return m.MaxStoreSizePerFamily
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "greenfield.virtualgroup.Params")
}

func init() {
	proto.RegisterFile("greenfield/virtualgroup/params.proto", fileDescriptor_d8ecf89dd5128885)
}

var fileDescriptor_d8ecf89dd5128885 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0x87, 0x1b, 0x77, 0x5d, 0x34, 0xb8, 0x88, 0x45, 0xb1, 0xbb, 0x87, 0xb6, 0xfe, 0x61, 0x9d,
	0xcb, 0xb4, 0x07, 0x3d, 0x88, 0x78, 0x1a, 0xc4, 0x65, 0x41, 0x96, 0xd2, 0x01, 0x0f, 0x5e, 0x42,
	0xda, 0x66, 0x33, 0x61, 0x9a, 0xa4, 0x24, 0xe9, 0xd0, 0xdd, 0x9b, 0xdf, 0xc0, 0xa3, 0xc7, 0xfd,
	0x10, 0x7e, 0x88, 0x3d, 0x2e, 0x9e, 0xc4, 0xc3, 0x20, 0x33, 0x17, 0x3f, 0x86, 0x24, 0x2d, 0x5a,
	0x11, 0x4f, 0x6d, 0x7f, 0x7d, 0xf2, 0xf0, 0xbe, 0x6f, 0x5e, 0xf8, 0x94, 0x2a, 0x42, 0xc4, 0x19,
	0x23, 0x75, 0x95, 0xae, 0x98, 0x32, 0x2d, 0xae, 0xa9, 0x92, 0x6d, 0x93, 0x36, 0x58, 0x61, 0xae,
	0x93, 0x46, 0x49, 0x23, 0xfd, 0x87, 0x7f, 0xa8, 0x64, 0x4c, 0x1d, 0x1e, 0x94, 0x52, 0x73, 0xa9,
	0x91, 0xc3, 0xd2, 0xfe, 0xa3, 0x3f, 0x73, 0x78, 0x9f, 0x4a, 0x2a, 0xfb, 0xdc, 0xbe, 0xf5, 0xe9,
	0xe3, 0x8f, 0x3b, 0x70, 0x2f, 0x73, 0x6a, 0xff, 0x09, 0xdc, 0xaf, 0x48, 0x23, 0x35, 0x33, 0xa8,
	0x22, 0x42, 0xf2, 0x00, 0xc4, 0x60, 0x72, 0x3b, 0xbf, 0x33, 0x84, 0x6f, 0x6c, 0xe6, 0x2f, 0xe0,
	0x3d, 0xba, 0xa2, 0x48, 0x1b, 0xbc, 0x64, 0x82, 0xa2, 0x46, 0xb1, 0x92, 0x04, 0x37, 0x2c, 0x38,
	0x7b, 0x7d, 0xb5, 0x8e, 0xbc, 0xef, 0xeb, 0xe8, 0x88, 0x32, 0xb3, 0x68, 0x8b, 0xa4, 0x94, 0x7c,
	0xa8, 0x60, 0x78, 0x4c, 0x75, 0xb5, 0x4c, 0xcd, 0x79, 0x43, 0x74, 0x72, 0x22, 0xcc, 0xd7, 0x2f,
	0x53, 0x38, 0x14, 0x78, 0x22, 0x4c, 0x7e, 0x97, 0xae, 0xe8, 0xbc, 0xb7, 0x66, 0x56, 0xea, 0x67,
	0xf0, 0x88, 0xe3, 0x0e, 0xd5, 0xb2, 0xc4, 0x35, 0x1a, 0x9a, 0x44, 0xae, 0x4b, 0x24, 0x5a, 0x8e,
	0x1a, 0xa2, 0x50, 0xd1, 0x96, 0x4b, 0x62, 0x82, 0x9d, 0x18, 0x4c, 0xf6, 0xf3, 0x98, 0xe3, 0xee,
	0x9d, 0x85, 0xdf, 0xf7, 0xec, 0xb1, 0x45, 0x4f, 0x5b, 0x9e, 0x11, 0x35, 0x73, 0x9c, 0x9f, 0xc3,
	0x67, 0xd6, 0x48, 0x6b, 0x59, 0xfc, 0x57, 0x79, 0x86, 0x39, 0xab, 0xcf, 0x83, 0x5d, 0xa7, 0x7c,
	0xc4, 0x71, 0x77, 0xec, 0xe8, 0x7f, 0x9d, 0x6f, 0x1d, 0xe8, 0xbf, 0x84, 0x07, 0xd6, 0xa9, 0x8d,
	0x54, 0x04, 0x69, 0x76, 0x41, 0xc6, 0x96, 0x9b, 0x31, 0x98, 0xec, 0xe6, 0x0f, 0x38, 0xee, 0xe6,
	0xf6, 0xff, 0x9c, 0x5d, 0x90, 0xdf, 0x27, 0x5f, 0xdd, 0xfa, 0x7c, 0x19, 0x79, 0x3f, 0x2f, 0x23,
	0x30, 0x3b, 0xbd, 0xda, 0x84, 0xe0, 0x7a, 0x13, 0x82, 0x1f, 0x9b, 0x10, 0x7c, 0xda, 0x86, 0xde,
	0xf5, 0x36, 0xf4, 0xbe, 0x6d, 0x43, 0xef, 0xc3, 0x8b, 0xd1, 0x28, 0x0b, 0x51, 0x4c, 0xcb, 0x05,
	0x66, 0x22, 0x1d, 0xad, 0x48, 0xf7, 0xf7, 0x92, 0xb8, 0xe1, 0x16, 0x7b, 0xee, 0x6a, 0x9f, 0xff,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x82, 0x0a, 0x7b, 0x8a, 0x4c, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.DepositDenom != that1.DepositDenom {
		return false
	}
	if !this.GvgStakingPrice.Equal(that1.GvgStakingPrice) {
		return false
	}
	if this.MaxLocalVirtualGroupNumPerBucket != that1.MaxLocalVirtualGroupNumPerBucket {
		return false
	}
	if this.MaxGlobalVirtualGroupNumPerFamily != that1.MaxGlobalVirtualGroupNumPerFamily {
		return false
	}
	if this.MaxStoreSizePerFamily != that1.MaxStoreSizePerFamily {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxStoreSizePerFamily != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxStoreSizePerFamily))
		i--
		dAtA[i] = 0x28
	}
	if m.MaxGlobalVirtualGroupNumPerFamily != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxGlobalVirtualGroupNumPerFamily))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxLocalVirtualGroupNumPerBucket != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxLocalVirtualGroupNumPerBucket))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.GvgStakingPrice.Size()
		i -= size
		if _, err := m.GvgStakingPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.DepositDenom) > 0 {
		i -= len(m.DepositDenom)
		copy(dAtA[i:], m.DepositDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DepositDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DepositDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.GvgStakingPrice.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MaxLocalVirtualGroupNumPerBucket != 0 {
		n += 1 + sovParams(uint64(m.MaxLocalVirtualGroupNumPerBucket))
	}
	if m.MaxGlobalVirtualGroupNumPerFamily != 0 {
		n += 1 + sovParams(uint64(m.MaxGlobalVirtualGroupNumPerFamily))
	}
	if m.MaxStoreSizePerFamily != 0 {
		n += 1 + sovParams(uint64(m.MaxStoreSizePerFamily))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GvgStakingPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GvgStakingPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLocalVirtualGroupNumPerBucket", wireType)
			}
			m.MaxLocalVirtualGroupNumPerBucket = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxLocalVirtualGroupNumPerBucket |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxGlobalVirtualGroupNumPerFamily", wireType)
			}
			m.MaxGlobalVirtualGroupNumPerFamily = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxGlobalVirtualGroupNumPerFamily |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxStoreSizePerFamily", wireType)
			}
			m.MaxStoreSizePerFamily = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxStoreSizePerFamily |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)