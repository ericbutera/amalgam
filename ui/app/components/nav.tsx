"use client";
import Link from "next/link";

import {
  ArrowPathIcon,
  BarsArrowDownIcon,
  SparklesIcon,
} from "@heroicons/react/24/solid";
import { getGraph } from "../lib/fetch";
import withToast from "../lib/toast";

const generateFeeds = async () => {
  await withToast(
    "Generating feeds...",
    "Feeds generated successfully!",
    "Failed to generate feeds.",
    async () => {
      await getGraph().GenerateFeeds();
    }
  );
};

const fetchFeeds = async () => {
  await withToast(
    "Fetching feeds...",
    "Feeds fetched successfully!",
    "Failed to fetch feeds.",
    async () => {
      await getGraph().FetchFeeds();
    }
  );
};

export default function Nav() {
  return (
    <div className="navbar bg-base-100">
      <div className="navbar-none">
        <Link href="/" className="btn btn-ghost text-xl">
          home
        </Link>
      </div>
      <div className="flex-none">
        <Link href="/feeds">feeds</Link>
      </div>
      <div className="flex-none">
        <div className="dropdown dropdown-end">
          <label tabIndex={0} className="btn btn-ghost">
            Actions
            <BarsArrowDownIcon className="w-6 h-6 text-default-500" />
          </label>
          <ul
            tabIndex={0}
            className="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-52"
          >
            <li>
              <a onClick={generateFeeds}>
                <SparklesIcon className="w-6 h-6 text-default-500" /> Generate
              </a>
            </li>
            <li>
              <a onClick={fetchFeeds}>
                <ArrowPathIcon className="w-6 h-6 text-default-500" /> Refresh
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
}