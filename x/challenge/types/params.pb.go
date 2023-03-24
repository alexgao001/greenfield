// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/challenge/params.proto

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
	// Challenges which will be emitted in each block, including user submitted or randomly triggered.
	ChallengeCountPerBlock uint64 `protobuf:"varint,1,opt,name=challenge_count_per_block,json=challengeCountPerBlock,proto3" json:"challenge_count_per_block,omitempty" yaml:"challenge_count_per_block"`
	// Challenges will be expired after the period, including user submitted or randomly triggered.
	ChallengeKeepAlivePeriod uint64 `protobuf:"varint,2,opt,name=challenge_keep_alive_period,json=challengeKeepAlivePeriod,proto3" json:"challenge_keep_alive_period,omitempty" yaml:"challenge_keep_alive_period"`
	// The count of blocks to stand for the period in which the same storage and object info cannot be slashed again.
	SlashCoolingOffPeriod uint64 `protobuf:"varint,3,opt,name=slash_cooling_off_period,json=slashCoolingOffPeriod,proto3" json:"slash_cooling_off_period,omitempty" yaml:"slash_cooling_off_period"`
	// The slash coin amount will be calculated from the size of object info, and adjusted by this rate.
	SlashAmountSizeRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=slash_amount_size_rate,json=slashAmountSizeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_amount_size_rate" yaml:"slash_amount_size_rate"`
	// The minimal slash amount.
	SlashAmountMin github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=slash_amount_min,json=slashAmountMin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"slash_amount_min"`
	// The maximum slash amount.
	SlashAmountMax github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=slash_amount_max,json=slashAmountMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"slash_amount_max"`
	// The ratio of slash amount to reward all current validators.
	RewardValidatorRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=reward_validator_ratio,json=rewardValidatorRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_validator_ratio" yaml:"reward_validator_ratio"`
	// The ratio of reward amount to reward attestation submitter.
	RewardSubmitterRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=reward_submitter_ratio,json=rewardSubmitterRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_submitter_ratio" yaml:"reward_challenger_ratio"`
	// The reward amount to submitter will be adjusted by the threshold.
	RewardSubmitterThreshold github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=reward_submitter_threshold,json=rewardSubmitterThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reward_submitter_threshold"`
	// Heartbeat interval, based on challenge id, defines the frequency of heartbeat attestation.
	HeartbeatInterval uint64 `protobuf:"varint,10,opt,name=heartbeat_interval,json=heartbeatInterval,proto3" json:"heartbeat_interval,omitempty" yaml:"heartbeat_interval"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2396367ee53edf57, []int{0}
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

func (m *Params) GetChallengeCountPerBlock() uint64 {
	if m != nil {
		return m.ChallengeCountPerBlock
	}
	return 0
}

func (m *Params) GetChallengeKeepAlivePeriod() uint64 {
	if m != nil {
		return m.ChallengeKeepAlivePeriod
	}
	return 0
}

func (m *Params) GetSlashCoolingOffPeriod() uint64 {
	if m != nil {
		return m.SlashCoolingOffPeriod
	}
	return 0
}

func (m *Params) GetHeartbeatInterval() uint64 {
	if m != nil {
		return m.HeartbeatInterval
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "bnbchain.greenfield.challenge.Params")
}

func init() { proto.RegisterFile("greenfield/challenge/params.proto", fileDescriptor_2396367ee53edf57) }

var fileDescriptor_2396367ee53edf57 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0x8d, 0xdf, 0x2b, 0xa5, 0x9d, 0x05, 0x02, 0x53, 0x2a, 0xb7, 0x28, 0x76, 0x30, 0xa8, 0xea,
	0x26, 0x89, 0x10, 0xbb, 0x8a, 0x4d, 0x53, 0x36, 0x51, 0x41, 0x8d, 0x5c, 0xc4, 0x02, 0x21, 0x59,
	0x63, 0xfb, 0xda, 0x1e, 0xc5, 0x9e, 0xb1, 0xc6, 0x93, 0x90, 0x66, 0x0d, 0xac, 0x59, 0xb2, 0xe4,
	0x23, 0xf8, 0x88, 0x2e, 0x2b, 0x56, 0x88, 0x85, 0x85, 0x12, 0x89, 0x0f, 0xc8, 0x17, 0x20, 0x8f,
	0x13, 0x27, 0x4d, 0xe8, 0xa2, 0xa2, 0xab, 0xc4, 0xe7, 0x9c, 0x7b, 0xcf, 0xb9, 0x9e, 0xf1, 0x45,
	0x8f, 0x02, 0x0e, 0x40, 0x7d, 0x02, 0x91, 0xd7, 0x74, 0x43, 0x1c, 0x45, 0x40, 0x03, 0x68, 0x26,
	0x98, 0xe3, 0x38, 0x6d, 0x24, 0x9c, 0x09, 0xa6, 0x56, 0x1d, 0xea, 0xb8, 0x21, 0x26, 0xb4, 0x31,
	0xd7, 0x36, 0x4a, 0xed, 0xee, 0x8e, 0xcb, 0xd2, 0x98, 0xa5, 0xb6, 0x14, 0x37, 0x8b, 0x87, 0xa2,
	0x72, 0x77, 0x2b, 0x60, 0x01, 0x2b, 0xf0, 0xfc, 0x5f, 0x81, 0x9a, 0xbf, 0x37, 0xd0, 0x7a, 0x47,
	0x1a, 0xa8, 0x36, 0xda, 0x29, 0x1b, 0xd9, 0x2e, 0xeb, 0x51, 0x61, 0x27, 0xc0, 0x6d, 0x27, 0x62,
	0x6e, 0x57, 0x53, 0x6a, 0xca, 0xfe, 0x5a, 0xeb, 0xc9, 0x24, 0x33, 0x6a, 0x67, 0x38, 0x8e, 0x0e,
	0xcc, 0x2b, 0xa5, 0xa6, 0xb5, 0x5d, 0x72, 0x47, 0x39, 0xd5, 0x01, 0xde, 0xca, 0x09, 0x15, 0xd0,
	0xc3, 0x79, 0x55, 0x17, 0x20, 0xb1, 0x71, 0x44, 0xfa, 0x90, 0x97, 0x12, 0xe6, 0x69, 0xff, 0x49,
	0x8b, 0xbd, 0x49, 0x66, 0x98, 0xcb, 0x16, 0x2b, 0x62, 0xd3, 0xd2, 0x4a, 0xf6, 0x18, 0x20, 0x39,
	0xcc, 0xb9, 0x8e, 0xa4, 0xd4, 0x77, 0x48, 0x4b, 0x23, 0x9c, 0x86, 0xb6, 0xcb, 0x58, 0x44, 0x68,
	0x60, 0x33, 0xdf, 0x9f, 0x79, 0xfc, 0x2f, 0x3d, 0x1e, 0x4f, 0x32, 0xc3, 0x28, 0x3c, 0xae, 0x52,
	0x9a, 0xd6, 0x03, 0x49, 0x1d, 0x15, 0xcc, 0x89, 0xef, 0x4f, 0xbb, 0x7f, 0x50, 0xd0, 0x76, 0x51,
	0x84, 0x63, 0x39, 0x78, 0x4a, 0x86, 0x60, 0x73, 0x2c, 0x40, 0x5b, 0xab, 0x29, 0xfb, 0x9b, 0xad,
	0x93, 0xf3, 0xcc, 0xa8, 0xfc, 0xcc, 0x8c, 0xbd, 0x80, 0x88, 0xb0, 0xe7, 0x34, 0x5c, 0x16, 0x4f,
	0x0f, 0x62, 0xfa, 0x53, 0x4f, 0xbd, 0x6e, 0x53, 0x9c, 0x25, 0x90, 0x36, 0x5e, 0x80, 0x3b, 0xc9,
	0x8c, 0xea, 0x62, 0x94, 0xe5, 0xae, 0xa6, 0x75, 0x5f, 0x12, 0x87, 0x12, 0x3f, 0x25, 0x43, 0xb0,
	0xb0, 0x00, 0xd5, 0x47, 0x77, 0x2f, 0xe9, 0x63, 0x42, 0xb5, 0x5b, 0xd2, 0xff, 0xf9, 0x35, 0xfc,
	0xdb, 0x54, 0x7c, 0xff, 0x56, 0x47, 0xd3, 0x7b, 0xd2, 0xa6, 0xc2, 0xba, 0xb3, 0x60, 0xf6, 0x8a,
	0xd0, 0x55, 0x1f, 0x3c, 0xd0, 0xd6, 0x6f, 0xda, 0x07, 0x0f, 0xd4, 0x8f, 0x0a, 0xda, 0xe6, 0xf0,
	0x1e, 0x73, 0xcf, 0xee, 0xe3, 0x88, 0x78, 0x58, 0x30, 0x9e, 0xcf, 0x4f, 0x98, 0x76, 0xfb, 0xdf,
	0x5e, 0xeb, 0xdf, 0xbb, 0x9a, 0xd6, 0x56, 0x41, 0xbc, 0x99, 0xe1, 0x56, 0x0e, 0xab, 0x9f, 0xe6,
	0x39, 0xd2, 0x9e, 0x13, 0x13, 0x21, 0x60, 0x96, 0x63, 0x43, 0xe6, 0xe8, 0x5c, 0x3b, 0x87, 0x7e,
	0x29, 0x47, 0x79, 0x6d, 0x97, 0x83, 0x9c, 0xce, 0xec, 0x8a, 0x20, 0x43, 0xb4, 0xbb, 0x92, 0x43,
	0x84, 0x1c, 0xd2, 0x90, 0x45, 0x9e, 0xb6, 0x79, 0x03, 0x47, 0xa0, 0x2d, 0xf9, 0xbe, 0x9e, 0x75,
	0x57, 0x5f, 0x22, 0x35, 0x04, 0xcc, 0x85, 0x03, 0x58, 0xd8, 0x84, 0x0a, 0xe0, 0x7d, 0x1c, 0x69,
	0x48, 0x7e, 0x3b, 0xd5, 0x49, 0x66, 0xec, 0x14, 0x13, 0xad, 0x6a, 0x4c, 0xeb, 0x5e, 0x09, 0xb6,
	0xa7, 0xd8, 0xc1, 0xda, 0x97, 0xaf, 0x46, 0xa5, 0x75, 0x7c, 0x3e, 0xd2, 0x95, 0x8b, 0x91, 0xae,
	0xfc, 0x1a, 0xe9, 0xca, 0xe7, 0xb1, 0x5e, 0xb9, 0x18, 0xeb, 0x95, 0x1f, 0x63, 0xbd, 0xf2, 0xf6,
	0xe9, 0x42, 0x7a, 0x87, 0x3a, 0x75, 0xb9, 0xde, 0x9a, 0x0b, 0xab, 0x70, 0xb0, 0xb0, 0x0c, 0xe5,
	0x30, 0xce, 0xba, 0x5c, 0x5e, 0xcf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x14, 0x09, 0xe1, 0x32,
	0x31, 0x05, 0x00, 0x00,
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
	if m.HeartbeatInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.HeartbeatInterval))
		i--
		dAtA[i] = 0x50
	}
	{
		size := m.RewardSubmitterThreshold.Size()
		i -= size
		if _, err := m.RewardSubmitterThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.RewardSubmitterRatio.Size()
		i -= size
		if _, err := m.RewardSubmitterRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.RewardValidatorRatio.Size()
		i -= size
		if _, err := m.RewardValidatorRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.SlashAmountMax.Size()
		i -= size
		if _, err := m.SlashAmountMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.SlashAmountMin.Size()
		i -= size
		if _, err := m.SlashAmountMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.SlashAmountSizeRate.Size()
		i -= size
		if _, err := m.SlashAmountSizeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.SlashCoolingOffPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SlashCoolingOffPeriod))
		i--
		dAtA[i] = 0x18
	}
	if m.ChallengeKeepAlivePeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ChallengeKeepAlivePeriod))
		i--
		dAtA[i] = 0x10
	}
	if m.ChallengeCountPerBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ChallengeCountPerBlock))
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
	if m.ChallengeCountPerBlock != 0 {
		n += 1 + sovParams(uint64(m.ChallengeCountPerBlock))
	}
	if m.ChallengeKeepAlivePeriod != 0 {
		n += 1 + sovParams(uint64(m.ChallengeKeepAlivePeriod))
	}
	if m.SlashCoolingOffPeriod != 0 {
		n += 1 + sovParams(uint64(m.SlashCoolingOffPeriod))
	}
	l = m.SlashAmountSizeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashAmountMin.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashAmountMax.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.RewardValidatorRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.RewardSubmitterRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.RewardSubmitterThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.HeartbeatInterval != 0 {
		n += 1 + sovParams(uint64(m.HeartbeatInterval))
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeCountPerBlock", wireType)
			}
			m.ChallengeCountPerBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengeCountPerBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeKeepAlivePeriod", wireType)
			}
			m.ChallengeKeepAlivePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengeKeepAlivePeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashCoolingOffPeriod", wireType)
			}
			m.SlashCoolingOffPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SlashCoolingOffPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmountSizeRate", wireType)
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
			if err := m.SlashAmountSizeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmountMin", wireType)
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
			if err := m.SlashAmountMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmountMax", wireType)
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
			if err := m.SlashAmountMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardValidatorRatio", wireType)
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
			if err := m.RewardValidatorRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardSubmitterRatio", wireType)
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
			if err := m.RewardSubmitterRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardSubmitterThreshold", wireType)
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
			if err := m.RewardSubmitterThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeartbeatInterval", wireType)
			}
			m.HeartbeatInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeartbeatInterval |= uint64(b&0x7F) << shift
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
