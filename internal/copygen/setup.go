package copygen

import (
	gql_model "github.com/ericbutera/amalgam/graph/graph/model"
	db_model "github.com/ericbutera/amalgam/internal/db/models"
	svc_model "github.com/ericbutera/amalgam/internal/service"
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

	CreateFeedToServiceFeed(*svc_model.Feed) *svc_model.Feed
}
