'use client'

import Link from 'next/link';
import useFeeds from '../../data/feeds';

export default function Feeds() {
  const { loading, error, feeds, mutate } = useFeeds();

  if (error) return <div>failed to load</div>
  if (loading) return <div>loading...</div>
  if (!feeds || feeds?.length === 0) return <div>no feeds found</div>

  return (
      <div>
        <ul>
          {feeds?.map((feed) => (
            <li key={feed.id}>
              <Link href={`/feed/${feed.id}/articles`}>{feed.name}</Link>
            </li>
          ))}
        </ul>
      </div>
  );
}