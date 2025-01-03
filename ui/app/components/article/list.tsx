"use client";

import Link from "next/link";
import useArticles from "../../data/articles";
import { Pagination } from "@/app/types/pagination";
import Nav from "@/app/components/article/pagination";

interface ArticlesProps {
  feedId: string;
  pagination: Pagination;
}

export default function Articles({ feedId, pagination }: ArticlesProps) {
  const { loading, error, articles } = useArticles(feedId, pagination);

  if (error) return <div>failed to load articles</div>;
  if (loading) return <div>loading...</div>;
  if (!articles || articles?.articles.length === 0)
    return <div>no articles found</div>;

  return (
    <div className="overflow-x-auto">
      <ul>
        {articles?.articles.map((article) => (
          <li
            key={article.id}
            className={`border-b border-base-300 last:border-0 ${
              article.userArticle?.viewedAt
                ? "bg-article-read"
                : "bg-article-new"
            } `}
          >
            <Link
              href={`/articles/${article.id}`}
              className="block text-sm md:text-base lg:text-lg hover:bg-primary hover:text-white transition-colors duration-200"
            >
              {article.title}
            </Link>
          </li>
        ))}
      </ul>

      <Nav base={`/feeds/${feedId}/articles`} cursor={articles.cursor} />
    </div>
  );
}
