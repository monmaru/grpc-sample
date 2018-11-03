// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ght.proto

package ght

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the language name.
type Req struct {
	Lang                 string   `protobuf:"bytes,1,opt,name=lang,proto3" json:"lang,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_5eae51fa75fcc0c2, []int{0}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

// The response message containing the repositories.
type Res struct {
	Repositories         []*Repository `protobuf:"bytes,1,rep,name=repositories,proto3" json:"repositories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Res) Reset()         { *m = Res{} }
func (m *Res) String() string { return proto.CompactTextString(m) }
func (*Res) ProtoMessage()    {}
func (*Res) Descriptor() ([]byte, []int) {
	return fileDescriptor_5eae51fa75fcc0c2, []int{1}
}

func (m *Res) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Res.Unmarshal(m, b)
}
func (m *Res) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Res.Marshal(b, m, deterministic)
}
func (m *Res) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Res.Merge(m, src)
}
func (m *Res) XXX_Size() int {
	return xxx_messageInfo_Res.Size(m)
}
func (m *Res) XXX_DiscardUnknown() {
	xxx_messageInfo_Res.DiscardUnknown(m)
}

var xxx_messageInfo_Res proto.InternalMessageInfo

func (m *Res) GetRepositories() []*Repository {
	if m != nil {
		return m.Repositories
	}
	return nil
}

type Repository struct {
	Rank                 int32    `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Lang                 string   `protobuf:"bytes,6,opt,name=lang,proto3" json:"lang,omitempty"`
	Star                 *Star    `protobuf:"bytes,7,opt,name=star,proto3" json:"star,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}
func (*Repository) Descriptor() ([]byte, []int) {
	return fileDescriptor_5eae51fa75fcc0c2, []int{2}
}

func (m *Repository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repository.Unmarshal(m, b)
}
func (m *Repository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repository.Marshal(b, m, deterministic)
}
func (m *Repository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repository.Merge(m, src)
}
func (m *Repository) XXX_Size() int {
	return xxx_messageInfo_Repository.Size(m)
}
func (m *Repository) XXX_DiscardUnknown() {
	xxx_messageInfo_Repository.DiscardUnknown(m)
}

var xxx_messageInfo_Repository proto.InternalMessageInfo

func (m *Repository) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *Repository) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Repository) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Repository) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Repository) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Repository) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *Repository) GetStar() *Star {
	if m != nil {
		return m.Star
	}
	return nil
}

type Star struct {
	Today                string   `protobuf:"bytes,1,opt,name=today,proto3" json:"today,omitempty"`
	Total                string   `protobuf:"bytes,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Star) Reset()         { *m = Star{} }
func (m *Star) String() string { return proto.CompactTextString(m) }
func (*Star) ProtoMessage()    {}
func (*Star) Descriptor() ([]byte, []int) {
	return fileDescriptor_5eae51fa75fcc0c2, []int{3}
}

func (m *Star) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Star.Unmarshal(m, b)
}
func (m *Star) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Star.Marshal(b, m, deterministic)
}
func (m *Star) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Star.Merge(m, src)
}
func (m *Star) XXX_Size() int {
	return xxx_messageInfo_Star.Size(m)
}
func (m *Star) XXX_DiscardUnknown() {
	xxx_messageInfo_Star.DiscardUnknown(m)
}

var xxx_messageInfo_Star proto.InternalMessageInfo

func (m *Star) GetToday() string {
	if m != nil {
		return m.Today
	}
	return ""
}

func (m *Star) GetTotal() string {
	if m != nil {
		return m.Total
	}
	return ""
}

func init() {
	proto.RegisterType((*Req)(nil), "ght.Req")
	proto.RegisterType((*Res)(nil), "ght.Res")
	proto.RegisterType((*Repository)(nil), "ght.Repository")
	proto.RegisterType((*Star)(nil), "ght.Star")
}

func init() { proto.RegisterFile("proto/ght.proto", fileDescriptor_5eae51fa75fcc0c2) }

var fileDescriptor_5eae51fa75fcc0c2 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x51, 0x4b, 0x4e, 0xc3, 0x30,
	0x10, 0x55, 0x9a, 0xa4, 0xd0, 0x29, 0xa2, 0xc8, 0xaa, 0x90, 0xa9, 0x8a, 0x14, 0x65, 0xd5, 0x55,
	0x23, 0xa5, 0x2b, 0x38, 0x00, 0x2c, 0xd8, 0x05, 0x2e, 0xe0, 0xb6, 0x96, 0x63, 0x11, 0xd9, 0xa9,
	0x3d, 0x15, 0xaa, 0x10, 0x1b, 0xae, 0xc0, 0x51, 0x38, 0x0a, 0x57, 0xe0, 0x20, 0xc8, 0x93, 0x96,
	0xc2, 0xee, 0x7d, 0xec, 0x99, 0x79, 0x33, 0x30, 0x6a, 0x9d, 0x45, 0x5b, 0xa8, 0x1a, 0xe7, 0x84,
	0x58, 0xac, 0x6a, 0x9c, 0x4c, 0x95, 0xb5, 0xaa, 0x91, 0x85, 0x68, 0x75, 0x21, 0x8c, 0xb1, 0x28,
	0x50, 0x5b, 0xe3, 0xbb, 0x27, 0xf9, 0x15, 0xc4, 0x95, 0xdc, 0x30, 0x06, 0x49, 0x23, 0x8c, 0xe2,
	0x51, 0x16, 0xcd, 0x06, 0x15, 0xe1, 0xfc, 0x36, 0x58, 0x9e, 0x2d, 0xe0, 0xcc, 0xc9, 0xd6, 0x7a,
	0x8d, 0xd6, 0x69, 0xe9, 0x79, 0x94, 0xc5, 0xb3, 0x61, 0x39, 0x9a, 0x87, 0x36, 0xd5, 0xc1, 0xd8,
	0x55, 0xff, 0x1e, 0xe5, 0x9f, 0x11, 0xc0, 0xd1, 0x0c, 0xe5, 0x9d, 0x30, 0xcf, 0x54, 0x3e, 0xad,
	0x08, 0xb3, 0x31, 0xa4, 0xf6, 0xc5, 0x48, 0xc7, 0x7b, 0xd4, 0xb3, 0x23, 0x41, 0x45, 0x8d, 0x8d,
	0xe4, 0x71, 0xa7, 0x12, 0x61, 0x17, 0x10, 0x6f, 0x5d, 0xc3, 0x13, 0xd2, 0x02, 0x64, 0x19, 0x0c,
	0xd7, 0xd2, 0xaf, 0x9c, 0x6e, 0x43, 0x1a, 0x9e, 0x92, 0xf3, 0x57, 0xfa, 0x8d, 0xd4, 0x3f, 0x46,
	0x62, 0xd7, 0x90, 0x78, 0x14, 0x8e, 0x9f, 0x64, 0xd1, 0x6c, 0x58, 0x0e, 0x28, 0xc3, 0x23, 0x0a,
	0x57, 0x91, 0x9c, 0x97, 0x90, 0x04, 0x46, 0x43, 0xd8, 0xb5, 0xd8, 0xed, 0xd7, 0xd1, 0x91, 0x4e,
	0x45, 0xd1, 0x1c, 0x06, 0x26, 0x52, 0x3e, 0xc0, 0xf9, 0xbd, 0xc6, 0x7a, 0xbb, 0x7c, 0x72, 0xd2,
	0xac, 0xb5, 0x51, 0xec, 0x06, 0xd2, 0x3b, 0x89, 0xab, 0x9a, 0x9d, 0xee, 0x77, 0xb4, 0x99, 0x1c,
	0x90, 0xcf, 0xa7, 0xef, 0x5f, 0xdf, 0x1f, 0xbd, 0x4b, 0x36, 0xa6, 0x7b, 0xe0, 0xfe, 0x4f, 0xf1,
	0x1a, 0xc6, 0x7b, 0x5b, 0xf6, 0xe9, 0x28, 0x8b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x90,
	0x17, 0x1a, 0xca, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GithubTrendingClient is the client API for GithubTrending service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GithubTrendingClient interface {
	// Sends a repositories.
	Fetch(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type githubTrendingClient struct {
	cc *grpc.ClientConn
}

func NewGithubTrendingClient(cc *grpc.ClientConn) GithubTrendingClient {
	return &githubTrendingClient{cc}
}

func (c *githubTrendingClient) Fetch(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/ght.GithubTrending/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GithubTrendingServer is the server API for GithubTrending service.
type GithubTrendingServer interface {
	// Sends a repositories.
	Fetch(context.Context, *Req) (*Res, error)
}

func RegisterGithubTrendingServer(s *grpc.Server, srv GithubTrendingServer) {
	s.RegisterService(&_GithubTrending_serviceDesc, srv)
}

func _GithubTrending_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubTrendingServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ght.GithubTrending/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubTrendingServer).Fetch(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _GithubTrending_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ght.GithubTrending",
	HandlerType: (*GithubTrendingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _GithubTrending_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ght.proto",
}
