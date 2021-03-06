// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Words
	WordSlice
	Empty
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Words struct {
	Count int32    `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Word  []string `protobuf:"bytes,2,rep,name=word" json:"word,omitempty"`
}

func (m *Words) Reset()                    { *m = Words{} }
func (m *Words) String() string            { return proto.CompactTextString(m) }
func (*Words) ProtoMessage()               {}
func (*Words) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Words) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Words) GetWord() []string {
	if m != nil {
		return m.Word
	}
	return nil
}

type WordSlice struct {
	Slice []*Words `protobuf:"bytes,1,rep,name=slice" json:"slice,omitempty"`
}

func (m *WordSlice) Reset()                    { *m = WordSlice{} }
func (m *WordSlice) String() string            { return proto.CompactTextString(m) }
func (*WordSlice) ProtoMessage()               {}
func (*WordSlice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WordSlice) GetSlice() []*Words {
	if m != nil {
		return m.Slice
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Words)(nil), "api.Words")
	proto.RegisterType((*WordSlice)(nil), "api.WordSlice")
	proto.RegisterType((*Empty)(nil), "api.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StoredWords service

type StoredWordsClient interface {
	GetWord(ctx context.Context, in *Words, opts ...grpc.CallOption) (*Words, error)
	UpdateWords(ctx context.Context, in *Words, opts ...grpc.CallOption) (*Words, error)
	TopFive(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WordSlice, error)
}

type storedWordsClient struct {
	cc *grpc.ClientConn
}

func NewStoredWordsClient(cc *grpc.ClientConn) StoredWordsClient {
	return &storedWordsClient{cc}
}

func (c *storedWordsClient) GetWord(ctx context.Context, in *Words, opts ...grpc.CallOption) (*Words, error) {
	out := new(Words)
	err := grpc.Invoke(ctx, "/api.StoredWords/GetWord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storedWordsClient) UpdateWords(ctx context.Context, in *Words, opts ...grpc.CallOption) (*Words, error) {
	out := new(Words)
	err := grpc.Invoke(ctx, "/api.StoredWords/UpdateWords", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storedWordsClient) TopFive(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WordSlice, error) {
	out := new(WordSlice)
	err := grpc.Invoke(ctx, "/api.StoredWords/TopFive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StoredWords service

type StoredWordsServer interface {
	GetWord(context.Context, *Words) (*Words, error)
	UpdateWords(context.Context, *Words) (*Words, error)
	TopFive(context.Context, *Empty) (*WordSlice, error)
}

func RegisterStoredWordsServer(s *grpc.Server, srv StoredWordsServer) {
	s.RegisterService(&_StoredWords_serviceDesc, srv)
}

func _StoredWords_GetWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Words)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoredWordsServer).GetWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StoredWords/GetWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoredWordsServer).GetWord(ctx, req.(*Words))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoredWords_UpdateWords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Words)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoredWordsServer).UpdateWords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StoredWords/UpdateWords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoredWordsServer).UpdateWords(ctx, req.(*Words))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoredWords_TopFive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoredWordsServer).TopFive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StoredWords/TopFive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoredWordsServer).TopFive(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoredWords_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.StoredWords",
	HandlerType: (*StoredWordsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWord",
			Handler:    _StoredWords_GetWord_Handler,
		},
		{
			MethodName: "UpdateWords",
			Handler:    _StoredWords_UpdateWords_Handler,
		},
		{
			MethodName: "TopFive",
			Handler:    _StoredWords_TopFive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x3d, 0xcb, 0xc2, 0x30,
	0x10, 0xc7, 0x9b, 0xa7, 0x4f, 0x2c, 0xbd, 0x82, 0x43, 0x70, 0x28, 0x9d, 0x42, 0x1c, 0xec, 0x62,
	0xc1, 0xfa, 0x19, 0xd4, 0xbd, 0x55, 0x9c, 0x6b, 0x93, 0x21, 0xa0, 0x26, 0xa4, 0x51, 0x71, 0xf7,
	0x83, 0x4b, 0x12, 0x91, 0x4e, 0x6e, 0xbf, 0xe3, 0xfe, 0x2f, 0x77, 0x90, 0x76, 0x5a, 0x56, 0xda,
	0x28, 0xab, 0x48, 0xdc, 0x69, 0xc9, 0x56, 0x80, 0x8f, 0xca, 0xf0, 0x81, 0xcc, 0x00, 0xf7, 0xea,
	0x76, 0xb5, 0x39, 0xa2, 0xa8, 0xc4, 0x4d, 0x18, 0x08, 0x81, 0xff, 0x87, 0x32, 0x3c, 0xff, 0xa3,
	0x71, 0x99, 0x36, 0x9e, 0xd9, 0x12, 0x52, 0x67, 0x69, 0xcf, 0xb2, 0x17, 0x84, 0x02, 0x1e, 0x1c,
	0xe4, 0x88, 0xc6, 0x65, 0x56, 0x43, 0xe5, 0xf2, 0x7d, 0x62, 0x13, 0x16, 0x2c, 0x01, 0xbc, 0xb9,
	0x68, 0xfb, 0xac, 0x5f, 0x08, 0xb2, 0xd6, 0x2a, 0x23, 0x78, 0x68, 0x9c, 0x43, 0xb2, 0x13, 0xd6,
	0x31, 0x19, 0xd9, 0x8a, 0x11, 0xb3, 0x88, 0x2c, 0x20, 0x3b, 0x68, 0xde, 0x59, 0x11, 0x3c, 0xbf,
	0x84, 0xc9, 0x5e, 0xe9, 0xad, 0xbc, 0x8b, 0x8f, 0xc8, 0x97, 0x16, 0xd3, 0xaf, 0xc8, 0xdf, 0xcb,
	0xa2, 0xd3, 0xc4, 0x7f, 0xbf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xfc, 0xda, 0x59, 0x0a,
	0x01, 0x00, 0x00,
}
