// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/exchange/v1beta1/exchange.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Params struct {
	MarketCreationFee   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=market_creation_fee,json=marketCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"market_creation_fee"`
	DefaultMakerFeeRate github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,2,opt,name=default_maker_fee_rate,json=defaultMakerFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"default_maker_fee_rate"`
	DefaultTakerFeeRate github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,3,opt,name=default_taker_fee_rate,json=defaultTakerFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"default_taker_fee_rate"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb2114aee993f375, []int{0}
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

type Market struct {
	Id            uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BaseDenom     string                                 `protobuf:"bytes,2,opt,name=base_denom,json=baseDenom,proto3" json:"base_denom,omitempty"`
	QuoteDenom    string                                 `protobuf:"bytes,3,opt,name=quote_denom,json=quoteDenom,proto3" json:"quote_denom,omitempty"`
	EscrowAddress string                                 `protobuf:"bytes,4,opt,name=escrow_address,json=escrowAddress,proto3" json:"escrow_address,omitempty"`
	MakerFeeRate  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=maker_fee_rate,json=makerFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"maker_fee_rate"`
	TakerFeeRate  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=taker_fee_rate,json=takerFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"taker_fee_rate"`
}

func (m *Market) Reset()         { *m = Market{} }
func (m *Market) String() string { return proto.CompactTextString(m) }
func (*Market) ProtoMessage()    {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb2114aee993f375, []int{1}
}
func (m *Market) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Market.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return m.Size()
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

type MarketState struct {
	LastPrice *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=last_price,json=lastPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"last_price,omitempty"`
}

func (m *MarketState) Reset()         { *m = MarketState{} }
func (m *MarketState) String() string { return proto.CompactTextString(m) }
func (*MarketState) ProtoMessage()    {}
func (*MarketState) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb2114aee993f375, []int{2}
}
func (m *MarketState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketState.Merge(m, src)
}
func (m *MarketState) XXX_Size() int {
	return m.Size()
}
func (m *MarketState) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketState.DiscardUnknown(m)
}

var xxx_messageInfo_MarketState proto.InternalMessageInfo

type Order struct {
	Id               uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Orderer          string                                 `protobuf:"bytes,2,opt,name=orderer,proto3" json:"orderer,omitempty"`
	MarketId         uint64                                 `protobuf:"varint,3,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	IsBuy            bool                                   `protobuf:"varint,4,opt,name=is_buy,json=isBuy,proto3" json:"is_buy,omitempty"`
	Price            github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
	Quantity         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=quantity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"quantity"`
	MsgHeight        int64                                  `protobuf:"varint,7,opt,name=msg_height,json=msgHeight,proto3" json:"msg_height,omitempty"`
	OpenQuantity     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=open_quantity,json=openQuantity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"open_quantity"`
	RemainingDeposit github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=remaining_deposit,json=remainingDeposit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"remaining_deposit"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb2114aee993f375, []int{3}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "crescent.exchange.v1beta1.Params")
	proto.RegisterType((*Market)(nil), "crescent.exchange.v1beta1.Market")
	proto.RegisterType((*MarketState)(nil), "crescent.exchange.v1beta1.MarketState")
	proto.RegisterType((*Order)(nil), "crescent.exchange.v1beta1.Order")
}

func init() {
	proto.RegisterFile("crescent/exchange/v1beta1/exchange.proto", fileDescriptor_bb2114aee993f375)
}

var fileDescriptor_bb2114aee993f375 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x9b, 0x76, 0xed, 0x1a, 0x6f, 0xab, 0x58, 0x06, 0x28, 0x1b, 0x22, 0xad, 0x26, 0x81,
	0x2a, 0xa4, 0x25, 0x0c, 0x84, 0xc4, 0x2d, 0x5d, 0x35, 0x31, 0xa4, 0x89, 0x91, 0x4d, 0x08, 0xc1,
	0x45, 0xe4, 0x26, 0x67, 0xa9, 0xd5, 0x25, 0xee, 0x6c, 0x77, 0x5b, 0xc5, 0x4b, 0xf0, 0x1c, 0xf0,
	0x22, 0xbd, 0x63, 0x97, 0x88, 0x8b, 0x01, 0xeb, 0x8b, 0x20, 0xdb, 0x69, 0xd7, 0x8d, 0x1b, 0x08,
	0x57, 0x6d, 0xfe, 0x73, 0xce, 0xff, 0xd9, 0xbf, 0x13, 0xa3, 0x66, 0xc8, 0x80, 0x87, 0x90, 0x0a,
	0x0f, 0xce, 0xc2, 0x2e, 0x4e, 0x63, 0xf0, 0x4e, 0x36, 0x3b, 0x20, 0xf0, 0xe6, 0x54, 0x70, 0xfb,
	0x8c, 0x0a, 0x6a, 0xad, 0x4e, 0x3a, 0xdd, 0x69, 0x21, 0xeb, 0x5c, 0xbb, 0x1d, 0xd3, 0x98, 0xaa,
	0x2e, 0x4f, 0xfe, 0xd3, 0x03, 0x6b, 0x4e, 0x48, 0x79, 0x42, 0xb9, 0xd7, 0xc1, 0xfc, 0xca, 0x34,
	0xa4, 0x24, 0xd5, 0xf5, 0xf5, 0xaf, 0x45, 0x54, 0xd9, 0xc3, 0x0c, 0x27, 0xdc, 0xfa, 0x88, 0x56,
	0x12, 0xcc, 0x7a, 0x20, 0x82, 0x90, 0x01, 0x16, 0x84, 0xa6, 0xc1, 0x21, 0x80, 0x6d, 0x34, 0x4a,
	0xcd, 0x85, 0x27, 0xab, 0xae, 0x36, 0x72, 0xa5, 0xd1, 0x84, 0xe9, 0x6e, 0x51, 0x92, 0xb6, 0x1e,
	0x8f, 0x2e, 0xea, 0x85, 0xcf, 0x3f, 0xea, 0xcd, 0x98, 0x88, 0xee, 0xa0, 0xe3, 0x86, 0x34, 0xf1,
	0x32, 0xaa, 0xfe, 0xd9, 0xe0, 0x51, 0xcf, 0x13, 0xc3, 0x3e, 0x70, 0x35, 0xc0, 0xfd, 0x65, 0xcd,
	0xd9, 0xca, 0x30, 0xdb, 0x00, 0x56, 0x88, 0xee, 0x46, 0x70, 0x88, 0x07, 0x47, 0x22, 0x48, 0x70,
	0x0f, 0x98, 0x44, 0x07, 0x0c, 0x0b, 0xb0, 0x8b, 0x0d, 0xa3, 0x69, 0xb6, 0x5c, 0x09, 0xf9, 0x7e,
	0x51, 0x7f, 0xf8, 0x17, 0x90, 0x36, 0x84, 0xfe, 0x4a, 0xe6, 0xb6, 0x2b, 0xcd, 0xb6, 0x01, 0x7c,
	0x2c, 0xae, 0x41, 0xc4, 0x75, 0x48, 0xe9, 0xbf, 0x20, 0x07, 0x33, 0x90, 0xf5, 0x2f, 0x45, 0x54,
	0xd9, 0x55, 0xfb, 0xb3, 0x6a, 0xa8, 0x48, 0x22, 0xdb, 0x68, 0x18, 0xcd, 0x39, 0xbf, 0x48, 0x22,
	0xeb, 0x3e, 0x42, 0x32, 0xbe, 0x20, 0x82, 0x94, 0x26, 0x7a, 0x63, 0xbe, 0x29, 0x95, 0xb6, 0x14,
	0xac, 0x3a, 0x5a, 0x38, 0x1e, 0x50, 0x31, 0xa9, 0xab, 0x35, 0xf9, 0x48, 0x49, 0xba, 0xe1, 0x01,
	0xaa, 0x01, 0x0f, 0x19, 0x3d, 0x0d, 0x70, 0x14, 0x31, 0xe0, 0xdc, 0x9e, 0x53, 0x3d, 0x4b, 0x5a,
	0x7d, 0xa1, 0x45, 0xeb, 0x00, 0xd5, 0x6e, 0x64, 0x58, 0xce, 0xb5, 0xbd, 0xc5, 0x64, 0x36, 0xbc,
	0x03, 0x54, 0xbb, 0x11, 0x5a, 0x25, 0x9f, 0xab, 0x98, 0x4d, 0xeb, 0x1d, 0x5a, 0xd0, 0x61, 0xed,
	0x0b, 0x09, 0xd9, 0x41, 0xe8, 0x08, 0x73, 0x11, 0xf4, 0x19, 0x09, 0x41, 0x25, 0x67, 0xb6, 0x1e,
	0xfd, 0x83, 0xb9, 0x29, 0xa7, 0xf7, 0xe4, 0xf0, 0xfa, 0xa8, 0x84, 0xca, 0xaf, 0x59, 0x04, 0xec,
	0x8f, 0x63, 0xb0, 0xd1, 0x3c, 0x95, 0x05, 0x60, 0xd9, 0x19, 0x4c, 0x1e, 0xad, 0x7b, 0xc8, 0xcc,
	0x3e, 0x01, 0x12, 0xa9, 0xfc, 0xe7, 0xfc, 0xaa, 0x16, 0x76, 0x22, 0xeb, 0x0e, 0xaa, 0x10, 0x1e,
	0x74, 0x06, 0x43, 0x95, 0x7a, 0xd5, 0x2f, 0x13, 0xde, 0x1a, 0x0c, 0xad, 0x36, 0x2a, 0xeb, 0xd5,
	0xe6, 0x0b, 0x59, 0x0f, 0x5b, 0xaf, 0x50, 0xf5, 0x78, 0x80, 0x53, 0x41, 0xc4, 0x30, 0x47, 0xae,
	0x3b, 0xa9, 0xf0, 0xa7, 0xf3, 0xf2, 0x35, 0x4b, 0x78, 0x1c, 0x74, 0x81, 0xc4, 0x5d, 0x61, 0xcf,
	0x37, 0x8c, 0x66, 0xc9, 0x37, 0x13, 0x1e, 0xbf, 0x54, 0x82, 0xb5, 0x8f, 0x96, 0x68, 0x1f, 0xd2,
	0x60, 0xca, 0xab, 0xe6, 0xe2, 0x2d, 0x4a, 0x93, 0x37, 0x13, 0xe6, 0x07, 0xb4, 0xcc, 0x20, 0xc1,
	0x24, 0x25, 0x69, 0x1c, 0x44, 0xd0, 0xa7, 0x9c, 0x08, 0xdb, 0xcc, 0x65, 0x7c, 0x6b, 0x6a, 0xd4,
	0xd6, 0x3e, 0xad, 0xb7, 0xa3, 0x5f, 0x4e, 0x61, 0x74, 0xe9, 0x18, 0xe7, 0x97, 0x8e, 0xf1, 0xf3,
	0xd2, 0x31, 0x3e, 0x8d, 0x9d, 0xc2, 0xf9, 0xd8, 0x29, 0x7c, 0x1b, 0x3b, 0x85, 0xf7, 0xcf, 0x67,
	0x7d, 0xb3, 0xeb, 0x71, 0x23, 0x05, 0x71, 0x4a, 0x59, 0x6f, 0x2a, 0x78, 0x27, 0xcf, 0xbc, 0xb3,
	0xab, 0xeb, 0x55, 0xd1, 0x3a, 0x15, 0x75, 0x07, 0x3e, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x26,
	0x3c, 0x31, 0x4b, 0x80, 0x05, 0x00, 0x00,
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
		size := m.DefaultTakerFeeRate.Size()
		i -= size
		if _, err := m.DefaultTakerFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExchange(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.DefaultMakerFeeRate.Size()
		i -= size
		if _, err := m.DefaultMakerFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExchange(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MarketCreationFee) > 0 {
		for iNdEx := len(m.MarketCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MarketCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExchange(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Market) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Market) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Market) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TakerFeeRate.Size()
		i -= size
		if _, err := m.TakerFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExchange(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.MakerFeeRate.Size()
		i -= size
		if _, err := m.MakerFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExchange(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.EscrowAddress) > 0 {
		i -= len(m.EscrowAddress)
		copy(dAtA[i:], m.EscrowAddress)
		i = encodeVarintExchange(dAtA, i, uint64(len(m.EscrowAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.QuoteDenom) > 0 {
		i -= len(m.QuoteDenom)
		copy(dAtA[i:], m.QuoteDenom)
		i = encodeVarintExchange(dAtA, i, uint64(len(m.QuoteDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintExchange(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintExchange(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MarketState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastPrice != nil {
		{
			size := m.LastPrice.Size()
			i -= size
			if _, err := m.LastPrice.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintExchange(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RemainingDeposit.Size()
		i -= size
		if _, err := m.RemainingDeposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExchange(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.OpenQuantity.Size()
		i -= size
		if _, err := m.OpenQuantity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExchange(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.MsgHeight != 0 {
		i = encodeVarintExchange(dAtA, i, uint64(m.MsgHeight))
		i--
		dAtA[i] = 0x38
	}
	{
		size := m.Quantity.Size()
		i -= size
		if _, err := m.Quantity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExchange(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExchange(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.IsBuy {
		i--
		if m.IsBuy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.MarketId != 0 {
		i = encodeVarintExchange(dAtA, i, uint64(m.MarketId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Orderer) > 0 {
		i -= len(m.Orderer)
		copy(dAtA[i:], m.Orderer)
		i = encodeVarintExchange(dAtA, i, uint64(len(m.Orderer)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintExchange(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintExchange(dAtA []byte, offset int, v uint64) int {
	offset -= sovExchange(v)
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
	if len(m.MarketCreationFee) > 0 {
		for _, e := range m.MarketCreationFee {
			l = e.Size()
			n += 1 + l + sovExchange(uint64(l))
		}
	}
	l = m.DefaultMakerFeeRate.Size()
	n += 1 + l + sovExchange(uint64(l))
	l = m.DefaultTakerFeeRate.Size()
	n += 1 + l + sovExchange(uint64(l))
	return n
}

func (m *Market) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovExchange(uint64(m.Id))
	}
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovExchange(uint64(l))
	}
	l = len(m.QuoteDenom)
	if l > 0 {
		n += 1 + l + sovExchange(uint64(l))
	}
	l = len(m.EscrowAddress)
	if l > 0 {
		n += 1 + l + sovExchange(uint64(l))
	}
	l = m.MakerFeeRate.Size()
	n += 1 + l + sovExchange(uint64(l))
	l = m.TakerFeeRate.Size()
	n += 1 + l + sovExchange(uint64(l))
	return n
}

func (m *MarketState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LastPrice != nil {
		l = m.LastPrice.Size()
		n += 1 + l + sovExchange(uint64(l))
	}
	return n
}

func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovExchange(uint64(m.Id))
	}
	l = len(m.Orderer)
	if l > 0 {
		n += 1 + l + sovExchange(uint64(l))
	}
	if m.MarketId != 0 {
		n += 1 + sovExchange(uint64(m.MarketId))
	}
	if m.IsBuy {
		n += 2
	}
	l = m.Price.Size()
	n += 1 + l + sovExchange(uint64(l))
	l = m.Quantity.Size()
	n += 1 + l + sovExchange(uint64(l))
	if m.MsgHeight != 0 {
		n += 1 + sovExchange(uint64(m.MsgHeight))
	}
	l = m.OpenQuantity.Size()
	n += 1 + l + sovExchange(uint64(l))
	l = m.RemainingDeposit.Size()
	n += 1 + l + sovExchange(uint64(l))
	return n
}

func sovExchange(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExchange(x uint64) (n int) {
	return sovExchange(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExchange
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
				return fmt.Errorf("proto: wrong wireType = %d for field MarketCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketCreationFee = append(m.MarketCreationFee, types.Coin{})
			if err := m.MarketCreationFee[len(m.MarketCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultMakerFeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DefaultMakerFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultTakerFeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DefaultTakerFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExchange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExchange
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
func (m *Market) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExchange
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
			return fmt.Errorf("proto: Market: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Market: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuoteDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EscrowAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerFeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MakerFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerFeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TakerFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExchange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExchange
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
func (m *MarketState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExchange
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
			return fmt.Errorf("proto: MarketState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.LastPrice = &v
			if err := m.LastPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExchange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExchange
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
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExchange
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
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orderer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Orderer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			m.MarketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MarketId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsBuy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsBuy = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgHeight", wireType)
			}
			m.MsgHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenQuantity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OpenQuantity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingDeposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchange
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
				return ErrInvalidLengthExchange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemainingDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExchange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExchange
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
func skipExchange(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExchange
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
					return 0, ErrIntOverflowExchange
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
					return 0, ErrIntOverflowExchange
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
				return 0, ErrInvalidLengthExchange
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExchange
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExchange
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExchange        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExchange          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExchange = fmt.Errorf("proto: unexpected end of group")
)
