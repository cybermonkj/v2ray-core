package internet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_serial "v2ray.com/core/common/serial"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransportProtocol int32

const (
	TransportProtocol_TCP       TransportProtocol = 0
	TransportProtocol_UDP       TransportProtocol = 1
	TransportProtocol_MKCP      TransportProtocol = 2
	TransportProtocol_WebSocket TransportProtocol = 3
	TransportProtocol_HTTP      TransportProtocol = 4
)

var TransportProtocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
	2: "MKCP",
	3: "WebSocket",
	4: "HTTP",
}
var TransportProtocol_value = map[string]int32{
	"TCP":       0,
	"UDP":       1,
	"MKCP":      2,
	"WebSocket": 3,
	"HTTP":      4,
}

func (x TransportProtocol) String() string {
	return proto.EnumName(TransportProtocol_name, int32(x))
}
func (TransportProtocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TransportConfig struct {
	// Type of network that this settings supports.
	Protocol TransportProtocol `protobuf:"varint,1,opt,name=protocol,enum=v2ray.core.transport.internet.TransportProtocol" json:"protocol,omitempty"`
	// Specific settings. Must be of the transports.
	Settings *v2ray_core_common_serial.TypedMessage `protobuf:"bytes,2,opt,name=settings" json:"settings,omitempty"`
}

func (m *TransportConfig) Reset()                    { *m = TransportConfig{} }
func (m *TransportConfig) String() string            { return proto.CompactTextString(m) }
func (*TransportConfig) ProtoMessage()               {}
func (*TransportConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TransportConfig) GetProtocol() TransportProtocol {
	if m != nil {
		return m.Protocol
	}
	return TransportProtocol_TCP
}

func (m *TransportConfig) GetSettings() *v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.Settings
	}
	return nil
}

type StreamConfig struct {
	// Effective network.
	Protocol          TransportProtocol  `protobuf:"varint,1,opt,name=protocol,enum=v2ray.core.transport.internet.TransportProtocol" json:"protocol,omitempty"`
	TransportSettings []*TransportConfig `protobuf:"bytes,2,rep,name=transport_settings,json=transportSettings" json:"transport_settings,omitempty"`
	// Type of security. Must be a message name of the settings proto.
	SecurityType string `protobuf:"bytes,3,opt,name=security_type,json=securityType" json:"security_type,omitempty"`
	// Settings for transport security. For now the only choice is TLS.
	SecuritySettings []*v2ray_core_common_serial.TypedMessage `protobuf:"bytes,4,rep,name=security_settings,json=securitySettings" json:"security_settings,omitempty"`
}

func (m *StreamConfig) Reset()                    { *m = StreamConfig{} }
func (m *StreamConfig) String() string            { return proto.CompactTextString(m) }
func (*StreamConfig) ProtoMessage()               {}
func (*StreamConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StreamConfig) GetProtocol() TransportProtocol {
	if m != nil {
		return m.Protocol
	}
	return TransportProtocol_TCP
}

func (m *StreamConfig) GetTransportSettings() []*TransportConfig {
	if m != nil {
		return m.TransportSettings
	}
	return nil
}

func (m *StreamConfig) GetSecurityType() string {
	if m != nil {
		return m.SecurityType
	}
	return ""
}

func (m *StreamConfig) GetSecuritySettings() []*v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.SecuritySettings
	}
	return nil
}

type ProxyConfig struct {
	Tag string `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
}

func (m *ProxyConfig) Reset()                    { *m = ProxyConfig{} }
func (m *ProxyConfig) String() string            { return proto.CompactTextString(m) }
func (*ProxyConfig) ProtoMessage()               {}
func (*ProxyConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProxyConfig) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func init() {
	proto.RegisterType((*TransportConfig)(nil), "v2ray.core.transport.internet.TransportConfig")
	proto.RegisterType((*StreamConfig)(nil), "v2ray.core.transport.internet.StreamConfig")
	proto.RegisterType((*ProxyConfig)(nil), "v2ray.core.transport.internet.ProxyConfig")
	proto.RegisterEnum("v2ray.core.transport.internet.TransportProtocol", TransportProtocol_name, TransportProtocol_value)
}

func init() { proto.RegisterFile("v2ray.com/core/transport/internet/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0x87, 0xdf, 0x24, 0xe5, 0x35, 0x9d, 0xb6, 0x9a, 0xee, 0xa9, 0x08, 0xc5, 0x5a, 0x41, 0x82,
	0x87, 0x4d, 0x89, 0xdf, 0xa0, 0xf1, 0x50, 0xd1, 0x62, 0x48, 0xa2, 0x82, 0x20, 0x25, 0x5d, 0xd7,
	0x10, 0x6c, 0xb2, 0x65, 0xb3, 0x8a, 0xf9, 0x3c, 0xde, 0xbc, 0xfb, 0xfd, 0x24, 0xff, 0x96, 0xa2,
	0x50, 0x7a, 0xf1, 0x36, 0x64, 0x7e, 0xf3, 0xcc, 0x93, 0x59, 0xc0, 0x6f, 0x36, 0x0f, 0x73, 0x4c,
	0x58, 0x62, 0x11, 0xc6, 0xa9, 0x25, 0x78, 0x98, 0x66, 0x6b, 0xc6, 0x85, 0x15, 0xa7, 0x82, 0xf2,
	0x94, 0x0a, 0x8b, 0xb0, 0xf4, 0x39, 0x8e, 0xf0, 0x9a, 0x33, 0xc1, 0xd0, 0xb0, 0xc9, 0x73, 0x8a,
	0x65, 0x16, 0x37, 0xd9, 0xc3, 0xc9, 0x0f, 0x1c, 0x61, 0x49, 0xc2, 0x52, 0x2b, 0xa3, 0x3c, 0x0e,
	0x57, 0x96, 0xc8, 0xd7, 0xf4, 0x69, 0x91, 0xd0, 0x2c, 0x0b, 0x23, 0x5a, 0x01, 0xc7, 0x1f, 0x0a,
	0x1c, 0x04, 0x0d, 0xc8, 0x29, 0x57, 0xa1, 0x6b, 0xd0, 0xcb, 0x26, 0x61, 0xab, 0x81, 0x32, 0x52,
	0xcc, 0x7d, 0x7b, 0x82, 0xb7, 0xee, 0xc5, 0x92, 0xe0, 0xd6, 0x73, 0x9e, 0x24, 0xa0, 0x29, 0xe8,
	0x19, 0x15, 0x22, 0x4e, 0xa3, 0x6c, 0xa0, 0x8e, 0x14, 0xb3, 0x63, 0x9f, 0x6e, 0xd2, 0x2a, 0x45,
	0x5c, 0x29, 0xe2, 0xa0, 0x50, 0x9c, 0x57, 0x86, 0x9e, 0x9c, 0x1b, 0x7f, 0xa9, 0xd0, 0xf5, 0x05,
	0xa7, 0x61, 0xf2, 0x27, 0x8a, 0x8f, 0x80, 0xe4, 0xc4, 0x62, 0x43, 0x56, 0x33, 0x3b, 0x36, 0xde,
	0x95, 0x5b, 0x99, 0x79, 0x7d, 0x99, 0xf1, 0x6b, 0x10, 0x3a, 0x81, 0x5e, 0x46, 0xc9, 0x2b, 0x8f,
	0x45, 0xbe, 0x28, 0xde, 0x60, 0xa0, 0x8d, 0x14, 0xb3, 0xed, 0x75, 0x9b, 0x8f, 0xc5, 0x4f, 0x23,
	0x1f, 0xfa, 0x32, 0x24, 0x15, 0x5a, 0xa5, 0xc2, 0xae, 0xf7, 0x32, 0x1a, 0x40, 0xb3, 0x79, 0x7c,
	0x04, 0x1d, 0x97, 0xb3, 0xf7, 0xbc, 0xbe, 0x9a, 0x01, 0x9a, 0x08, 0xa3, 0xf2, 0x60, 0x6d, 0xaf,
	0x28, 0xcf, 0x66, 0xd0, 0xff, 0x75, 0x18, 0xb4, 0x07, 0x5a, 0xe0, 0xb8, 0xc6, 0xbf, 0xa2, 0xb8,
	0xbd, 0x70, 0x0d, 0x05, 0xe9, 0xd0, 0x9a, 0x5f, 0x39, 0xae, 0xa1, 0xa2, 0x1e, 0xb4, 0xef, 0xe9,
	0xd2, 0x67, 0xe4, 0x85, 0x0a, 0x43, 0x2b, 0x1a, 0xb3, 0x20, 0x70, 0x8d, 0xd6, 0xf4, 0x06, 0x8e,
	0x09, 0x4b, 0xb6, 0x1f, 0xcb, 0x55, 0x1e, 0xf4, 0xa6, 0xfe, 0x54, 0x87, 0x77, 0xb6, 0x17, 0xe6,
	0xd8, 0x29, 0xb2, 0x52, 0x02, 0x5f, 0xd6, 0xfd, 0xe5, 0xff, 0xf2, 0x79, 0xce, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xf3, 0x7b, 0xd5, 0x57, 0x23, 0x03, 0x00, 0x00,
}
