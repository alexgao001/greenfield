// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/payment/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	// Time duration which the buffer balance need to be reserved for NetOutFlow e.g. 6 month
	ReserveTime uint64 `protobuf:"varint,1,opt,name=reserve_time,json=reserveTime,proto3" json:"reserve_time,omitempty" yaml:"reserve_time"`
	// The maximum number of payment accounts that can be created by one user
	PaymentAccountCountLimit uint64 `protobuf:"varint,2,opt,name=payment_account_count_limit,json=paymentAccountCountLimit,proto3" json:"payment_account_count_limit,omitempty" yaml:"payment_account_count_limit"`
	// Time duration threshold of forced settlement.
	// If dynamic balance is less than NetOutFlowRate * forcedSettleTime, the account can be forced settled.
	ForcedSettleTime uint64 `protobuf:"varint,3,opt,name=forced_settle_time,json=forcedSettleTime,proto3" json:"forced_settle_time,omitempty" yaml:"forced_settle_time"`
	// the maximum number of accounts that will be forced settled in one block
	MaxAutoForceSettleNum uint64 `protobuf:"varint,4,opt,name=max_auto_force_settle_num,json=maxAutoForceSettleNum,proto3" json:"max_auto_force_settle_num,omitempty" yaml:"max_auto_force_settle_num"`
	// The denom of fee charged in payment module
	FeeDenom string `protobuf:"bytes,5,opt,name=fee_denom,json=feeDenom,proto3" json:"fee_denom,omitempty" yaml:"fee_denom"`
	// The tax rate to pay for validators in storage payment. The default value is 1%(0.01)
	ValidatorTaxRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=validator_tax_rate,json=validatorTaxRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"validator_tax_rate"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd7d37632356c8f4, []int{0}
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

func (m *Params) GetReserveTime() uint64 {
	if m != nil {
		return m.ReserveTime
	}
	return 0
}

func (m *Params) GetPaymentAccountCountLimit() uint64 {
	if m != nil {
		return m.PaymentAccountCountLimit
	}
	return 0
}

func (m *Params) GetForcedSettleTime() uint64 {
	if m != nil {
		return m.ForcedSettleTime
	}
	return 0
}

func (m *Params) GetMaxAutoForceSettleNum() uint64 {
	if m != nil {
		return m.MaxAutoForceSettleNum
	}
	return 0
}

func (m *Params) GetFeeDenom() string {
	if m != nil {
		return m.FeeDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "bnbchain.greenfield.payment.Params")
}

func init() { proto.RegisterFile("greenfield/payment/params.proto", fileDescriptor_bd7d37632356c8f4) }

var fileDescriptor_bd7d37632356c8f4 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x37, 0xb6, 0x2e, 0x76, 0xf4, 0xb0, 0xc4, 0x8a, 0xd9, 0x16, 0x93, 0x25, 0x48, 0xe9,
	0x65, 0x13, 0xc4, 0x5b, 0xf1, 0xd2, 0x75, 0x11, 0x44, 0x11, 0x89, 0x3d, 0x79, 0x70, 0x98, 0x24,
	0x6f, 0xb7, 0xa3, 0x99, 0x99, 0x30, 0x99, 0x94, 0xec, 0x7f, 0xe1, 0x1f, 0xe3, 0x1f, 0xd1, 0x63,
	0xf1, 0x24, 0x1e, 0x82, 0xec, 0xfe, 0x07, 0xb9, 0x78, 0x95, 0xcc, 0x4c, 0xeb, 0x82, 0xd8, 0xcb,
	0xfc, 0x78, 0xdf, 0xef, 0xfb, 0xbc, 0x07, 0xef, 0xa1, 0x60, 0x29, 0x01, 0xf8, 0x82, 0x42, 0x91,
	0xc7, 0x25, 0x59, 0x31, 0xe0, 0x2a, 0x2e, 0x89, 0x24, 0xac, 0x8a, 0x4a, 0x29, 0x94, 0x70, 0x0f,
	0x53, 0x9e, 0x66, 0xe7, 0x84, 0xf2, 0xe8, 0xaf, 0x33, 0xb2, 0xce, 0x83, 0x71, 0x26, 0x2a, 0x26,
	0x2a, 0xac, 0xad, 0xb1, 0xf9, 0x98, 0xbc, 0x83, 0xfd, 0xa5, 0x58, 0x0a, 0x13, 0xef, 0x5f, 0x26,
	0x1a, 0xfe, 0xde, 0x41, 0xc3, 0xf7, 0x1a, 0xef, 0x9e, 0xa0, 0x07, 0x12, 0x2a, 0x90, 0x17, 0x80,
	0x15, 0x65, 0xe0, 0x39, 0x13, 0xe7, 0x78, 0x77, 0xf6, 0xb8, 0x6b, 0x83, 0x87, 0x2b, 0xc2, 0x8a,
	0x93, 0x70, 0x5b, 0x0d, 0x93, 0xfb, 0xf6, 0x7b, 0x46, 0x19, 0xb8, 0x80, 0x0e, 0x6d, 0x0b, 0x98,
	0x64, 0x99, 0xa8, 0xb9, 0xc2, 0xe6, 0x2c, 0x28, 0xa3, 0xca, 0xbb, 0xa3, 0x51, 0x47, 0x5d, 0x1b,
	0x84, 0x06, 0x75, 0x8b, 0x39, 0x4c, 0x3c, 0xab, 0x9e, 0x1a, 0xf1, 0x65, 0x7f, 0xbc, 0xed, 0x25,
	0xf7, 0x0d, 0x72, 0x17, 0x42, 0x66, 0x90, 0xe3, 0x0a, 0x94, 0x2a, 0x6c, 0xa3, 0x3b, 0x9a, 0xfe,
	0xa4, 0x6b, 0x83, 0xb1, 0xa1, 0xff, 0xeb, 0x09, 0x93, 0x91, 0x09, 0x7e, 0xd0, 0x31, 0xdd, 0xf3,
	0x27, 0x34, 0x66, 0xa4, 0xc1, 0xa4, 0x56, 0x02, 0x6b, 0xf1, 0x3a, 0x81, 0xd7, 0xcc, 0xdb, 0xd5,
	0xcc, 0xa7, 0x5d, 0x1b, 0x4c, 0x0c, 0xf3, 0xbf, 0xd6, 0x30, 0x79, 0xc4, 0x48, 0x73, 0x5a, 0x2b,
	0xf1, 0xaa, 0x57, 0x4c, 0x81, 0x77, 0x35, 0x73, 0x9f, 0xa1, 0xbd, 0x05, 0x00, 0xce, 0x81, 0x0b,
	0xe6, 0xdd, 0x9d, 0x38, 0xc7, 0x7b, 0xb3, 0xfd, 0xae, 0x0d, 0x46, 0xb6, 0xc7, 0x6b, 0x29, 0x4c,
	0xee, 0x2d, 0x00, 0xe6, 0xfd, 0xd3, 0xfd, 0x8c, 0xdc, 0x0b, 0x52, 0xd0, 0x9c, 0x28, 0x21, 0xb1,
	0x22, 0x0d, 0x96, 0x44, 0x81, 0x37, 0xd4, 0xb9, 0x2f, 0x2e, 0xdb, 0x60, 0xf0, 0xb3, 0x0d, 0x8e,
	0x96, 0x54, 0x9d, 0xd7, 0x69, 0x94, 0x09, 0x66, 0x07, 0x6c, 0xaf, 0x69, 0x95, 0x7f, 0x89, 0xd5,
	0xaa, 0x84, 0x2a, 0x9a, 0x43, 0xf6, 0xfd, 0xdb, 0x14, 0xd9, 0xf9, 0xcf, 0x21, 0x4b, 0x46, 0x37,
	0xdc, 0x33, 0xd2, 0x24, 0x44, 0xc1, 0xec, 0xf5, 0xe5, 0xda, 0x77, 0xae, 0xd6, 0xbe, 0xf3, 0x6b,
	0xed, 0x3b, 0x5f, 0x37, 0xfe, 0xe0, 0x6a, 0xe3, 0x0f, 0x7e, 0x6c, 0xfc, 0xc1, 0xc7, 0x78, 0xab,
	0x42, 0xca, 0xd3, 0xa9, 0xde, 0xb6, 0x78, 0x6b, 0x2f, 0x9b, 0x9b, 0xcd, 0xd4, 0xe5, 0xd2, 0xa1,
	0xde, 0xa5, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x06, 0x43, 0x4e, 0x46, 0xbc, 0x02, 0x00,
	0x00,
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
	{
		size := m.ValidatorTaxRate.Size()
		i -= size
		if _, err := m.ValidatorTaxRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.FeeDenom) > 0 {
		i -= len(m.FeeDenom)
		copy(dAtA[i:], m.FeeDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.FeeDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if m.MaxAutoForceSettleNum != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxAutoForceSettleNum))
		i--
		dAtA[i] = 0x20
	}
	if m.ForcedSettleTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ForcedSettleTime))
		i--
		dAtA[i] = 0x18
	}
	if m.PaymentAccountCountLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PaymentAccountCountLimit))
		i--
		dAtA[i] = 0x10
	}
	if m.ReserveTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ReserveTime))
		i--
		dAtA[i] = 0x8
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
	if m.ReserveTime != 0 {
		n += 1 + sovParams(uint64(m.ReserveTime))
	}
	if m.PaymentAccountCountLimit != 0 {
		n += 1 + sovParams(uint64(m.PaymentAccountCountLimit))
	}
	if m.ForcedSettleTime != 0 {
		n += 1 + sovParams(uint64(m.ForcedSettleTime))
	}
	if m.MaxAutoForceSettleNum != 0 {
		n += 1 + sovParams(uint64(m.MaxAutoForceSettleNum))
	}
	l = len(m.FeeDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.ValidatorTaxRate.Size()
	n += 1 + l + sovParams(uint64(l))
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveTime", wireType)
			}
			m.ReserveTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReserveTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentAccountCountLimit", wireType)
			}
			m.PaymentAccountCountLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PaymentAccountCountLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForcedSettleTime", wireType)
			}
			m.ForcedSettleTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ForcedSettleTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAutoForceSettleNum", wireType)
			}
			m.MaxAutoForceSettleNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAutoForceSettleNum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeDenom", wireType)
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
			m.FeeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorTaxRate", wireType)
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
			if err := m.ValidatorTaxRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
