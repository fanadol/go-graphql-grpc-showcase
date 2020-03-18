// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/user/proto/user.proto

package userpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74a4dd0105d3cb3, []int{0}
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

func (m *User) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetRequest struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74a4dd0105d3cb3, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type GetResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74a4dd0105d3cb3, []int{2}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type GetUserFriendsResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserFriendsResponse) Reset()         { *m = GetUserFriendsResponse{} }
func (m *GetUserFriendsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserFriendsResponse) ProtoMessage()    {}
func (*GetUserFriendsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74a4dd0105d3cb3, []int{3}
}

func (m *GetUserFriendsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserFriendsResponse.Unmarshal(m, b)
}
func (m *GetUserFriendsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserFriendsResponse.Marshal(b, m, deterministic)
}
func (m *GetUserFriendsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserFriendsResponse.Merge(m, src)
}
func (m *GetUserFriendsResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserFriendsResponse.Size(m)
}
func (m *GetUserFriendsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserFriendsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserFriendsResponse proto.InternalMessageInfo

func (m *GetUserFriendsResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "userpb.User")
	proto.RegisterType((*GetRequest)(nil), "userpb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "userpb.GetResponse")
	proto.RegisterType((*GetUserFriendsResponse)(nil), "userpb.GetUserFriendsResponse")
}

func init() { proto.RegisterFile("internal/user/proto/user.proto", fileDescriptor_e74a4dd0105d3cb3) }

var fileDescriptor_e74a4dd0105d3cb3 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x4f, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0xec, 0xa6, 0x4d, 0xc1, 0x57, 0xe9, 0xe1, 0x09, 0x25, 0x44, 0x29, 0x61, 0x4f, 0xc1, 0xc3,
	0x06, 0xab, 0x47, 0x8f, 0xb5, 0xa1, 0x17, 0x0f, 0x01, 0x3f, 0xa0, 0xc1, 0x67, 0x29, 0xb4, 0xd9,
	0x75, 0xdf, 0xe6, 0xe0, 0x07, 0xf8, 0xdf, 0xb2, 0xbb, 0x14, 0x1a, 0xf1, 0xe2, 0x6d, 0xde, 0x9b,
	0x19, 0x66, 0x06, 0x96, 0x87, 0xce, 0x91, 0xed, 0x76, 0xc7, 0xaa, 0x67, 0xb2, 0x95, 0xb1, 0xda,
	0xe9, 0x00, 0x55, 0x80, 0x38, 0xf5, 0xd8, 0xb4, 0xf9, 0xed, 0x5e, 0xeb, 0xfd, 0x91, 0xa2, 0xa0,
	0xed, 0x3f, 0x2a, 0x3a, 0x19, 0xf7, 0x15, 0x45, 0xf2, 0x1e, 0x26, 0x6f, 0x4c, 0x16, 0xe7, 0x90,
	0x6c, 0xd7, 0x99, 0x28, 0x44, 0x99, 0x36, 0xc9, 0x76, 0x8d, 0x08, 0x93, 0xd7, 0xdd, 0x89, 0xb2,
	0xa4, 0x10, 0xe5, 0x55, 0x13, 0xb0, 0xbc, 0x03, 0xa8, 0xc9, 0x35, 0xf4, 0xd9, 0x13, 0xbb, 0xdf,
	0x0e, 0xf9, 0x00, 0xb3, 0xc0, 0xb2, 0xd1, 0x1d, 0x13, 0x4a, 0x48, 0x7d, 0x3e, 0x67, 0xa2, 0x18,
	0x97, 0xb3, 0xd5, 0xb5, 0x8a, 0x6d, 0x94, 0x4f, 0x6b, 0x22, 0x25, 0x9f, 0x61, 0x51, 0x93, 0xf3,
	0x9f, 0x8d, 0x3d, 0x50, 0xf7, 0xce, 0xff, 0x71, 0xaf, 0xbe, 0x05, 0xa4, 0xfe, 0x66, 0x7c, 0x82,
	0x71, 0x4d, 0x0e, 0x17, 0x2a, 0x2e, 0x55, 0xe7, 0xa5, 0xea, 0xc5, 0x2f, 0xcd, 0x6f, 0xce, 0xee,
	0x8b, 0x7e, 0x72, 0x84, 0x1b, 0x98, 0x0f, 0xd3, 0x11, 0x07, 0xc2, 0x30, 0x33, 0x5f, 0x5e, 0xfc,
	0xfe, 0x68, 0x2a, 0x47, 0xed, 0x34, 0xc4, 0x3d, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x06, 0xe0,
	0xae, 0x1b, 0x90, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetResponse, error)
	GetUserFriends(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetUserFriendsResponse, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Get(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/userpb.Users/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserFriends(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetUserFriendsResponse, error) {
	out := new(GetUserFriendsResponse)
	err := c.cc.Invoke(ctx, "/userpb.Users/GetUserFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	Get(context.Context, *empty.Empty) (*GetResponse, error)
	GetUserFriends(context.Context, *GetRequest) (*GetUserFriendsResponse, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) Get(ctx context.Context, req *empty.Empty) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedUsersServer) GetUserFriends(ctx context.Context, req *GetRequest) (*GetUserFriendsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFriends not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.Users/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Get(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.Users/GetUserFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserFriends(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userpb.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Users_Get_Handler,
		},
		{
			MethodName: "GetUserFriends",
			Handler:    _Users_GetUserFriends_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/user/proto/user.proto",
}