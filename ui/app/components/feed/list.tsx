"use client";

import Link from "next/link";
import useFeeds from "../../data/feeds";

interface FeedsProps {
  feedId?: string;
}

export default function Feeds({ feedId }: FeedsProps) {
  const { loading, error, feeds } = useFeeds();

  if (error) return <div>failed to load</div>;
  if (loading) return <div>loading...</div>;
  if (!feeds || feeds?.length === 0) return <div>no feeds found</div>;

  const isHighlight = (id: string) => (id == feedId ? "active" : "");

  return (
    <div>
      <ul className="menu bg-base-200 w-56">
        {feeds?.map((feed) => (
          <li key={feed.id}>
            <Link
              href={`/feeds/${feed.id}/articles`}
              className={isHighlight(feed.id)}
            >
              {feed.name || feed.url}
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
}
