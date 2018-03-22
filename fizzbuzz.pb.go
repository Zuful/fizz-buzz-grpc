// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fizzbuzz.proto

/*
Package fizzbuzz is a generated protocol buffer package.

It is generated from these files:
	fizzbuzz.proto

It has these top-level messages:
	FizzbuzzRequest
	FizzbuzzReply
*/
package fizzbuzz

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

// The request containing the fizzbuzz params.
type FizzbuzzRequest struct {
	String1 string `protobuf:"bytes,1,opt,name=string1" json:"string1,omitempty"`
	String2 string `protobuf:"bytes,2,opt,name=string2" json:"string2,omitempty"`
	Int1    int32  `protobuf:"varint,3,opt,name=int1" json:"int1,omitempty"`
	Int2    int32  `protobuf:"varint,4,opt,name=int2" json:"int2,omitempty"`
	Limit   int32  `protobuf:"varint,5,opt,name=limit" json:"limit,omitempty"`
}

func (m *FizzbuzzRequest) Reset()                    { *m = FizzbuzzRequest{} }
func (m *FizzbuzzRequest) String() string            { return proto.CompactTextString(m) }
func (*FizzbuzzRequest) ProtoMessage()               {}
func (*FizzbuzzRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FizzbuzzRequest) GetString1() string {
	if m != nil {
		return m.String1
	}
	return ""
}

func (m *FizzbuzzRequest) GetString2() string {
	if m != nil {
		return m.String2
	}
	return ""
}

func (m *FizzbuzzRequest) GetInt1() int32 {
	if m != nil {
		return m.Int1
	}
	return 0
}

func (m *FizzbuzzRequest) GetInt2() int32 {
	if m != nil {
		return m.Int2
	}
	return 0
}

func (m *FizzbuzzRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// The response message containing the result
type FizzbuzzReply struct {
	Reply []string `protobuf:"bytes,1,rep,name=reply" json:"reply,omitempty"`
}

func (m *FizzbuzzReply) Reset()                    { *m = FizzbuzzReply{} }
func (m *FizzbuzzReply) String() string            { return proto.CompactTextString(m) }
func (*FizzbuzzReply) ProtoMessage()               {}
func (*FizzbuzzReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FizzbuzzReply) GetReply() []string {
	if m != nil {
		return m.Reply
	}
	return nil
}

func init() {
	proto.RegisterType((*FizzbuzzRequest)(nil), "fizzbuzz.FizzbuzzRequest")
	proto.RegisterType((*FizzbuzzReply)(nil), "fizzbuzz.FizzbuzzReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Fizzbuzz service

type FizzbuzzClient interface {
	// Processes a fizzbuzz request
	DoFizzbuzz(ctx context.Context, in *FizzbuzzRequest, opts ...grpc.CallOption) (*FizzbuzzReply, error)
}

type fizzbuzzClient struct {
	cc *grpc.ClientConn
}

func NewFizzbuzzClient(cc *grpc.ClientConn) FizzbuzzClient {
	return &fizzbuzzClient{cc}
}

func (c *fizzbuzzClient) DoFizzbuzz(ctx context.Context, in *FizzbuzzRequest, opts ...grpc.CallOption) (*FizzbuzzReply, error) {
	out := new(FizzbuzzReply)
	err := grpc.Invoke(ctx, "/fizzbuzz.Fizzbuzz/DoFizzbuzz", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Fizzbuzz service

type FizzbuzzServer interface {
	// Processes a fizzbuzz request
	DoFizzbuzz(context.Context, *FizzbuzzRequest) (*FizzbuzzReply, error)
}

func RegisterFizzbuzzServer(s *grpc.Server, srv FizzbuzzServer) {
	s.RegisterService(&_Fizzbuzz_serviceDesc, srv)
}

func _Fizzbuzz_DoFizzbuzz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FizzbuzzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FizzbuzzServer).DoFizzbuzz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fizzbuzz.Fizzbuzz/DoFizzbuzz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FizzbuzzServer).DoFizzbuzz(ctx, req.(*FizzbuzzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fizzbuzz_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fizzbuzz.Fizzbuzz",
	HandlerType: (*FizzbuzzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoFizzbuzz",
			Handler:    _Fizzbuzz_DoFizzbuzz_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fizzbuzz.proto",
}

func init() { proto.RegisterFile("fizzbuzz.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcb, 0xac, 0xaa,
	0x4a, 0x2a, 0xad, 0xaa, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x9a,
	0x19, 0xb9, 0xf8, 0xdd, 0xa0, 0x9c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x09, 0x2e,
	0xf6, 0xe2, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x43, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18,
	0x17, 0x21, 0x63, 0x24, 0xc1, 0x84, 0x2c, 0x63, 0x24, 0x24, 0xc4, 0xc5, 0x92, 0x99, 0x57, 0x62,
	0x28, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0x66, 0x43, 0xc5, 0x8c, 0x24, 0x58, 0xe0, 0x62,
	0x46, 0x42, 0x22, 0x5c, 0xac, 0x39, 0x99, 0xb9, 0x99, 0x25, 0x12, 0xac, 0x60, 0x41, 0x08, 0x47,
	0x49, 0x95, 0x8b, 0x17, 0xe1, 0x88, 0x82, 0x9c, 0x4a, 0x90, 0xb2, 0x22, 0x10, 0x43, 0x82, 0x51,
	0x81, 0x59, 0x83, 0x33, 0x08, 0xc2, 0x31, 0xf2, 0xe3, 0xe2, 0x80, 0x29, 0x13, 0x72, 0xe2, 0xe2,
	0x72, 0xc9, 0x87, 0xf3, 0x24, 0xf5, 0xe0, 0x3e, 0x44, 0xf3, 0x8d, 0x94, 0x38, 0x36, 0xa9, 0x82,
	0x9c, 0x4a, 0x25, 0x86, 0x24, 0x36, 0x70, 0x68, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x00,
	0x6d, 0x0e, 0x9f, 0x1f, 0x01, 0x00, 0x00,
}
