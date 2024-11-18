"use client";

import useArticle from "../../data/article";
import Link from "next/link";

interface ArticleDetailsProps {
  id: string;
}

export default function ArticleDetails({ id }: ArticleDetailsProps) {
  const { loading, error, article } = useArticle(id);

  if (error) return <div>An error has occurred while loading the article.</div>;
  if (loading) return <div>Loading...</div>;
  if (!article) return <div>Article not found.</div>;

  return (
    <div className="p-4">
      <h1 className="text-3xl font-bold mb-4">{article.title}</h1>
      <p className="text-gray-600 mb-6">
        2024-11-11
        {/* Published on {new Date(article.date).toLocaleDateString()} */}
      </p>
      <div className="mb-6">{article.content}</div>
      <Link
        href={`/feeds/${article.feedId}/articles`}
        className="text-blue-500 hover:underline"
      >
        Back to Feed
      </Link>
    </div>
  );
}
