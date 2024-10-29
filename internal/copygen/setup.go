package copygen

import (
	// TODO: add API models
	gql_model "github.com/ericbutera/amalgam/graph/graph/model"
	gql_client "github.com/ericbutera/amalgam/graph/pkg/client"
	db_model "github.com/ericbutera/amalgam/internal/db/models"
	svc_model "github.com/ericbutera/amalgam/internal/service/models"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
)

type Copygen interface {
	DbToServiceFeed(*db_model.Feed) *svc_model.Feed
	ServiceToDbFeed(*svc_model.Feed) *db_model.Feed
	DbToServiceArticle(*db_model.Article) *svc_model.Article
	ServiceToDbArticle(*svc_model.Article) *db_model.Article

	GraphToServiceFeed(*gql_model.Feed) *svc_model.Feed
	ServiceToGraphFeed(*svc_model.Feed) *gql_model.Feed
	GraphToServiceArticle(*gql_model.Article) *svc_model.Article
	ServiceToGraphArticle(*svc_model.Article) *gql_model.Article

	// tag .* json
	GraphClientToApiFeedGet(*gql_client.GetFeedFeed) *svc_model.Feed
	// tag .* json
	GraphClientToApiArticle(*gql_client.GetArticleArticle) *svc_model.Article
	// tag .* json
	GraphClientToApiArticleList(*gql_client.ListArticlesArticlesArticle) *svc_model.Article

	// tag .* json
	ProtoCreateFeedToService(*pb.CreateFeedRequest_Feed) *svc_model.Feed
	// tag .* json
	ProtoUpdateFeedToService(*pb.UpdateFeedRequest_Feed) *svc_model.Feed
	// tag .* json
	ServiceToProtoFeed(*svc_model.Feed) *pb.Feed
	// TODO: feed busted (json is feed_id -> feedId)
	// tag .* json
	ProtoToServiceArticle(*pb.Article) *svc_model.Article
	// tag .* json
	ServiceToProtoArticle(*svc_model.Article) *pb.Article
}
