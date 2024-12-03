"use client";
import { getGraph } from "@/app/lib/fetch";
import useArticle from "@/app/data/article";
import ArticleDetails from "@/app/components/article/details";

interface PageProps {
  params: {
    id: string;
  };
}

export default function Page({ params }: PageProps) {
  const articleId = params.id;

  const { loading, error, article } = useArticle(articleId);

  if (error) return <div>An error has occurred.</div>;
  if (loading) <div>Loading...</div>;
  if (!article) return <div>Article not found.</div>;

  if (article && !article?.userArticle?.viewedAt) {
    getGraph().MarkArticleRead({ id: article.id });
  }

  return <ArticleDetails id={articleId} />;
}
