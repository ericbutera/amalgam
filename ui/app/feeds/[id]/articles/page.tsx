"use client";

import Feeds from "@/app/components/feed/list";
import Articles from "@/app/components/article/list";
import Header from "@/app/components/article/header";

interface PageProps {
  params: {
    id: string;
  };
}

export default function Page({ params }: PageProps) {
  const feedId = params.id;

  return (
    <div className="flex w-full">
      <Feeds feedId={feedId} />
      <div className="flex-1 w-full">
        <Header id={feedId} />
        <Articles id={feedId} />
      </div>
    </div>
  );
}
