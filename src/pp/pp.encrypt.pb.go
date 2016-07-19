// Code generated by protoc-gen-go.
// source: pp.encrypt.proto
// DO NOT EDIT!

package pp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EncryptReq struct {
	Pubkey           []byte `protobuf:"bytes,1,opt,name=pubkey" json:"pubkey,omitempty"`
	Nonce            []byte `protobuf:"bytes,2,opt,name=nonce" json:"nonce,omitempty"`
	Encryptdata      []byte `protobuf:"bytes,3,opt,name=encryptdata" json:"encryptdata,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EncryptReq) Reset()                    { *m = EncryptReq{} }
func (m *EncryptReq) String() string            { return proto.CompactTextString(m) }
func (*EncryptReq) ProtoMessage()               {}
func (*EncryptReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *EncryptReq) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *EncryptReq) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *EncryptReq) GetEncryptdata() []byte {
	if m != nil {
		return m.Encryptdata
	}
	return nil
}

type EncryptRes struct {
	Result           *Result `protobuf:"bytes,1,req,name=result" json:"result,omitempty"`
	Nonce            []byte  `protobuf:"bytes,10,opt,name=nonce" json:"nonce,omitempty"`
	Encryptdata      []byte  `protobuf:"bytes,20,opt,name=encryptdata" json:"encryptdata,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EncryptRes) Reset()                    { *m = EncryptRes{} }
func (m *EncryptRes) String() string            { return proto.CompactTextString(m) }
func (*EncryptRes) ProtoMessage()               {}
func (*EncryptRes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *EncryptRes) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *EncryptRes) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *EncryptRes) GetEncryptdata() []byte {
	if m != nil {
		return m.Encryptdata
	}
	return nil
}

func init() {
	proto.RegisterType((*EncryptReq)(nil), "pp.EncryptReq")
	proto.RegisterType((*EncryptRes)(nil), "pp.EncryptRes")
}

func init() { proto.RegisterFile("pp.encrypt.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xd0, 0x4b,
	0xcd, 0x4b, 0x2e, 0xaa, 0x2c, 0x28, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x28,
	0x90, 0xe2, 0x07, 0x8a, 0x26, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7, 0x41, 0x04, 0x95, 0x1c, 0xb8, 0xb8,
	0x5c, 0x21, 0xaa, 0x82, 0x52, 0x0b, 0x85, 0xf8, 0xb8, 0xd8, 0x0a, 0x4a, 0x93, 0xb2, 0x53, 0x2b,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x84, 0x78, 0xb9, 0x58, 0xf3, 0xf2, 0xf3, 0x92, 0x53, 0x25,
	0x98, 0xc0, 0x5c, 0x61, 0x2e, 0x6e, 0xa8, 0x91, 0x29, 0x89, 0x25, 0x89, 0x12, 0xcc, 0x20, 0x41,
	0x25, 0x1f, 0x24, 0x13, 0x8a, 0x85, 0xa4, 0xb8, 0xd8, 0x8a, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x80,
	0x26, 0x30, 0x69, 0x70, 0x1b, 0x71, 0xe9, 0x01, 0x6d, 0x0c, 0x02, 0x8b, 0x20, 0x4c, 0xe3, 0xc2,
	0x66, 0x9a, 0x08, 0x48, 0x10, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x55, 0x53, 0xbf, 0xb7, 0x00,
	0x00, 0x00,
}
