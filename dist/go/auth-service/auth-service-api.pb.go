// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth-service-api.proto

package auth_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_go "github.com/influenzanet/api/dist/go"
	user_management "github.com/influenzanet/api/dist/go/user-management"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EncodedToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncodedToken) Reset()         { *m = EncodedToken{} }
func (m *EncodedToken) String() string { return proto.CompactTextString(m) }
func (*EncodedToken) ProtoMessage()    {}
func (*EncodedToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8e86308594dc230, []int{0}
}

func (m *EncodedToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncodedToken.Unmarshal(m, b)
}
func (m *EncodedToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncodedToken.Marshal(b, m, deterministic)
}
func (m *EncodedToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodedToken.Merge(m, src)
}
func (m *EncodedToken) XXX_Size() int {
	return xxx_messageInfo_EncodedToken.Size(m)
}
func (m *EncodedToken) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodedToken.DiscardUnknown(m)
}

var xxx_messageInfo_EncodedToken proto.InternalMessageInfo

func (m *EncodedToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*EncodedToken)(nil), "influenzanet.auth_service_api.EncodedToken")
}

func init() { proto.RegisterFile("auth-service-api.proto", fileDescriptor_e8e86308594dc230) }

var fileDescriptor_e8e86308594dc230 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x5f, 0x4b, 0xe3, 0x40,
	0x14, 0xc5, 0x29, 0xcb, 0x96, 0xdd, 0xd9, 0xa5, 0x42, 0x28, 0xc5, 0x46, 0x0a, 0x22, 0x3e, 0x28,
	0x92, 0x89, 0x54, 0xf0, 0xbd, 0x4a, 0x5f, 0x44, 0x44, 0xfa, 0xc7, 0x82, 0x08, 0x65, 0xd2, 0xdc,
	0x4e, 0x2e, 0x26, 0x33, 0x43, 0xe6, 0x8e, 0xa5, 0x7e, 0x48, 0x3f, 0x93, 0xa4, 0x69, 0x69, 0xf3,
	0x22, 0xe4, 0x2d, 0x97, 0xfc, 0xce, 0x39, 0x77, 0xce, 0x0c, 0xeb, 0x08, 0x47, 0x49, 0x60, 0x21,
	0xff, 0xc0, 0x05, 0x04, 0xc2, 0x20, 0x37, 0xb9, 0x26, 0xed, 0xf5, 0x50, 0x2d, 0x53, 0x07, 0xea,
	0x53, 0x28, 0x20, 0x5e, 0x40, 0xf3, 0x2d, 0x34, 0x17, 0x06, 0x7d, 0x4f, 0xa6, 0x3a, 0x12, 0x69,
	0x40, 0x6b, 0x03, 0xb6, 0x94, 0xf8, 0x5d, 0x67, 0x21, 0x0f, 0x32, 0xa1, 0x84, 0x84, 0x0c, 0x14,
	0xed, 0xdd, 0xfc, 0x13, 0xa9, 0xb5, 0x4c, 0x21, 0xdc, 0x4c, 0x91, 0x5b, 0x86, 0x90, 0x19, 0x5a,
	0x97, 0x3f, 0xcf, 0xce, 0xd9, 0xff, 0xa1, 0x5a, 0xe8, 0x18, 0xe2, 0x89, 0x7e, 0x07, 0xe5, 0xb5,
	0xd9, 0x6f, 0x2a, 0x3e, 0x8e, 0x1b, 0xa7, 0x8d, 0x8b, 0xbf, 0xa3, 0x72, 0xe8, 0x7f, 0xfd, 0x62,
	0xad, 0x81, 0xa3, 0x64, 0x5c, 0x6e, 0x31, 0x30, 0xe8, 0xdd, 0xb2, 0xe6, 0x98, 0x04, 0x39, 0xeb,
	0x75, 0x78, 0x19, 0xc0, 0x77, 0x01, 0x7c, 0x58, 0x04, 0xf8, 0x6d, 0x5e, 0x39, 0xc6, 0x96, 0x7e,
	0x63, 0xad, 0x47, 0x2d, 0x51, 0xcd, 0x90, 0x92, 0x61, 0x26, 0x30, 0xf5, 0x7a, 0x55, 0x6e, 0x6a,
	0x21, 0xbf, 0xcf, 0x21, 0x06, 0x45, 0x28, 0x52, 0xeb, 0x5f, 0xf1, 0x1f, 0xdb, 0xe0, 0x95, 0xf5,
	0x91, 0x1d, 0x8d, 0x51, 0x2a, 0x67, 0xf6, 0xf6, 0x97, 0x55, 0x7d, 0xd1, 0xd3, 0x7c, 0xdf, 0xd3,
	0xc6, 0xe2, 0x09, 0x56, 0x45, 0x6a, 0xbd, 0xa8, 0x29, 0xfb, 0xf7, 0x22, 0x52, 0x8c, 0x05, 0xc1,
	0xc3, 0x6c, 0xe2, 0xd5, 0xd1, 0xfa, 0xdd, 0x2a, 0xfc, 0x2c, 0x72, 0xbb, 0xb3, 0x8d, 0xd9, 0x9f,
	0x11, 0x28, 0x58, 0xd5, 0xf6, 0xac, 0x03, 0xdf, 0xf5, 0x5f, 0xaf, 0x25, 0x52, 0xe2, 0x22, 0xbe,
	0xd0, 0x59, 0x78, 0x28, 0x0c, 0x85, 0xc1, 0x30, 0x46, 0x4b, 0xa1, 0xd4, 0xe1, 0xe1, 0xfb, 0x8c,
	0x9a, 0x9b, 0xfb, 0xbd, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x99, 0x09, 0x39, 0x92, 0xb6, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceApiClient is the client API for AuthServiceApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceApiClient interface {
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*_go.Status, error)
	// Auth:
	LoginWithEmail(ctx context.Context, in *_go.UserCredentials, opts ...grpc.CallOption) (*EncodedToken, error)
	SignupWithEmail(ctx context.Context, in *user_management.NewUser, opts ...grpc.CallOption) (*EncodedToken, error)
	// Token handling:
	ValidateJWT(ctx context.Context, in *EncodedToken, opts ...grpc.CallOption) (*_go.ParsedToken, error)
	RenewJWT(ctx context.Context, in *EncodedToken, opts ...grpc.CallOption) (*EncodedToken, error)
}

type authServiceApiClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceApiClient(cc *grpc.ClientConn) AuthServiceApiClient {
	return &authServiceApiClient{cc}
}

func (c *authServiceApiClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*_go.Status, error) {
	out := new(_go.Status)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) LoginWithEmail(ctx context.Context, in *_go.UserCredentials, opts ...grpc.CallOption) (*EncodedToken, error) {
	out := new(EncodedToken)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/LoginWithEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) SignupWithEmail(ctx context.Context, in *user_management.NewUser, opts ...grpc.CallOption) (*EncodedToken, error) {
	out := new(EncodedToken)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/SignupWithEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) ValidateJWT(ctx context.Context, in *EncodedToken, opts ...grpc.CallOption) (*_go.ParsedToken, error) {
	out := new(_go.ParsedToken)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/ValidateJWT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) RenewJWT(ctx context.Context, in *EncodedToken, opts ...grpc.CallOption) (*EncodedToken, error) {
	out := new(EncodedToken)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/RenewJWT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceApiServer is the server API for AuthServiceApi service.
type AuthServiceApiServer interface {
	Status(context.Context, *empty.Empty) (*_go.Status, error)
	// Auth:
	LoginWithEmail(context.Context, *_go.UserCredentials) (*EncodedToken, error)
	SignupWithEmail(context.Context, *user_management.NewUser) (*EncodedToken, error)
	// Token handling:
	ValidateJWT(context.Context, *EncodedToken) (*_go.ParsedToken, error)
	RenewJWT(context.Context, *EncodedToken) (*EncodedToken, error)
}

func RegisterAuthServiceApiServer(s *grpc.Server, srv AuthServiceApiServer) {
	s.RegisterService(&_AuthServiceApi_serviceDesc, srv)
}

func _AuthServiceApi_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_LoginWithEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.UserCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).LoginWithEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/LoginWithEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).LoginWithEmail(ctx, req.(*_go.UserCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_SignupWithEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user_management.NewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).SignupWithEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/SignupWithEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).SignupWithEmail(ctx, req.(*user_management.NewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_ValidateJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncodedToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).ValidateJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/ValidateJWT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).ValidateJWT(ctx, req.(*EncodedToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_RenewJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncodedToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).RenewJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/RenewJWT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).RenewJWT(ctx, req.(*EncodedToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthServiceApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "influenzanet.auth_service_api.AuthServiceApi",
	HandlerType: (*AuthServiceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _AuthServiceApi_Status_Handler,
		},
		{
			MethodName: "LoginWithEmail",
			Handler:    _AuthServiceApi_LoginWithEmail_Handler,
		},
		{
			MethodName: "SignupWithEmail",
			Handler:    _AuthServiceApi_SignupWithEmail_Handler,
		},
		{
			MethodName: "ValidateJWT",
			Handler:    _AuthServiceApi_ValidateJWT_Handler,
		},
		{
			MethodName: "RenewJWT",
			Handler:    _AuthServiceApi_RenewJWT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth-service-api.proto",
}
