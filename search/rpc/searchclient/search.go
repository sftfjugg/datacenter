// Code generated by goctl. DO NOT EDIT!
// Source: search.proto

//go:generate mockgen -destination ./search_mock.go -package searchclient -source $GOFILE

package searchclient

import (
	"context"

	"datacenter/search/rpc/search"

	"github.com/zeromicro/go-zero/zrpc"
)

type (
	LimitReq    = search.LimitReq
	SearchReq   = search.SearchReq
	ArticleResp = search.ArticleResp
	ArticleReq  = search.ArticleReq
	Request     = search.Request
	Response    = search.Response

	Search interface {
		ArticleInit(ctx context.Context, in *Request) (*Response, error)
		ArticleStore(ctx context.Context, in *ArticleReq) (*Response, error)
		ArticleSearch(ctx context.Context, in *SearchReq) (*ArticleResp, error)
	}

	defaultSearch struct {
		cli zrpc.Client
	}
)

func NewSearch(cli zrpc.Client) Search {
	return &defaultSearch{
		cli: cli,
	}
}

func (m *defaultSearch) ArticleInit(ctx context.Context, in *Request) (*Response, error) {
	client := search.NewSearchClient(m.cli.Conn())
	return client.ArticleInit(ctx, in)
}

func (m *defaultSearch) ArticleStore(ctx context.Context, in *ArticleReq) (*Response, error) {
	client := search.NewSearchClient(m.cli.Conn())
	return client.ArticleStore(ctx, in)
}

func (m *defaultSearch) ArticleSearch(ctx context.Context, in *SearchReq) (*ArticleResp, error) {
	client := search.NewSearchClient(m.cli.Conn())
	return client.ArticleSearch(ctx, in)
}
