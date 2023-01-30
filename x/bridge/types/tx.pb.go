// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/bridge/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgTransferOut struct {
	From   string      `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To     string      `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount *types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *MsgTransferOut) Reset()         { *m = MsgTransferOut{} }
func (m *MsgTransferOut) String() string { return proto.CompactTextString(m) }
func (*MsgTransferOut) ProtoMessage()    {}
func (*MsgTransferOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_5360e58e7e095845, []int{0}
}
func (m *MsgTransferOut) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferOut.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferOut.Merge(m, src)
}
func (m *MsgTransferOut) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferOut) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferOut.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferOut proto.InternalMessageInfo

func (m *MsgTransferOut) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *MsgTransferOut) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *MsgTransferOut) GetAmount() *types.Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

type MsgTransferOutResponse struct {
}

func (m *MsgTransferOutResponse) Reset()         { *m = MsgTransferOutResponse{} }
func (m *MsgTransferOutResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTransferOutResponse) ProtoMessage()    {}
func (*MsgTransferOutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5360e58e7e095845, []int{1}
}
func (m *MsgTransferOutResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferOutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferOutResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferOutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferOutResponse.Merge(m, src)
}
func (m *MsgTransferOutResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferOutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferOutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferOutResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgTransferOut)(nil), "bnbchain.greenfield.bridge.MsgTransferOut")
	proto.RegisterType((*MsgTransferOutResponse)(nil), "bnbchain.greenfield.bridge.MsgTransferOutResponse")
}

func init() { proto.RegisterFile("greenfield/bridge/tx.proto", fileDescriptor_5360e58e7e095845) }

var fileDescriptor_5360e58e7e095845 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x6a, 0xeb, 0x30,
	0x14, 0x86, 0xa3, 0xe4, 0x12, 0xb8, 0x0a, 0x64, 0xd0, 0x50, 0x5c, 0x0f, 0x22, 0x64, 0x0a, 0x85,
	0x1e, 0x61, 0xf7, 0x0d, 0xda, 0xa5, 0x4b, 0x28, 0x98, 0x4e, 0xdd, 0x2c, 0x47, 0x56, 0x04, 0xb5,
	0x8e, 0x91, 0xe4, 0x92, 0xbe, 0x45, 0x1f, 0xab, 0x63, 0xc6, 0x8e, 0xc5, 0x7e, 0x91, 0x52, 0x3b,
	0x6d, 0x93, 0xa1, 0xd0, 0x4d, 0x88, 0x4f, 0xff, 0xa7, 0xf3, 0x1f, 0x1a, 0x6b, 0xa7, 0x94, 0x2d,
	0x8d, 0x7a, 0xdc, 0x08, 0xe9, 0xcc, 0x46, 0x2b, 0x11, 0x76, 0x50, 0x3b, 0x0c, 0xc8, 0x62, 0x69,
	0x65, 0xb1, 0xcd, 0x8d, 0x85, 0x1f, 0x08, 0x06, 0x28, 0xe6, 0x05, 0xfa, 0x0a, 0xbd, 0x90, 0xb9,
	0x57, 0xe2, 0x29, 0x91, 0x2a, 0xe4, 0x89, 0x28, 0xd0, 0xd8, 0xe1, 0xed, 0x52, 0xd3, 0xf9, 0xda,
	0xeb, 0x7b, 0x97, 0x5b, 0x5f, 0x2a, 0x77, 0xd7, 0x04, 0xc6, 0xe8, 0xbf, 0xd2, 0x61, 0x15, 0x91,
	0x05, 0x59, 0xfd, 0xcf, 0xfa, 0x33, 0x9b, 0xd3, 0x71, 0xc0, 0x68, 0xdc, 0xdf, 0x8c, 0x03, 0xb2,
	0x84, 0x4e, 0xf3, 0x0a, 0x1b, 0x1b, 0xa2, 0xc9, 0x82, 0xac, 0x66, 0xe9, 0x39, 0x0c, 0x1a, 0xf8,
	0xd4, 0xc0, 0x41, 0x03, 0x37, 0x68, 0x6c, 0x76, 0x00, 0x97, 0x11, 0x3d, 0x3b, 0x15, 0x65, 0xca,
	0xd7, 0x68, 0xbd, 0x4a, 0x03, 0x9d, 0xac, 0xbd, 0x66, 0x15, 0x9d, 0x1d, 0x7f, 0xe3, 0x02, 0x7e,
	0x9f, 0x0a, 0x4e, 0x93, 0xe2, 0xf4, 0xef, 0xec, 0x97, 0xf5, 0xfa, 0xf6, 0xb5, 0xe5, 0x64, 0xdf,
	0x72, 0xf2, 0xde, 0x72, 0xf2, 0xd2, 0xf1, 0xd1, 0xbe, 0xe3, 0xa3, 0xb7, 0x8e, 0x8f, 0x1e, 0x40,
	0x9b, 0xb0, 0x6d, 0x24, 0x14, 0x58, 0x09, 0x69, 0xe5, 0x65, 0x1f, 0x2c, 0x8e, 0xfa, 0xdf, 0x7d,
	0x6f, 0xe0, 0xb9, 0x56, 0x5e, 0x4e, 0xfb, 0x26, 0xaf, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x44,
	0xe6, 0xd7, 0x34, 0xa3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	TransferOut(ctx context.Context, in *MsgTransferOut, opts ...grpc.CallOption) (*MsgTransferOutResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) TransferOut(ctx context.Context, in *MsgTransferOut, opts ...grpc.CallOption) (*MsgTransferOutResponse, error) {
	out := new(MsgTransferOutResponse)
	err := c.cc.Invoke(ctx, "/bnbchain.greenfield.bridge.Msg/TransferOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	TransferOut(context.Context, *MsgTransferOut) (*MsgTransferOutResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) TransferOut(ctx context.Context, req *MsgTransferOut) (*MsgTransferOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferOut not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_TransferOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferOut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bnbchain.greenfield.bridge.Msg/TransferOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferOut(ctx, req.(*MsgTransferOut))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bnbchain.greenfield.bridge.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransferOut",
			Handler:    _Msg_TransferOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greenfield/bridge/tx.proto",
}

func (m *MsgTransferOut) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferOut) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferOut) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != nil {
		{
			size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintTx(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransferOutResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferOutResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferOutResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgTransferOut) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != nil {
		l = m.Amount.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgTransferOutResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgTransferOut) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTransferOut: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferOut: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Amount == nil {
				m.Amount = &types.Coin{}
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgTransferOutResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTransferOutResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferOutResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
