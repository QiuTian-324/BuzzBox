// Code generated by goctl. DO NOT EDIT.
// Source: article.proto

package server

import (
	"context"

	"BuzzBox/service/article/rpc/internal/logic"
	"BuzzBox/service/article/rpc/internal/svc"
	"BuzzBox/service/article/rpc/pb"
)

type ArticleServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedArticleServer
}

func NewArticleServer(svcCtx *svc.ServiceContext) *ArticleServer {
	return &ArticleServer{
		svcCtx: svcCtx,
	}
}

func (s *ArticleServer) PublishArticle(ctx context.Context, in *pb.PublishArticleRequest) (*pb.PublishArticleResponse, error) {
	l := logic.NewPublishArticleLogic(ctx, s.svcCtx)
	return l.PublishArticle(in)
}

func (s *ArticleServer) GetArticleList(ctx context.Context, in *pb.GetArticleListRequest) (*pb.GetArticleListResponse, error) {
	l := logic.NewGetArticleListLogic(ctx, s.svcCtx)
	return l.GetArticleList(in)
}

func (s *ArticleServer) ArticleDelete(ctx context.Context, in *pb.ArticleDeleteRequest) (*pb.ArticleDeleteResponse, error) {
	l := logic.NewArticleDeleteLogic(ctx, s.svcCtx)
	return l.ArticleDelete(in)
}

func (s *ArticleServer) ArticleDetail(ctx context.Context, in *pb.ArticleDetailRequest) (*pb.ArticleDetailResponse, error) {
	l := logic.NewArticleDetailLogic(ctx, s.svcCtx)
	return l.ArticleDetail(in)
}