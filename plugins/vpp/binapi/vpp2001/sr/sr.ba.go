// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/sr.api.json

/*
Package sr is a generated VPP binary API for 'sr' module.

It consists of:
	  3 types
	 20 messages
	 10 services
*/
package sr

import (
	bytes "bytes"
	context "context"
	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
	io "io"
	strconv "strconv"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "sr"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xbf277f96
)

// SrIP6Address represents VPP binary API type 'sr_ip6_address'.
type SrIP6Address struct {
	Data []byte `struc:"[16]byte"`
}

func (*SrIP6Address) GetTypeName() string { return "sr_ip6_address" }

// Srv6Sid represents VPP binary API type 'srv6_sid'.
type Srv6Sid struct {
	Addr []byte `struc:"[16]byte"`
}

func (*Srv6Sid) GetTypeName() string { return "srv6_sid" }

// Srv6SidList represents VPP binary API type 'srv6_sid_list'.
type Srv6SidList struct {
	NumSids uint8 `struc:"sizeof=Sids"`
	Weight  uint32
	Sids    []Srv6Sid
}

func (*Srv6SidList) GetTypeName() string { return "srv6_sid_list" }

// SrLocalsidAddDel represents VPP binary API message 'sr_localsid_add_del'.
type SrLocalsidAddDel struct {
	IsDel     uint8
	Localsid  Srv6Sid
	EndPsp    uint8
	Behavior  uint8
	SwIfIndex uint32
	VlanIndex uint32
	FibTable  uint32
	NhAddr6   []byte `struc:"[16]byte"`
	NhAddr4   []byte `struc:"[4]byte"`
}

func (m *SrLocalsidAddDel) Reset()                        { *m = SrLocalsidAddDel{} }
func (*SrLocalsidAddDel) GetMessageName() string          { return "sr_localsid_add_del" }
func (*SrLocalsidAddDel) GetCrcString() string            { return "b30489eb" }
func (*SrLocalsidAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// SrLocalsidAddDelReply represents VPP binary API message 'sr_localsid_add_del_reply'.
type SrLocalsidAddDelReply struct {
	Retval int32
}

func (m *SrLocalsidAddDelReply) Reset()                        { *m = SrLocalsidAddDelReply{} }
func (*SrLocalsidAddDelReply) GetMessageName() string          { return "sr_localsid_add_del_reply" }
func (*SrLocalsidAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*SrLocalsidAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SrLocalsidsDetails represents VPP binary API message 'sr_localsids_details'.
type SrLocalsidsDetails struct {
	Addr                    Srv6Sid
	EndPsp                  uint8
	Behavior                uint16
	FibTable                uint32
	VlanIndex               uint32
	XconnectNhAddr6         []byte `struc:"[16]byte"`
	XconnectNhAddr4         []byte `struc:"[4]byte"`
	XconnectIfaceOrVrfTable uint32
}

func (m *SrLocalsidsDetails) Reset()                        { *m = SrLocalsidsDetails{} }
func (*SrLocalsidsDetails) GetMessageName() string          { return "sr_localsids_details" }
func (*SrLocalsidsDetails) GetCrcString() string            { return "0791babc" }
func (*SrLocalsidsDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// SrLocalsidsDump represents VPP binary API message 'sr_localsids_dump'.
type SrLocalsidsDump struct{}

func (m *SrLocalsidsDump) Reset()                        { *m = SrLocalsidsDump{} }
func (*SrLocalsidsDump) GetMessageName() string          { return "sr_localsids_dump" }
func (*SrLocalsidsDump) GetCrcString() string            { return "51077d14" }
func (*SrLocalsidsDump) GetMessageType() api.MessageType { return api.RequestMessage }

// SrPoliciesDetails represents VPP binary API message 'sr_policies_details'.
type SrPoliciesDetails struct {
	Bsid        Srv6Sid
	Type        uint8
	IsEncap     uint8
	FibTable    uint32
	NumSidLists uint8 `struc:"sizeof=SidLists"`
	SidLists    []Srv6SidList
}

func (m *SrPoliciesDetails) Reset()                        { *m = SrPoliciesDetails{} }
func (*SrPoliciesDetails) GetMessageName() string          { return "sr_policies_details" }
func (*SrPoliciesDetails) GetCrcString() string            { return "5087f460" }
func (*SrPoliciesDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// SrPoliciesDump represents VPP binary API message 'sr_policies_dump'.
type SrPoliciesDump struct{}

func (m *SrPoliciesDump) Reset()                        { *m = SrPoliciesDump{} }
func (*SrPoliciesDump) GetMessageName() string          { return "sr_policies_dump" }
func (*SrPoliciesDump) GetCrcString() string            { return "51077d14" }
func (*SrPoliciesDump) GetMessageType() api.MessageType { return api.RequestMessage }

// SrPolicyAdd represents VPP binary API message 'sr_policy_add'.
type SrPolicyAdd struct {
	BsidAddr []byte `struc:"[16]byte"`
	Weight   uint32
	IsEncap  uint8
	Type     uint8
	FibTable uint32
	Sids     Srv6SidList
}

func (m *SrPolicyAdd) Reset()                        { *m = SrPolicyAdd{} }
func (*SrPolicyAdd) GetMessageName() string          { return "sr_policy_add" }
func (*SrPolicyAdd) GetCrcString() string            { return "4b6e2484" }
func (*SrPolicyAdd) GetMessageType() api.MessageType { return api.RequestMessage }

// SrPolicyAddReply represents VPP binary API message 'sr_policy_add_reply'.
type SrPolicyAddReply struct {
	Retval int32
}

func (m *SrPolicyAddReply) Reset()                        { *m = SrPolicyAddReply{} }
func (*SrPolicyAddReply) GetMessageName() string          { return "sr_policy_add_reply" }
func (*SrPolicyAddReply) GetCrcString() string            { return "e8d4e804" }
func (*SrPolicyAddReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SrPolicyDel represents VPP binary API message 'sr_policy_del'.
type SrPolicyDel struct {
	BsidAddr      Srv6Sid
	SrPolicyIndex uint32
}

func (m *SrPolicyDel) Reset()                        { *m = SrPolicyDel{} }
func (*SrPolicyDel) GetMessageName() string          { return "sr_policy_del" }
func (*SrPolicyDel) GetCrcString() string            { return "e4133171" }
func (*SrPolicyDel) GetMessageType() api.MessageType { return api.RequestMessage }

// SrPolicyDelReply represents VPP binary API message 'sr_policy_del_reply'.
type SrPolicyDelReply struct {
	Retval int32
}

func (m *SrPolicyDelReply) Reset()                        { *m = SrPolicyDelReply{} }
func (*SrPolicyDelReply) GetMessageName() string          { return "sr_policy_del_reply" }
func (*SrPolicyDelReply) GetCrcString() string            { return "e8d4e804" }
func (*SrPolicyDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SrPolicyMod represents VPP binary API message 'sr_policy_mod'.
type SrPolicyMod struct {
	BsidAddr      []byte `struc:"[16]byte"`
	SrPolicyIndex uint32
	FibTable      uint32
	Operation     uint8
	SlIndex       uint32
	Weight        uint32
	Sids          Srv6SidList
}

func (m *SrPolicyMod) Reset()                        { *m = SrPolicyMod{} }
func (*SrPolicyMod) GetMessageName() string          { return "sr_policy_mod" }
func (*SrPolicyMod) GetCrcString() string            { return "c1dfaee0" }
func (*SrPolicyMod) GetMessageType() api.MessageType { return api.RequestMessage }

// SrPolicyModReply represents VPP binary API message 'sr_policy_mod_reply'.
type SrPolicyModReply struct {
	Retval int32
}

func (m *SrPolicyModReply) Reset()                        { *m = SrPolicyModReply{} }
func (*SrPolicyModReply) GetMessageName() string          { return "sr_policy_mod_reply" }
func (*SrPolicyModReply) GetCrcString() string            { return "e8d4e804" }
func (*SrPolicyModReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SrSetEncapHopLimit represents VPP binary API message 'sr_set_encap_hop_limit'.
type SrSetEncapHopLimit struct {
	HopLimit uint8
}

func (m *SrSetEncapHopLimit) Reset()                        { *m = SrSetEncapHopLimit{} }
func (*SrSetEncapHopLimit) GetMessageName() string          { return "sr_set_encap_hop_limit" }
func (*SrSetEncapHopLimit) GetCrcString() string            { return "aa75d7d0" }
func (*SrSetEncapHopLimit) GetMessageType() api.MessageType { return api.RequestMessage }

// SrSetEncapHopLimitReply represents VPP binary API message 'sr_set_encap_hop_limit_reply'.
type SrSetEncapHopLimitReply struct {
	Retval int32
}

func (m *SrSetEncapHopLimitReply) Reset()                        { *m = SrSetEncapHopLimitReply{} }
func (*SrSetEncapHopLimitReply) GetMessageName() string          { return "sr_set_encap_hop_limit_reply" }
func (*SrSetEncapHopLimitReply) GetCrcString() string            { return "e8d4e804" }
func (*SrSetEncapHopLimitReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SrSetEncapSource represents VPP binary API message 'sr_set_encap_source'.
type SrSetEncapSource struct {
	EncapsSource []byte `struc:"[16]byte"`
}

func (m *SrSetEncapSource) Reset()                        { *m = SrSetEncapSource{} }
func (*SrSetEncapSource) GetMessageName() string          { return "sr_set_encap_source" }
func (*SrSetEncapSource) GetCrcString() string            { return "d05bb4de" }
func (*SrSetEncapSource) GetMessageType() api.MessageType { return api.RequestMessage }

// SrSetEncapSourceReply represents VPP binary API message 'sr_set_encap_source_reply'.
type SrSetEncapSourceReply struct {
	Retval int32
}

func (m *SrSetEncapSourceReply) Reset()                        { *m = SrSetEncapSourceReply{} }
func (*SrSetEncapSourceReply) GetMessageName() string          { return "sr_set_encap_source_reply" }
func (*SrSetEncapSourceReply) GetCrcString() string            { return "e8d4e804" }
func (*SrSetEncapSourceReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SrSteeringAddDel represents VPP binary API message 'sr_steering_add_del'.
type SrSteeringAddDel struct {
	IsDel         uint8
	BsidAddr      []byte `struc:"[16]byte"`
	SrPolicyIndex uint32
	TableID       uint32
	PrefixAddr    []byte `struc:"[16]byte"`
	MaskWidth     uint32
	SwIfIndex     uint32
	TrafficType   uint8
}

func (m *SrSteeringAddDel) Reset()                        { *m = SrSteeringAddDel{} }
func (*SrSteeringAddDel) GetMessageName() string          { return "sr_steering_add_del" }
func (*SrSteeringAddDel) GetCrcString() string            { return "28b5dcab" }
func (*SrSteeringAddDel) GetMessageType() api.MessageType { return api.RequestMessage }

// SrSteeringAddDelReply represents VPP binary API message 'sr_steering_add_del_reply'.
type SrSteeringAddDelReply struct {
	Retval int32
}

func (m *SrSteeringAddDelReply) Reset()                        { *m = SrSteeringAddDelReply{} }
func (*SrSteeringAddDelReply) GetMessageName() string          { return "sr_steering_add_del_reply" }
func (*SrSteeringAddDelReply) GetCrcString() string            { return "e8d4e804" }
func (*SrSteeringAddDelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SrSteeringPolDetails represents VPP binary API message 'sr_steering_pol_details'.
type SrSteeringPolDetails struct {
	TrafficType uint8
	FibTable    uint32
	PrefixAddr  []byte `struc:"[16]byte"`
	MaskWidth   uint32
	SwIfIndex   uint32
	Bsid        Srv6Sid
}

func (m *SrSteeringPolDetails) Reset()                        { *m = SrSteeringPolDetails{} }
func (*SrSteeringPolDetails) GetMessageName() string          { return "sr_steering_pol_details" }
func (*SrSteeringPolDetails) GetCrcString() string            { return "5627d483" }
func (*SrSteeringPolDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// SrSteeringPolDump represents VPP binary API message 'sr_steering_pol_dump'.
type SrSteeringPolDump struct{}

func (m *SrSteeringPolDump) Reset()                        { *m = SrSteeringPolDump{} }
func (*SrSteeringPolDump) GetMessageName() string          { return "sr_steering_pol_dump" }
func (*SrSteeringPolDump) GetCrcString() string            { return "51077d14" }
func (*SrSteeringPolDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*SrLocalsidAddDel)(nil), "sr.SrLocalsidAddDel")
	api.RegisterMessage((*SrLocalsidAddDelReply)(nil), "sr.SrLocalsidAddDelReply")
	api.RegisterMessage((*SrLocalsidsDetails)(nil), "sr.SrLocalsidsDetails")
	api.RegisterMessage((*SrLocalsidsDump)(nil), "sr.SrLocalsidsDump")
	api.RegisterMessage((*SrPoliciesDetails)(nil), "sr.SrPoliciesDetails")
	api.RegisterMessage((*SrPoliciesDump)(nil), "sr.SrPoliciesDump")
	api.RegisterMessage((*SrPolicyAdd)(nil), "sr.SrPolicyAdd")
	api.RegisterMessage((*SrPolicyAddReply)(nil), "sr.SrPolicyAddReply")
	api.RegisterMessage((*SrPolicyDel)(nil), "sr.SrPolicyDel")
	api.RegisterMessage((*SrPolicyDelReply)(nil), "sr.SrPolicyDelReply")
	api.RegisterMessage((*SrPolicyMod)(nil), "sr.SrPolicyMod")
	api.RegisterMessage((*SrPolicyModReply)(nil), "sr.SrPolicyModReply")
	api.RegisterMessage((*SrSetEncapHopLimit)(nil), "sr.SrSetEncapHopLimit")
	api.RegisterMessage((*SrSetEncapHopLimitReply)(nil), "sr.SrSetEncapHopLimitReply")
	api.RegisterMessage((*SrSetEncapSource)(nil), "sr.SrSetEncapSource")
	api.RegisterMessage((*SrSetEncapSourceReply)(nil), "sr.SrSetEncapSourceReply")
	api.RegisterMessage((*SrSteeringAddDel)(nil), "sr.SrSteeringAddDel")
	api.RegisterMessage((*SrSteeringAddDelReply)(nil), "sr.SrSteeringAddDelReply")
	api.RegisterMessage((*SrSteeringPolDetails)(nil), "sr.SrSteeringPolDetails")
	api.RegisterMessage((*SrSteeringPolDump)(nil), "sr.SrSteeringPolDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SrLocalsidAddDel)(nil),
		(*SrLocalsidAddDelReply)(nil),
		(*SrLocalsidsDetails)(nil),
		(*SrLocalsidsDump)(nil),
		(*SrPoliciesDetails)(nil),
		(*SrPoliciesDump)(nil),
		(*SrPolicyAdd)(nil),
		(*SrPolicyAddReply)(nil),
		(*SrPolicyDel)(nil),
		(*SrPolicyDelReply)(nil),
		(*SrPolicyMod)(nil),
		(*SrPolicyModReply)(nil),
		(*SrSetEncapHopLimit)(nil),
		(*SrSetEncapHopLimitReply)(nil),
		(*SrSetEncapSource)(nil),
		(*SrSetEncapSourceReply)(nil),
		(*SrSteeringAddDel)(nil),
		(*SrSteeringAddDelReply)(nil),
		(*SrSteeringPolDetails)(nil),
		(*SrSteeringPolDump)(nil),
	}
}

// RPCService represents RPC service API for sr module.
type RPCService interface {
	DumpSrLocalsids(ctx context.Context, in *SrLocalsidsDump) (RPCService_DumpSrLocalsidsClient, error)
	DumpSrPolicies(ctx context.Context, in *SrPoliciesDump) (RPCService_DumpSrPoliciesClient, error)
	DumpSrSteeringPol(ctx context.Context, in *SrSteeringPolDump) (RPCService_DumpSrSteeringPolClient, error)
	SrLocalsidAddDel(ctx context.Context, in *SrLocalsidAddDel) (*SrLocalsidAddDelReply, error)
	SrPolicyAdd(ctx context.Context, in *SrPolicyAdd) (*SrPolicyAddReply, error)
	SrPolicyDel(ctx context.Context, in *SrPolicyDel) (*SrPolicyDelReply, error)
	SrPolicyMod(ctx context.Context, in *SrPolicyMod) (*SrPolicyModReply, error)
	SrSetEncapHopLimit(ctx context.Context, in *SrSetEncapHopLimit) (*SrSetEncapHopLimitReply, error)
	SrSetEncapSource(ctx context.Context, in *SrSetEncapSource) (*SrSetEncapSourceReply, error)
	SrSteeringAddDel(ctx context.Context, in *SrSteeringAddDel) (*SrSteeringAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpSrLocalsids(ctx context.Context, in *SrLocalsidsDump) (RPCService_DumpSrLocalsidsClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpSrLocalsidsClient{stream}
	return x, nil
}

type RPCService_DumpSrLocalsidsClient interface {
	Recv() (*SrLocalsidsDetails, error)
}

type serviceClient_DumpSrLocalsidsClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpSrLocalsidsClient) Recv() (*SrLocalsidsDetails, error) {
	m := new(SrLocalsidsDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpSrPolicies(ctx context.Context, in *SrPoliciesDump) (RPCService_DumpSrPoliciesClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpSrPoliciesClient{stream}
	return x, nil
}

type RPCService_DumpSrPoliciesClient interface {
	Recv() (*SrPoliciesDetails, error)
}

type serviceClient_DumpSrPoliciesClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpSrPoliciesClient) Recv() (*SrPoliciesDetails, error) {
	m := new(SrPoliciesDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpSrSteeringPol(ctx context.Context, in *SrSteeringPolDump) (RPCService_DumpSrSteeringPolClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpSrSteeringPolClient{stream}
	return x, nil
}

type RPCService_DumpSrSteeringPolClient interface {
	Recv() (*SrSteeringPolDetails, error)
}

type serviceClient_DumpSrSteeringPolClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpSrSteeringPolClient) Recv() (*SrSteeringPolDetails, error) {
	m := new(SrSteeringPolDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) SrLocalsidAddDel(ctx context.Context, in *SrLocalsidAddDel) (*SrLocalsidAddDelReply, error) {
	out := new(SrLocalsidAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SrPolicyAdd(ctx context.Context, in *SrPolicyAdd) (*SrPolicyAddReply, error) {
	out := new(SrPolicyAddReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SrPolicyDel(ctx context.Context, in *SrPolicyDel) (*SrPolicyDelReply, error) {
	out := new(SrPolicyDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SrPolicyMod(ctx context.Context, in *SrPolicyMod) (*SrPolicyModReply, error) {
	out := new(SrPolicyModReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SrSetEncapHopLimit(ctx context.Context, in *SrSetEncapHopLimit) (*SrSetEncapHopLimitReply, error) {
	out := new(SrSetEncapHopLimitReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SrSetEncapSource(ctx context.Context, in *SrSetEncapSource) (*SrSetEncapSourceReply, error) {
	out := new(SrSetEncapSourceReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SrSteeringAddDel(ctx context.Context, in *SrSteeringAddDel) (*SrSteeringAddDelReply, error) {
	out := new(SrSteeringAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
