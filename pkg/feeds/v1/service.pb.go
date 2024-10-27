// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: feeds/v1/service.proto

package feeds

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Feed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Feed) Reset() {
	*x = Feed{}
	mi := &file_feeds_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Feed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feed) ProtoMessage() {}

func (x *Feed) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feed.ProtoReflect.Descriptor instead.
func (*Feed) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *Feed) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Feed) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Feed) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content     string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	FeedId      string `protobuf:"bytes,4,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
	Url         string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	ImageUrl    string `protobuf:"bytes,6,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Preview     string `protobuf:"bytes,7,opt,name=preview,proto3" json:"preview,omitempty"`
	Guid        string `protobuf:"bytes,8,opt,name=guid,proto3" json:"guid,omitempty"`
	AuthorName  string `protobuf:"bytes,9,opt,name=authorName,proto3" json:"authorName,omitempty"`
	AuthorEmail string `protobuf:"bytes,10,opt,name=authorEmail,proto3" json:"authorEmail,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	mi := &file_feeds_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *Article) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Article) GetFeedId() string {
	if x != nil {
		return x.FeedId
	}
	return ""
}

func (x *Article) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Article) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Article) GetPreview() string {
	if x != nil {
		return x.Preview
	}
	return ""
}

func (x *Article) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *Article) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *Article) GetAuthorEmail() string {
	if x != nil {
		return x.AuthorEmail
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_feeds_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{2}
}

type GetFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFeedRequest) Reset() {
	*x = GetFeedRequest{}
	mi := &file_feeds_v1_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedRequest) ProtoMessage() {}

func (x *GetFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedRequest.ProtoReflect.Descriptor instead.
func (*GetFeedRequest) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetFeedRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feed *Feed `protobuf:"bytes,1,opt,name=feed,proto3" json:"feed,omitempty"`
}

func (x *GetFeedResponse) Reset() {
	*x = GetFeedResponse{}
	mi := &file_feeds_v1_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedResponse) ProtoMessage() {}

func (x *GetFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedResponse.ProtoReflect.Descriptor instead.
func (*GetFeedResponse) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetFeedResponse) GetFeed() *Feed {
	if x != nil {
		return x.Feed
	}
	return nil
}

type ListFeedsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListFeedsRequest) Reset() {
	*x = ListFeedsRequest{}
	mi := &file_feeds_v1_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFeedsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeedsRequest) ProtoMessage() {}

func (x *ListFeedsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeedsRequest.ProtoReflect.Descriptor instead.
func (*ListFeedsRequest) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{5}
}

type ListFeedsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feeds []*Feed `protobuf:"bytes,1,rep,name=feeds,proto3" json:"feeds,omitempty"`
}

func (x *ListFeedsResponse) Reset() {
	*x = ListFeedsResponse{}
	mi := &file_feeds_v1_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFeedsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeedsResponse) ProtoMessage() {}

func (x *ListFeedsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeedsResponse.ProtoReflect.Descriptor instead.
func (*ListFeedsResponse) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListFeedsResponse) GetFeeds() []*Feed {
	if x != nil {
		return x.Feeds
	}
	return nil
}

type CreateFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feed *CreateFeedRequest_Feed `protobuf:"bytes,1,opt,name=feed,proto3" json:"feed,omitempty"`
}

func (x *CreateFeedRequest) Reset() {
	*x = CreateFeedRequest{}
	mi := &file_feeds_v1_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFeedRequest) ProtoMessage() {}

func (x *CreateFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFeedRequest.ProtoReflect.Descriptor instead.
func (*CreateFeedRequest) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *CreateFeedRequest) GetFeed() *CreateFeedRequest_Feed {
	if x != nil {
		return x.Feed
	}
	return nil
}

type CreateFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateFeedResponse) Reset() {
	*x = CreateFeedResponse{}
	mi := &file_feeds_v1_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFeedResponse) ProtoMessage() {}

func (x *CreateFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFeedResponse.ProtoReflect.Descriptor instead.
func (*CreateFeedResponse) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *CreateFeedResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feed *UpdateFeedRequest_Feed `protobuf:"bytes,1,opt,name=feed,proto3" json:"feed,omitempty"`
}

func (x *UpdateFeedRequest) Reset() {
	*x = UpdateFeedRequest{}
	mi := &file_feeds_v1_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeedRequest) ProtoMessage() {}

func (x *UpdateFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeedRequest.ProtoReflect.Descriptor instead.
func (*UpdateFeedRequest) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateFeedRequest) GetFeed() *UpdateFeedRequest_Feed {
	if x != nil {
		return x.Feed
	}
	return nil
}

type UpdateFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateFeedResponse) Reset() {
	*x = UpdateFeedResponse{}
	mi := &file_feeds_v1_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeedResponse) ProtoMessage() {}

func (x *UpdateFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeedResponse.ProtoReflect.Descriptor instead.
func (*UpdateFeedResponse) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{10}
}

type ListArticlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeedId string `protobuf:"bytes,1,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
}

func (x *ListArticlesRequest) Reset() {
	*x = ListArticlesRequest{}
	mi := &file_feeds_v1_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticlesRequest) ProtoMessage() {}

func (x *ListArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticlesRequest.ProtoReflect.Descriptor instead.
func (*ListArticlesRequest) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{11}
}

func (x *ListArticlesRequest) GetFeedId() string {
	if x != nil {
		return x.FeedId
	}
	return ""
}

type ListArticlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles []*Article `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
}

func (x *ListArticlesResponse) Reset() {
	*x = ListArticlesResponse{}
	mi := &file_feeds_v1_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListArticlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticlesResponse) ProtoMessage() {}

func (x *ListArticlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticlesResponse.ProtoReflect.Descriptor instead.
func (*ListArticlesResponse) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{12}
}

func (x *ListArticlesResponse) GetArticles() []*Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

type GetArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetArticleRequest) Reset() {
	*x = GetArticleRequest{}
	mi := &file_feeds_v1_service_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleRequest) ProtoMessage() {}

func (x *GetArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleRequest.ProtoReflect.Descriptor instead.
func (*GetArticleRequest) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{13}
}

func (x *GetArticleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *GetArticleResponse) Reset() {
	*x = GetArticleResponse{}
	mi := &file_feeds_v1_service_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleResponse) ProtoMessage() {}

func (x *GetArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleResponse.ProtoReflect.Descriptor instead.
func (*GetArticleResponse) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{14}
}

func (x *GetArticleResponse) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type CreateFeedRequest_Feed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url  string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateFeedRequest_Feed) Reset() {
	*x = CreateFeedRequest_Feed{}
	mi := &file_feeds_v1_service_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFeedRequest_Feed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFeedRequest_Feed) ProtoMessage() {}

func (x *CreateFeedRequest_Feed) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFeedRequest_Feed.ProtoReflect.Descriptor instead.
func (*CreateFeedRequest_Feed) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{7, 0}
}

func (x *CreateFeedRequest_Feed) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateFeedRequest_Feed) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateFeedRequest_Feed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Deprecated: Marked as deprecated in feeds/v1/service.proto.
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateFeedRequest_Feed) Reset() {
	*x = UpdateFeedRequest_Feed{}
	mi := &file_feeds_v1_service_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFeedRequest_Feed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeedRequest_Feed) ProtoMessage() {}

func (x *UpdateFeedRequest_Feed) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_v1_service_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeedRequest_Feed.ProtoReflect.Descriptor instead.
func (*UpdateFeedRequest_Feed) Descriptor() ([]byte, []int) {
	return file_feeds_v1_service_proto_rawDescGZIP(), []int{9, 0}
}

func (x *UpdateFeedRequest_Feed) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Deprecated: Marked as deprecated in feeds/v1/service.proto.
func (x *UpdateFeedRequest_Feed) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdateFeedRequest_Feed) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_feeds_v1_service_proto protoreflect.FileDescriptor

var file_feeds_v1_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x22, 0x3c, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x80, 0x02, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x65, 0x65, 0x64, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x12, 0x0a,
	0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52,
	0x04, 0x66, 0x65, 0x65, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x65,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x05, 0x66, 0x65, 0x65, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x05, 0x66,
	0x65, 0x65, 0x64, 0x73, 0x22, 0x77, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x66, 0x65, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x04, 0x66, 0x65, 0x65, 0x64, 0x1a,
	0x2c, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x66, 0x65, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x04, 0x66, 0x65, 0x65, 0x64, 0x1a,
	0x40, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x65, 0x65, 0x64, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x23,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x32, 0xbd, 0x03, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65,
	0x64, 0x12, 0x18, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65,
	0x65, 0x64, 0x73, 0x12, 0x1a, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x12, 0x1b, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x65, 0x65, 0x64, 0x12, 0x1b, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x1d,
	0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x69, 0x63, 0x62, 0x75, 0x74, 0x65, 0x72, 0x61, 0x2f,
	0x61, 0x6d, 0x61, 0x6c, 0x67, 0x61, 0x6d, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feeds_v1_service_proto_rawDescOnce sync.Once
	file_feeds_v1_service_proto_rawDescData = file_feeds_v1_service_proto_rawDesc
)

func file_feeds_v1_service_proto_rawDescGZIP() []byte {
	file_feeds_v1_service_proto_rawDescOnce.Do(func() {
		file_feeds_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_feeds_v1_service_proto_rawDescData)
	})
	return file_feeds_v1_service_proto_rawDescData
}

var file_feeds_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_feeds_v1_service_proto_goTypes = []any{
	(*Feed)(nil),                   // 0: feeds.v1.Feed
	(*Article)(nil),                // 1: feeds.v1.Article
	(*Empty)(nil),                  // 2: feeds.v1.Empty
	(*GetFeedRequest)(nil),         // 3: feeds.v1.GetFeedRequest
	(*GetFeedResponse)(nil),        // 4: feeds.v1.GetFeedResponse
	(*ListFeedsRequest)(nil),       // 5: feeds.v1.ListFeedsRequest
	(*ListFeedsResponse)(nil),      // 6: feeds.v1.ListFeedsResponse
	(*CreateFeedRequest)(nil),      // 7: feeds.v1.CreateFeedRequest
	(*CreateFeedResponse)(nil),     // 8: feeds.v1.CreateFeedResponse
	(*UpdateFeedRequest)(nil),      // 9: feeds.v1.UpdateFeedRequest
	(*UpdateFeedResponse)(nil),     // 10: feeds.v1.UpdateFeedResponse
	(*ListArticlesRequest)(nil),    // 11: feeds.v1.ListArticlesRequest
	(*ListArticlesResponse)(nil),   // 12: feeds.v1.ListArticlesResponse
	(*GetArticleRequest)(nil),      // 13: feeds.v1.GetArticleRequest
	(*GetArticleResponse)(nil),     // 14: feeds.v1.GetArticleResponse
	(*CreateFeedRequest_Feed)(nil), // 15: feeds.v1.CreateFeedRequest.Feed
	(*UpdateFeedRequest_Feed)(nil), // 16: feeds.v1.UpdateFeedRequest.Feed
}
var file_feeds_v1_service_proto_depIdxs = []int32{
	0,  // 0: feeds.v1.GetFeedResponse.feed:type_name -> feeds.v1.Feed
	0,  // 1: feeds.v1.ListFeedsResponse.feeds:type_name -> feeds.v1.Feed
	15, // 2: feeds.v1.CreateFeedRequest.feed:type_name -> feeds.v1.CreateFeedRequest.Feed
	16, // 3: feeds.v1.UpdateFeedRequest.feed:type_name -> feeds.v1.UpdateFeedRequest.Feed
	1,  // 4: feeds.v1.ListArticlesResponse.articles:type_name -> feeds.v1.Article
	1,  // 5: feeds.v1.GetArticleResponse.article:type_name -> feeds.v1.Article
	3,  // 6: feeds.v1.FeedService.GetFeed:input_type -> feeds.v1.GetFeedRequest
	5,  // 7: feeds.v1.FeedService.ListFeeds:input_type -> feeds.v1.ListFeedsRequest
	7,  // 8: feeds.v1.FeedService.CreateFeed:input_type -> feeds.v1.CreateFeedRequest
	9,  // 9: feeds.v1.FeedService.UpdateFeed:input_type -> feeds.v1.UpdateFeedRequest
	11, // 10: feeds.v1.FeedService.ListArticles:input_type -> feeds.v1.ListArticlesRequest
	13, // 11: feeds.v1.FeedService.GetArticle:input_type -> feeds.v1.GetArticleRequest
	4,  // 12: feeds.v1.FeedService.GetFeed:output_type -> feeds.v1.GetFeedResponse
	6,  // 13: feeds.v1.FeedService.ListFeeds:output_type -> feeds.v1.ListFeedsResponse
	8,  // 14: feeds.v1.FeedService.CreateFeed:output_type -> feeds.v1.CreateFeedResponse
	10, // 15: feeds.v1.FeedService.UpdateFeed:output_type -> feeds.v1.UpdateFeedResponse
	12, // 16: feeds.v1.FeedService.ListArticles:output_type -> feeds.v1.ListArticlesResponse
	14, // 17: feeds.v1.FeedService.GetArticle:output_type -> feeds.v1.GetArticleResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_feeds_v1_service_proto_init() }
func file_feeds_v1_service_proto_init() {
	if File_feeds_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feeds_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feeds_v1_service_proto_goTypes,
		DependencyIndexes: file_feeds_v1_service_proto_depIdxs,
		MessageInfos:      file_feeds_v1_service_proto_msgTypes,
	}.Build()
	File_feeds_v1_service_proto = out.File
	file_feeds_v1_service_proto_rawDesc = nil
	file_feeds_v1_service_proto_goTypes = nil
	file_feeds_v1_service_proto_depIdxs = nil
}
