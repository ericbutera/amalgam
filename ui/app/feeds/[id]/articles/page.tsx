"use client";

import Articles from "@/app/components/article/list";
import Header from "@/app/components/article/header";
import { useSearchParams } from "next/navigation";

interface PageProps {
  params: {
    id: string;
  };
}

export default function Page({ params }: PageProps) {
  const searchParams = useSearchParams();
  const cursor = searchParams.get("cursor") || undefined;
  const feedId = params.id;

  return (
    <div>
      <Header id={feedId} />
      <Articles feedId={feedId} cursor={cursor} />
    </div>
  );
}
