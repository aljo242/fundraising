// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fundraising/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// GenesisState defines the fundraising module's genesis state.
type GenesisState struct {
	// params defines all the parameters for the module
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// auctions define the auction interface for genesis state; the module
	// supports FixedPriceAuction or BatchAuction
	Auctions []*types.Any `protobuf:"bytes,2,rep,name=auctions,proto3" json:"auctions,omitempty"`
	// allowed_bidder_records define the allowed bidder records for the auction
	AllowedBidderRecords []AllowedBidderRecord `protobuf:"bytes,3,rep,name=allowed_bidder_records,json=allowedBidderRecords,proto3" json:"allowed_bidder_records"`
	// bids define the bid records used for genesis state
	Bids []Bid `protobuf:"bytes,4,rep,name=bids,proto3" json:"bids"`
	// vesting_queues define the vesting queue records used for genesis
	// state
	VestingQueues []VestingQueue `protobuf:"bytes,5,rep,name=vesting_queues,json=vestingQueues,proto3" json:"vesting_queues"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a35424efc9855161, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

type AllowedBidderRecord struct {
	// auction_ id specifies index of the auction
	AuctionId uint64 `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	// allowed_bidder specifies allowed bidder for the auction
	AllowedBidder AllowedBidder `protobuf:"bytes,2,opt,name=allowed_bidder,json=allowedBidder,proto3" json:"allowed_bidder"`
}

func (m *AllowedBidderRecord) Reset()         { *m = AllowedBidderRecord{} }
func (m *AllowedBidderRecord) String() string { return proto.CompactTextString(m) }
func (*AllowedBidderRecord) ProtoMessage()    {}
func (*AllowedBidderRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_a35424efc9855161, []int{1}
}
func (m *AllowedBidderRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedBidderRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedBidderRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedBidderRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedBidderRecord.Merge(m, src)
}
func (m *AllowedBidderRecord) XXX_Size() int {
	return m.Size()
}
func (m *AllowedBidderRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedBidderRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedBidderRecord proto.InternalMessageInfo

func (m *AllowedBidderRecord) GetAuctionId() uint64 {
	if m != nil {
		return m.AuctionId
	}
	return 0
}

func (m *AllowedBidderRecord) GetAllowedBidder() AllowedBidder {
	if m != nil {
		return m.AllowedBidder
	}
	return AllowedBidder{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tendermint.fundraising.GenesisState")
	proto.RegisterType((*AllowedBidderRecord)(nil), "tendermint.fundraising.AllowedBidderRecord")
}

func init() { proto.RegisterFile("fundraising/genesis.proto", fileDescriptor_a35424efc9855161) }

var fileDescriptor_a35424efc9855161 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0x4e, 0xd6, 0x30, 0x15, 0x6f, 0xec, 0x60, 0xaa, 0x29, 0x1d, 0x5a, 0x3a, 0x4d, 0x20, 0x4d,
	0x42, 0x38, 0xd2, 0xd0, 0x2e, 0x08, 0x21, 0x35, 0x17, 0xb4, 0x13, 0x2c, 0x48, 0x1c, 0xb8, 0x44,
	0x4e, 0xec, 0x19, 0x4b, 0x8d, 0x5d, 0x62, 0xa7, 0xd0, 0x37, 0x28, 0x37, 0x1e, 0xa1, 0x0f, 0xc1,
	0x43, 0x54, 0x9c, 0x7a, 0xe4, 0x84, 0x50, 0x7b, 0xe1, 0x31, 0x50, 0x6d, 0x17, 0x52, 0xd1, 0x4a,
	0xdc, 0xfc, 0xff, 0xff, 0xf7, 0x7d, 0xff, 0x97, 0xef, 0x0f, 0xe8, 0xde, 0xd6, 0x82, 0x54, 0x98,
	0x2b, 0x2e, 0x58, 0xcc, 0xa8, 0xa0, 0x8a, 0x2b, 0x34, 0xac, 0xa4, 0x96, 0xf0, 0x58, 0x53, 0x41,
	0x68, 0x55, 0x72, 0xa1, 0x51, 0x03, 0x75, 0xd2, 0x2d, 0xa4, 0x2a, 0xa5, 0xca, 0x0c, 0x2a, 0xb6,
	0x85, 0xa5, 0x9c, 0x74, 0x98, 0x64, 0xd2, 0xf6, 0x57, 0x2f, 0xd7, 0xed, 0x32, 0x29, 0xd9, 0x80,
	0xc6, 0xa6, 0xca, 0xeb, 0xdb, 0x18, 0x8b, 0xb1, 0x1b, 0x9d, 0x36, 0xd7, 0x37, 0xde, 0x6e, 0x1c,
	0x36, 0xc7, 0x43, 0x5c, 0xe1, 0xd2, 0x6d, 0x3a, 0xff, 0xdc, 0x02, 0x87, 0x2f, 0xad, 0xdd, 0x37,
	0x1a, 0x6b, 0x0a, 0x9f, 0x83, 0x7d, 0x0b, 0x08, 0xfd, 0x33, 0xff, 0xe2, 0xe0, 0x32, 0x42, 0xdb,
	0xed, 0xa3, 0xd7, 0x06, 0x95, 0x04, 0xb3, 0x1f, 0x3d, 0x2f, 0x75, 0x1c, 0xf8, 0x02, 0xb4, 0x71,
	0x5d, 0x68, 0x2e, 0x85, 0x0a, 0xf7, 0xce, 0x5a, 0x17, 0x07, 0x97, 0x1d, 0x64, 0x5d, 0xa3, 0xb5,
	0x6b, 0xd4, 0x17, 0xe3, 0xe4, 0xf0, 0xdb, 0xd7, 0x27, 0xed, 0xbe, 0x45, 0x5e, 0xa7, 0x7f, 0x38,
	0x90, 0x81, 0x63, 0x3c, 0x18, 0xc8, 0x8f, 0x94, 0x64, 0x39, 0x27, 0x84, 0x56, 0x59, 0x45, 0x0b,
	0x59, 0x11, 0x15, 0xb6, 0x8c, 0xda, 0xe3, 0x5d, 0x6e, 0xfa, 0x96, 0x95, 0x18, 0x52, 0x6a, 0x38,
	0xce, 0x5a, 0x07, 0xff, 0x3b, 0x52, 0xf0, 0x0a, 0x04, 0x39, 0x27, 0x2a, 0x0c, 0x8c, 0xec, 0x83,
	0x5d, 0xb2, 0x09, 0x5f, 0xcb, 0x18, 0x38, 0xbc, 0x01, 0x47, 0x23, 0xaa, 0x34, 0x17, 0x2c, 0xfb,
	0x50, 0xd3, 0x9a, 0xaa, 0xf0, 0x8e, 0x11, 0x78, 0xb8, 0x4b, 0xe0, 0xad, 0x45, 0xdf, 0xac, 0xc0,
	0x4e, 0xe9, 0xde, 0xa8, 0xd1, 0x53, 0xcf, 0xda, 0x93, 0x69, 0xcf, 0xfb, 0x35, 0xed, 0x79, 0xe7,
	0x13, 0x1f, 0xdc, 0xdf, 0xf2, 0x1d, 0xf0, 0x14, 0x00, 0x17, 0x50, 0xc6, 0x89, 0x39, 0x4b, 0x90,
	0xde, 0x75, 0x9d, 0x6b, 0x02, 0x53, 0x70, 0xb4, 0x99, 0x59, 0xb8, 0x67, 0x2e, 0xf7, 0xe8, 0xbf,
	0xb2, 0x5a, 0x9b, 0xda, 0x48, 0x29, 0x79, 0x35, 0x5b, 0x44, 0xfe, 0x7c, 0x11, 0xf9, 0x3f, 0x17,
	0x91, 0xff, 0x65, 0x19, 0x79, 0xf3, 0x65, 0xe4, 0x7d, 0x5f, 0x46, 0xde, 0xbb, 0x2b, 0xc6, 0xf5,
	0xfb, 0x3a, 0x47, 0x85, 0x2c, 0xe3, 0xbf, 0xfa, 0xcd, 0x7f, 0x2e, 0xfe, 0xb4, 0x51, 0xe9, 0xf1,
	0x90, 0xaa, 0x7c, 0xdf, 0x9c, 0xff, 0xe9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xa2, 0xf3,
	0x06, 0x28, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingQueues) > 0 {
		for iNdEx := len(m.VestingQueues) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingQueues[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Bids) > 0 {
		for iNdEx := len(m.Bids) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Bids[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AllowedBidderRecords) > 0 {
		for iNdEx := len(m.AllowedBidderRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AllowedBidderRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Auctions) > 0 {
		for iNdEx := len(m.Auctions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Auctions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *AllowedBidderRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedBidderRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedBidderRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AllowedBidder.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AuctionId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Auctions) > 0 {
		for _, e := range m.Auctions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AllowedBidderRecords) > 0 {
		for _, e := range m.AllowedBidderRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Bids) > 0 {
		for _, e := range m.Bids {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.VestingQueues) > 0 {
		for _, e := range m.VestingQueues {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *AllowedBidderRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionId != 0 {
		n += 1 + sovGenesis(uint64(m.AuctionId))
	}
	l = m.AllowedBidder.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Auctions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Auctions = append(m.Auctions, &types.Any{})
			if err := m.Auctions[len(m.Auctions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedBidderRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedBidderRecords = append(m.AllowedBidderRecords, AllowedBidderRecord{})
			if err := m.AllowedBidderRecords[len(m.AllowedBidderRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bids", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bids = append(m.Bids, Bid{})
			if err := m.Bids[len(m.Bids)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingQueues", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingQueues = append(m.VestingQueues, VestingQueue{})
			if err := m.VestingQueues[len(m.VestingQueues)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *AllowedBidderRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: AllowedBidderRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedBidderRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedBidder", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AllowedBidder.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
