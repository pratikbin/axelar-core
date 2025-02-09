// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/nexus/v1beta1/events.proto

package types

import (
	fmt "fmt"
	github_com_axelarnetwork_axelar_core_x_nexus_exported "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type FeeDeducted struct {
	TransferID       github_com_axelarnetwork_axelar_core_x_nexus_exported.TransferID `protobuf:"varint,1,opt,name=transfer_id,json=transferId,proto3,casttype=github.com/axelarnetwork/axelar-core/x/nexus/exported.TransferID" json:"transfer_id,omitempty"`
	RecipientChain   github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName  `protobuf:"bytes,2,opt,name=recipient_chain,json=recipientChain,proto3,casttype=github.com/axelarnetwork/axelar-core/x/nexus/exported.ChainName" json:"recipient_chain,omitempty"`
	RecipientAddress string                                                           `protobuf:"bytes,3,opt,name=recipient_address,json=recipientAddress,proto3" json:"recipient_address,omitempty"`
	Amount           types.Coin                                                       `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount"`
	Fee              types.Coin                                                       `protobuf:"bytes,5,opt,name=fee,proto3" json:"fee"`
}

func (m *FeeDeducted) Reset()         { *m = FeeDeducted{} }
func (m *FeeDeducted) String() string { return proto.CompactTextString(m) }
func (*FeeDeducted) ProtoMessage()    {}
func (*FeeDeducted) Descriptor() ([]byte, []int) {
	return fileDescriptor_4433ea5171b09eb9, []int{0}
}
func (m *FeeDeducted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeDeducted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeDeducted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeDeducted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeDeducted.Merge(m, src)
}
func (m *FeeDeducted) XXX_Size() int {
	return m.Size()
}
func (m *FeeDeducted) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeDeducted.DiscardUnknown(m)
}

var xxx_messageInfo_FeeDeducted proto.InternalMessageInfo

func (m *FeeDeducted) GetTransferID() github_com_axelarnetwork_axelar_core_x_nexus_exported.TransferID {
	if m != nil {
		return m.TransferID
	}
	return 0
}

func (m *FeeDeducted) GetRecipientChain() github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName {
	if m != nil {
		return m.RecipientChain
	}
	return ""
}

func (m *FeeDeducted) GetRecipientAddress() string {
	if m != nil {
		return m.RecipientAddress
	}
	return ""
}

func (m *FeeDeducted) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *FeeDeducted) GetFee() types.Coin {
	if m != nil {
		return m.Fee
	}
	return types.Coin{}
}

type InsufficientFee struct {
	TransferID       github_com_axelarnetwork_axelar_core_x_nexus_exported.TransferID `protobuf:"varint,1,opt,name=transfer_id,json=transferId,proto3,casttype=github.com/axelarnetwork/axelar-core/x/nexus/exported.TransferID" json:"transfer_id,omitempty"`
	RecipientChain   github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName  `protobuf:"bytes,2,opt,name=recipient_chain,json=recipientChain,proto3,casttype=github.com/axelarnetwork/axelar-core/x/nexus/exported.ChainName" json:"recipient_chain,omitempty"`
	RecipientAddress string                                                           `protobuf:"bytes,3,opt,name=recipient_address,json=recipientAddress,proto3" json:"recipient_address,omitempty"`
	Amount           types.Coin                                                       `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount"`
	Fee              types.Coin                                                       `protobuf:"bytes,5,opt,name=fee,proto3" json:"fee"`
}

func (m *InsufficientFee) Reset()         { *m = InsufficientFee{} }
func (m *InsufficientFee) String() string { return proto.CompactTextString(m) }
func (*InsufficientFee) ProtoMessage()    {}
func (*InsufficientFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_4433ea5171b09eb9, []int{1}
}
func (m *InsufficientFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InsufficientFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InsufficientFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InsufficientFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsufficientFee.Merge(m, src)
}
func (m *InsufficientFee) XXX_Size() int {
	return m.Size()
}
func (m *InsufficientFee) XXX_DiscardUnknown() {
	xxx_messageInfo_InsufficientFee.DiscardUnknown(m)
}

var xxx_messageInfo_InsufficientFee proto.InternalMessageInfo

func (m *InsufficientFee) GetTransferID() github_com_axelarnetwork_axelar_core_x_nexus_exported.TransferID {
	if m != nil {
		return m.TransferID
	}
	return 0
}

func (m *InsufficientFee) GetRecipientChain() github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName {
	if m != nil {
		return m.RecipientChain
	}
	return ""
}

func (m *InsufficientFee) GetRecipientAddress() string {
	if m != nil {
		return m.RecipientAddress
	}
	return ""
}

func (m *InsufficientFee) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *InsufficientFee) GetFee() types.Coin {
	if m != nil {
		return m.Fee
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*FeeDeducted)(nil), "axelar.nexus.v1beta1.FeeDeducted")
	proto.RegisterType((*InsufficientFee)(nil), "axelar.nexus.v1beta1.InsufficientFee")
}

func init() { proto.RegisterFile("axelar/nexus/v1beta1/events.proto", fileDescriptor_4433ea5171b09eb9) }

var fileDescriptor_4433ea5171b09eb9 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x93, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0xc6, 0x13, 0x52, 0xae, 0x84, 0x2b, 0x71, 0x21, 0xba, 0x43, 0xe8, 0x90, 0x94, 0x4e, 0x95,
	0x10, 0x8e, 0x5a, 0x06, 0x46, 0x20, 0xad, 0x2a, 0x55, 0x42, 0x0c, 0x11, 0x13, 0x4b, 0xe5, 0x38,
	0x27, 0xad, 0x45, 0x63, 0x47, 0xb6, 0x53, 0xc2, 0x23, 0x30, 0x20, 0xf1, 0x58, 0x1d, 0x3b, 0x32,
	0x45, 0x28, 0x7d, 0x8b, 0x4e, 0x28, 0x7f, 0xda, 0xae, 0xc0, 0xc2, 0x72, 0xb7, 0xe4, 0x9c, 0x9f,
	0xbf, 0xef, 0xf8, 0x3b, 0x32, 0x7a, 0x4e, 0x0a, 0xd8, 0x12, 0xe9, 0x73, 0x28, 0x72, 0xe5, 0xef,
	0x26, 0x11, 0x68, 0x32, 0xf1, 0x61, 0x07, 0x5c, 0x2b, 0x9c, 0x49, 0xa1, 0x85, 0x7d, 0xd7, 0x22,
	0xb8, 0x41, 0x70, 0x87, 0x0c, 0xee, 0xd6, 0x62, 0x2d, 0x1a, 0xc0, 0xaf, 0xbf, 0x5a, 0x76, 0xe0,
	0x52, 0xa1, 0x52, 0xa1, 0xfc, 0x88, 0x28, 0xb8, 0xa8, 0x51, 0xc1, 0x78, 0xdb, 0x1f, 0x7d, 0xb3,
	0x50, 0x7f, 0x01, 0x30, 0x87, 0x38, 0xa7, 0x1a, 0x62, 0x5b, 0xa1, 0xbe, 0x96, 0x84, 0xab, 0x04,
	0xe4, 0x8a, 0xc5, 0x8e, 0x39, 0x34, 0xc7, 0xbd, 0x20, 0xac, 0x4a, 0x0f, 0x7d, 0xec, 0xca, 0xcb,
	0xf9, 0xa9, 0xf4, 0xde, 0xae, 0x99, 0xde, 0xe4, 0x11, 0xa6, 0x22, 0xf5, 0xdb, 0x69, 0x38, 0xe8,
	0x2f, 0x42, 0x7e, 0xee, 0xfe, 0x5e, 0x52, 0x21, 0xc1, 0x2f, 0xba, 0x5b, 0x40, 0x91, 0x09, 0xa9,
	0x21, 0xc6, 0x57, 0x8d, 0x10, 0x9d, 0x6d, 0x96, 0xb1, 0xbd, 0x45, 0xb7, 0x12, 0x28, 0xcb, 0x18,
	0x70, 0xbd, 0xa2, 0x1b, 0xc2, 0xb8, 0xf3, 0x60, 0x68, 0x8e, 0x1f, 0x05, 0xb3, 0x53, 0xe9, 0xbd,
	0xf9, 0x37, 0xab, 0x59, 0x2d, 0xf3, 0x81, 0xa4, 0x10, 0x3e, 0xbe, 0x68, 0x37, 0x35, 0xfb, 0x05,
	0x7a, 0x7a, 0x75, 0x23, 0x71, 0x2c, 0x41, 0x29, 0xc7, 0xaa, 0xfd, 0xc2, 0x27, 0x97, 0xc6, 0xbb,
	0xb6, 0x6e, 0xbf, 0x46, 0x37, 0x24, 0x15, 0x39, 0xd7, 0x4e, 0x6f, 0x68, 0x8e, 0xfb, 0xd3, 0x67,
	0xb8, 0x0d, 0x14, 0xd7, 0x81, 0x9e, 0xb3, 0xc7, 0x33, 0xc1, 0x78, 0xd0, 0xdb, 0x97, 0x9e, 0x11,
	0x76, 0xb8, 0x3d, 0x41, 0x56, 0x02, 0xe0, 0x3c, 0xfc, 0xb3, 0x53, 0x35, 0x3b, 0xfa, 0x6e, 0xa1,
	0xdb, 0x25, 0x57, 0x79, 0x92, 0x30, 0x5a, 0xcf, 0xb0, 0x00, 0xb8, 0xdf, 0xc7, 0xff, 0xdb, 0x47,
	0xf0, 0x7e, 0x5f, 0xb9, 0xe6, 0xa1, 0x72, 0xcd, 0x5f, 0x95, 0x6b, 0xfe, 0x38, 0xba, 0xc6, 0xe1,
	0xe8, 0x1a, 0x3f, 0x8f, 0xae, 0xf1, 0x69, 0xfa, 0x57, 0x19, 0xe8, 0xaf, 0x19, 0xa8, 0xe8, 0xa6,
	0x79, 0x70, 0xaf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x9c, 0x8d, 0x7a, 0xe1, 0x03, 0x00,
	0x00,
}

func (m *FeeDeducted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeDeducted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeDeducted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.RecipientAddress) > 0 {
		i -= len(m.RecipientAddress)
		copy(dAtA[i:], m.RecipientAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RecipientAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RecipientChain) > 0 {
		i -= len(m.RecipientChain)
		copy(dAtA[i:], m.RecipientChain)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RecipientChain)))
		i--
		dAtA[i] = 0x12
	}
	if m.TransferID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.TransferID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *InsufficientFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InsufficientFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InsufficientFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.RecipientAddress) > 0 {
		i -= len(m.RecipientAddress)
		copy(dAtA[i:], m.RecipientAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RecipientAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RecipientChain) > 0 {
		i -= len(m.RecipientChain)
		copy(dAtA[i:], m.RecipientChain)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RecipientChain)))
		i--
		dAtA[i] = 0x12
	}
	if m.TransferID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.TransferID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FeeDeducted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TransferID != 0 {
		n += 1 + sovEvents(uint64(m.TransferID))
	}
	l = len(m.RecipientChain)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.RecipientAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovEvents(uint64(l))
	l = m.Fee.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *InsufficientFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TransferID != 0 {
		n += 1 + sovEvents(uint64(m.TransferID))
	}
	l = len(m.RecipientChain)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.RecipientAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovEvents(uint64(l))
	l = m.Fee.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FeeDeducted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: FeeDeducted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeDeducted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferID", wireType)
			}
			m.TransferID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransferID |= github_com_axelarnetwork_axelar_core_x_nexus_exported.TransferID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientChain = github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *InsufficientFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: InsufficientFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InsufficientFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferID", wireType)
			}
			m.TransferID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransferID |= github_com_axelarnetwork_axelar_core_x_nexus_exported.TransferID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientChain = github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
