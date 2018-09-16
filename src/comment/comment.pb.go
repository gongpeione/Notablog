// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comment.proto

package comment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

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

type User struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type CommentRequest struct {
	CommentId            int32    `protobuf:"varint,1,opt,name=commentId,proto3" json:"commentId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentRequest) Reset()         { *m = CommentRequest{} }
func (m *CommentRequest) String() string { return proto.CompactTextString(m) }
func (*CommentRequest) ProtoMessage()    {}
func (*CommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{1}
}
func (m *CommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentRequest.Unmarshal(m, b)
}
func (m *CommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentRequest.Marshal(b, m, deterministic)
}
func (m *CommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentRequest.Merge(m, src)
}
func (m *CommentRequest) XXX_Size() int {
	return xxx_messageInfo_CommentRequest.Size(m)
}
func (m *CommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommentRequest proto.InternalMessageInfo

func (m *CommentRequest) GetCommentId() int32 {
	if m != nil {
		return m.CommentId
	}
	return 0
}

type CommentResponse struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentResponse) Reset()         { *m = CommentResponse{} }
func (m *CommentResponse) String() string { return proto.CompactTextString(m) }
func (*CommentResponse) ProtoMessage()    {}
func (*CommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{2}
}
func (m *CommentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentResponse.Unmarshal(m, b)
}
func (m *CommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentResponse.Marshal(b, m, deterministic)
}
func (m *CommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentResponse.Merge(m, src)
}
func (m *CommentResponse) XXX_Size() int {
	return xxx_messageInfo_CommentResponse.Size(m)
}
func (m *CommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommentResponse proto.InternalMessageInfo

func (m *CommentResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "comment.User")
	proto.RegisterType((*CommentRequest)(nil), "comment.CommentRequest")
	proto.RegisterType((*CommentResponse)(nil), "comment.CommentResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommentServiceClient interface {
	AddComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
}

type commentServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommentServiceClient(cc *grpc.ClientConn) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) AddComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/comment.CommentService/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
type CommentServiceServer interface {
	AddComment(context.Context, *CommentRequest) (*CommentResponse, error)
}

func RegisterCommentServiceServer(s *grpc.Server, srv CommentServiceServer) {
	s.RegisterService(&_CommentService_serviceDesc, srv)
}

func _CommentService_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.CommentService/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).AddComment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddComment",
			Handler:    _CommentService_AddComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}

func init() { proto.RegisterFile("comment.proto", fileDescriptor_749aee09ea917828) }

var fileDescriptor_749aee09ea917828 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0xcd,
	0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x64, 0xd2,
	0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12,
	0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xca, 0x94, 0xbc, 0xb8, 0x58, 0x42, 0x8b, 0x53, 0x8b, 0x84,
	0xc4, 0xb8, 0xd8, 0x4a, 0x8b, 0x53, 0x8b, 0x3c, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83,
	0xa0, 0x3c, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x30, 0x1b, 0xa4, 0x36, 0xb1, 0x2c, 0xb1, 0x24, 0xb1, 0x48, 0x82, 0x19, 0x2c, 0x0a, 0xe5,
	0x29, 0xe9, 0x71, 0xf1, 0x39, 0x43, 0x2c, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92,
	0xe1, 0xe2, 0x84, 0x3a, 0x03, 0x6e, 0x30, 0x42, 0x40, 0x49, 0x95, 0x8b, 0x1f, 0xae, 0xbe, 0xb8,
	0x20, 0x3f, 0xaf, 0x38, 0x15, 0x64, 0x5d, 0x72, 0x7e, 0x4a, 0x2a, 0x58, 0x2d, 0x67, 0x10, 0x98,
	0x6d, 0x94, 0x06, 0x37, 0x36, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x84, 0x8b, 0xcb,
	0x31, 0x25, 0x05, 0x2a, 0x28, 0x24, 0xae, 0x07, 0xf3, 0x39, 0xaa, 0xed, 0x52, 0x12, 0x98, 0x12,
	0x10, 0x6b, 0x94, 0x84, 0x9b, 0x2e, 0x3f, 0x99, 0xcc, 0xc4, 0x2b, 0xc4, 0xad, 0x5f, 0x66, 0xa8,
	0x0f, 0x55, 0x94, 0xc4, 0x06, 0x0e, 0x11, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0x33,
	0xcd, 0x3c, 0x49, 0x01, 0x00, 0x00,
}
