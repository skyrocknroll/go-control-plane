// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/address.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/address.proto
	api/auth.proto
	api/base.proto
	api/bootstrap.proto
	api/cds.proto
	api/discovery.proto
	api/eds.proto
	api/hds.proto
	api/health_check.proto
	api/lds.proto
	api/protocol.proto
	api/rds.proto
	api/rls.proto
	api/sds.proto

It has these top-level messages:
	Pipe
	SocketAddress
	BindConfig
	Address
	CidrRange
	AuthAction
	Locality
	Node
	Endpoint
	Metadata
	RuntimeUInt32
	HeaderValue
	HeaderValueOption
	ApiConfigSource
	AggregatedConfigSource
	ConfigSource
	TransportSocket
	LightstepConfig
	ZipkinConfig
	Tracing
	Admin
	ClusterManager
	StatsdSink
	StatsSink
	TagSpecifier
	StatsConfig
	Watchdog
	Runtime
	RateLimitServiceConfig
	Bootstrap
	UpstreamBindConfig
	CircuitBreakers
	Cluster
	DiscoveryRequest
	DiscoveryResponse
	LbEndpoint
	LocalityLbEndpoints
	EndpointLoadMetricStats
	UpstreamLocalityStats
	ClusterStats
	LoadStatsRequest
	ClusterLoadAssignment
	LoadStatsResponse
	Capability
	HealthCheckRequest
	EndpointHealth
	EndpointHealthResponse
	HealthCheckRequestOrEndpointHealthResponse
	LocalityEndpoints
	ClusterHealthCheck
	HealthCheckSpecifier
	HealthCheck
	Filter
	FilterChainMatch
	FilterChain
	Listener
	TcpProtocolOptions
	Http1ProtocolOptions
	Http2ProtocolOptions
	GrpcProtocolOptions
	WeightedCluster
	RouteMatch
	CorsPolicy
	RouteAction
	RedirectAction
	Decorator
	Route
	VirtualCluster
	RateLimit
	HeaderMatcher
	VirtualHost
	RouteConfiguration
	RateLimitRequest
	RateLimitDescriptor
	RateLimitResponse
	DataSource
	TlsParameters
	TlsCertificate
	TlsSessionTicketKeys
	CertificateValidationContext
	CommonTlsContext
	UpstreamTlsContext
	DownstreamTlsContext
	SdsSecretConfig
	Secret
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SocketAddress_Protocol int32

const (
	SocketAddress_TCP SocketAddress_Protocol = 0
	SocketAddress_UDP SocketAddress_Protocol = 1
)

var SocketAddress_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}
var SocketAddress_Protocol_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x SocketAddress_Protocol) String() string {
	return proto.EnumName(SocketAddress_Protocol_name, int32(x))
}
func (SocketAddress_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Pipe struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *Pipe) Reset()                    { *m = Pipe{} }
func (m *Pipe) String() string            { return proto.CompactTextString(m) }
func (*Pipe) ProtoMessage()               {}
func (*Pipe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Pipe) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type SocketAddress struct {
	Protocol SocketAddress_Protocol `protobuf:"varint,1,opt,name=protocol,enum=envoy.api.v2.SocketAddress_Protocol" json:"protocol,omitempty"`
	Address  string                 `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	// Types that are valid to be assigned to PortSpecifier:
	//	*SocketAddress_PortValue
	//	*SocketAddress_NamedPort
	PortSpecifier isSocketAddress_PortSpecifier `protobuf_oneof:"port_specifier"`
	ResolverName  string                        `protobuf:"bytes,5,opt,name=resolver_name,json=resolverName" json:"resolver_name,omitempty"`
}

func (m *SocketAddress) Reset()                    { *m = SocketAddress{} }
func (m *SocketAddress) String() string            { return proto.CompactTextString(m) }
func (*SocketAddress) ProtoMessage()               {}
func (*SocketAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isSocketAddress_PortSpecifier interface {
	isSocketAddress_PortSpecifier()
}

type SocketAddress_PortValue struct {
	PortValue uint32 `protobuf:"varint,3,opt,name=port_value,json=portValue,oneof"`
}
type SocketAddress_NamedPort struct {
	NamedPort string `protobuf:"bytes,4,opt,name=named_port,json=namedPort,oneof"`
}

func (*SocketAddress_PortValue) isSocketAddress_PortSpecifier() {}
func (*SocketAddress_NamedPort) isSocketAddress_PortSpecifier() {}

func (m *SocketAddress) GetPortSpecifier() isSocketAddress_PortSpecifier {
	if m != nil {
		return m.PortSpecifier
	}
	return nil
}

func (m *SocketAddress) GetProtocol() SocketAddress_Protocol {
	if m != nil {
		return m.Protocol
	}
	return SocketAddress_TCP
}

func (m *SocketAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SocketAddress) GetPortValue() uint32 {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_PortValue); ok {
		return x.PortValue
	}
	return 0
}

func (m *SocketAddress) GetNamedPort() string {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_NamedPort); ok {
		return x.NamedPort
	}
	return ""
}

func (m *SocketAddress) GetResolverName() string {
	if m != nil {
		return m.ResolverName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SocketAddress) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SocketAddress_OneofMarshaler, _SocketAddress_OneofUnmarshaler, _SocketAddress_OneofSizer, []interface{}{
		(*SocketAddress_PortValue)(nil),
		(*SocketAddress_NamedPort)(nil),
	}
}

func _SocketAddress_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SocketAddress)
	// port_specifier
	switch x := m.PortSpecifier.(type) {
	case *SocketAddress_PortValue:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.PortValue))
	case *SocketAddress_NamedPort:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.NamedPort)
	case nil:
	default:
		return fmt.Errorf("SocketAddress.PortSpecifier has unexpected type %T", x)
	}
	return nil
}

func _SocketAddress_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SocketAddress)
	switch tag {
	case 3: // port_specifier.port_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.PortSpecifier = &SocketAddress_PortValue{uint32(x)}
		return true, err
	case 4: // port_specifier.named_port
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PortSpecifier = &SocketAddress_NamedPort{x}
		return true, err
	default:
		return false, nil
	}
}

func _SocketAddress_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SocketAddress)
	// port_specifier
	switch x := m.PortSpecifier.(type) {
	case *SocketAddress_PortValue:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.PortValue))
	case *SocketAddress_NamedPort:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.NamedPort)))
		n += len(x.NamedPort)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type BindConfig struct {
	SourceAddress *SocketAddress `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
}

func (m *BindConfig) Reset()                    { *m = BindConfig{} }
func (m *BindConfig) String() string            { return proto.CompactTextString(m) }
func (*BindConfig) ProtoMessage()               {}
func (*BindConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BindConfig) GetSourceAddress() *SocketAddress {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

type Address struct {
	// Types that are valid to be assigned to Address:
	//	*Address_SocketAddress
	//	*Address_Pipe
	Address isAddress_Address `protobuf_oneof:"address"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isAddress_Address interface {
	isAddress_Address()
}

type Address_SocketAddress struct {
	SocketAddress *SocketAddress `protobuf:"bytes,1,opt,name=socket_address,json=socketAddress,oneof"`
}
type Address_Pipe struct {
	Pipe *Pipe `protobuf:"bytes,2,opt,name=pipe,oneof"`
}

func (*Address_SocketAddress) isAddress_Address() {}
func (*Address_Pipe) isAddress_Address()          {}

func (m *Address) GetAddress() isAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Address) GetSocketAddress() *SocketAddress {
	if x, ok := m.GetAddress().(*Address_SocketAddress); ok {
		return x.SocketAddress
	}
	return nil
}

func (m *Address) GetPipe() *Pipe {
	if x, ok := m.GetAddress().(*Address_Pipe); ok {
		return x.Pipe
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Address) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Address_OneofMarshaler, _Address_OneofUnmarshaler, _Address_OneofSizer, []interface{}{
		(*Address_SocketAddress)(nil),
		(*Address_Pipe)(nil),
	}
}

func _Address_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Address)
	// address
	switch x := m.Address.(type) {
	case *Address_SocketAddress:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SocketAddress); err != nil {
			return err
		}
	case *Address_Pipe:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pipe); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Address.Address has unexpected type %T", x)
	}
	return nil
}

func _Address_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Address)
	switch tag {
	case 1: // address.socket_address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SocketAddress)
		err := b.DecodeMessage(msg)
		m.Address = &Address_SocketAddress{msg}
		return true, err
	case 2: // address.pipe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Pipe)
		err := b.DecodeMessage(msg)
		m.Address = &Address_Pipe{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Address_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Address)
	// address
	switch x := m.Address.(type) {
	case *Address_SocketAddress:
		s := proto.Size(x.SocketAddress)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Address_Pipe:
		s := proto.Size(x.Pipe)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CidrRange struct {
	AddressPrefix string                       `protobuf:"bytes,1,opt,name=address_prefix,json=addressPrefix" json:"address_prefix,omitempty"`
	PrefixLen     *google_protobuf.UInt32Value `protobuf:"bytes,2,opt,name=prefix_len,json=prefixLen" json:"prefix_len,omitempty"`
}

func (m *CidrRange) Reset()                    { *m = CidrRange{} }
func (m *CidrRange) String() string            { return proto.CompactTextString(m) }
func (*CidrRange) ProtoMessage()               {}
func (*CidrRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CidrRange) GetAddressPrefix() string {
	if m != nil {
		return m.AddressPrefix
	}
	return ""
}

func (m *CidrRange) GetPrefixLen() *google_protobuf.UInt32Value {
	if m != nil {
		return m.PrefixLen
	}
	return nil
}

func init() {
	proto.RegisterType((*Pipe)(nil), "envoy.api.v2.Pipe")
	proto.RegisterType((*SocketAddress)(nil), "envoy.api.v2.SocketAddress")
	proto.RegisterType((*BindConfig)(nil), "envoy.api.v2.BindConfig")
	proto.RegisterType((*Address)(nil), "envoy.api.v2.Address")
	proto.RegisterType((*CidrRange)(nil), "envoy.api.v2.CidrRange")
	proto.RegisterEnum("envoy.api.v2.SocketAddress_Protocol", SocketAddress_Protocol_name, SocketAddress_Protocol_value)
}

func init() { proto.RegisterFile("api/address.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x93, 0x94, 0xd4, 0xd3, 0x3a, 0x0a, 0x7b, 0xc1, 0x0a, 0x15, 0x54, 0x06, 0xa4, 0x9c,
	0x36, 0x92, 0x7b, 0xe4, 0x02, 0x49, 0x0f, 0x41, 0x42, 0x95, 0x65, 0x28, 0x27, 0x24, 0x6b, 0x6b,
	0x8f, 0xcd, 0x0a, 0x77, 0x77, 0xb5, 0x76, 0x0c, 0x5c, 0x11, 0x27, 0x3e, 0x89, 0x13, 0xbf, 0xc3,
	0x5f, 0xa0, 0x5d, 0xaf, 0xab, 0xe6, 0x82, 0xc4, 0x6d, 0xf6, 0xcd, 0x7b, 0x33, 0xb3, 0xf3, 0x06,
	0x1e, 0x32, 0xc5, 0xd7, 0xac, 0x28, 0x34, 0x36, 0x0d, 0x55, 0x5a, 0xb6, 0x92, 0x9c, 0xa2, 0xe8,
	0xe4, 0x37, 0xca, 0x14, 0xa7, 0x5d, 0xbc, 0x7c, 0x52, 0x49, 0x59, 0xd5, 0xb8, 0xb6, 0xb9, 0x9b,
	0x7d, 0xb9, 0xfe, 0xa2, 0x99, 0x52, 0xa8, 0x1d, 0x7b, 0xf9, 0xa8, 0x63, 0x35, 0x2f, 0x58, 0x8b,
	0xeb, 0x21, 0xe8, 0x13, 0xd1, 0x12, 0xa6, 0x09, 0x57, 0x48, 0x08, 0x4c, 0x15, 0x6b, 0x3f, 0x85,
	0xde, 0xb9, 0xb7, 0xf2, 0x53, 0x1b, 0x47, 0xdf, 0xc7, 0x10, 0xbc, 0x93, 0xf9, 0x67, 0x6c, 0x5f,
	0xf7, 0xad, 0xc9, 0x2b, 0x38, 0xb6, 0xb2, 0x5c, 0xd6, 0x96, 0x39, 0x8f, 0x9f, 0xd3, 0xfb, 0x73,
	0xd0, 0x03, 0x3a, 0x4d, 0x1c, 0x37, 0xbd, 0x53, 0x91, 0x10, 0x66, 0xee, 0x1f, 0xe1, 0xd8, 0xb6,
	0x1a, 0x9e, 0xe4, 0x29, 0x80, 0x92, 0xba, 0xcd, 0x3a, 0x56, 0xef, 0x31, 0x9c, 0x9c, 0x7b, 0xab,
	0x60, 0x37, 0x4a, 0x7d, 0x83, 0x7d, 0x30, 0x90, 0x21, 0x08, 0x76, 0x8b, 0x45, 0x66, 0xa0, 0x70,
	0x6a, 0xd4, 0x86, 0x60, 0xb1, 0x44, 0xea, 0x96, 0x3c, 0x83, 0x40, 0x63, 0x23, 0xeb, 0x0e, 0x75,
	0x66, 0xd0, 0xf0, 0xc8, 0x76, 0x38, 0x1d, 0xc0, 0x2b, 0x76, 0x8b, 0xd1, 0x19, 0x1c, 0x0f, 0x63,
	0x91, 0x19, 0x4c, 0xde, 0x6f, 0x93, 0xc5, 0xc8, 0x04, 0xd7, 0x97, 0xc9, 0xc2, 0xdb, 0x2c, 0x60,
	0x6e, 0x87, 0x68, 0x14, 0xe6, 0xbc, 0xe4, 0xa8, 0xa3, 0x8f, 0x00, 0x1b, 0x2e, 0x8a, 0xad, 0x14,
	0x25, 0xaf, 0xc8, 0x15, 0xcc, 0x1b, 0xb9, 0xd7, 0x39, 0x66, 0xc3, 0x2f, 0xcc, 0x1a, 0x4e, 0xe2,
	0xc7, 0xff, 0x58, 0xc3, 0x06, 0x7e, 0xfd, 0xf9, 0x3d, 0x39, 0xfa, 0xe9, 0x8d, 0x17, 0x5e, 0x1a,
	0xf4, 0x72, 0x97, 0x8a, 0x7e, 0x78, 0x30, 0x1b, 0x96, 0x7b, 0x69, 0x6a, 0x1b, 0xdd, 0x7f, 0xd4,
	0xde, 0x8d, 0x4c, 0xc5, 0xfb, 0x16, 0xad, 0x60, 0xaa, 0xb8, 0x42, 0xbb, 0xdd, 0x93, 0x98, 0x1c,
	0x6a, 0x8d, 0xd5, 0xbb, 0x51, 0x6a, 0x19, 0x1b, 0xff, 0xce, 0x8a, 0x48, 0x82, 0xbf, 0xe5, 0x85,
	0x4e, 0x99, 0xa8, 0x90, 0xbc, 0x80, 0xb9, 0xc3, 0x33, 0xa5, 0xb1, 0xe4, 0x5f, 0xdd, 0x51, 0x04,
	0x0e, 0x4d, 0x2c, 0x48, 0x5e, 0x02, 0xf4, 0xe9, 0xac, 0x46, 0xe1, 0xda, 0x9d, 0xd1, 0xfe, 0x0e,
	0xe9, 0x70, 0x87, 0xf4, 0xfa, 0x8d, 0x68, 0x2f, 0x62, 0x6b, 0x60, 0xea, 0xf7, 0xfc, 0xb7, 0x28,
	0x6e, 0x1e, 0x58, 0xc2, 0xc5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xee, 0x8e, 0x31, 0xd9,
	0x02, 0x00, 0x00,
}
