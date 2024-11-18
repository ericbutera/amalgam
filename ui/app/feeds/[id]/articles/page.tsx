"use client";

import Feeds from "@/app/components/feed/list";
import Articles from "../../../components/article/list";

interface PageProps {
  params: {
    id: string;
  };
}

export default function Page({ params }: PageProps) {
  const feedId = params.id;
  return (
    <div className="flex">
      <Feeds feedId={feedId} />
      <div className="flex-1">
        <h1 className="text-2xl font-bold">Articles for Feed {feedId}</h1>
        <Articles id={feedId} />
      </div>
    </div>
  );
}
