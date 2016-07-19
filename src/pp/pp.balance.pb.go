// Code generated by protoc-gen-go.
// source: pp.balance.proto
// DO NOT EDIT!

package pp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetBalanceReq struct {
	AccountId        []byte  `protobuf:"bytes,1,opt,name=account_id" json:"account_id,omitempty"`
	CoinType         *string `protobuf:"bytes,2,opt,name=coin_type" json:"coin_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetBalanceReq) Reset()                    { *m = GetBalanceReq{} }
func (m *GetBalanceReq) String() string            { return proto.CompactTextString(m) }
func (*GetBalanceReq) ProtoMessage()               {}
func (*GetBalanceReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *GetBalanceReq) GetAccountId() []byte {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *GetBalanceReq) GetCoinType() string {
	if m != nil && m.CoinType != nil {
		return *m.CoinType
	}
	return ""
}

type GetBalanceRes struct {
	AccountId        []byte  `protobuf:"bytes,10,opt,name=account_id" json:"account_id,omitempty"`
	CoinType         *string `protobuf:"bytes,11,opt,name=coin_type" json:"coin_type,omitempty"`
	Balance          *uint64 `protobuf:"varint,12,opt,name=balance" json:"balance,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetBalanceRes) Reset()                    { *m = GetBalanceRes{} }
func (m *GetBalanceRes) String() string            { return proto.CompactTextString(m) }
func (*GetBalanceRes) ProtoMessage()               {}
func (*GetBalanceRes) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *GetBalanceRes) GetAccountId() []byte {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *GetBalanceRes) GetCoinType() string {
	if m != nil && m.CoinType != nil {
		return *m.CoinType
	}
	return ""
}

func (m *GetBalanceRes) GetBalance() uint64 {
	if m != nil && m.Balance != nil {
		return *m.Balance
	}
	return 0
}

func init() {
	proto.RegisterType((*GetBalanceReq)(nil), "pp.GetBalanceReq")
	proto.RegisterType((*GetBalanceRes)(nil), "pp.GetBalanceRes")
}

func init() { proto.RegisterFile("pp.balance.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xd0, 0x4b,
	0x4a, 0xcc, 0x49, 0xcc, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x28,
	0x50, 0x32, 0xe3, 0xe2, 0x75, 0x4f, 0x2d, 0x71, 0x82, 0x88, 0x07, 0xa5, 0x16, 0x0a, 0x09, 0x71,
	0x71, 0x25, 0x26, 0x27, 0xe7, 0x97, 0xe6, 0x95, 0xc4, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0xf0, 0x08, 0x09, 0x72, 0x71, 0x26, 0xe7, 0x67, 0xe6, 0xc5, 0x97, 0x54, 0x16, 0xa4, 0x4a, 0x30,
	0x01, 0x85, 0x38, 0x95, 0xdc, 0x51, 0xf5, 0x15, 0xa3, 0xe9, 0xe3, 0xc2, 0xd4, 0xc7, 0x0d, 0xd2,
	0x27, 0xc4, 0xcf, 0xc5, 0x0e, 0x75, 0x84, 0x04, 0x0f, 0x50, 0x80, 0x05, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xcb, 0xae, 0x56, 0x32, 0x97, 0x00, 0x00, 0x00,
}
