package goverter

// TODO: research https://goverter.jmattheis.de/reference/extend
// this might be a way to lower the amount of custom conversions

import (
	"time"

	db_model "github.com/ericbutera/amalgam/internal/db/models"
	svc_model "github.com/ericbutera/amalgam/internal/service/models"
	gql_client "github.com/ericbutera/amalgam/pkg/clients/graphql"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	gql_model "github.com/ericbutera/amalgam/services/graph/graph/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:output:file ./generated.gen.go
// goverter:output:package github.com/ericbutera/amalgam/internal/goverter
type Converter interface {
	// goverter:ignoreMissing
	// goverter:map ID ID
	// goverter:map UpdatedAt | Time
	ConvertBase(struct {
		ID        string
		UpdatedAt time.Time
	}) *db_model.Base

	// goverter:map Base.ID ID
	DbToServiceFeed(*db_model.Feed) *svc_model.Feed
	// goverter:ignore Base
	ServiceToDbFeed(*svc_model.Feed) *db_model.Feed
	// goverter:map Base.ID ID
	// goverter:map Base.UpdatedAt UpdatedAt | Time
	DbToServiceArticle(*db_model.Article) *svc_model.Article
	// goverter:matchIgnoreCase
	// goverter:ignore Base Feed
	ServiceToDbArticle(*svc_model.Article) *db_model.Article

	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	// goverter:map FeedID ID
	// goverter:map CreatedAt CreatedAt | Time
	// goverter:map ViewedAt ViewedAt | Time
	// goverter:map UnreadStartAt UnreadStartAt | Time
	// goverter:map UnreadCount UnreadCount | Int32ToInt
	ServiceToGraphFeed(*svc_model.UserFeed) *gql_model.Feed
	// goverter:ignoreMissing
	// goverter:map UpdatedAt UpdatedAt | Time
	ServiceToGraphArticle(*svc_model.Article) *gql_model.Article
	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	GraphClientToApiFeedGet(*gql_client.GetFeedFeed) *svc_model.Feed
	// goverter:matchIgnoreCase
	// goverter:map UpdatedAt UpdatedAt | Time
	GraphClientToApiArticle(*gql_client.GetArticleArticle) *svc_model.Article
	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	// goverter:map UpdatedAt UpdatedAt | Time
	GraphClientToApiArticleList(*gql_client.ListArticlesArticlesArticlesResponseArticlesArticle) *svc_model.Article

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
	// goverter:ignoreMissing
	ProtoToServiceFeed(*pb.Feed) *svc_model.Feed
	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	// goverter:map CreatedAt | ProtoTimestampToTime
	// goverter:map ViewedAt | ProtoTimestampToTime
	// goverter:map UnreadStartAt | ProtoTimestampToTime
	ProtoToServiceUserFeed(*pb.UserFeed) *svc_model.UserFeed
	// goverter:matchIgnoreCase
	// goverter:ignore state sizeCache unknownFields
	// goverter:map CreatedAt | TimeToProtoTimestamp
	// goverter:map UnreadStartAt | TimeToProtoTimestamp
	// goverter:map ViewedAt | TimeToProtoTimestamp
	ServiceToProtoUserFeed(*svc_model.UserFeed) *pb.UserFeed
	// goverter:matchIgnoreCase
	// goverter:map UpdatedAt | ProtoTimestampToTime
	ProtoToServiceArticle(*pb.Article) *svc_model.Article
	// goverter:matchIgnoreCase
	// goverter:ignore state sizeCache unknownFields
	// goverter:map UpdatedAt | TimeToProtoTimestamp
	ServiceToProtoArticle(*svc_model.Article) *pb.Article
	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	// goverter:map FeedId ID
	// goverter:map CreatedAt | ProtoTimestampToTime
	// goverter:map ViewedAt | ProtoTimestampToTime
	// goverter:map UnreadStartAt | ProtoTimestampToTime
	// goverter:map UnreadCount | Int32ToInt
	ProtoUserFeedToGraphUserFeed(*pb.UserFeed) *gql_model.Feed
	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	// goverter:map UpdatedAt | ProtoTimestampToTime
	ProtoToGraphArticle(*pb.Article) *gql_model.Article
	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	// goverter:map ViewedAt | ProtoTimestampToTime
	ProtoToGraphUserArticle(*pb.GetUserArticlesResponse_UserArticle) *gql_model.UserArticle
	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	ProtoToGraphPagination(*pb.Pagination) *gql_model.Pagination
	// goverter:useZeroValueOnPointerInconsistency
	// goverter:matchIgnoreCase
	// goverter:ignoreMissing
	// goverter:ignore state sizeCache unknownFields
	// goverter:map Limit Limit | IntPtrToInt32
	GraphToProtoListOptions(*gql_model.ListOptions) *pb.ListOptions
	// goverter:ignore state sizeCache unknownFields
	// goverter:map ViewedAt | NillableTimeToProtoTimestamp
	ServiceToProtoUserArticle(*svc_model.UserArticle) *pb.GetUserArticlesResponse_UserArticle
}

func Time(t time.Time) time.Time {
	return t
}

func ProtoTimestampToTime(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{} // Return zero value if nil
	}
	return ts.AsTime()
}

func TimeToProtoTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func NillableTimeToProtoTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func Int32ToInt(i int32) int {
	return int(i)
}

func IntToInt32(i int) int32 {
	return int32(i)
}

func IntPtrToInt32(i *int) int32 {
	if i == nil {
		return 0 // Or provide a default value
	}
	return IntToInt32(*i)
}
