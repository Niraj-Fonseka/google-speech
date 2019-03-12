// Code generated by protoc-gen-go. DO NOT EDIT.
// source: text-to-speech.proto

package v1

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

type TextToSpeechMessage struct {
	Data     string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Response string `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
	Gender   string `protobuf:"bytes,3,opt,name=gender" json:"gender,omitempty"`
	Play     bool   `protobuf:"varint,4,opt,name=play" json:"play,omitempty"`
}

func (m *TextToSpeechMessage) Reset()                    { *m = TextToSpeechMessage{} }
func (m *TextToSpeechMessage) String() string            { return proto.CompactTextString(m) }
func (*TextToSpeechMessage) ProtoMessage()               {}
func (*TextToSpeechMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *TextToSpeechMessage) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *TextToSpeechMessage) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func (m *TextToSpeechMessage) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *TextToSpeechMessage) GetPlay() bool {
	if m != nil {
		return m.Play
	}
	return false
}

func init() {
	proto.RegisterType((*TextToSpeechMessage)(nil), "v1.TextToSpeechMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TextToSpeech service

type TextToSpeechClient interface {
	GenerateSpeech(ctx context.Context, in *TextToSpeechMessage, opts ...grpc.CallOption) (*TextToSpeechMessage, error)
}

type textToSpeechClient struct {
	cc *grpc.ClientConn
}

func NewTextToSpeechClient(cc *grpc.ClientConn) TextToSpeechClient {
	return &textToSpeechClient{cc}
}

func (c *textToSpeechClient) GenerateSpeech(ctx context.Context, in *TextToSpeechMessage, opts ...grpc.CallOption) (*TextToSpeechMessage, error) {
	out := new(TextToSpeechMessage)
	err := grpc.Invoke(ctx, "/v1.TextToSpeech/GenerateSpeech", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TextToSpeech service

type TextToSpeechServer interface {
	GenerateSpeech(context.Context, *TextToSpeechMessage) (*TextToSpeechMessage, error)
}

func RegisterTextToSpeechServer(s *grpc.Server, srv TextToSpeechServer) {
	s.RegisterService(&_TextToSpeech_serviceDesc, srv)
}

func _TextToSpeech_GenerateSpeech_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextToSpeechMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextToSpeechServer).GenerateSpeech(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.TextToSpeech/GenerateSpeech",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextToSpeechServer).GenerateSpeech(ctx, req.(*TextToSpeechMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _TextToSpeech_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.TextToSpeech",
	HandlerType: (*TextToSpeechServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateSpeech",
			Handler:    _TextToSpeech_GenerateSpeech_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "text-to-speech.proto",
}

func init() { proto.RegisterFile("text-to-speech.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x49, 0xad, 0x28,
	0xd1, 0x2d, 0xc9, 0xd7, 0x2d, 0x2e, 0x48, 0x4d, 0x4d, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x2a, 0x33, 0x54, 0x2a, 0xe4, 0x12, 0x0e, 0x49, 0xad, 0x28, 0x09, 0xc9, 0x0f, 0x06,
	0xcb, 0xf8, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a, 0x09, 0x71, 0xb1, 0xa4, 0x24, 0x96, 0x24,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x52, 0x5c, 0x1c, 0x45, 0xa9, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x12, 0x4c, 0x60, 0x71, 0x38, 0x5f, 0x48, 0x8c, 0x8b, 0x2d, 0x3d,
	0x35, 0x2f, 0x25, 0xb5, 0x48, 0x82, 0x19, 0x2c, 0x03, 0xe5, 0x81, 0xcc, 0x29, 0xc8, 0x49, 0xac,
	0x94, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x08, 0x02, 0xb3, 0x8d, 0x42, 0xb8, 0x78, 0x90, 0xad, 0x14,
	0x72, 0xe1, 0xe2, 0x73, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0x49, 0x85, 0x8a, 0x88, 0xeb, 0x95,
	0x19, 0xea, 0x61, 0x71, 0x96, 0x14, 0x2e, 0x09, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x9f, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x32, 0x9d, 0x1b, 0x5d, 0xeb, 0x00, 0x00, 0x00,
}
