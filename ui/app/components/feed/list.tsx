'use client'

import Link from 'next/link';
import useFeeds from '../../data/feeds';

export default function Feeds() {
  const { loading, error, feeds, mutate } = useFeeds();

  if (error) return <div>failed to load</div>
  if (loading) return <div>loading...</div>

  return (
      <div>
        <ul>
          {feeds.map((feed) => (
            <li key={feed.ID}>
              <Link href={`/feed/${feed.ID}/articles`}>{feed.Name}</Link>
            </li>
          ))}
        </ul>
      </div>
  );
}