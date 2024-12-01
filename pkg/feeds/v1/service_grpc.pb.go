// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: feeds/v1/service.proto

package feeds

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FeedService_GetFeed_FullMethodName       = "/feeds.v1.FeedService/GetFeed"
	FeedService_GetUserFeed_FullMethodName   = "/feeds.v1.FeedService/GetUserFeed"
	FeedService_ListFeeds_FullMethodName     = "/feeds.v1.FeedService/ListFeeds"
	FeedService_ListUserFeeds_FullMethodName = "/feeds.v1.FeedService/ListUserFeeds"
	FeedService_CreateFeed_FullMethodName    = "/feeds.v1.FeedService/CreateFeed"
	FeedService_UpdateFeed_FullMethodName    = "/feeds.v1.FeedService/UpdateFeed"
	FeedService_ListArticles_FullMethodName  = "/feeds.v1.FeedService/ListArticles"
	FeedService_GetArticle_FullMethodName    = "/feeds.v1.FeedService/GetArticle"
	FeedService_SaveArticle_FullMethodName   = "/feeds.v1.FeedService/SaveArticle"
	FeedService_UpdateStats_FullMethodName   = "/feeds.v1.FeedService/UpdateStats"
	FeedService_Ready_FullMethodName         = "/feeds.v1.FeedService/Ready"
	FeedService_FeedTask_FullMethodName      = "/feeds.v1.FeedService/FeedTask"
)

// FeedServiceClient is the client API for FeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedServiceClient interface {
	GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*GetFeedResponse, error)
	GetUserFeed(ctx context.Context, in *GetUserFeedRequest, opts ...grpc.CallOption) (*GetUserFeedResponse, error)
	ListFeeds(ctx context.Context, in *ListFeedsRequest, opts ...grpc.CallOption) (*ListFeedsResponse, error)
	ListUserFeeds(ctx context.Context, in *ListUserFeedsRequest, opts ...grpc.CallOption) (*ListUserFeedsResponse, error)
	CreateFeed(ctx context.Context, in *CreateFeedRequest, opts ...grpc.CallOption) (*CreateFeedResponse, error)
	UpdateFeed(ctx context.Context, in *UpdateFeedRequest, opts ...grpc.CallOption) (*UpdateFeedResponse, error)
	ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*ListArticlesResponse, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleResponse, error)
	SaveArticle(ctx context.Context, in *SaveArticleRequest, opts ...grpc.CallOption) (*SaveArticleResponse, error)
	UpdateStats(ctx context.Context, in *UpdateStatsRequest, opts ...grpc.CallOption) (*UpdateStatsResponse, error)
	Ready(ctx context.Context, in *ReadyRequest, opts ...grpc.CallOption) (*ReadyResponse, error)
	// Deprecated: Do not use.
	// Deprecated: use graph service
	FeedTask(ctx context.Context, in *FeedTaskRequest, opts ...grpc.CallOption) (*FeedTaskResponse, error)
}

type feedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedServiceClient(cc grpc.ClientConnInterface) FeedServiceClient {
	return &feedServiceClient{cc}
}

func (c *feedServiceClient) GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*GetFeedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFeedResponse)
	err := c.cc.Invoke(ctx, FeedService_GetFeed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) GetUserFeed(ctx context.Context, in *GetUserFeedRequest, opts ...grpc.CallOption) (*GetUserFeedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserFeedResponse)
	err := c.cc.Invoke(ctx, FeedService_GetUserFeed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) ListFeeds(ctx context.Context, in *ListFeedsRequest, opts ...grpc.CallOption) (*ListFeedsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFeedsResponse)
	err := c.cc.Invoke(ctx, FeedService_ListFeeds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) ListUserFeeds(ctx context.Context, in *ListUserFeedsRequest, opts ...grpc.CallOption) (*ListUserFeedsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserFeedsResponse)
	err := c.cc.Invoke(ctx, FeedService_ListUserFeeds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) CreateFeed(ctx context.Context, in *CreateFeedRequest, opts ...grpc.CallOption) (*CreateFeedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateFeedResponse)
	err := c.cc.Invoke(ctx, FeedService_CreateFeed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) UpdateFeed(ctx context.Context, in *UpdateFeedRequest, opts ...grpc.CallOption) (*UpdateFeedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateFeedResponse)
	err := c.cc.Invoke(ctx, FeedService_UpdateFeed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*ListArticlesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListArticlesResponse)
	err := c.cc.Invoke(ctx, FeedService_ListArticles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticleResponse)
	err := c.cc.Invoke(ctx, FeedService_GetArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) SaveArticle(ctx context.Context, in *SaveArticleRequest, opts ...grpc.CallOption) (*SaveArticleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveArticleResponse)
	err := c.cc.Invoke(ctx, FeedService_SaveArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) UpdateStats(ctx context.Context, in *UpdateStatsRequest, opts ...grpc.CallOption) (*UpdateStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateStatsResponse)
	err := c.cc.Invoke(ctx, FeedService_UpdateStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) Ready(ctx context.Context, in *ReadyRequest, opts ...grpc.CallOption) (*ReadyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadyResponse)
	err := c.cc.Invoke(ctx, FeedService_Ready_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *feedServiceClient) FeedTask(ctx context.Context, in *FeedTaskRequest, opts ...grpc.CallOption) (*FeedTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FeedTaskResponse)
	err := c.cc.Invoke(ctx, FeedService_FeedTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServiceServer is the server API for FeedService service.
// All implementations must embed UnimplementedFeedServiceServer
// for forward compatibility.
type FeedServiceServer interface {
	GetFeed(context.Context, *GetFeedRequest) (*GetFeedResponse, error)
	GetUserFeed(context.Context, *GetUserFeedRequest) (*GetUserFeedResponse, error)
	ListFeeds(context.Context, *ListFeedsRequest) (*ListFeedsResponse, error)
	ListUserFeeds(context.Context, *ListUserFeedsRequest) (*ListUserFeedsResponse, error)
	CreateFeed(context.Context, *CreateFeedRequest) (*CreateFeedResponse, error)
	UpdateFeed(context.Context, *UpdateFeedRequest) (*UpdateFeedResponse, error)
	ListArticles(context.Context, *ListArticlesRequest) (*ListArticlesResponse, error)
	GetArticle(context.Context, *GetArticleRequest) (*GetArticleResponse, error)
	SaveArticle(context.Context, *SaveArticleRequest) (*SaveArticleResponse, error)
	UpdateStats(context.Context, *UpdateStatsRequest) (*UpdateStatsResponse, error)
	Ready(context.Context, *ReadyRequest) (*ReadyResponse, error)
	// Deprecated: Do not use.
	// Deprecated: use graph service
	FeedTask(context.Context, *FeedTaskRequest) (*FeedTaskResponse, error)
	mustEmbedUnimplementedFeedServiceServer()
}

// UnimplementedFeedServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFeedServiceServer struct{}

func (UnimplementedFeedServiceServer) GetFeed(context.Context, *GetFeedRequest) (*GetFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
func (UnimplementedFeedServiceServer) GetUserFeed(context.Context, *GetUserFeedRequest) (*GetUserFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFeed not implemented")
}
func (UnimplementedFeedServiceServer) ListFeeds(context.Context, *ListFeedsRequest) (*ListFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeeds not implemented")
}
func (UnimplementedFeedServiceServer) ListUserFeeds(context.Context, *ListUserFeedsRequest) (*ListUserFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserFeeds not implemented")
}
func (UnimplementedFeedServiceServer) CreateFeed(context.Context, *CreateFeedRequest) (*CreateFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeed not implemented")
}
func (UnimplementedFeedServiceServer) UpdateFeed(context.Context, *UpdateFeedRequest) (*UpdateFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFeed not implemented")
}
func (UnimplementedFeedServiceServer) ListArticles(context.Context, *ListArticlesRequest) (*ListArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}
func (UnimplementedFeedServiceServer) GetArticle(context.Context, *GetArticleRequest) (*GetArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedFeedServiceServer) SaveArticle(context.Context, *SaveArticleRequest) (*SaveArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveArticle not implemented")
}
func (UnimplementedFeedServiceServer) UpdateStats(context.Context, *UpdateStatsRequest) (*UpdateStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStats not implemented")
}
func (UnimplementedFeedServiceServer) Ready(context.Context, *ReadyRequest) (*ReadyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ready not implemented")
}
func (UnimplementedFeedServiceServer) FeedTask(context.Context, *FeedTaskRequest) (*FeedTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedTask not implemented")
}
func (UnimplementedFeedServiceServer) mustEmbedUnimplementedFeedServiceServer() {}
func (UnimplementedFeedServiceServer) testEmbeddedByValue()                     {}

// UnsafeFeedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedServiceServer will
// result in compilation errors.
type UnsafeFeedServiceServer interface {
	mustEmbedUnimplementedFeedServiceServer()
}

func RegisterFeedServiceServer(s grpc.ServiceRegistrar, srv FeedServiceServer) {
	// If the following call pancis, it indicates UnimplementedFeedServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FeedService_ServiceDesc, srv)
}

func _FeedService_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_GetFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).GetFeed(ctx, req.(*GetFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_GetUserFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).GetUserFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_GetUserFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).GetUserFeed(ctx, req.(*GetUserFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_ListFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).ListFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_ListFeeds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).ListFeeds(ctx, req.(*ListFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_ListUserFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).ListUserFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_ListUserFeeds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).ListUserFeeds(ctx, req.(*ListUserFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_CreateFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).CreateFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_CreateFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).CreateFeed(ctx, req.(*CreateFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_UpdateFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).UpdateFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_UpdateFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).UpdateFeed(ctx, req.(*UpdateFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_ListArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).ListArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_ListArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).ListArticles(ctx, req.(*ListArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_GetArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_SaveArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).SaveArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_SaveArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).SaveArticle(ctx, req.(*SaveArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_UpdateStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).UpdateStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_UpdateStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).UpdateStats(ctx, req.(*UpdateStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_Ready_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).Ready(ctx, req.(*ReadyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_FeedTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).FeedTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_FeedTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).FeedTask(ctx, req.(*FeedTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedService_ServiceDesc is the grpc.ServiceDesc for FeedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feeds.v1.FeedService",
	HandlerType: (*FeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeed",
			Handler:    _FeedService_GetFeed_Handler,
		},
		{
			MethodName: "GetUserFeed",
			Handler:    _FeedService_GetUserFeed_Handler,
		},
		{
			MethodName: "ListFeeds",
			Handler:    _FeedService_ListFeeds_Handler,
		},
		{
			MethodName: "ListUserFeeds",
			Handler:    _FeedService_ListUserFeeds_Handler,
		},
		{
			MethodName: "CreateFeed",
			Handler:    _FeedService_CreateFeed_Handler,
		},
		{
			MethodName: "UpdateFeed",
			Handler:    _FeedService_UpdateFeed_Handler,
		},
		{
			MethodName: "ListArticles",
			Handler:    _FeedService_ListArticles_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _FeedService_GetArticle_Handler,
		},
		{
			MethodName: "SaveArticle",
			Handler:    _FeedService_SaveArticle_Handler,
		},
		{
			MethodName: "UpdateStats",
			Handler:    _FeedService_UpdateStats_Handler,
		},
		{
			MethodName: "Ready",
			Handler:    _FeedService_Ready_Handler,
		},
		{
			MethodName: "FeedTask",
			Handler:    _FeedService_FeedTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feeds/v1/service.proto",
}
