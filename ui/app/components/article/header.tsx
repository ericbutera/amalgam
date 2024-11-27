"use client";

import useFeed from "@/app/data/feed";

interface PageProps {
  id: string;
}

export default function Header({ id }: PageProps) {
  const { loading, feed, error } = useFeed(id);

  if (loading) return <div>loading...</div>;
  if (error) return <div>failed to load feed</div>;

  return <h1 className="text-2xl font-bold">{feed?.name}</h1>;
}
