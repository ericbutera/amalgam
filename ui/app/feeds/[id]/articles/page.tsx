"use client";

import Articles from "@/app/components/article/list";
import Header from "@/app/components/article/header";
import { usePagination } from "@/app/hooks/usePagination";

interface PageProps {
  params: {
    id: string;
  };
}

export default function Page({ params }: PageProps) {
  const feedId = params.id;
  const pagination = usePagination();

  return (
    <div>
      <Header id={feedId} />
      <Articles feedId={feedId} pagination={pagination} />
    </div>
  );
}
