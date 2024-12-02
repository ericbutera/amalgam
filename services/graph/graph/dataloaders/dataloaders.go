package dataloaders

import (
	"context"

	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	// "github.com/graph-gophers/dataloader"
	dataloader "github.com/graph-gophers/dataloader/v7"
)

// https://github.com/graphql/dataloader
// > DataLoader serves requests to many different users with different access permissions. It may be dangerous to use one cache across many users, and is encouraged to create a new DataLoader per request
// https://github.com/graph-gophers/dataloader

type (
	PbUserArticle = pb.GetUserArticlesResponse_UserArticle
	ArticleRes    = dataloader.Result[*PbUserArticle]
	ArticleID     = string
)

func NewUserArticleLoader(rpcClient pb.FeedServiceClient, userID string) *dataloader.Loader[ArticleID, *PbUserArticle] {
	batchFunc := func(ctx context.Context, articleIDs []ArticleID) []*ArticleRes {
		var articles map[ArticleID]*PbUserArticle
		resp, err := rpcClient.GetUserArticles(ctx, &pb.GetUserArticlesRequest{
			User:       &pb.User{Id: userID},
			ArticleIds: articleIDs,
		})
		if err == nil {
			articles = resp.GetArticles()
		}

		results := make([]*ArticleRes, len(articleIDs))
		for i, articleID := range articleIDs {
			if article, ok := articles[articleID]; ok {
				results[i] = &ArticleRes{Data: article}
			} else {
				results[i] = &ArticleRes{Data: nil}
			}
		}
		return results
	}

	cache := &dataloader.NoCache[ArticleID, *PbUserArticle]{}
	return dataloader.NewBatchedLoader(batchFunc, dataloader.WithCache(cache))
}
