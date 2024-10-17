package registrar

import (
	"context"

	"github.com/opensearch-project/opensearch-go/v4"
	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
)

var _ Client = (*opensearchapi.Client)(nil)

type Client interface {
	SearchShards(ctx context.Context, req *opensearchapi.SearchShardsReq) (*opensearchapi.SearchShardsResp, error)
	UpdateByQuery(ctx context.Context, req opensearchapi.UpdateByQueryReq) (*opensearchapi.UpdateByQueryResp, error)
	Ping(ctx context.Context, req *opensearchapi.PingReq) (*opensearch.Response, error)
	MSearch(ctx context.Context, req opensearchapi.MSearchReq) (*opensearchapi.MSearchResp, error)
	Info(ctx context.Context, req *opensearchapi.InfoReq) (*opensearchapi.InfoResp, error)
	Update(ctx context.Context, req opensearchapi.UpdateReq) (*opensearchapi.UpdateResp, error)
	Reindex(ctx context.Context, req opensearchapi.ReindexReq) (*opensearchapi.ReindexResp, error)
	MTermvectors(ctx context.Context, req opensearchapi.MTermvectorsReq) (*opensearchapi.MTermvectorsResp, error)
	Termvectors(ctx context.Context, req opensearchapi.TermvectorsReq) (*opensearchapi.TermvectorsResp, error)
	Index(ctx context.Context, req opensearchapi.IndexReq) (*opensearchapi.IndexResp, error)
	MGet(ctx context.Context, req opensearchapi.MGetReq) (*opensearchapi.MGetResp, error)
	ReindexRethrottle(ctx context.Context, req opensearchapi.ReindexRethrottleReq) (*opensearchapi.ReindexRethrottleResp, error)
	Bulk(ctx context.Context, req opensearchapi.BulkReq) (*opensearchapi.BulkResp, error)
	MSearchTemplate(ctx context.Context, req opensearchapi.MSearchTemplateReq) (*opensearchapi.MSearchTemplateResp, error)
	Search(ctx context.Context, req *opensearchapi.SearchReq) (*opensearchapi.SearchResp, error)
	RankEval(ctx context.Context, req opensearchapi.RankEvalReq) (*opensearchapi.RankEvalResp, error)
	Aliases(ctx context.Context, req opensearchapi.AliasesReq) (*opensearchapi.AliasesResp, error)
	UpdateByQueryRethrottle(ctx context.Context, req opensearchapi.UpdateByQueryRethrottleReq) (*opensearchapi.UpdateByQueryRethrottleResp, error)
	RenderSearchTemplate(ctx context.Context, req opensearchapi.RenderSearchTemplateReq) (*opensearchapi.RenderSearchTemplateResp, error)
}
