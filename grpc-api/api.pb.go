// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Words
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

func init() {
	proto.RegisterType((*Words)(nil), "api.Words")
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

// Server API for StoredWords service

type StoredWordsServer interface {
	GetWord(context.Context, *Words) (*Words, error)
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

var _StoredWords_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.StoredWords",
	HandlerType: (*StoredWordsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWord",
			Handler:    _StoredWords_GetWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x32, 0xe4, 0x62, 0x0d, 0xcf,
	0x2f, 0x4a, 0x29, 0x16, 0x12, 0xe1, 0x62, 0x4d, 0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0d, 0x82, 0x70, 0x84, 0x84, 0xb8, 0x58, 0xca, 0xf3, 0x8b, 0x52, 0x24, 0x98, 0x14,
	0x98, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x23, 0x23, 0x2e, 0xee, 0xe0, 0x92, 0xfc, 0xa2, 0xd4, 0x14,
	0x88, 0x46, 0x65, 0x2e, 0x76, 0xf7, 0xd4, 0x12, 0x10, 0x5b, 0x88, 0x4b, 0x0f, 0x64, 0x3a, 0x58,
	0x58, 0x0a, 0x89, 0xad, 0xc4, 0x90, 0xc4, 0x06, 0xb6, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x6c, 0x45, 0x0b, 0x4d, 0x7f, 0x00, 0x00, 0x00,
}