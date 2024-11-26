"use client";

import Link from "next/link";
import useArticles from "../../data/articles";

interface ArticlesProps {
  id: string;
}

export default function Articles({ id }: ArticlesProps) {
  const { loading, error, articles } = useArticles(id);

  if (error) return <div>An error has occurred.</div>;
  if (loading) <div>Loading...</div>;
  if (!articles || articles?.articles.length === 0)
    return <div>no articles found</div>;

  return (
    <div className="overflow-x-auto">
      <ul>
        {articles?.articles.map((article) => (
          <li
            key={article.id}
            className="border-b border-base-300 last:border-0"
          >
            <Link
              href={`/article/${article.id}`}
              className="block text-sm md:text-base lg:text-lg hover:bg-primary hover:text-white transition-colors duration-200"
            >
              {article.title}
            </Link>
          </li>
        ))}
      </ul>
      {/* TODO: fetch more button
      {articles.pagination?.next && (
        <button className="btn btn-primary mt-4">Load more</button>
      )} */}
    </div>
  );
}
