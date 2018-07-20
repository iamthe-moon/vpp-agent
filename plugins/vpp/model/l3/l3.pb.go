// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: l3.proto

/*
Package l3 is a generated protocol buffer package.

It is generated from these files:
	l3.proto

It has these top-level messages:
	StaticRoutes
	ArpTable
	ProxyArpRanges
	ProxyArpInterfaces
	STNTable
*/
package l3

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type StaticRoutes_Route_RouteType int32

const (
	StaticRoutes_Route_INTRA_VRF StaticRoutes_Route_RouteType = 0
	StaticRoutes_Route_INTER_VRF StaticRoutes_Route_RouteType = 1
)

var StaticRoutes_Route_RouteType_name = map[int32]string{
	0: "INTRA_VRF",
	1: "INTER_VRF",
}
var StaticRoutes_Route_RouteType_value = map[string]int32{
	"INTRA_VRF": 0,
	"INTER_VRF": 1,
}

func (x StaticRoutes_Route_RouteType) String() string {
	return proto.EnumName(StaticRoutes_Route_RouteType_name, int32(x))
}
func (StaticRoutes_Route_RouteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorL3, []int{0, 0, 0}
}

// Static routes
type StaticRoutes struct {
	Routes []*StaticRoutes_Route `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty"`
}

func (m *StaticRoutes) Reset()                    { *m = StaticRoutes{} }
func (m *StaticRoutes) String() string            { return proto.CompactTextString(m) }
func (*StaticRoutes) ProtoMessage()               {}
func (*StaticRoutes) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{0} }

func (m *StaticRoutes) GetRoutes() []*StaticRoutes_Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type StaticRoutes_Route struct {
	Type              StaticRoutes_Route_RouteType `protobuf:"varint,10,opt,name=type,proto3,enum=l3.StaticRoutes_Route_RouteType" json:"type,omitempty"`
	VrfId             uint32                       `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	Description       string                       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	DstIpAddr         string                       `protobuf:"bytes,3,opt,name=dst_ip_addr,json=dstIpAddr,proto3" json:"dst_ip_addr,omitempty"`
	NextHopAddr       string                       `protobuf:"bytes,4,opt,name=next_hop_addr,json=nextHopAddr,proto3" json:"next_hop_addr,omitempty"`
	OutgoingInterface string                       `protobuf:"bytes,5,opt,name=outgoing_interface,json=outgoingInterface,proto3" json:"outgoing_interface,omitempty"`
	Weight            uint32                       `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	Preference        uint32                       `protobuf:"varint,7,opt,name=preference,proto3" json:"preference,omitempty"`
	// (a poor man's primary and backup)
	ViaVrfId uint32 `protobuf:"varint,8,opt,name=via_vrf_id,json=viaVrfId,proto3" json:"via_vrf_id,omitempty"`
}

func (m *StaticRoutes_Route) Reset()                    { *m = StaticRoutes_Route{} }
func (m *StaticRoutes_Route) String() string            { return proto.CompactTextString(m) }
func (*StaticRoutes_Route) ProtoMessage()               {}
func (*StaticRoutes_Route) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{0, 0} }

func (m *StaticRoutes_Route) GetType() StaticRoutes_Route_RouteType {
	if m != nil {
		return m.Type
	}
	return StaticRoutes_Route_INTRA_VRF
}

func (m *StaticRoutes_Route) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *StaticRoutes_Route) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *StaticRoutes_Route) GetDstIpAddr() string {
	if m != nil {
		return m.DstIpAddr
	}
	return ""
}

func (m *StaticRoutes_Route) GetNextHopAddr() string {
	if m != nil {
		return m.NextHopAddr
	}
	return ""
}

func (m *StaticRoutes_Route) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *StaticRoutes_Route) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *StaticRoutes_Route) GetPreference() uint32 {
	if m != nil {
		return m.Preference
	}
	return 0
}

func (m *StaticRoutes_Route) GetViaVrfId() uint32 {
	if m != nil {
		return m.ViaVrfId
	}
	return 0
}

// IP ARP entries
type ArpTable struct {
	ArpEntries []*ArpTable_ArpEntry `protobuf:"bytes,1,rep,name=arp_entries,json=arpEntries" json:"arp_entries,omitempty"`
}

func (m *ArpTable) Reset()                    { *m = ArpTable{} }
func (m *ArpTable) String() string            { return proto.CompactTextString(m) }
func (*ArpTable) ProtoMessage()               {}
func (*ArpTable) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{1} }

func (m *ArpTable) GetArpEntries() []*ArpTable_ArpEntry {
	if m != nil {
		return m.ArpEntries
	}
	return nil
}

type ArpTable_ArpEntry struct {
	Interface   string `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	IpAddress   string `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	PhysAddress string `protobuf:"bytes,3,opt,name=phys_address,json=physAddress,proto3" json:"phys_address,omitempty"`
	Static      bool   `protobuf:"varint,4,opt,name=static,proto3" json:"static,omitempty"`
}

func (m *ArpTable_ArpEntry) Reset()                    { *m = ArpTable_ArpEntry{} }
func (m *ArpTable_ArpEntry) String() string            { return proto.CompactTextString(m) }
func (*ArpTable_ArpEntry) ProtoMessage()               {}
func (*ArpTable_ArpEntry) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{1, 0} }

func (m *ArpTable_ArpEntry) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *ArpTable_ArpEntry) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *ArpTable_ArpEntry) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func (m *ArpTable_ArpEntry) GetStatic() bool {
	if m != nil {
		return m.Static
	}
	return false
}

// Proxy ARP ranges
type ProxyArpRanges struct {
	RangeLists []*ProxyArpRanges_RangeList `protobuf:"bytes,1,rep,name=range_lists,json=rangeLists" json:"range_lists,omitempty"`
}

func (m *ProxyArpRanges) Reset()                    { *m = ProxyArpRanges{} }
func (m *ProxyArpRanges) String() string            { return proto.CompactTextString(m) }
func (*ProxyArpRanges) ProtoMessage()               {}
func (*ProxyArpRanges) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{2} }

func (m *ProxyArpRanges) GetRangeLists() []*ProxyArpRanges_RangeList {
	if m != nil {
		return m.RangeLists
	}
	return nil
}

type ProxyArpRanges_RangeList struct {
	Label  string                            `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Ranges []*ProxyArpRanges_RangeList_Range `protobuf:"bytes,2,rep,name=ranges" json:"ranges,omitempty"`
}

func (m *ProxyArpRanges_RangeList) Reset()                    { *m = ProxyArpRanges_RangeList{} }
func (m *ProxyArpRanges_RangeList) String() string            { return proto.CompactTextString(m) }
func (*ProxyArpRanges_RangeList) ProtoMessage()               {}
func (*ProxyArpRanges_RangeList) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{2, 0} }

func (m *ProxyArpRanges_RangeList) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ProxyArpRanges_RangeList) GetRanges() []*ProxyArpRanges_RangeList_Range {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type ProxyArpRanges_RangeList_Range struct {
	FirstIp string `protobuf:"bytes,1,opt,name=first_ip,json=firstIp,proto3" json:"first_ip,omitempty"`
	LastIp  string `protobuf:"bytes,2,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
}

func (m *ProxyArpRanges_RangeList_Range) Reset()         { *m = ProxyArpRanges_RangeList_Range{} }
func (m *ProxyArpRanges_RangeList_Range) String() string { return proto.CompactTextString(m) }
func (*ProxyArpRanges_RangeList_Range) ProtoMessage()    {}
func (*ProxyArpRanges_RangeList_Range) Descriptor() ([]byte, []int) {
	return fileDescriptorL3, []int{2, 0, 0}
}

func (m *ProxyArpRanges_RangeList_Range) GetFirstIp() string {
	if m != nil {
		return m.FirstIp
	}
	return ""
}

func (m *ProxyArpRanges_RangeList_Range) GetLastIp() string {
	if m != nil {
		return m.LastIp
	}
	return ""
}

// Proxy ARP interfaces
type ProxyArpInterfaces struct {
	InterfaceLists []*ProxyArpInterfaces_InterfaceList `protobuf:"bytes,1,rep,name=interface_lists,json=interfaceLists" json:"interface_lists,omitempty"`
}

func (m *ProxyArpInterfaces) Reset()                    { *m = ProxyArpInterfaces{} }
func (m *ProxyArpInterfaces) String() string            { return proto.CompactTextString(m) }
func (*ProxyArpInterfaces) ProtoMessage()               {}
func (*ProxyArpInterfaces) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{3} }

func (m *ProxyArpInterfaces) GetInterfaceLists() []*ProxyArpInterfaces_InterfaceList {
	if m != nil {
		return m.InterfaceLists
	}
	return nil
}

type ProxyArpInterfaces_InterfaceList struct {
	Label      string                                        `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Interfaces []*ProxyArpInterfaces_InterfaceList_Interface `protobuf:"bytes,2,rep,name=interfaces" json:"interfaces,omitempty"`
}

func (m *ProxyArpInterfaces_InterfaceList) Reset()         { *m = ProxyArpInterfaces_InterfaceList{} }
func (m *ProxyArpInterfaces_InterfaceList) String() string { return proto.CompactTextString(m) }
func (*ProxyArpInterfaces_InterfaceList) ProtoMessage()    {}
func (*ProxyArpInterfaces_InterfaceList) Descriptor() ([]byte, []int) {
	return fileDescriptorL3, []int{3, 0}
}

func (m *ProxyArpInterfaces_InterfaceList) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ProxyArpInterfaces_InterfaceList) GetInterfaces() []*ProxyArpInterfaces_InterfaceList_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type ProxyArpInterfaces_InterfaceList_Interface struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ProxyArpInterfaces_InterfaceList_Interface) Reset() {
	*m = ProxyArpInterfaces_InterfaceList_Interface{}
}
func (m *ProxyArpInterfaces_InterfaceList_Interface) String() string {
	return proto.CompactTextString(m)
}
func (*ProxyArpInterfaces_InterfaceList_Interface) ProtoMessage() {}
func (*ProxyArpInterfaces_InterfaceList_Interface) Descriptor() ([]byte, []int) {
	return fileDescriptorL3, []int{3, 0, 0}
}

func (m *ProxyArpInterfaces_InterfaceList_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// STN (Steal The NIC) feature table
type STNTable struct {
	StnEntries []*STNTable_STNTableEntry `protobuf:"bytes,1,rep,name=stn_entries,json=stnEntries" json:"stn_entries,omitempty"`
}

func (m *STNTable) Reset()                    { *m = STNTable{} }
func (m *STNTable) String() string            { return proto.CompactTextString(m) }
func (*STNTable) ProtoMessage()               {}
func (*STNTable) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{4} }

func (m *STNTable) GetStnEntries() []*STNTable_STNTableEntry {
	if m != nil {
		return m.StnEntries
	}
	return nil
}

type STNTable_STNTableEntry struct {
	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Interface string `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
}

func (m *STNTable_STNTableEntry) Reset()                    { *m = STNTable_STNTableEntry{} }
func (m *STNTable_STNTableEntry) String() string            { return proto.CompactTextString(m) }
func (*STNTable_STNTableEntry) ProtoMessage()               {}
func (*STNTable_STNTableEntry) Descriptor() ([]byte, []int) { return fileDescriptorL3, []int{4, 0} }

func (m *STNTable_STNTableEntry) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *STNTable_STNTableEntry) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func init() {
	proto.RegisterType((*StaticRoutes)(nil), "l3.StaticRoutes")
	proto.RegisterType((*StaticRoutes_Route)(nil), "l3.StaticRoutes.Route")
	proto.RegisterType((*ArpTable)(nil), "l3.ArpTable")
	proto.RegisterType((*ArpTable_ArpEntry)(nil), "l3.ArpTable.ArpEntry")
	proto.RegisterType((*ProxyArpRanges)(nil), "l3.ProxyArpRanges")
	proto.RegisterType((*ProxyArpRanges_RangeList)(nil), "l3.ProxyArpRanges.RangeList")
	proto.RegisterType((*ProxyArpRanges_RangeList_Range)(nil), "l3.ProxyArpRanges.RangeList.Range")
	proto.RegisterType((*ProxyArpInterfaces)(nil), "l3.ProxyArpInterfaces")
	proto.RegisterType((*ProxyArpInterfaces_InterfaceList)(nil), "l3.ProxyArpInterfaces.InterfaceList")
	proto.RegisterType((*ProxyArpInterfaces_InterfaceList_Interface)(nil), "l3.ProxyArpInterfaces.InterfaceList.Interface")
	proto.RegisterType((*STNTable)(nil), "l3.STNTable")
	proto.RegisterType((*STNTable_STNTableEntry)(nil), "l3.STNTable.STNTableEntry")
	proto.RegisterEnum("l3.StaticRoutes_Route_RouteType", StaticRoutes_Route_RouteType_name, StaticRoutes_Route_RouteType_value)
}

func init() { proto.RegisterFile("l3.proto", fileDescriptorL3) }

var fileDescriptorL3 = []byte{
	// 636 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x26, 0x6d, 0x5c, 0xfb, 0xba, 0xe9, 0x07, 0x23, 0x5a, 0x8c, 0x55, 0x4a, 0xb0, 0x58,
	0x84, 0x05, 0x5e, 0x34, 0x88, 0x05, 0x15, 0x8b, 0x2c, 0x8a, 0x88, 0x54, 0x2a, 0x34, 0x8d, 0xba,
	0xb5, 0xdc, 0x78, 0x92, 0x8e, 0x64, 0xec, 0xd1, 0xcc, 0x34, 0x34, 0x5b, 0xde, 0x01, 0x16, 0xbc,
	0x02, 0x0f, 0xc2, 0xbb, 0x20, 0x1e, 0x81, 0x05, 0xf2, 0xcc, 0xd8, 0x71, 0x22, 0x40, 0x6c, 0xe2,
	0xfb, 0x73, 0xee, 0xb5, 0xcf, 0x99, 0x93, 0x01, 0x37, 0x1f, 0xc6, 0x5c, 0x94, 0xaa, 0xc4, 0x9d,
	0x7c, 0x18, 0x7d, 0xdd, 0x82, 0xdd, 0x0b, 0x95, 0x2a, 0x36, 0x25, 0xe5, 0x8d, 0xa2, 0x12, 0xc7,
	0xe0, 0x08, 0x1d, 0x05, 0xa8, 0xbf, 0x35, 0xf0, 0x8f, 0x0f, 0xe2, 0x7c, 0x18, 0xb7, 0x11, 0xb1,
	0x7e, 0x10, 0x8b, 0x0a, 0xbf, 0x77, 0xa0, 0xab, 0x2b, 0xf8, 0x39, 0x6c, 0xab, 0x25, 0xa7, 0x01,
	0xf4, 0xd1, 0x60, 0xef, 0xb8, 0xff, 0xfb, 0x39, 0xf3, 0x3b, 0x59, 0x72, 0x4a, 0x34, 0x1a, 0xef,
	0x83, 0xb3, 0x10, 0xb3, 0x84, 0x65, 0x01, 0xea, 0xa3, 0x41, 0x8f, 0x74, 0x17, 0x62, 0x36, 0xce,
	0x70, 0x1f, 0xfc, 0x8c, 0xca, 0xa9, 0x60, 0x5c, 0xb1, 0xb2, 0x08, 0x3a, 0x7d, 0x34, 0xf0, 0x48,
	0xbb, 0x84, 0x8f, 0xc0, 0xcf, 0xa4, 0x4a, 0x18, 0x4f, 0xd2, 0x2c, 0x13, 0xc1, 0x96, 0x46, 0x78,
	0x99, 0x54, 0x63, 0x3e, 0xca, 0x32, 0x81, 0x23, 0xe8, 0x15, 0xf4, 0x56, 0x25, 0xd7, 0xa5, 0x45,
	0x6c, 0x9b, 0x1d, 0x55, 0xf1, 0x4d, 0x69, 0x30, 0xcf, 0x00, 0x97, 0x37, 0x6a, 0x5e, 0xb2, 0x62,
	0x9e, 0xb0, 0x42, 0x51, 0x31, 0x4b, 0xa7, 0x34, 0xe8, 0x6a, 0xe0, 0xdd, 0xba, 0x33, 0xae, 0x1b,
	0xf8, 0x00, 0x9c, 0x0f, 0x94, 0xcd, 0xaf, 0x55, 0xe0, 0xe8, 0x6f, 0xb5, 0x19, 0x3e, 0x02, 0xe0,
	0x82, 0xce, 0xa8, 0xa0, 0xc5, 0x94, 0x06, 0x3b, 0xba, 0xd7, 0xaa, 0xe0, 0x43, 0x80, 0x05, 0x4b,
	0x13, 0xcb, 0xd3, 0xd5, 0x7d, 0x77, 0xc1, 0xd2, 0xcb, 0x8a, 0x6a, 0xf4, 0x14, 0xbc, 0x46, 0x14,
	0xdc, 0x03, 0x6f, 0x7c, 0x3e, 0x21, 0xa3, 0xe4, 0x92, 0xbc, 0xbe, 0xf3, 0x9f, 0x4d, 0x4f, 0x89,
	0x4e, 0x51, 0xf4, 0x0d, 0x81, 0x3b, 0x12, 0x7c, 0x92, 0x5e, 0xe5, 0x14, 0xbf, 0x00, 0x3f, 0x15,
	0x3c, 0xa1, 0x85, 0x12, 0xac, 0x39, 0xae, 0xfd, 0x4a, 0xf6, 0x1a, 0x52, 0x05, 0xa7, 0x85, 0x12,
	0x4b, 0x02, 0xa9, 0x89, 0x18, 0x95, 0xe1, 0x47, 0xb3, 0x44, 0x37, 0xf0, 0x21, 0x78, 0x2b, 0xe2,
	0xc8, 0x68, 0xd8, 0x14, 0xf0, 0x43, 0x00, 0xab, 0x2f, 0x95, 0xd2, 0x1e, 0x82, 0xc7, 0xb4, 0x76,
	0x54, 0x4a, 0xfc, 0x18, 0x76, 0xf9, 0xf5, 0x52, 0x36, 0x00, 0x73, 0x06, 0x7e, 0x55, 0xab, 0x21,
	0x07, 0xe0, 0x48, 0x6d, 0x02, 0x2d, 0xbf, 0x4b, 0x6c, 0x16, 0xfd, 0x40, 0xb0, 0xf7, 0x4e, 0x94,
	0xb7, 0xcb, 0x91, 0xe0, 0x24, 0x2d, 0xe6, 0x54, 0xe2, 0x57, 0xe0, 0x8b, 0x2a, 0x4a, 0x72, 0x26,
	0x55, 0xcd, 0xe7, 0xb0, 0xe2, 0xb3, 0x0e, 0x8c, 0xf5, 0xe3, 0x8c, 0x49, 0x45, 0x40, 0xd4, 0xa1,
	0x0c, 0xbf, 0x20, 0xf0, 0x9a, 0x0e, 0xbe, 0x07, 0xdd, 0x3c, 0xbd, 0xa2, 0xb9, 0xe5, 0x64, 0x12,
	0xfc, 0x12, 0x1c, 0x3d, 0x51, 0x71, 0xa9, 0xb6, 0x47, 0x7f, 0xdb, 0x6e, 0x22, 0x62, 0x27, 0xc2,
	0x13, 0xe8, 0xea, 0x02, 0x7e, 0x00, 0xee, 0x8c, 0x09, 0x6d, 0x3d, 0xbb, 0x7d, 0x47, 0xe7, 0x63,
	0x8e, 0xef, 0xc3, 0x4e, 0x9e, 0x9a, 0x8e, 0x11, 0xcb, 0xa9, 0xd2, 0x31, 0x8f, 0x7e, 0x22, 0xc0,
	0xf5, 0x7b, 0x1a, 0x3f, 0x49, 0xfc, 0x16, 0xfe, 0x6f, 0xc4, 0x5e, 0xa3, 0xfd, 0xa4, 0xfd, 0x61,
	0xab, 0x81, 0xb8, 0x09, 0x35, 0xfd, 0x3d, 0xd6, 0x4e, 0x65, 0xf8, 0x19, 0x41, 0x6f, 0x0d, 0xf1,
	0x07, 0x19, 0xce, 0x01, 0x9a, 0xc9, 0x5a, 0x8a, 0xf8, 0x5f, 0xde, 0xb8, 0xca, 0x48, 0x6b, 0x43,
	0xf8, 0x08, 0xbc, 0xd5, 0x9f, 0x04, 0xc3, 0x76, 0x91, 0xbe, 0xaf, 0xcd, 0xa4, 0xe3, 0xe8, 0x13,
	0x02, 0xf7, 0x62, 0x72, 0x6e, 0x7c, 0x7b, 0x02, 0xbe, 0x54, 0xc5, 0x86, 0x6f, 0x43, 0x7d, 0x5d,
	0x58, 0x48, 0x13, 0x58, 0xf3, 0x4a, 0x55, 0xd4, 0xe6, 0x3d, 0x83, 0xde, 0x5a, 0x73, 0xc3, 0xa2,
	0x68, 0xd3, 0xa2, 0x6b, 0xfe, 0xee, 0x6c, 0xf8, 0xfb, 0xca, 0xd1, 0x17, 0xe1, 0xf0, 0x57, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x41, 0x97, 0x51, 0xa6, 0x14, 0x05, 0x00, 0x00,
}
