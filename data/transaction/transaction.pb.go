// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transaction.proto

package transaction

import (
	bytes "bytes"
	fmt "fmt"
	github_com_Dharitri_org_sk_core_data "github.com/Dharitri-org/sk-core/data"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// Transaction holds all the data needed for a value transfer or SC call
type Transaction struct {
	Nonce       uint64        `protobuf:"varint,1,opt,name=Nonce,proto3" json:"nonce"`
	Value       *math_big.Int `protobuf:"bytes,2,opt,name=Value,proto3,casttypewith=math/big.Int;github.com/Dharitri-org/sk-core/data.BigIntCaster" json:"value"`
	RcvAddr     []byte        `protobuf:"bytes,3,opt,name=RcvAddr,proto3" json:"receiver"`
	RcvUserName []byte        `protobuf:"bytes,4,opt,name=RcvUserName,proto3" json:"rcvUserName,omitempty"`
	SndAddr     []byte        `protobuf:"bytes,5,opt,name=SndAddr,proto3" json:"sender"`
	SndUserName []byte        `protobuf:"bytes,6,opt,name=SndUserName,proto3" json:"sndUserName,omitempty"`
	GasPrice    uint64        `protobuf:"varint,7,opt,name=GasPrice,proto3" json:"gasPrice,omitempty"`
	GasLimit    uint64        `protobuf:"varint,8,opt,name=GasLimit,proto3" json:"gasLimit,omitempty"`
	Data        []byte        `protobuf:"bytes,9,opt,name=Data,proto3" json:"data,omitempty"`
	ChainID     []byte        `protobuf:"bytes,10,opt,name=ChainID,proto3" json:"chainID"`
	Version     uint32        `protobuf:"varint,11,opt,name=Version,proto3" json:"version"`
	Signature   []byte        `protobuf:"bytes,12,opt,name=Signature,proto3" json:"signature,omitempty"`
	Options     uint32        `protobuf:"varint,13,opt,name=Options,proto3" json:"options,omitempty"`
}

func (m *Transaction) Reset()      { *m = Transaction{} }
func (*Transaction) ProtoMessage() {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return m.Size()
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetValue() *math_big.Int {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Transaction) GetRcvAddr() []byte {
	if m != nil {
		return m.RcvAddr
	}
	return nil
}

func (m *Transaction) GetRcvUserName() []byte {
	if m != nil {
		return m.RcvUserName
	}
	return nil
}

func (m *Transaction) GetSndAddr() []byte {
	if m != nil {
		return m.SndAddr
	}
	return nil
}

func (m *Transaction) GetSndUserName() []byte {
	if m != nil {
		return m.SndUserName
	}
	return nil
}

func (m *Transaction) GetGasPrice() uint64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *Transaction) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *Transaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Transaction) GetChainID() []byte {
	if m != nil {
		return m.ChainID
	}
	return nil
}

func (m *Transaction) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Transaction) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Transaction) GetOptions() uint32 {
	if m != nil {
		return m.Options
	}
	return 0
}

func init() {
	proto.RegisterType((*Transaction)(nil), "proto.Transaction")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490) }

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xb1, 0x8e, 0xd3, 0x30,
	0x1c, 0xc6, 0x63, 0xb8, 0x5c, 0xae, 0x6e, 0x0f, 0xe9, 0x8c, 0x0e, 0x0c, 0x83, 0x5d, 0x21, 0x38,
	0x75, 0xa0, 0x8d, 0x04, 0x62, 0x3a, 0x09, 0x89, 0x5e, 0x11, 0xaa, 0x40, 0x07, 0x4a, 0xe1, 0x06,
	0x36, 0x37, 0x31, 0xa9, 0x05, 0xb1, 0x2b, 0xc7, 0xad, 0xc4, 0xc6, 0x23, 0xf0, 0x06, 0xac, 0x88,
	0x27, 0x61, 0xec, 0xd8, 0x29, 0xd0, 0x74, 0x41, 0x99, 0xee, 0x11, 0x50, 0x9c, 0xf6, 0x1a, 0x98,
	0x12, 0xff, 0xbe, 0xef, 0xfb, 0x7f, 0x96, 0x65, 0xc3, 0x23, 0xa3, 0x99, 0x4c, 0x59, 0x68, 0x84,
	0x92, 0xbd, 0xa9, 0x56, 0x46, 0x21, 0xd7, 0x7e, 0xee, 0x76, 0x63, 0x61, 0x26, 0xb3, 0x71, 0x2f,
	0x54, 0x89, 0x1f, 0xab, 0x58, 0xf9, 0x16, 0x8f, 0x67, 0x1f, 0xec, 0xca, 0x2e, 0xec, 0x5f, 0x95,
	0xba, 0xf7, 0xcd, 0x85, 0xcd, 0xb7, 0xbb, 0x59, 0x88, 0x42, 0xf7, 0x5c, 0xc9, 0x90, 0x63, 0xd0,
	0x06, 0x9d, 0xbd, 0x7e, 0xa3, 0xc8, 0xa8, 0x2b, 0x4b, 0x10, 0x54, 0x1c, 0x31, 0xe8, 0x5e, 0xb0,
	0x4f, 0x33, 0x8e, 0xaf, 0xb5, 0x41, 0xa7, 0xd5, 0x7f, 0x59, 0x1a, 0xe6, 0x25, 0xf8, 0xf1, 0x8b,
	0x3e, 0x4d, 0x98, 0x99, 0xf8, 0x63, 0x11, 0xf7, 0x86, 0xd2, 0x9c, 0xd6, 0x36, 0x32, 0x98, 0x30,
	0x2d, 0x8c, 0x16, 0x5d, 0xa5, 0x63, 0x3f, 0xfd, 0xd8, 0x0d, 0x95, 0xe6, 0x7e, 0xc4, 0x0c, 0xeb,
	0xf5, 0x45, 0x3c, 0x94, 0xe6, 0x8c, 0xa5, 0x86, 0xeb, 0xa0, 0x9a, 0x8c, 0x4e, 0xa0, 0x17, 0x84,
	0xf3, 0x67, 0x51, 0xa4, 0xf1, 0x75, 0x5b, 0xd2, 0x2a, 0x32, 0x7a, 0xa0, 0x79, 0xc8, 0xc5, 0x9c,
	0xeb, 0x60, 0x2b, 0xa2, 0x53, 0xd8, 0x0c, 0xc2, 0xf9, 0xbb, 0x94, 0xeb, 0x73, 0x96, 0x70, 0xbc,
	0x67, 0xbd, 0x77, 0x8a, 0x8c, 0x1e, 0xeb, 0x1d, 0x7e, 0xa8, 0x12, 0x61, 0x78, 0x32, 0x35, 0x9f,
	0x83, 0xba, 0x1b, 0xdd, 0x87, 0xde, 0x48, 0x46, 0xb6, 0xc4, 0xb5, 0x41, 0x58, 0x64, 0x74, 0x3f,
	0xe5, 0x32, 0x2a, 0x2b, 0x36, 0x52, 0x59, 0x31, 0x92, 0xd1, 0x55, 0xc5, 0xfe, 0xae, 0x22, 0xdd,
	0xe1, 0x7a, 0x45, 0xcd, 0x8d, 0x1e, 0xc1, 0x83, 0x17, 0x2c, 0x7d, 0xa3, 0x45, 0xc8, 0xb1, 0x67,
	0x8f, 0xf3, 0x56, 0x91, 0x51, 0x14, 0x6f, 0x58, 0x2d, 0x76, 0xe5, 0xdb, 0x64, 0x5e, 0x89, 0x44,
	0x18, 0x7c, 0xf0, 0x4f, 0xc6, 0xb2, 0xff, 0x32, 0x96, 0xa1, 0x13, 0xb8, 0x37, 0x60, 0x86, 0xe1,
	0x86, 0xdd, 0x1d, 0x2a, 0x32, 0x7a, 0xa3, 0x3c, 0xdb, 0x9a, 0xd7, 0xea, 0xe8, 0x01, 0xf4, 0xce,
	0x26, 0x4c, 0xc8, 0xe1, 0x00, 0x43, 0x6b, 0x6d, 0x16, 0x19, 0xf5, 0xc2, 0x0a, 0x05, 0x5b, 0xad,
	0xb4, 0x5d, 0x70, 0x9d, 0x0a, 0x25, 0x71, 0xb3, 0x0d, 0x3a, 0x87, 0x95, 0x6d, 0x5e, 0xa1, 0x60,
	0xab, 0xa1, 0x27, 0xb0, 0x31, 0x12, 0xb1, 0x64, 0x66, 0xa6, 0x39, 0x6e, 0xd9, 0x79, 0xb7, 0x8b,
	0x8c, 0xde, 0x4c, 0xb7, 0xb0, 0xd6, 0xbf, 0x73, 0x22, 0x1f, 0x7a, 0xaf, 0xa7, 0xe5, 0x55, 0x4b,
	0xf1, 0xa1, 0x9d, 0x7e, 0x5c, 0x64, 0xf4, 0x48, 0x55, 0xa8, 0x16, 0xd9, 0xba, 0xfa, 0xcf, 0x17,
	0x2b, 0xe2, 0x2c, 0x57, 0xc4, 0xb9, 0x5c, 0x11, 0xf0, 0x25, 0x27, 0xe0, 0x7b, 0x4e, 0xc0, 0xcf,
	0x9c, 0x80, 0x45, 0x4e, 0xc0, 0x32, 0x27, 0xe0, 0x77, 0x4e, 0xc0, 0x9f, 0x9c, 0x38, 0x97, 0x39,
	0x01, 0x5f, 0xd7, 0xc4, 0x59, 0xac, 0x89, 0xb3, 0x5c, 0x13, 0xe7, 0x7d, 0xb3, 0xf6, 0x48, 0xc6,
	0xfb, 0xf6, 0xbe, 0x3f, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x31, 0x52, 0x34, 0x71, 0x3a, 0x03,
	0x00, 0x00,
}

func (this *Transaction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Transaction)
	if !ok {
		that2, ok := that.(Transaction)
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
	if this.Nonce != that1.Nonce {
		return false
	}
	{
		__caster := &github_com_Dharitri_org_sk_core_data.BigIntCaster{}
		if !__caster.Equal(this.Value, that1.Value) {
			return false
		}
	}
	if !bytes.Equal(this.RcvAddr, that1.RcvAddr) {
		return false
	}
	if !bytes.Equal(this.RcvUserName, that1.RcvUserName) {
		return false
	}
	if !bytes.Equal(this.SndAddr, that1.SndAddr) {
		return false
	}
	if !bytes.Equal(this.SndUserName, that1.SndUserName) {
		return false
	}
	if this.GasPrice != that1.GasPrice {
		return false
	}
	if this.GasLimit != that1.GasLimit {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if !bytes.Equal(this.ChainID, that1.ChainID) {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if !bytes.Equal(this.Signature, that1.Signature) {
		return false
	}
	if this.Options != that1.Options {
		return false
	}
	return true
}
func (this *Transaction) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 17)
	s = append(s, "&transaction.Transaction{")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "RcvAddr: "+fmt.Sprintf("%#v", this.RcvAddr)+",\n")
	s = append(s, "RcvUserName: "+fmt.Sprintf("%#v", this.RcvUserName)+",\n")
	s = append(s, "SndAddr: "+fmt.Sprintf("%#v", this.SndAddr)+",\n")
	s = append(s, "SndUserName: "+fmt.Sprintf("%#v", this.SndUserName)+",\n")
	s = append(s, "GasPrice: "+fmt.Sprintf("%#v", this.GasPrice)+",\n")
	s = append(s, "GasLimit: "+fmt.Sprintf("%#v", this.GasLimit)+",\n")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "ChainID: "+fmt.Sprintf("%#v", this.ChainID)+",\n")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "Signature: "+fmt.Sprintf("%#v", this.Signature)+",\n")
	s = append(s, "Options: "+fmt.Sprintf("%#v", this.Options)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTransaction(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Transaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Transaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Options != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.Options))
		i--
		dAtA[i] = 0x68
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x62
	}
	if m.Version != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x58
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x4a
	}
	if m.GasLimit != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x40
	}
	if m.GasPrice != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.GasPrice))
		i--
		dAtA[i] = 0x38
	}
	if len(m.SndUserName) > 0 {
		i -= len(m.SndUserName)
		copy(dAtA[i:], m.SndUserName)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.SndUserName)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SndAddr) > 0 {
		i -= len(m.SndAddr)
		copy(dAtA[i:], m.SndAddr)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.SndAddr)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RcvUserName) > 0 {
		i -= len(m.RcvUserName)
		copy(dAtA[i:], m.RcvUserName)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.RcvUserName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RcvAddr) > 0 {
		i -= len(m.RcvAddr)
		copy(dAtA[i:], m.RcvAddr)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.RcvAddr)))
		i--
		dAtA[i] = 0x1a
	}
	{
		__caster := &github_com_Dharitri_org_sk_core_data.BigIntCaster{}
		size := __caster.Size(m.Value)
		i -= size
		if _, err := __caster.MarshalTo(m.Value, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTransaction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Nonce != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransaction(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransaction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Transaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovTransaction(uint64(m.Nonce))
	}
	{
		__caster := &github_com_Dharitri_org_sk_core_data.BigIntCaster{}
		l = __caster.Size(m.Value)
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.RcvAddr)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.RcvUserName)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.SndAddr)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.SndUserName)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.GasPrice != 0 {
		n += 1 + sovTransaction(uint64(m.GasPrice))
	}
	if m.GasLimit != 0 {
		n += 1 + sovTransaction(uint64(m.GasLimit))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovTransaction(uint64(m.Version))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.Options != 0 {
		n += 1 + sovTransaction(uint64(m.Options))
	}
	return n
}

func sovTransaction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransaction(x uint64) (n int) {
	return sovTransaction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Transaction) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Transaction{`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`RcvAddr:` + fmt.Sprintf("%v", this.RcvAddr) + `,`,
		`RcvUserName:` + fmt.Sprintf("%v", this.RcvUserName) + `,`,
		`SndAddr:` + fmt.Sprintf("%v", this.SndAddr) + `,`,
		`SndUserName:` + fmt.Sprintf("%v", this.SndUserName) + `,`,
		`GasPrice:` + fmt.Sprintf("%v", this.GasPrice) + `,`,
		`GasLimit:` + fmt.Sprintf("%v", this.GasLimit) + `,`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`ChainID:` + fmt.Sprintf("%v", this.ChainID) + `,`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`Signature:` + fmt.Sprintf("%v", this.Signature) + `,`,
		`Options:` + fmt.Sprintf("%v", this.Options) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTransaction(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Transaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: Transaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_Dharitri_org_sk_core_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Value = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RcvAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RcvAddr = append(m.RcvAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.RcvAddr == nil {
				m.RcvAddr = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RcvUserName", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RcvUserName = append(m.RcvUserName[:0], dAtA[iNdEx:postIndex]...)
			if m.RcvUserName == nil {
				m.RcvUserName = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SndAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SndAddr = append(m.SndAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.SndAddr == nil {
				m.SndAddr = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SndUserName", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SndUserName = append(m.SndUserName[:0], dAtA[iNdEx:postIndex]...)
			if m.SndUserName == nil {
				m.SndUserName = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			m.GasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = append(m.ChainID[:0], dAtA[iNdEx:postIndex]...)
			if m.ChainID == nil {
				m.ChainID = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			m.Options = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Options |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransaction
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
func skipTransaction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
				return 0, ErrInvalidLengthTransaction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransaction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransaction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransaction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransaction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransaction = fmt.Errorf("proto: unexpected end of group")
)
