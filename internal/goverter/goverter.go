package goverter

import (
	gql_model "github.com/ericbutera/amalgam/graph/graph/model"
	db_model "github.com/ericbutera/amalgam/internal/db/models"
	svc_model "github.com/ericbutera/amalgam/internal/service/models"
	gql_client "github.com/ericbutera/amalgam/pkg/clients/graphql"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
)

// goverter:converter
// goverter:output:file ./generated.gen.go
// goverter:output:package goverter
type Converter interface {
	// goverter:map Base.ID ID
	DbToServiceFeed(*db_model.Feed) *svc_model.Feed
	// goverter:ignore Base
	ServiceToDbFeed(*svc_model.Feed) *db_model.Feed
	// goverter:ignoreMissing
	// goverter:map ID ID
	ConvertBase(struct{ ID string }) *db_model.Base
	// goverter:map Base.ID ID
	DbToServiceArticle(*db_model.Article) *svc_model.Article
	// goverter:ignore Base Feed
	ServiceToDbArticle(*svc_model.Article) *db_model.Article
	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	GraphToServiceFeed(*gql_model.Feed) *svc_model.Feed
	ServiceToGraphFeed(*svc_model.Feed) *gql_model.Feed
	// goverter:useZeroValueOnPointerInconsistency
	GraphToServiceArticle(*gql_model.Article) *svc_model.Article
	ServiceToGraphArticle(*svc_model.Article) *gql_model.Article
	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	GraphClientToApiFeedGet(*gql_client.GetFeedFeed) *svc_model.Feed
	// goverter:matchIgnoreCase
	GraphClientToApiArticle(*gql_client.GetArticleArticle) *svc_model.Article
	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	GraphClientToApiArticleList(*gql_client.ListArticlesArticlesArticle) *svc_model.Article
	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	ProtoCreateFeedToService(*pb.CreateFeedRequest_Feed) *svc_model.Feed
	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	ProtoUpdateFeedToService(*pb.UpdateFeedRequest_Feed) *svc_model.Feed
	// goverter:matchIgnoreCase
	// goverter:ignore state sizeCache unknownFields
	ServiceToProtoFeed(*svc_model.Feed) *pb.Feed
	// goverter:matchIgnoreCase
	ProtoToServiceArticle(*pb.Article) *svc_model.Article
	// goverter:matchIgnoreCase
	// goverter:ignore state sizeCache unknownFields
	ServiceToProtoArticle(*svc_model.Article) *pb.Article
}
