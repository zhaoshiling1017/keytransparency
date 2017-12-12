// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/v1/keytransparency_proto/admin.proto

package keytransparency_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf4 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf2 "github.com/golang/protobuf/ptypes/duration"
import trillian "github.com/google/trillian"
import keyspb "github.com/google/trillian/crypto/keyspb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Domain contains information on a single domain
type Domain struct {
	// DomainId can be any URL safe string.
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	// Log contains the Log-Tree's info.
	Log *trillian.Tree `protobuf:"bytes,2,opt,name=log" json:"log,omitempty"`
	// Map contains the Map-Tree's info.
	Map *trillian.Tree `protobuf:"bytes,3,opt,name=map" json:"map,omitempty"`
	// Vrf contains the VRF public key.
	Vrf *keyspb.PublicKey `protobuf:"bytes,4,opt,name=vrf" json:"vrf,omitempty"`
	// min_interval is the minimum time between epochs.
	MinInterval *google_protobuf2.Duration `protobuf:"bytes,5,opt,name=min_interval,json=minInterval" json:"min_interval,omitempty"`
	// max_interval is the maximum time between epochs.
	MaxInterval *google_protobuf2.Duration `protobuf:"bytes,6,opt,name=max_interval,json=maxInterval" json:"max_interval,omitempty"`
	// Deleted indicates whether the domain has been marked as deleted.
	// By its presence in a response, this domain has not been garbage collected.
	Deleted bool `protobuf:"varint,7,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *Domain) Reset()                    { *m = Domain{} }
func (m *Domain) String() string            { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()               {}
func (*Domain) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Domain) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *Domain) GetLog() *trillian.Tree {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Domain) GetMap() *trillian.Tree {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *Domain) GetVrf() *keyspb.PublicKey {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *Domain) GetMinInterval() *google_protobuf2.Duration {
	if m != nil {
		return m.MinInterval
	}
	return nil
}

func (m *Domain) GetMaxInterval() *google_protobuf2.Duration {
	if m != nil {
		return m.MaxInterval
	}
	return nil
}

func (m *Domain) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

// ListDomains request.
// No pagination options are provided.
type ListDomainsRequest struct {
	// showDeleted requests domains that have been marked for deletion
	// but have not been garbage collected.
	ShowDeleted bool `protobuf:"varint,1,opt,name=show_deleted,json=showDeleted" json:"show_deleted,omitempty"`
}

func (m *ListDomainsRequest) Reset()                    { *m = ListDomainsRequest{} }
func (m *ListDomainsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDomainsRequest) ProtoMessage()               {}
func (*ListDomainsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ListDomainsRequest) GetShowDeleted() bool {
	if m != nil {
		return m.ShowDeleted
	}
	return false
}

// ListDomains response contains domains.
type ListDomainsResponse struct {
	Domains []*Domain `protobuf:"bytes,1,rep,name=domains" json:"domains,omitempty"`
}

func (m *ListDomainsResponse) Reset()                    { *m = ListDomainsResponse{} }
func (m *ListDomainsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDomainsResponse) ProtoMessage()               {}
func (*ListDomainsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ListDomainsResponse) GetDomains() []*Domain {
	if m != nil {
		return m.Domains
	}
	return nil
}

// GetDomainRequest specifies the domain to retrieve information for.
type GetDomainRequest struct {
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	// showDeleted requests domains that have been marked for deletion
	// but have not been garbage collected.
	ShowDeleted bool `protobuf:"varint,2,opt,name=show_deleted,json=showDeleted" json:"show_deleted,omitempty"`
}

func (m *GetDomainRequest) Reset()                    { *m = GetDomainRequest{} }
func (m *GetDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDomainRequest) ProtoMessage()               {}
func (*GetDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetDomainRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *GetDomainRequest) GetShowDeleted() bool {
	if m != nil {
		return m.ShowDeleted
	}
	return false
}

// GetDomainResponse contains the configuration info for one domain.
type GetDomainResponse struct {
	Domain *Domain `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
}

func (m *GetDomainResponse) Reset()                    { *m = GetDomainResponse{} }
func (m *GetDomainResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDomainResponse) ProtoMessage()               {}
func (*GetDomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *GetDomainResponse) GetDomain() *Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

// CreateDomainRequest creates a new domain
type CreateDomainRequest struct {
	DomainId    string                     `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	MinInterval *google_protobuf2.Duration `protobuf:"bytes,2,opt,name=min_interval,json=minInterval" json:"min_interval,omitempty"`
	MaxInterval *google_protobuf2.Duration `protobuf:"bytes,3,opt,name=max_interval,json=maxInterval" json:"max_interval,omitempty"`
}

func (m *CreateDomainRequest) Reset()                    { *m = CreateDomainRequest{} }
func (m *CreateDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDomainRequest) ProtoMessage()               {}
func (*CreateDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *CreateDomainRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *CreateDomainRequest) GetMinInterval() *google_protobuf2.Duration {
	if m != nil {
		return m.MinInterval
	}
	return nil
}

func (m *CreateDomainRequest) GetMaxInterval() *google_protobuf2.Duration {
	if m != nil {
		return m.MaxInterval
	}
	return nil
}

// CreateDomainResponse contains the configuration info for the new domain.
type CreateDomainResponse struct {
	Domain *Domain `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
}

func (m *CreateDomainResponse) Reset()                    { *m = CreateDomainResponse{} }
func (m *CreateDomainResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateDomainResponse) ProtoMessage()               {}
func (*CreateDomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *CreateDomainResponse) GetDomain() *Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

// DeleteDomainRequest deletes a domain
type DeleteDomainRequest struct {
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
}

func (m *DeleteDomainRequest) Reset()                    { *m = DeleteDomainRequest{} }
func (m *DeleteDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDomainRequest) ProtoMessage()               {}
func (*DeleteDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *DeleteDomainRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

// UndeleteDomainRequest deletes a domain
type UndeleteDomainRequest struct {
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
}

func (m *UndeleteDomainRequest) Reset()                    { *m = UndeleteDomainRequest{} }
func (m *UndeleteDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*UndeleteDomainRequest) ProtoMessage()               {}
func (*UndeleteDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *UndeleteDomainRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func init() {
	proto.RegisterType((*Domain)(nil), "google.keytransparency.v1.Domain")
	proto.RegisterType((*ListDomainsRequest)(nil), "google.keytransparency.v1.ListDomainsRequest")
	proto.RegisterType((*ListDomainsResponse)(nil), "google.keytransparency.v1.ListDomainsResponse")
	proto.RegisterType((*GetDomainRequest)(nil), "google.keytransparency.v1.GetDomainRequest")
	proto.RegisterType((*GetDomainResponse)(nil), "google.keytransparency.v1.GetDomainResponse")
	proto.RegisterType((*CreateDomainRequest)(nil), "google.keytransparency.v1.CreateDomainRequest")
	proto.RegisterType((*CreateDomainResponse)(nil), "google.keytransparency.v1.CreateDomainResponse")
	proto.RegisterType((*DeleteDomainRequest)(nil), "google.keytransparency.v1.DeleteDomainRequest")
	proto.RegisterType((*UndeleteDomainRequest)(nil), "google.keytransparency.v1.UndeleteDomainRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KeyTransparencyAdminService service

type KeyTransparencyAdminServiceClient interface {
	// ListDomains returns a list of all domains this Key Transparency server
	// operates on.
	ListDomains(ctx context.Context, in *ListDomainsRequest, opts ...grpc.CallOption) (*ListDomainsResponse, error)
	// GetDomain returns the confiuration information for a given domain.
	GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*GetDomainResponse, error)
	// CreateDomain creates a new Trillian log/map pair.  A unique domainId must
	// be provided.  To create a new domain with the same name as a previously
	// deleted domain, a user must wait X days until the domain is garbage
	// collected.
	CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*CreateDomainResponse, error)
	// DeleteDomain marks a domain as deleted.  Domains will be garbage collected
	// after X days.
	DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error)
	// UndeleteDomain marks a previously deleted domain as active if it has not
	// already been garbage collected.
	UndeleteDomain(ctx context.Context, in *UndeleteDomainRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error)
}

type keyTransparencyAdminServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeyTransparencyAdminServiceClient(cc *grpc.ClientConn) KeyTransparencyAdminServiceClient {
	return &keyTransparencyAdminServiceClient{cc}
}

func (c *keyTransparencyAdminServiceClient) ListDomains(ctx context.Context, in *ListDomainsRequest, opts ...grpc.CallOption) (*ListDomainsResponse, error) {
	out := new(ListDomainsResponse)
	err := grpc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdminService/ListDomains", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminServiceClient) GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*GetDomainResponse, error) {
	out := new(GetDomainResponse)
	err := grpc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdminService/GetDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminServiceClient) CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*CreateDomainResponse, error) {
	out := new(CreateDomainResponse)
	err := grpc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdminService/CreateDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminServiceClient) DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error) {
	out := new(google_protobuf4.Empty)
	err := grpc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdminService/DeleteDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminServiceClient) UndeleteDomain(ctx context.Context, in *UndeleteDomainRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error) {
	out := new(google_protobuf4.Empty)
	err := grpc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdminService/UndeleteDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KeyTransparencyAdminService service

type KeyTransparencyAdminServiceServer interface {
	// ListDomains returns a list of all domains this Key Transparency server
	// operates on.
	ListDomains(context.Context, *ListDomainsRequest) (*ListDomainsResponse, error)
	// GetDomain returns the confiuration information for a given domain.
	GetDomain(context.Context, *GetDomainRequest) (*GetDomainResponse, error)
	// CreateDomain creates a new Trillian log/map pair.  A unique domainId must
	// be provided.  To create a new domain with the same name as a previously
	// deleted domain, a user must wait X days until the domain is garbage
	// collected.
	CreateDomain(context.Context, *CreateDomainRequest) (*CreateDomainResponse, error)
	// DeleteDomain marks a domain as deleted.  Domains will be garbage collected
	// after X days.
	DeleteDomain(context.Context, *DeleteDomainRequest) (*google_protobuf4.Empty, error)
	// UndeleteDomain marks a previously deleted domain as active if it has not
	// already been garbage collected.
	UndeleteDomain(context.Context, *UndeleteDomainRequest) (*google_protobuf4.Empty, error)
}

func RegisterKeyTransparencyAdminServiceServer(s *grpc.Server, srv KeyTransparencyAdminServiceServer) {
	s.RegisterService(&_KeyTransparencyAdminService_serviceDesc, srv)
}

func _KeyTransparencyAdminService_ListDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServiceServer).ListDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdminService/ListDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServiceServer).ListDomains(ctx, req.(*ListDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdminService_GetDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServiceServer).GetDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdminService/GetDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServiceServer).GetDomain(ctx, req.(*GetDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdminService_CreateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServiceServer).CreateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdminService/CreateDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServiceServer).CreateDomain(ctx, req.(*CreateDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdminService_DeleteDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServiceServer).DeleteDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdminService/DeleteDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServiceServer).DeleteDomain(ctx, req.(*DeleteDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdminService_UndeleteDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UndeleteDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServiceServer).UndeleteDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdminService/UndeleteDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServiceServer).UndeleteDomain(ctx, req.(*UndeleteDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyTransparencyAdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.keytransparency.v1.KeyTransparencyAdminService",
	HandlerType: (*KeyTransparencyAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDomains",
			Handler:    _KeyTransparencyAdminService_ListDomains_Handler,
		},
		{
			MethodName: "GetDomain",
			Handler:    _KeyTransparencyAdminService_GetDomain_Handler,
		},
		{
			MethodName: "CreateDomain",
			Handler:    _KeyTransparencyAdminService_CreateDomain_Handler,
		},
		{
			MethodName: "DeleteDomain",
			Handler:    _KeyTransparencyAdminService_DeleteDomain_Handler,
		},
		{
			MethodName: "UndeleteDomain",
			Handler:    _KeyTransparencyAdminService_UndeleteDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/keytransparency_proto/admin.proto",
}

func init() { proto.RegisterFile("api/v1/keytransparency_proto/admin.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xc1, 0x4e, 0xd4, 0x40,
	0x18, 0x4e, 0x77, 0x65, 0x81, 0xd9, 0x95, 0xc0, 0xac, 0x62, 0x29, 0x46, 0x4b, 0xbd, 0x34, 0xa8,
	0xad, 0xac, 0x26, 0x46, 0xf4, 0xa2, 0x40, 0x0c, 0xc1, 0x18, 0xad, 0x78, 0xf1, 0x42, 0x66, 0xdb,
	0x61, 0x99, 0xd8, 0x76, 0xea, 0x74, 0xba, 0xd0, 0x18, 0x0f, 0x1a, 0x13, 0x1f, 0xc0, 0x57, 0xf0,
	0xec, 0xcd, 0x27, 0xf1, 0x15, 0x7c, 0x10, 0xd3, 0xe9, 0x74, 0x2d, 0x65, 0xb7, 0x2e, 0xe1, 0xd4,
	0x4e, 0xff, 0xff, 0xfb, 0xbf, 0x6f, 0xbe, 0xf9, 0xd2, 0x01, 0x26, 0x8a, 0x88, 0x3d, 0xdc, 0xb0,
	0xdf, 0xe3, 0x94, 0x33, 0x14, 0xc6, 0x11, 0x62, 0x38, 0x74, 0xd3, 0x83, 0x88, 0x51, 0x4e, 0x6d,
	0xe4, 0x05, 0x24, 0xb4, 0xc4, 0x3b, 0x5c, 0x19, 0x50, 0x3a, 0xf0, 0xb1, 0x55, 0xe9, 0xb4, 0x86,
	0x1b, 0xda, 0xf5, 0xbc, 0x64, 0x67, 0xb3, 0x50, 0x18, 0x52, 0x8e, 0x38, 0xa1, 0x61, 0x9c, 0x03,
	0xb5, 0x55, 0x59, 0x15, 0xab, 0x7e, 0x72, 0x68, 0xe3, 0x20, 0xe2, 0xa9, 0x2c, 0xde, 0xa8, 0x16,
	0xbd, 0x84, 0x09, 0xb4, 0xac, 0x2f, 0x70, 0x46, 0x7c, 0x9f, 0xa0, 0x62, 0xad, 0xb9, 0x2c, 0x8d,
	0x38, 0xcd, 0xf4, 0xc6, 0x51, 0x5f, 0x3e, 0xf2, 0x9a, 0xf1, 0xa3, 0x01, 0x5a, 0xdb, 0x34, 0x40,
	0x24, 0x84, 0xab, 0x60, 0xde, 0x13, 0x6f, 0x07, 0xc4, 0x53, 0x15, 0x5d, 0x31, 0xe7, 0x9d, 0xb9,
	0xfc, 0xc3, 0xae, 0x07, 0x75, 0xd0, 0xf4, 0xe9, 0x40, 0x6d, 0xe8, 0x8a, 0xd9, 0xee, 0x2d, 0x58,
	0x23, 0x86, 0x7d, 0x86, 0xb1, 0x93, 0x95, 0xb2, 0x8e, 0x00, 0x45, 0x6a, 0x73, 0x7c, 0x47, 0x80,
	0x22, 0x78, 0x0b, 0x34, 0x87, 0xec, 0x50, 0xbd, 0x24, 0x3a, 0x96, 0x2c, 0xa9, 0xe3, 0x55, 0xd2,
	0xf7, 0x89, 0xbb, 0x87, 0x53, 0x27, 0xab, 0xc2, 0x27, 0xa0, 0x13, 0x64, 0x12, 0x42, 0x8e, 0xd9,
	0x10, 0xf9, 0xea, 0x8c, 0xe8, 0x5e, 0xb1, 0xa4, 0x93, 0xc5, 0x9e, 0xad, 0x6d, 0xb9, 0x67, 0xa7,
	0x1d, 0x90, 0x70, 0x57, 0x76, 0x0b, 0x34, 0x3a, 0xf9, 0x87, 0x6e, 0xfd, 0x1f, 0x8d, 0x4e, 0x46,
	0x68, 0x15, 0xcc, 0x7a, 0xd8, 0xc7, 0x1c, 0x7b, 0xea, 0xac, 0xae, 0x98, 0x73, 0x4e, 0xb1, 0x34,
	0x1e, 0x02, 0xf8, 0x82, 0xc4, 0x3c, 0x77, 0x2a, 0x76, 0xf0, 0x87, 0x04, 0xc7, 0x1c, 0xae, 0x81,
	0x4e, 0x7c, 0x44, 0x8f, 0x0f, 0x0a, 0x90, 0x22, 0x40, 0xed, 0xec, 0xdb, 0xb6, 0x04, 0x3a, 0xa0,
	0x7b, 0x0a, 0x18, 0x47, 0x34, 0x8c, 0x31, 0x7c, 0x0c, 0x66, 0x73, 0x6b, 0x63, 0x55, 0xd1, 0x9b,
	0x66, 0xbb, 0xb7, 0x66, 0x4d, 0x8c, 0x8a, 0x95, 0x83, 0x9d, 0x02, 0x61, 0x38, 0x60, 0xf1, 0x39,
	0x96, 0x23, 0x0b, 0x29, 0xb5, 0x87, 0x57, 0xd5, 0xd9, 0x38, 0xab, 0xf3, 0x25, 0x58, 0x2a, 0xcd,
	0x94, 0x2a, 0x1f, 0x81, 0x56, 0x3e, 0x43, 0x4c, 0x9c, 0x4a, 0xa4, 0x04, 0x18, 0x3f, 0x15, 0xd0,
	0xdd, 0x62, 0x18, 0x71, 0x7c, 0x0e, 0x9d, 0xd5, 0xb3, 0x6f, 0x5c, 0xe8, 0xec, 0x9b, 0xe7, 0x39,
	0x7b, 0xe3, 0x35, 0xb8, 0x72, 0x5a, 0xef, 0xc5, 0x3d, 0xe8, 0x81, 0x6e, 0x6e, 0xef, 0xf4, 0x16,
	0x18, 0x0f, 0xc0, 0xd5, 0xb7, 0xa1, 0x77, 0x4e, 0x54, 0xef, 0xd7, 0x0c, 0x58, 0xdd, 0xc3, 0xe9,
	0x7e, 0x49, 0xcf, 0xd3, 0xec, 0x37, 0xf4, 0x06, 0xb3, 0x21, 0x71, 0x31, 0xfc, 0xac, 0x80, 0x76,
	0x29, 0x86, 0xf0, 0x6e, 0xcd, 0x26, 0xce, 0xe6, 0x5c, 0xb3, 0xa6, 0x6d, 0xcf, 0x3d, 0x33, 0xba,
	0x5f, 0x7e, 0xff, 0xf9, 0xde, 0xb8, 0x0c, 0xdb, 0xd9, 0x5f, 0x52, 0xa6, 0x16, 0x7e, 0x53, 0xc0,
	0xfc, 0x28, 0x62, 0xf0, 0x76, 0xcd, 0xc8, 0x6a, 0xb8, 0xb5, 0x3b, 0xd3, 0x35, 0x4b, 0xf6, 0x9b,
	0x82, 0x7d, 0x05, 0x5e, 0x2b, 0xb1, 0xdb, 0x1f, 0x47, 0xe6, 0x7d, 0xca, 0x94, 0x74, 0xca, 0x67,
	0x0d, 0xeb, 0xf6, 0x37, 0x26, 0xc4, 0x9a, 0x3d, 0x75, 0xbf, 0x94, 0xb4, 0x2c, 0x24, 0x2d, 0x1a,
	0x65, 0x43, 0x36, 0x95, 0x75, 0x78, 0x0c, 0x3a, 0xe5, 0x84, 0xd4, 0x0a, 0x19, 0x13, 0x25, 0x6d,
	0xf9, 0x4c, 0xb8, 0x77, 0xb2, 0x7b, 0xa2, 0xb0, 0x60, 0x7d, 0xa2, 0x05, 0x5f, 0x15, 0xb0, 0x70,
	0x3a, 0x67, 0xf0, 0x5e, 0x0d, 0xf7, 0xd8, 0x48, 0x4e, 0x64, 0x37, 0x05, 0xbb, 0xb1, 0xae, 0x4f,
	0x60, 0xdf, 0x4c, 0xe4, 0xb8, 0x67, 0x3b, 0xef, 0xb6, 0x06, 0x84, 0x1f, 0x25, 0x7d, 0xcb, 0xa5,
	0x81, 0x2d, 0xaf, 0xb5, 0x0a, 0xbf, 0xed, 0x52, 0x96, 0x5f, 0x93, 0x93, 0xae, 0xdc, 0x7e, 0x4b,
	0x3c, 0xee, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x97, 0xfa, 0xce, 0x99, 0x07, 0x00, 0x00,
}