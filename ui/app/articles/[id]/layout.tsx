"use client";

import Feeds from "@/app/components/feed/list";
import useArticle from "@/app/data/article";

export default function ArticleLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { id: string };
}) {
  const { loading, error, article } = useArticle(params.id);
  if (loading) return <div>Loading...</div>;
  if (error) return <div>An error has occurred.</div>;

  return (
    <div className="drawer lg:drawer-open">
      <input id="amalgam-sidebar" type="checkbox" className="drawer-toggle" />
      <div className="drawer-content">
        {/* Page content */}
        {children}

        <label
          htmlFor="amalgam-sidebar"
          className="btn btn-primary drawer-button lg:hidden"
        >
          Open drawer
        </label>
      </div>
      <div className="drawer-side">
        <label
          htmlFor="amalgam-sidebar"
          aria-label="close sidebar"
          className="drawer-overlay"
        ></label>
        {/* Sidebar content */}
        <Feeds feedId={article?.feedId} />
      </div>
    </div>
  );
}
