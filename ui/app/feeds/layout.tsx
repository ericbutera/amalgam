"use client";

import { useParams } from "next/navigation";
import Feeds from "@/app/components/feed/list";

export default function FeedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const params = useParams();
  const feedId = Array.isArray(params.id) ? params.id[0] : (params.id ?? "");

  return (
    <div className="drawer lg:drawer-open">
      <input id="amalgam-sidebar" type="checkbox" className="drawer-toggle" />
      <div className="drawer-content">
        <div className="p-4">
          {/* Page content */}
          {children}

          <label
            htmlFor="amalgam-sidebar"
            className="btn btn-primary drawer-button lg:hidden"
          >
            Open drawer
          </label>
        </div>
      </div>
      <div className="drawer-side">
        <label
          htmlFor="amalgam-sidebar"
          aria-label="close sidebar"
          className="drawer-overlay"
        ></label>
        {/* Sidebar content */}
        <Feeds feedId={feedId} />
      </div>
    </div>
  );
}
