// Code generated by protoc-gen-go.
// source: pp.account.proto
// DO NOT EDIT!

package pp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateAccountReq struct {
	Pubkey           *string `protobuf:"bytes,10,opt,name=pubkey" json:"pubkey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateAccountReq) Reset()                    { *m = CreateAccountReq{} }
func (m *CreateAccountReq) String() string            { return proto.CompactTextString(m) }
func (*CreateAccountReq) ProtoMessage()               {}
func (*CreateAccountReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *CreateAccountReq) GetPubkey() string {
	if m != nil && m.Pubkey != nil {
		return *m.Pubkey
	}
	return ""
}

type CreateAccountRes struct {
	Result           *Result `protobuf:"bytes,1,req,name=result" json:"result,omitempty"`
	AccountId        *string `protobuf:"bytes,10,opt,name=account_id" json:"account_id,omitempty"`
	CreatedAt        *int64  `protobuf:"varint,20,opt,name=created_at" json:"created_at,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateAccountRes) Reset()                    { *m = CreateAccountRes{} }
func (m *CreateAccountRes) String() string            { return proto.CompactTextString(m) }
func (*CreateAccountRes) ProtoMessage()               {}
func (*CreateAccountRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *CreateAccountRes) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *CreateAccountRes) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *CreateAccountRes) GetCreatedAt() int64 {
	if m != nil && m.CreatedAt != nil {
		return *m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateAccountReq)(nil), "pp.CreateAccountReq")
	proto.RegisterType((*CreateAccountRes)(nil), "pp.CreateAccountRes")
}

func init() { proto.RegisterFile("pp.account.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xd0, 0x4b,
	0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x28,
	0x90, 0xe2, 0x07, 0x8a, 0x26, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7, 0x41, 0x04, 0x95, 0x94, 0xb8, 0x04,
	0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x1d, 0x21, 0x6a, 0x83, 0x52, 0x0b, 0x85, 0xf8, 0xb8, 0xd8,
	0x0a, 0x4a, 0x93, 0xb2, 0x53, 0x2b, 0x25, 0xb8, 0x14, 0x18, 0x35, 0x38, 0x95, 0xc2, 0x30, 0xd4,
	0x14, 0x0b, 0x49, 0x71, 0xb1, 0x15, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x48, 0x30, 0x2a, 0x30, 0x69,
	0x70, 0x1b, 0x71, 0xe9, 0x01, 0x4d, 0x0e, 0x02, 0x8b, 0x08, 0x09, 0x71, 0x71, 0x41, 0x6d, 0x8e,
	0xcf, 0x4c, 0x81, 0x98, 0x01, 0x12, 0x4b, 0x06, 0x9b, 0x91, 0x12, 0x9f, 0x58, 0x22, 0x21, 0x02,
	0x14, 0x63, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x45, 0x13, 0xe2, 0xa3, 0x00, 0x00, 0x00,
}
