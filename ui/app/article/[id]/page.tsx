"use client";

import useArticle from "../../data/article";
import Feeds from "@/app/components/feed/list";
import ArticleDetails from "@/app/components/article/details";

interface PageProps {
  params: {
    id: string;
  };
}

export default function Page({ params }: PageProps) {
  const articleId = params.id;

  const { loading, error, article } = useArticle(params.id);

  if (error) return <div>An error has occurred.</div>;
  if (loading) <div>Loading...</div>;
  if (!article) return <div>Article not found.</div>;

  return (
    <div className="flex">
      <Feeds feedId={article.feedId} />
      <div className="flex-1">
        <ArticleDetails id={articleId} />
      </div>
    </div>
  );
}
