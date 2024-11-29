"use client";

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
    <div>
      <Header id={feedId} />
      <Articles feedId={feedId} />
    </div>
  );
}
