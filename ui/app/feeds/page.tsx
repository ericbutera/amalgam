'use client'

import useSWR from 'swr'
import Link from 'next/link'

export default function Feeds() {
  const fetcher = (url) => fetch(url).then((res) => res.json());
  const { data, error, isLoading } = useSWR('http://localhost:8080/feeds', fetcher)

  if (error) return <div>failed to load</div>
  if (isLoading) return <div>loading...</div>

  return (
    <div>
      <div>
        <h1>feed list</h1>
      </div>
      <div>
        <ul>
          {data.feeds.map((feed) => (
            <li key={feed.ID}>
              <Link href={`/feed/${feed.ID}/articles`}>{feed.Name}</Link>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}