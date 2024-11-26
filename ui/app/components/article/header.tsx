"use client";

import useFeed from "@/app/data/feed";

interface PageProps {
  id: string;
}

export default function Page({ id }: PageProps) {
  const { loading, feed, error } = useFeed(id);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error loading feed</div>;

  return <h1 className="text-2xl font-bold">{feed?.name}</h1>;
}
