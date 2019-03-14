// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zregister.proto

package zmet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// XXX is this used? Not in client.go; deprecate
type ZRegisterResult int32

const (
	ZRegisterResult_ZRegNone        ZRegisterResult = 0
	ZRegisterResult_ZRegSuccess     ZRegisterResult = 1
	ZRegisterResult_ZRegNotActive   ZRegisterResult = 2
	ZRegisterResult_ZRegAlreadyDone ZRegisterResult = 3
	ZRegisterResult_ZRegDeviceNA    ZRegisterResult = 4
	ZRegisterResult_ZRegFailed      ZRegisterResult = 5
)

var ZRegisterResult_name = map[int32]string{
	0: "ZRegNone",
	1: "ZRegSuccess",
	2: "ZRegNotActive",
	3: "ZRegAlreadyDone",
	4: "ZRegDeviceNA",
	5: "ZRegFailed",
}
var ZRegisterResult_value = map[string]int32{
	"ZRegNone":        0,
	"ZRegSuccess":     1,
	"ZRegNotActive":   2,
	"ZRegAlreadyDone": 3,
	"ZRegDeviceNA":    4,
	"ZRegFailed":      5,
}

func (x ZRegisterResult) String() string {
	return proto.EnumName(ZRegisterResult_name, int32(x))
}
func (ZRegisterResult) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// XXX is this used? Not in client.go; deprecate
type ZRegisterResp struct {
	Result ZRegisterResult `protobuf:"varint,2,opt,name=result,enum=ZRegisterResult" json:"result,omitempty"`
}

func (m *ZRegisterResp) Reset()                    { *m = ZRegisterResp{} }
func (m *ZRegisterResp) String() string            { return proto.CompactTextString(m) }
func (*ZRegisterResp) ProtoMessage()               {}
func (*ZRegisterResp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ZRegisterResp) GetResult() ZRegisterResult {
	if m != nil {
		return m.Result
	}
	return ZRegisterResult_ZRegNone
}

type ZRegisterMsg struct {
	OnBoardKey string `protobuf:"bytes,1,opt,name=onBoardKey" json:"onBoardKey,omitempty"`
	PemCert    []byte `protobuf:"bytes,2,opt,name=pemCert,proto3" json:"pemCert,omitempty"`
	Serial     string `protobuf:"bytes,3,opt,name=serial" json:"serial,omitempty"`
}

func (m *ZRegisterMsg) Reset()                    { *m = ZRegisterMsg{} }
func (m *ZRegisterMsg) String() string            { return proto.CompactTextString(m) }
func (*ZRegisterMsg) ProtoMessage()               {}
func (*ZRegisterMsg) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ZRegisterMsg) GetOnBoardKey() string {
	if m != nil {
		return m.OnBoardKey
	}
	return ""
}

func (m *ZRegisterMsg) GetPemCert() []byte {
	if m != nil {
		return m.PemCert
	}
	return nil
}

func (m *ZRegisterMsg) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func init() {
	proto.RegisterType((*ZRegisterResp)(nil), "ZRegisterResp")
	proto.RegisterType((*ZRegisterMsg)(nil), "ZRegisterMsg")
	proto.RegisterEnum("ZRegisterResult", ZRegisterResult_name, ZRegisterResult_value)
}

func init() { proto.RegisterFile("zregister.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0x5f, 0xda, 0x67, 0xd4, 0x6b, 0x6c, 0xc6, 0x2b, 0x48, 0x10, 0x91, 0xd2, 0x55, 0x70,
	0x91, 0x80, 0xae, 0x5c, 0xa6, 0x16, 0x37, 0x62, 0x16, 0x71, 0xd7, 0x95, 0xd3, 0xcc, 0x25, 0x0e,
	0x4c, 0x3a, 0x61, 0x66, 0x52, 0x68, 0x3e, 0xbd, 0xa4, 0x89, 0x52, 0x5c, 0x9e, 0xdf, 0xe1, 0x9c,
	0xfb, 0x07, 0xc2, 0xce, 0x50, 0x25, 0xad, 0x23, 0x93, 0x34, 0x46, 0x3b, 0xbd, 0x78, 0x86, 0xcb,
	0x75, 0x31, 0xa2, 0x82, 0x6c, 0x83, 0x31, 0xf8, 0x86, 0x6c, 0xab, 0x5c, 0x34, 0x99, 0x7b, 0xf1,
	0xec, 0x91, 0x25, 0xc7, 0x7e, 0xab, 0x5c, 0x31, 0xfa, 0x8b, 0x4f, 0x08, 0x7e, 0xad, 0x77, 0x5b,
	0xe1, 0x3d, 0x80, 0xde, 0x2e, 0x35, 0x37, 0xe2, 0x8d, 0xf6, 0x91, 0x37, 0xf7, 0xe2, 0xf3, 0xe2,
	0x88, 0x60, 0x04, 0xa7, 0x0d, 0xd5, 0x2f, 0x64, 0x86, 0xea, 0xa0, 0xf8, 0x91, 0x78, 0x03, 0xbe,
	0x25, 0x23, 0xb9, 0x8a, 0xa6, 0x87, 0xd4, 0xa8, 0x1e, 0x3a, 0x08, 0xff, 0x0c, 0xc7, 0x00, 0xce,
	0x7a, 0x94, 0xeb, 0x2d, 0xb1, 0x7f, 0x18, 0xc2, 0x45, 0xaf, 0x3e, 0xda, 0xb2, 0x24, 0x6b, 0x99,
	0x87, 0x57, 0xc3, 0x39, 0xb9, 0x76, 0x59, 0xe9, 0xe4, 0x8e, 0xd8, 0x04, 0xaf, 0x87, 0x92, 0x4c,
	0x19, 0xe2, 0x62, 0xbf, 0xea, 0x83, 0x53, 0x64, 0xc3, 0xee, 0x2b, 0xda, 0xc9, 0x92, 0xf2, 0x8c,
	0xfd, 0xc7, 0x19, 0x40, 0x4f, 0x5e, 0xb9, 0x54, 0x24, 0xd8, 0xc9, 0xf2, 0x6e, 0x7d, 0x5b, 0x49,
	0xf7, 0xd5, 0x6e, 0x92, 0x52, 0xd7, 0x69, 0x47, 0x82, 0x04, 0x4f, 0x79, 0x23, 0xd3, 0xae, 0x26,
	0xb7, 0xf1, 0x0f, 0xdf, 0x7b, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x70, 0x58, 0x9b, 0x50,
	0x01, 0x00, 0x00,
}
