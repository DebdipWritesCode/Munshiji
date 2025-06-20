// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: service_munshiji.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Munshiji_CreateUser_FullMethodName             = "/pb.Munshiji/CreateUser"
	Munshiji_LoginUser_FullMethodName              = "/pb.Munshiji/LoginUser"
	Munshiji_RefreshAccessToken_FullMethodName     = "/pb.Munshiji/RefreshAccessToken"
	Munshiji_GetUserById_FullMethodName            = "/pb.Munshiji/GetUserById"
	Munshiji_CreateScoreSheet_FullMethodName       = "/pb.Munshiji/CreateScoreSheet"
	Munshiji_UpdateScoreSheet_FullMethodName       = "/pb.Munshiji/UpdateScoreSheet"
	Munshiji_GetScoreSheetByUserId_FullMethodName  = "/pb.Munshiji/GetScoreSheetByUserId"
	Munshiji_DeleteScoreSheet_FullMethodName       = "/pb.Munshiji/DeleteScoreSheet"
	Munshiji_GetScoreSheetDetails_FullMethodName   = "/pb.Munshiji/GetScoreSheetDetails"
	Munshiji_CreateDelegate_FullMethodName         = "/pb.Munshiji/CreateDelegate"
	Munshiji_GetDelegateById_FullMethodName        = "/pb.Munshiji/GetDelegateById"
	Munshiji_UpdateDelegateNameByID_FullMethodName = "/pb.Munshiji/UpdateDelegateNameByID"
	Munshiji_DeleteDelegate_FullMethodName         = "/pb.Munshiji/DeleteDelegate"
	Munshiji_CreateScore_FullMethodName            = "/pb.Munshiji/CreateScore"
	Munshiji_UpdateScore_FullMethodName            = "/pb.Munshiji/UpdateScore"
	Munshiji_DeleteScore_FullMethodName            = "/pb.Munshiji/DeleteScore"
	Munshiji_CreateParameter_FullMethodName        = "/pb.Munshiji/CreateParameter"
	Munshiji_GetParameterById_FullMethodName       = "/pb.Munshiji/GetParameterById"
	Munshiji_UpdateParameter_FullMethodName        = "/pb.Munshiji/UpdateParameter"
	Munshiji_DeleteParameter_FullMethodName        = "/pb.Munshiji/DeleteParameter"
	Munshiji_GetFeedbackByLLM_FullMethodName       = "/pb.Munshiji/GetFeedbackByLLM"
	Munshiji_VerifyEmail_FullMethodName            = "/pb.Munshiji/VerifyEmail"
)

// MunshijiClient is the client API for Munshiji service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MunshijiClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	RefreshAccessToken(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LoginUserResponse, error)
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error)
	CreateScoreSheet(ctx context.Context, in *CreateScoreSheetRequest, opts ...grpc.CallOption) (*CreateScoreSheetResponse, error)
	UpdateScoreSheet(ctx context.Context, in *UpdateScoreSheetRequest, opts ...grpc.CallOption) (*UpdateScoreSheetResponse, error)
	GetScoreSheetByUserId(ctx context.Context, in *GetScoreSheetByUserIdRequest, opts ...grpc.CallOption) (*GetScoreSheetByUserIdResponse, error)
	DeleteScoreSheet(ctx context.Context, in *DeleteScoreSheetRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetScoreSheetDetails(ctx context.Context, in *GetScoreSheetDetailsRequest, opts ...grpc.CallOption) (*GetScoreSheetDetailsResponse, error)
	CreateDelegate(ctx context.Context, in *CreateDelegateRequest, opts ...grpc.CallOption) (*CreateDelegateResponse, error)
	GetDelegateById(ctx context.Context, in *GetDelegateByIdRequest, opts ...grpc.CallOption) (*GetDelegateByIdResponse, error)
	UpdateDelegateNameByID(ctx context.Context, in *UpdateDelegateNameByIDRequest, opts ...grpc.CallOption) (*UpdateDelegateNameByIDResponse, error)
	DeleteDelegate(ctx context.Context, in *DeleteDelegateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateScore(ctx context.Context, in *CreateScoreRequest, opts ...grpc.CallOption) (*CreateScoreResponse, error)
	UpdateScore(ctx context.Context, in *UpdateScoreRequest, opts ...grpc.CallOption) (*UpdateScoreResponse, error)
	DeleteScore(ctx context.Context, in *DeleteScoreRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateParameter(ctx context.Context, in *CreateParameterRequest, opts ...grpc.CallOption) (*CreateParameterResponse, error)
	GetParameterById(ctx context.Context, in *GetParameterByIdRequest, opts ...grpc.CallOption) (*GetParameterByIdResponse, error)
	UpdateParameter(ctx context.Context, in *UpdateParameterRequest, opts ...grpc.CallOption) (*UpdateParameterResponse, error)
	DeleteParameter(ctx context.Context, in *DeleteParameterRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetFeedbackByLLM(ctx context.Context, in *GetFeedbackByLLMRequest, opts ...grpc.CallOption) (*GetFeedbackByLLMResponse, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error)
}

type munshijiClient struct {
	cc grpc.ClientConnInterface
}

func NewMunshijiClient(cc grpc.ClientConnInterface) MunshijiClient {
	return &munshijiClient{cc}
}

func (c *munshijiClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, Munshiji_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, Munshiji_LoginUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) RefreshAccessToken(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, Munshiji_RefreshAccessToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserByIdResponse)
	err := c.cc.Invoke(ctx, Munshiji_GetUserById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) CreateScoreSheet(ctx context.Context, in *CreateScoreSheetRequest, opts ...grpc.CallOption) (*CreateScoreSheetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateScoreSheetResponse)
	err := c.cc.Invoke(ctx, Munshiji_CreateScoreSheet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) UpdateScoreSheet(ctx context.Context, in *UpdateScoreSheetRequest, opts ...grpc.CallOption) (*UpdateScoreSheetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateScoreSheetResponse)
	err := c.cc.Invoke(ctx, Munshiji_UpdateScoreSheet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) GetScoreSheetByUserId(ctx context.Context, in *GetScoreSheetByUserIdRequest, opts ...grpc.CallOption) (*GetScoreSheetByUserIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetScoreSheetByUserIdResponse)
	err := c.cc.Invoke(ctx, Munshiji_GetScoreSheetByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) DeleteScoreSheet(ctx context.Context, in *DeleteScoreSheetRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Munshiji_DeleteScoreSheet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) GetScoreSheetDetails(ctx context.Context, in *GetScoreSheetDetailsRequest, opts ...grpc.CallOption) (*GetScoreSheetDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetScoreSheetDetailsResponse)
	err := c.cc.Invoke(ctx, Munshiji_GetScoreSheetDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) CreateDelegate(ctx context.Context, in *CreateDelegateRequest, opts ...grpc.CallOption) (*CreateDelegateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDelegateResponse)
	err := c.cc.Invoke(ctx, Munshiji_CreateDelegate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) GetDelegateById(ctx context.Context, in *GetDelegateByIdRequest, opts ...grpc.CallOption) (*GetDelegateByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDelegateByIdResponse)
	err := c.cc.Invoke(ctx, Munshiji_GetDelegateById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) UpdateDelegateNameByID(ctx context.Context, in *UpdateDelegateNameByIDRequest, opts ...grpc.CallOption) (*UpdateDelegateNameByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDelegateNameByIDResponse)
	err := c.cc.Invoke(ctx, Munshiji_UpdateDelegateNameByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) DeleteDelegate(ctx context.Context, in *DeleteDelegateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Munshiji_DeleteDelegate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) CreateScore(ctx context.Context, in *CreateScoreRequest, opts ...grpc.CallOption) (*CreateScoreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateScoreResponse)
	err := c.cc.Invoke(ctx, Munshiji_CreateScore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) UpdateScore(ctx context.Context, in *UpdateScoreRequest, opts ...grpc.CallOption) (*UpdateScoreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateScoreResponse)
	err := c.cc.Invoke(ctx, Munshiji_UpdateScore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) DeleteScore(ctx context.Context, in *DeleteScoreRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Munshiji_DeleteScore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) CreateParameter(ctx context.Context, in *CreateParameterRequest, opts ...grpc.CallOption) (*CreateParameterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateParameterResponse)
	err := c.cc.Invoke(ctx, Munshiji_CreateParameter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) GetParameterById(ctx context.Context, in *GetParameterByIdRequest, opts ...grpc.CallOption) (*GetParameterByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetParameterByIdResponse)
	err := c.cc.Invoke(ctx, Munshiji_GetParameterById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) UpdateParameter(ctx context.Context, in *UpdateParameterRequest, opts ...grpc.CallOption) (*UpdateParameterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateParameterResponse)
	err := c.cc.Invoke(ctx, Munshiji_UpdateParameter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) DeleteParameter(ctx context.Context, in *DeleteParameterRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Munshiji_DeleteParameter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) GetFeedbackByLLM(ctx context.Context, in *GetFeedbackByLLMRequest, opts ...grpc.CallOption) (*GetFeedbackByLLMResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFeedbackByLLMResponse)
	err := c.cc.Invoke(ctx, Munshiji_GetFeedbackByLLM_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *munshijiClient) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyEmailResponse)
	err := c.cc.Invoke(ctx, Munshiji_VerifyEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MunshijiServer is the server API for Munshiji service.
// All implementations must embed UnimplementedMunshijiServer
// for forward compatibility.
type MunshijiServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	RefreshAccessToken(context.Context, *empty.Empty) (*LoginUserResponse, error)
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error)
	CreateScoreSheet(context.Context, *CreateScoreSheetRequest) (*CreateScoreSheetResponse, error)
	UpdateScoreSheet(context.Context, *UpdateScoreSheetRequest) (*UpdateScoreSheetResponse, error)
	GetScoreSheetByUserId(context.Context, *GetScoreSheetByUserIdRequest) (*GetScoreSheetByUserIdResponse, error)
	DeleteScoreSheet(context.Context, *DeleteScoreSheetRequest) (*empty.Empty, error)
	GetScoreSheetDetails(context.Context, *GetScoreSheetDetailsRequest) (*GetScoreSheetDetailsResponse, error)
	CreateDelegate(context.Context, *CreateDelegateRequest) (*CreateDelegateResponse, error)
	GetDelegateById(context.Context, *GetDelegateByIdRequest) (*GetDelegateByIdResponse, error)
	UpdateDelegateNameByID(context.Context, *UpdateDelegateNameByIDRequest) (*UpdateDelegateNameByIDResponse, error)
	DeleteDelegate(context.Context, *DeleteDelegateRequest) (*empty.Empty, error)
	CreateScore(context.Context, *CreateScoreRequest) (*CreateScoreResponse, error)
	UpdateScore(context.Context, *UpdateScoreRequest) (*UpdateScoreResponse, error)
	DeleteScore(context.Context, *DeleteScoreRequest) (*empty.Empty, error)
	CreateParameter(context.Context, *CreateParameterRequest) (*CreateParameterResponse, error)
	GetParameterById(context.Context, *GetParameterByIdRequest) (*GetParameterByIdResponse, error)
	UpdateParameter(context.Context, *UpdateParameterRequest) (*UpdateParameterResponse, error)
	DeleteParameter(context.Context, *DeleteParameterRequest) (*empty.Empty, error)
	GetFeedbackByLLM(context.Context, *GetFeedbackByLLMRequest) (*GetFeedbackByLLMResponse, error)
	VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error)
	mustEmbedUnimplementedMunshijiServer()
}

// UnimplementedMunshijiServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMunshijiServer struct{}

func (UnimplementedMunshijiServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMunshijiServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedMunshijiServer) RefreshAccessToken(context.Context, *empty.Empty) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshAccessToken not implemented")
}
func (UnimplementedMunshijiServer) GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedMunshijiServer) CreateScoreSheet(context.Context, *CreateScoreSheetRequest) (*CreateScoreSheetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScoreSheet not implemented")
}
func (UnimplementedMunshijiServer) UpdateScoreSheet(context.Context, *UpdateScoreSheetRequest) (*UpdateScoreSheetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScoreSheet not implemented")
}
func (UnimplementedMunshijiServer) GetScoreSheetByUserId(context.Context, *GetScoreSheetByUserIdRequest) (*GetScoreSheetByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScoreSheetByUserId not implemented")
}
func (UnimplementedMunshijiServer) DeleteScoreSheet(context.Context, *DeleteScoreSheetRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScoreSheet not implemented")
}
func (UnimplementedMunshijiServer) GetScoreSheetDetails(context.Context, *GetScoreSheetDetailsRequest) (*GetScoreSheetDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScoreSheetDetails not implemented")
}
func (UnimplementedMunshijiServer) CreateDelegate(context.Context, *CreateDelegateRequest) (*CreateDelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDelegate not implemented")
}
func (UnimplementedMunshijiServer) GetDelegateById(context.Context, *GetDelegateByIdRequest) (*GetDelegateByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDelegateById not implemented")
}
func (UnimplementedMunshijiServer) UpdateDelegateNameByID(context.Context, *UpdateDelegateNameByIDRequest) (*UpdateDelegateNameByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDelegateNameByID not implemented")
}
func (UnimplementedMunshijiServer) DeleteDelegate(context.Context, *DeleteDelegateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDelegate not implemented")
}
func (UnimplementedMunshijiServer) CreateScore(context.Context, *CreateScoreRequest) (*CreateScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScore not implemented")
}
func (UnimplementedMunshijiServer) UpdateScore(context.Context, *UpdateScoreRequest) (*UpdateScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScore not implemented")
}
func (UnimplementedMunshijiServer) DeleteScore(context.Context, *DeleteScoreRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScore not implemented")
}
func (UnimplementedMunshijiServer) CreateParameter(context.Context, *CreateParameterRequest) (*CreateParameterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateParameter not implemented")
}
func (UnimplementedMunshijiServer) GetParameterById(context.Context, *GetParameterByIdRequest) (*GetParameterByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParameterById not implemented")
}
func (UnimplementedMunshijiServer) UpdateParameter(context.Context, *UpdateParameterRequest) (*UpdateParameterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParameter not implemented")
}
func (UnimplementedMunshijiServer) DeleteParameter(context.Context, *DeleteParameterRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteParameter not implemented")
}
func (UnimplementedMunshijiServer) GetFeedbackByLLM(context.Context, *GetFeedbackByLLMRequest) (*GetFeedbackByLLMResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedbackByLLM not implemented")
}
func (UnimplementedMunshijiServer) VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedMunshijiServer) mustEmbedUnimplementedMunshijiServer() {}
func (UnimplementedMunshijiServer) testEmbeddedByValue()                  {}

// UnsafeMunshijiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MunshijiServer will
// result in compilation errors.
type UnsafeMunshijiServer interface {
	mustEmbedUnimplementedMunshijiServer()
}

func RegisterMunshijiServer(s grpc.ServiceRegistrar, srv MunshijiServer) {
	// If the following call pancis, it indicates UnimplementedMunshijiServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Munshiji_ServiceDesc, srv)
}

func _Munshiji_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_RefreshAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).RefreshAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_RefreshAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).RefreshAccessToken(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_CreateScoreSheet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScoreSheetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).CreateScoreSheet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_CreateScoreSheet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).CreateScoreSheet(ctx, req.(*CreateScoreSheetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_UpdateScoreSheet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScoreSheetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).UpdateScoreSheet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_UpdateScoreSheet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).UpdateScoreSheet(ctx, req.(*UpdateScoreSheetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_GetScoreSheetByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScoreSheetByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).GetScoreSheetByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_GetScoreSheetByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).GetScoreSheetByUserId(ctx, req.(*GetScoreSheetByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_DeleteScoreSheet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScoreSheetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).DeleteScoreSheet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_DeleteScoreSheet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).DeleteScoreSheet(ctx, req.(*DeleteScoreSheetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_GetScoreSheetDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScoreSheetDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).GetScoreSheetDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_GetScoreSheetDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).GetScoreSheetDetails(ctx, req.(*GetScoreSheetDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_CreateDelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDelegateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).CreateDelegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_CreateDelegate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).CreateDelegate(ctx, req.(*CreateDelegateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_GetDelegateById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDelegateByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).GetDelegateById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_GetDelegateById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).GetDelegateById(ctx, req.(*GetDelegateByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_UpdateDelegateNameByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDelegateNameByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).UpdateDelegateNameByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_UpdateDelegateNameByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).UpdateDelegateNameByID(ctx, req.(*UpdateDelegateNameByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_DeleteDelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDelegateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).DeleteDelegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_DeleteDelegate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).DeleteDelegate(ctx, req.(*DeleteDelegateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_CreateScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).CreateScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_CreateScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).CreateScore(ctx, req.(*CreateScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_UpdateScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).UpdateScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_UpdateScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).UpdateScore(ctx, req.(*UpdateScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_DeleteScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).DeleteScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_DeleteScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).DeleteScore(ctx, req.(*DeleteScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_CreateParameter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateParameterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).CreateParameter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_CreateParameter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).CreateParameter(ctx, req.(*CreateParameterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_GetParameterById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParameterByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).GetParameterById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_GetParameterById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).GetParameterById(ctx, req.(*GetParameterByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_UpdateParameter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateParameterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).UpdateParameter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_UpdateParameter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).UpdateParameter(ctx, req.(*UpdateParameterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_DeleteParameter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteParameterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).DeleteParameter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_DeleteParameter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).DeleteParameter(ctx, req.(*DeleteParameterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_GetFeedbackByLLM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedbackByLLMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).GetFeedbackByLLM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_GetFeedbackByLLM_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).GetFeedbackByLLM(ctx, req.(*GetFeedbackByLLMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Munshiji_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MunshijiServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Munshiji_VerifyEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MunshijiServer).VerifyEmail(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Munshiji_ServiceDesc is the grpc.ServiceDesc for Munshiji service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Munshiji_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Munshiji",
	HandlerType: (*MunshijiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Munshiji_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Munshiji_LoginUser_Handler,
		},
		{
			MethodName: "RefreshAccessToken",
			Handler:    _Munshiji_RefreshAccessToken_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _Munshiji_GetUserById_Handler,
		},
		{
			MethodName: "CreateScoreSheet",
			Handler:    _Munshiji_CreateScoreSheet_Handler,
		},
		{
			MethodName: "UpdateScoreSheet",
			Handler:    _Munshiji_UpdateScoreSheet_Handler,
		},
		{
			MethodName: "GetScoreSheetByUserId",
			Handler:    _Munshiji_GetScoreSheetByUserId_Handler,
		},
		{
			MethodName: "DeleteScoreSheet",
			Handler:    _Munshiji_DeleteScoreSheet_Handler,
		},
		{
			MethodName: "GetScoreSheetDetails",
			Handler:    _Munshiji_GetScoreSheetDetails_Handler,
		},
		{
			MethodName: "CreateDelegate",
			Handler:    _Munshiji_CreateDelegate_Handler,
		},
		{
			MethodName: "GetDelegateById",
			Handler:    _Munshiji_GetDelegateById_Handler,
		},
		{
			MethodName: "UpdateDelegateNameByID",
			Handler:    _Munshiji_UpdateDelegateNameByID_Handler,
		},
		{
			MethodName: "DeleteDelegate",
			Handler:    _Munshiji_DeleteDelegate_Handler,
		},
		{
			MethodName: "CreateScore",
			Handler:    _Munshiji_CreateScore_Handler,
		},
		{
			MethodName: "UpdateScore",
			Handler:    _Munshiji_UpdateScore_Handler,
		},
		{
			MethodName: "DeleteScore",
			Handler:    _Munshiji_DeleteScore_Handler,
		},
		{
			MethodName: "CreateParameter",
			Handler:    _Munshiji_CreateParameter_Handler,
		},
		{
			MethodName: "GetParameterById",
			Handler:    _Munshiji_GetParameterById_Handler,
		},
		{
			MethodName: "UpdateParameter",
			Handler:    _Munshiji_UpdateParameter_Handler,
		},
		{
			MethodName: "DeleteParameter",
			Handler:    _Munshiji_DeleteParameter_Handler,
		},
		{
			MethodName: "GetFeedbackByLLM",
			Handler:    _Munshiji_GetFeedbackByLLM_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _Munshiji_VerifyEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_munshiji.proto",
}
