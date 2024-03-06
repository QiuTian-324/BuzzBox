// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: article.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Article_PublishArticle_FullMethodName = "/pb.Article/PublishArticle"
	Article_GetArticleList_FullMethodName = "/pb.Article/GetArticleList"
	Article_ArticleDelete_FullMethodName  = "/pb.Article/ArticleDelete"
	Article_ArticleDetail_FullMethodName  = "/pb.Article/ArticleDetail"
)

// ArticleClient is the client API for Article service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleClient interface {
	PublishArticle(ctx context.Context, in *PublishArticleRequest, opts ...grpc.CallOption) (*PublishArticleResponse, error)
	GetArticleList(ctx context.Context, in *GetArticleListRequest, opts ...grpc.CallOption) (*GetArticleListResponse, error)
	ArticleDelete(ctx context.Context, in *ArticleDeleteRequest, opts ...grpc.CallOption) (*ArticleDeleteResponse, error)
	ArticleDetail(ctx context.Context, in *ArticleDetailRequest, opts ...grpc.CallOption) (*ArticleDetailResponse, error)
}

type articleClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleClient(cc grpc.ClientConnInterface) ArticleClient {
	return &articleClient{cc}
}

func (c *articleClient) PublishArticle(ctx context.Context, in *PublishArticleRequest, opts ...grpc.CallOption) (*PublishArticleResponse, error) {
	out := new(PublishArticleResponse)
	err := c.cc.Invoke(ctx, Article_PublishArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) GetArticleList(ctx context.Context, in *GetArticleListRequest, opts ...grpc.CallOption) (*GetArticleListResponse, error) {
	out := new(GetArticleListResponse)
	err := c.cc.Invoke(ctx, Article_GetArticleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) ArticleDelete(ctx context.Context, in *ArticleDeleteRequest, opts ...grpc.CallOption) (*ArticleDeleteResponse, error) {
	out := new(ArticleDeleteResponse)
	err := c.cc.Invoke(ctx, Article_ArticleDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) ArticleDetail(ctx context.Context, in *ArticleDetailRequest, opts ...grpc.CallOption) (*ArticleDetailResponse, error) {
	out := new(ArticleDetailResponse)
	err := c.cc.Invoke(ctx, Article_ArticleDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServer is the server API for Article service.
// All implementations must embed UnimplementedArticleServer
// for forward compatibility
type ArticleServer interface {
	PublishArticle(context.Context, *PublishArticleRequest) (*PublishArticleResponse, error)
	GetArticleList(context.Context, *GetArticleListRequest) (*GetArticleListResponse, error)
	ArticleDelete(context.Context, *ArticleDeleteRequest) (*ArticleDeleteResponse, error)
	ArticleDetail(context.Context, *ArticleDetailRequest) (*ArticleDetailResponse, error)
	mustEmbedUnimplementedArticleServer()
}

// UnimplementedArticleServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServer struct {
}

func (UnimplementedArticleServer) PublishArticle(context.Context, *PublishArticleRequest) (*PublishArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishArticle not implemented")
}
func (UnimplementedArticleServer) GetArticleList(context.Context, *GetArticleListRequest) (*GetArticleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleList not implemented")
}
func (UnimplementedArticleServer) ArticleDelete(context.Context, *ArticleDeleteRequest) (*ArticleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleDelete not implemented")
}
func (UnimplementedArticleServer) ArticleDetail(context.Context, *ArticleDetailRequest) (*ArticleDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleDetail not implemented")
}
func (UnimplementedArticleServer) mustEmbedUnimplementedArticleServer() {}

// UnsafeArticleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServer will
// result in compilation errors.
type UnsafeArticleServer interface {
	mustEmbedUnimplementedArticleServer()
}

func RegisterArticleServer(s grpc.ServiceRegistrar, srv ArticleServer) {
	s.RegisterService(&Article_ServiceDesc, srv)
}

func _Article_PublishArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).PublishArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_PublishArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).PublishArticle(ctx, req.(*PublishArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_GetArticleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).GetArticleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_GetArticleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).GetArticleList(ctx, req.(*GetArticleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_ArticleDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).ArticleDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_ArticleDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).ArticleDelete(ctx, req.(*ArticleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_ArticleDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).ArticleDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_ArticleDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).ArticleDetail(ctx, req.(*ArticleDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Article_ServiceDesc is the grpc.ServiceDesc for Article service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Article_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Article",
	HandlerType: (*ArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishArticle",
			Handler:    _Article_PublishArticle_Handler,
		},
		{
			MethodName: "GetArticleList",
			Handler:    _Article_GetArticleList_Handler,
		},
		{
			MethodName: "ArticleDelete",
			Handler:    _Article_ArticleDelete_Handler,
		},
		{
			MethodName: "ArticleDetail",
			Handler:    _Article_ArticleDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}